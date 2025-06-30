package repository

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/model/expense"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"MyCar/internal/utils"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"time"
)

type PostgresRepository struct {
	db     *sql.DB
	redis  *redis.Client
	logTTL time.Duration
}

func NewPostgresRepository(pgDsn, redisAddr string, logTTL time.Duration) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", pgDsn)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	// Проверка соединения с Redis
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db:     db,
		redis:  rdb,
		logTTL: logTTL,
	}, nil
}

func (r *PostgresRepository) logChange(ctx context.Context, action string, entityType string, entityId interface{}, data interface{}) {
	key := fmt.Sprintf("log:%s:%v:%d", entityType, entityId, time.Now().UnixNano())
	entry := map[string]interface{}{
		"action":     action,
		"entityType": entityType,
		"entityId":   entityId,
		"data":       data,
		"timestamp":  time.Now().UTC(),
	}
	val, _ := json.Marshal(entry)
	r.redis.Set(ctx, key, val, r.logTTL)
}

func (r *PostgresRepository) SaveEntity(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
	ctx := context.Background()
	switch e := entity.(type) {
	case *auth.User:
		id := e.Id
		if id == uuid.Nil {
			id = uuid.New()
		}
		_, err := r.db.Exec(`
			INSERT INTO users (id, login, password, created_at_utc)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (id) DO UPDATE SET
				login = EXCLUDED.login,
				password = EXCLUDED.password
		`, id, e.Login, e.Password, time.Now().UTC())
		if err != nil {
			return uuid.Nil, model.NewApplicationError(model.ErrorTypeDatabase, "SaveEntity error", err)
		}
		r.logChange(ctx, "upsert", "user", id, e)
		return id, nil
	case *car.Car:
		id := e.Id
		if id == uuid.Nil {
			id = uuid.New()
		}
		_, err := r.db.Exec(`
			INSERT INTO cars (id, user_id, fuel_type_id, vehicle_type_id, year, plate, vin, brand_id, drive_type_id, body_type_id, transmission_type_id, created_at_utc)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
			ON CONFLICT (id) DO UPDATE SET
				user_id = EXCLUDED.user_id,
				fuel_type_id = EXCLUDED.fuel_type_id,
				vehicle_type_id = EXCLUDED.vehicle_type_id,
				year = EXCLUDED.year,
				plate = EXCLUDED.plate,
				vin = EXCLUDED.vin,
				brand_id = EXCLUDED.brand_id,
				drive_type_id = EXCLUDED.drive_type_id,
				body_type_id = EXCLUDED.body_type_id,
				transmission_type_id = EXCLUDED.transmission_type_id
		`, id, e.UserId, e.FuelType, e.VehicleType, e.Year, e.Plate, e.Vin, e.Brand.Id, e.DriveType.Id, e.BodyType.Id, e.TransmissionType.Id, time.Now().UTC())
		if err != nil {
			return uuid.Nil, model.NewApplicationError(model.ErrorTypeDatabase, "SaveEntity error", err)
		}
		r.logChange(ctx, "upsert", "car", id, e)
		return id, nil
	case *moto.Moto:
		id := e.Id
		if id == uuid.Nil {
			id = uuid.New()
		}
		_, err := r.db.Exec(`
			INSERT INTO motos (id, user_id, fuel_type_id, vehicle_type_id, year, plate, vin, brand_id, category_id, transmission_type_id, created_at_utc)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
			ON CONFLICT (id) DO UPDATE SET
				user_id = EXCLUDED.user_id,
				fuel_type_id = EXCLUDED.fuel_type_id,
				vehicle_type_id = EXCLUDED.vehicle_type_id,
				year = EXCLUDED.year,
				plate = EXCLUDED.plate,
				vin = EXCLUDED.vin,
				brand_id = EXCLUDED.brand_id,
				category_id = EXCLUDED.category_id,
				transmission_type_id = EXCLUDED.transmission_type_id
		`, id, e.UserId, e.FuelType, e.VehicleType, e.Year, e.Plate, e.Vin, e.Brand.Id, e.Category.Id, e.TransmissionType.Id, time.Now().UTC())
		if err != nil {
			return uuid.Nil, model.NewApplicationError(model.ErrorTypeDatabase, "SaveEntity error", err)
		}
		r.logChange(ctx, "upsert", "moto", id, e)
		return id, nil
	case *expense.Expense:
		id := e.Id
		if id == uuid.Nil {
			id = uuid.New()
		}
		_, err := r.db.Exec(`
			INSERT INTO expenses (id, vehicle_type_id, vehicle_id, category_id, amount, currency, exchange_rate, date, note, created_at_utc)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
			ON CONFLICT (id) DO UPDATE SET
				vehicle_type_id = EXCLUDED.vehicle_type_id,
				vehicle_id = EXCLUDED.vehicle_id,
				category_id = EXCLUDED.category_id,
				amount = EXCLUDED.amount,
				currency = EXCLUDED.currency,
				exchange_rate = EXCLUDED.exchange_rate,
				date = EXCLUDED.date,
				note = EXCLUDED.note
		`, id, e.VehicleType, e.VehicleId, e.Category, e.Amount, e.Currency, e.ExchangeRate, e.Date, e.Note, time.Now().UTC())
		if err != nil {
			return uuid.Nil, model.NewApplicationError(model.ErrorTypeDatabase, "SaveEntity error", err)
		}
		r.logChange(ctx, "upsert", "expense", id, e)
		return id, nil
	default:
		return uuid.Nil, model.NewApplicationError(model.ErrorTypeDatabase, "Unknown entity type", nil)
	}
}

func (r *PostgresRepository) DeleteEntity(entity model.BusinessEntity) *model.ApplicationError {
	ctx := context.Background()
	switch e := entity.(type) {
	case *auth.User:
		_, err := r.db.Exec(`DELETE FROM users WHERE id = $1`, e.Id)
		if err != nil {
			return model.NewApplicationError(model.ErrorTypeDatabase, "DeleteEntity error", err)
		}
		r.logChange(ctx, "delete", "user", e.Id, e)
	case *car.Car:
		_, err := r.db.Exec(`DELETE FROM cars WHERE id = $1`, e.Id)
		if err != nil {
			return model.NewApplicationError(model.ErrorTypeDatabase, "DeleteEntity error", err)
		}
		r.logChange(ctx, "delete", "car", e.Id, e)
	case *moto.Moto:
		_, err := r.db.Exec(`DELETE FROM motos WHERE id = $1`, e.Id)
		if err != nil {
			return model.NewApplicationError(model.ErrorTypeDatabase, "DeleteEntity error", err)
		}
		r.logChange(ctx, "delete", "moto", e.Id, e)
	case *expense.Expense:
		_, err := r.db.Exec(`DELETE FROM expenses WHERE id = $1`, e.Id)
		if err != nil {
			return model.NewApplicationError(model.ErrorTypeDatabase, "DeleteEntity error", err)
		}
		r.logChange(ctx, "delete", "expense", e.Id, e)
	case *model.Attachment:
		_, err := r.db.Exec(`DELETE FROM attachments WHERE id = $1`, e.Id)
		if err != nil {
			return model.NewApplicationError(model.ErrorTypeDatabase, "DeleteEntity error", err)
		}
		r.logChange(ctx, "delete", "attachment", e.Id, e)
	default:
		return model.NewApplicationError(model.ErrorTypeDatabase, "Unknown entity type", nil)
	}
	return nil
}

func (r *PostgresRepository) GetUsers() []*auth.User {
	rows, err := r.db.Query(`SELECT id, login, password FROM users`)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var users []*auth.User
	for rows.Next() {
		var u auth.User
		if err := rows.Scan(&u.Id, &u.Login, &u.Password); err == nil {
			users = append(users, &u)
		}
	}
	return users
}

func (r *PostgresRepository) GetCars() []*car.Car {
	rows, err := r.db.Query(`SELECT id, user_id, fuel_type_id, vehicle_type_id, year, plate, vin, brand_id, drive_type_id, body_type_id, transmission_type_id FROM cars`)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var cars []*car.Car
	for rows.Next() {
		var c car.Car
		if err := rows.Scan(&c.Id, &c.UserId, &c.FuelType, &c.VehicleType, &c.Year, &c.Plate, &c.Vin, &c.Brand.Id, &c.DriveType.Id, &c.BodyType.Id, &c.TransmissionType.Id); err == nil {
			cars = append(cars, &c)
		}
	}
	return cars
}

func (r *PostgresRepository) GetMotos() []*moto.Moto {
	rows, err := r.db.Query(`SELECT id, user_id, fuel_type_id, vehicle_type_id, year, plate, vin, brand_id, category_id, transmission_type_id FROM motos`)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var motos []*moto.Moto
	for rows.Next() {
		var m moto.Moto
		if err := rows.Scan(&m.Id, &m.UserId, &m.FuelType, &m.VehicleType, &m.Year, &m.Plate, &m.Vin, &m.Brand.Id, &m.Category.Id, &m.TransmissionType.Id); err == nil {
			motos = append(motos, &m)
		}
	}
	return motos
}

func (r *PostgresRepository) GetExpenses() []*expense.Expense {
	rows, err := r.db.Query(`SELECT id, vehicle_type_id, vehicle_id, category_id, amount, currency, exchange_rate, date, note FROM expenses`)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var expenses []*expense.Expense
	for rows.Next() {
		var e expense.Expense
		if err := rows.Scan(&e.Id, &e.VehicleType, &e.VehicleId, &e.Category, &e.Amount, &e.Currency, &e.ExchangeRate, &e.Date, &e.Note); err == nil {
			expenses = append(expenses, &e)
		}
	}
	return expenses
}

func (r *PostgresRepository) GetAttachments() []*model.Attachment {
	rows, err := r.db.Query(`SELECT id, user_id, file_name, file_path FROM attachments`)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var attachments []*model.Attachment
	for rows.Next() {
		var a model.Attachment
		if err := rows.Scan(&a.Id, &a.UserId, &a.FileName, &a.FilePath); err == nil {
			attachments = append(attachments, &a)
		}
	}
	return attachments
}

func (r *PostgresRepository) GetUserById(id uuid.UUID) (*auth.User, *model.ApplicationError) {
	row := r.db.QueryRow(`SELECT id, login, password FROM users WHERE id = $1`, id)
	var u auth.User
	if err := row.Scan(&u.Id, &u.Login, &u.Password); err != nil {
		return nil, model.NewApplicationError(model.ErrorTypeDatabase, "GetUserById error", err)
	}
	return &u, nil
}

func (r *PostgresRepository) GetCarById(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
	row := r.db.QueryRow(`SELECT id, user_id, fuel_type_id, vehicle_type_id, year, plate, vin, brand_id, drive_type_id, body_type_id, transmission_type_id FROM cars WHERE id = $1 AND user_id = $2`, id, userId)
	var c car.Car
	if err := row.Scan(&c.Id, &c.UserId, &c.FuelType, &c.VehicleType, &c.Year, &c.Plate, &c.Vin, &c.Brand.Id, &c.DriveType.Id, &c.BodyType.Id, &c.TransmissionType.Id); err != nil {
		return nil, model.NewApplicationError(model.ErrorTypeDatabase, "GetCarById error", err)
	}
	return &c, nil
}

func (r *PostgresRepository) GetMotoById(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
	row := r.db.QueryRow(`SELECT id, user_id, fuel_type_id, vehicle_type_id, year, plate, vin, brand_id, category_id, transmission_type_id FROM motos WHERE id = $1 AND user_id = $2`, id, userId)
	var m moto.Moto
	if err := row.Scan(&m.Id, &m.UserId, &m.FuelType, &m.VehicleType, &m.Year, &m.Plate, &m.Vin, &m.Brand.Id, &m.Category.Id, &m.TransmissionType.Id); err != nil {
		return nil, model.NewApplicationError(model.ErrorTypeDatabase, "GetMotoById error", err)
	}
	return &m, nil
}

func (r *PostgresRepository) GetExpenseById(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
	row := r.db.QueryRow(`SELECT id, vehicle_type_id, vehicle_id, category_id, amount, currency, exchange_rate, date, note FROM expenses WHERE id = $1`, id)
	var e expense.Expense
	if err := row.Scan(&e.Id, &e.VehicleType, &e.VehicleId, &e.Category, &e.Amount, &e.Currency, &e.ExchangeRate, &e.Date, &e.Note); err != nil {
		return nil, model.NewApplicationError(model.ErrorTypeDatabase, "GetExpenseById error", err)
	}
	return &e, nil
}

func (r *PostgresRepository) GetAttachmentById(id uuid.UUID, userId uuid.UUID) (*model.Attachment, *model.ApplicationError) {
	row := r.db.QueryRow(`SELECT id, user_id, file_name, file_path FROM attachments WHERE id = $1 AND user_id = $2`, id, userId)
	var a model.Attachment
	if err := row.Scan(&a.Id, &a.UserId, &a.FileName, &a.FilePath); err != nil {
		return nil, model.NewApplicationError(model.ErrorTypeDatabase, "GetAttachmentById error", err)
	}
	return &a, nil
}

func (r *PostgresRepository) GetUsersCount() int {
	var count int
	_ = r.db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&count)
	return count
}

func (r *PostgresRepository) GetCarsCount() int {
	var count int
	_ = r.db.QueryRow(`SELECT COUNT(*) FROM cars`).Scan(&count)
	return count
}

func (r *PostgresRepository) GetMotosCount() int {
	var count int
	_ = r.db.QueryRow(`SELECT COUNT(*) FROM motos`).Scan(&count)
	return count
}

func (r *PostgresRepository) GetExpensesCount() int {
	var count int
	_ = r.db.QueryRow(`SELECT COUNT(*) FROM expenses`).Scan(&count)
	return count
}

func (r *PostgresRepository) GetAttachmentsCount() int {
	var count int
	_ = r.db.QueryRow(`SELECT COUNT(*) FROM attachments`).Scan(&count)
	return count
}

func (r *PostgresRepository) GetUser(login, password string) (*auth.User, *model.ApplicationError) {
	row := r.db.QueryRow(`SELECT id, login, password FROM users WHERE login = $1`, login)
	var u auth.User
	if err := row.Scan(&u.Id, &u.Login, &u.Password); err != nil {
		return nil, model.NewApplicationError(model.ErrorTypeDatabase, "GetUser error", err)
	}

	arePasswordsEqual, errCmp := utils.CompareHashAndPassword(u.Password, password)
	if errCmp != nil || !arePasswordsEqual {
		return nil, model.NewApplicationError(model.ErrorTypeAuth, "Пользователь не найден или неверный пароль", nil)
	}
	return &u, nil
}

func (r *PostgresRepository) GetCarsByUserId(userId uuid.UUID) []*car.Car {
	rows, err := r.db.Query(`SELECT id, user_id, fuel_type_id, vehicle_type_id, year, plate, vin, brand_id, drive_type_id, body_type_id, transmission_type_id FROM cars WHERE user_id = $1`, userId)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var cars []*car.Car
	for rows.Next() {
		var c car.Car
		if err := rows.Scan(&c.Id, &c.UserId, &c.FuelType, &c.VehicleType, &c.Year, &c.Plate, &c.Vin, &c.Brand.Id, &c.DriveType.Id, &c.BodyType.Id, &c.TransmissionType.Id); err == nil {
			cars = append(cars, &c)
		}
	}
	return cars
}

func (r *PostgresRepository) GetMotosByUserId(userId uuid.UUID) []*moto.Moto {
	rows, err := r.db.Query(`SELECT id, user_id, fuel_type_id, vehicle_type_id, year, plate, vin, brand_id, category_id, transmission_type_id FROM motos WHERE user_id = $1`, userId)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var motos []*moto.Moto
	for rows.Next() {
		var m moto.Moto
		if err := rows.Scan(&m.Id, &m.UserId, &m.FuelType, &m.VehicleType, &m.Year, &m.Plate, &m.Vin, &m.Brand.Id, &m.Category.Id, &m.TransmissionType.Id); err == nil {
			motos = append(motos, &m)
		}
	}
	return motos
}

func (r *PostgresRepository) GetExpensesByUserId(userId uuid.UUID) []*expense.Expense {
	rows, err := r.db.Query(`SELECT e.id, e.vehicle_type_id, e.vehicle_id, e.category_id, e.amount, e.currency, e.exchange_rate, e.date, e.note
		FROM expenses e
		JOIN cars c ON e.vehicle_id = c.id
		WHERE c.user_id = $1
		UNION
		SELECT e.id, e.vehicle_type_id, e.vehicle_id, e.category_id, e.amount, e.currency, e.exchange_rate, e.date, e.note
		FROM expenses e
		JOIN motos m ON e.vehicle_id = m.id
		WHERE m.user_id = $1`, userId)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var expenses []*expense.Expense
	for rows.Next() {
		var e expense.Expense
		if err := rows.Scan(&e.Id, &e.VehicleType, &e.VehicleId, &e.Category, &e.Amount, &e.Currency, &e.ExchangeRate, &e.Date, &e.Note); err == nil {
			expenses = append(expenses, &e)
		}
	}
	return expenses
}

func (r *PostgresRepository) GetAttachmentsByUserId(userId uuid.UUID) []*model.Attachment {
	rows, err := r.db.Query(`SELECT id, user_id, file_name, file_path FROM attachments WHERE user_id = $1`, userId)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var attachments []*model.Attachment
	for rows.Next() {
		var a model.Attachment
		if err := rows.Scan(&a.Id, &a.UserId, &a.FileName, &a.FilePath); err == nil {
			attachments = append(attachments, &a)
		}
	}
	return attachments
}
