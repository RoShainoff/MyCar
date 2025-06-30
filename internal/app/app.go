package app

import (
	"MyCar/config"
	_ "MyCar/docs"
	"MyCar/internal/http/handler"
	"MyCar/internal/http/middleware"
	"MyCar/internal/repository"
	"MyCar/internal/service"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	fmt.Println("System running... Press Ctrl+C to stop")

	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	pgDsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.SqlDatabase.Host, cfg.SqlDatabase.User, cfg.SqlDatabase.Password, cfg.SqlDatabase.Name, cfg.SqlDatabase.Sslmode)
	redisAddr := cfg.Redis.Addr
	logTTL := time.Duration(cfg.Redis.LogTTLSeconds) * time.Second

	pgRepo, err := repository.NewPostgresRepository(pgDsn, redisAddr, logTTL)
	if err != nil {
		log.Fatalf("PostgresRepository error: %s", err)
	}

	mongoRepo, err := repository.NewMongoRepository(cfg.NoSqlDatabase.Uri, redisAddr, logTTL)
	if err != nil {
		log.Fatalf("MongoRepository error: %s", err)
	}

	repo := repository.NewCombinedRepository(pgRepo, mongoRepo)

	jwtService := service.NewJwtService(cfg)
	authService := service.NewAuthService(repo, jwtService)
	carService := service.NewCarService(repo)
	motoService := service.NewMotoService(repo)
	expenseService := service.NewExpenseService(repo)
	userService := service.NewUserService(repo)
	attachmentService := service.NewAttachmentService(repo)

	authHandler := handler.NewAuthHandler(authService)
	carHandler := handler.NewCarHandler(carService)
	motoHandler := handler.NewMotoHandler(motoService)
	expenseHandler := handler.NewExpenseHandler(expenseService)
	userHandler := handler.NewUserHandler(userService)
	attachmentHandler := handler.NewAttachmentHandler(attachmentService)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authMiddleware := middleware.AuthMiddleware(authService)

	protected := r.Group("/api")
	protected.Use(authMiddleware)
	{
		protected.GET("/user", userHandler.GetUser)
		protected.PUT("/user", userHandler.UpdateUser)
		protected.DELETE("/user", userHandler.DeleteUser)

		protected.POST("/car", carHandler.CreateCar)
		protected.GET("/car/:id", carHandler.GetCarById)
		protected.PUT("/car/:id", carHandler.UpdateCar)
		protected.DELETE("/car/:id", carHandler.DeleteCar)

		protected.POST("/moto", motoHandler.CreateMoto)
		protected.GET("/moto/:id", motoHandler.GetMotoById)
		protected.PUT("/moto/:id", motoHandler.UpdateMoto)
		protected.DELETE("/moto/:id", motoHandler.DeleteMoto)

		protected.POST("/expense", expenseHandler.CreateExpense)
		protected.GET("/expense/:id", expenseHandler.GetExpenseById)
		protected.PUT("/expense/:id", expenseHandler.UpdateExpense)
		protected.DELETE("/expense/:id", expenseHandler.DeleteExpense)

		protected.POST("/attachment", attachmentHandler.UploadAttachment)
		protected.GET("/attachment/:id", attachmentHandler.GetAttachmentById)
		protected.DELETE("/attachment/:id", attachmentHandler.DeleteAttachment)
	}

	r.POST("/api/auth/login", authHandler.Login)
	r.POST("/api/user", userHandler.CreateUser)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %s", err)
		}
	}()

	log.Printf("Server started on port %d", cfg.Server.Port)

	fmt.Println("System running... Press Ctrl+C to stop")

	<-ctx.Done()
	fmt.Println("\nShutting down gracefully...")

	time.Sleep(1 * time.Second)
}
