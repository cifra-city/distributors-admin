package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/cifra-city/distributors-admin/internal/data/nosql/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlacesEmployees interface {
	FilterById(id uuid.UUID) PlacesEmployees
	FilterByPlaceId(placeId uuid.UUID) PlacesEmployees
	FilterByEmployeeId(employeeId uuid.UUID) PlacesEmployees
	FilterByUserId(userId uuid.UUID) PlacesEmployees
	FilterByRole(role string) PlacesEmployees
	FilterByCreatedAt(start, end time.Time) PlacesEmployees
	FilterByUpdatedAt(start, end time.Time) PlacesEmployees
	SortBy(field string, ascending bool) PlacesEmployees
	Limit(limit int64) PlacesEmployees
	Skip(skip int64) PlacesEmployees

	Get(ctx context.Context) ([]models.PlaceEmployee, error)
	Delete(ctx context.Context) (int64, error)              // Возвращает количество удалённых записей
	Update(ctx context.Context, role string) (int64, error) // Возвращает количество обновлённых записей

	Create(ctx context.Context, placeID, employeeID, userID uuid.UUID, role string) (models.PlaceEmployee, error)
}

type placesEmployees struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func NewPlacesEmployees(uri, dbName, collectionName string) (PlacesEmployees, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	coll := database.Collection(collectionName)

	return &placesEmployees{
		client:     client,
		database:   database,
		collection: coll,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}, nil
}

// Фильтрация
func (p *placesEmployees) FilterById(id uuid.UUID) PlacesEmployees {
	p.filters["_id"] = id
	return p
}

func (p *placesEmployees) FilterByPlaceId(placeId uuid.UUID) PlacesEmployees {
	p.filters["place_id"] = placeId
	return p
}

func (p *placesEmployees) FilterByEmployeeId(employeeId uuid.UUID) PlacesEmployees {
	p.filters["employee_id"] = employeeId
	return p
}

func (p *placesEmployees) FilterByUserId(userId uuid.UUID) PlacesEmployees {
	p.filters["user_id"] = userId
	return p
}

func (p *placesEmployees) FilterByRole(role string) PlacesEmployees {
	p.filters["role"] = role
	return p
}

func (p *placesEmployees) FilterByCreatedAt(start, end time.Time) PlacesEmployees {
	p.filters["created_at"] = bson.M{"$gte": start, "$lte": end}
	return p
}

func (p *placesEmployees) FilterByUpdatedAt(start, end time.Time) PlacesEmployees {
	p.filters["updated_at"] = bson.M{"$gte": start, "$lte": end}
	return p
}

func (p *placesEmployees) SortBy(field string, ascending bool) PlacesEmployees {
	direction := 1
	if !ascending {
		direction = -1
	}
	p.sort = append(p.sort, bson.E{Key: field, Value: direction})
	return p
}

func (p *placesEmployees) Limit(limit int64) PlacesEmployees {
	p.limit = limit
	return p
}

func (p *placesEmployees) Skip(skip int64) PlacesEmployees {
	p.skip = skip
	return p
}

func (p *placesEmployees) Get(ctx context.Context) ([]models.PlaceEmployee, error) {
	findOptions := options.Find()
	if len(p.sort) > 0 {
		findOptions.SetSort(p.sort)
	}
	if p.limit > 0 {
		findOptions.SetLimit(p.limit)
	}
	if p.skip > 0 {
		findOptions.SetSkip(p.skip)
	}

	cursor, err := p.collection.Find(ctx, p.filters, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer cursor.Close(ctx)

	var results []models.PlaceEmployee
	if err := cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode results: %w", err)
	}

	// Сбрасываем фильтры после выполнения запроса
	p.filters = bson.M{}
	p.sort = bson.D{}
	p.limit = 0
	p.skip = 0

	return results, nil
}

func (p *placesEmployees) Update(ctx context.Context, role string) (int64, error) {
	if len(p.filters) == 0 {
		return 0, fmt.Errorf("no filters specified for update")
	}

	update := bson.M{
		"$set": bson.M{
			"role":       role,
			"updated_at": time.Now().UTC(),
		},
	}

	result, err := p.collection.UpdateMany(ctx, p.filters, update)
	if err != nil {
		return 0, fmt.Errorf("failed to update records: %w", err)
	}

	// Сбрасываем фильтры после выполнения операции
	p.filters = bson.M{}

	return result.ModifiedCount, nil
}

func (p *placesEmployees) Delete(ctx context.Context) (int64, error) {
	if len(p.filters) == 0 {
		return 0, fmt.Errorf("no filters specified for delete")
	}

	result, err := p.collection.DeleteMany(ctx, p.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to delete records: %w", err)
	}

	// Сбрасываем фильтры после выполнения операции
	p.filters = bson.M{}

	return result.DeletedCount, nil
}

func (p *placesEmployees) Create(ctx context.Context, placeID, employeeID, userID uuid.UUID, role string) (models.PlaceEmployee, error) {
	placeEmployee := models.PlaceEmployee{
		ID:         uuid.New(),
		PlaceID:    placeID,
		EmployeeID: employeeID,
		UserID:     userID,
		Role:       role,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}

	_, err := p.collection.InsertOne(ctx, placeEmployee)
	if err != nil {
		return models.PlaceEmployee{}, fmt.Errorf("failed to create place employee: %w", err)
	}

	return placeEmployee, nil
}
