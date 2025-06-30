package repository

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/model/expense"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
	redis      *redis.Client
	logTTL     time.Duration
}

func NewMongoRepository(mongoURI, redisAddr string, logTTL time.Duration) (*MongoRepository, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	coll := client.Database("mycar").Collection("attachments")
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	return &MongoRepository{
		client:     client,
		collection: coll,
		redis:      redisClient,
		logTTL:     logTTL,
	}, nil
}

func (r *MongoRepository) logChange(action string, att *model.Attachment) {
	logEntry := map[string]interface{}{
		"action":     action,
		"attachment": att,
		"timestamp":  time.Now().UTC(),
	}
	data, _ := json.Marshal(logEntry)
	key := "attachment_log:" + att.GetId().String() + ":" + time.Now().Format(time.RFC3339Nano)
	r.redis.Set(context.Background(), key, data, r.logTTL)
}

func (r *MongoRepository) SaveEntity(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
	att, ok := entity.(*model.Attachment)
	if !ok {
		return uuid.Nil, model.NewApplicationError(model.ErrorTypeDatabase, "Unsupported entity type", nil)
	}
	isNew := att.GetId() == uuid.Nil
	if isNew {
		att.SetId(uuid.New())
	}
	filter := bson.M{"id": att.GetId().String()}
	update := bson.M{"$set": att}
	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return uuid.Nil, model.NewApplicationError(model.ErrorTypeDatabase, "Mongo save error", err)
	}
	if isNew {
		r.logChange("create", att)
	} else {
		r.logChange("update", att)
	}
	return att.GetId(), nil
}

func (r *MongoRepository) DeleteEntity(entity model.BusinessEntity) *model.ApplicationError {
	att, ok := entity.(*model.Attachment)
	if !ok {
		return model.NewApplicationError(model.ErrorTypeDatabase, "Unsupported entity type", nil)
	}
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"id": att.GetId().String()})
	if err != nil {
		return model.NewApplicationError(model.ErrorTypeDatabase, "Mongo delete error", err)
	}
	r.logChange("delete", att)
	return nil
}

func (r *MongoRepository) GetAttachments() []*model.Attachment {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())
	var result []*model.Attachment
	for cursor.Next(context.Background()) {
		var att model.Attachment
		if err := cursor.Decode(&att); err == nil {
			result = append(result, &att)
		}
	}
	return result
}

func (r *MongoRepository) GetAttachmentById(id uuid.UUID, userId uuid.UUID) (*model.Attachment, *model.ApplicationError) {
	var att model.Attachment
	filter := bson.D{
		{"id", primitive.Binary{Subtype: 0, Data: id[:]}},
	}
	err := r.collection.FindOne(context.Background(), filter).Decode(&att)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, model.NewApplicationError(model.ErrorTypeNotFound, "Вложение не найдено", nil)
		}
		return nil, model.NewApplicationError(model.ErrorTypeDatabase, "Mongo find error", err)
	}
	return &att, nil
}

func (r *MongoRepository) GetAttachmentsByUserId(userId uuid.UUID) []*model.Attachment {
	cursor, err := r.collection.Find(context.Background(), bson.M{"auditfields.createdby": userId})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())
	var result []*model.Attachment
	for cursor.Next(context.Background()) {
		var att model.Attachment
		if err := cursor.Decode(&att); err == nil {
			result = append(result, &att)
		}
	}
	return result
}

func (r *MongoRepository) GetAttachmentsCount() int {
	count, err := r.collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return 0
	}
	return int(count)
}

// --- Методы-заглушки для остальных сущностей ---

func (r *MongoRepository) GetUsers() []*auth.User          { return nil }
func (r *MongoRepository) GetCars() []*car.Car             { return nil }
func (r *MongoRepository) GetMotos() []*moto.Moto          { return nil }
func (r *MongoRepository) GetExpenses() []*expense.Expense { return nil }
func (r *MongoRepository) GetUserById(id uuid.UUID) (*auth.User, *model.ApplicationError) {
	return nil, model.NewApplicationError(model.ErrorTypeDatabase, "Not supported", nil)
}
func (r *MongoRepository) GetCarById(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
	return nil, model.NewApplicationError(model.ErrorTypeDatabase, "Not supported", nil)
}
func (r *MongoRepository) GetMotoById(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
	return nil, model.NewApplicationError(model.ErrorTypeDatabase, "Not supported", nil)
}
func (r *MongoRepository) GetExpenseById(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
	return nil, model.NewApplicationError(model.ErrorTypeDatabase, "Not supported", nil)
}
func (r *MongoRepository) GetUsersCount() int    { return 0 }
func (r *MongoRepository) GetCarsCount() int     { return 0 }
func (r *MongoRepository) GetMotosCount() int    { return 0 }
func (r *MongoRepository) GetExpensesCount() int { return 0 }
func (r *MongoRepository) GetUser(login, password string) (*auth.User, *model.ApplicationError) {
	return nil, model.NewApplicationError(model.ErrorTypeDatabase, "Not supported", nil)
}
func (r *MongoRepository) GetCarsByUserId(userId uuid.UUID) []*car.Car    { return nil }
func (r *MongoRepository) GetMotosByUserId(userId uuid.UUID) []*moto.Moto { return nil }
func (r *MongoRepository) GetExpensesByUserId(userId uuid.UUID) []*expense.Expense {
	return nil
}
