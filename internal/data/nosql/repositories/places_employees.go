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
	Add(ctx context.Context,
		placeID uuid.UUID,
		employeeID uuid.UUID,
		userID uuid.UUID,
		role string,
	) (models.PlaceEmployee, error)
	Update(ctx context.Context,
		id uuid.UUID,
		role string,
	) (models.PlaceEmployee, error)
	Delete(ctx context.Context, id uuid.UUID) error

	Filter() *QueryBuilder
}

type placesEmployees struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
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
	}, nil
}

func (p *placesEmployees) Filter() *QueryBuilder {
	return &QueryBuilder{
		Collection: p.collection,
		Filters:    bson.M{},
		Sort:       bson.D{},
	}
}

func (p *placesEmployees) Add(ctx context.Context,
	placeID uuid.UUID,
	employeeID uuid.UUID,
	userID uuid.UUID,
	role string,
) (models.PlaceEmployee, error) {
	place := models.PlaceEmployee{
		ID:         uuid.New(),
		PlaceID:    placeID,
		EmployeeID: employeeID,
		UserID:     userID,
		Role:       role,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}
	_, err := p.collection.InsertOne(ctx, place)
	if err != nil {
		return models.PlaceEmployee{}, fmt.Errorf("failed to add place: %w", err)
	}

	return place, nil
}

func (p *placesEmployees) Update(ctx context.Context,
	id uuid.UUID,
	role string,
) (models.PlaceEmployee, error) {
	place := models.PlaceEmployee{
		Role:      role,
		UpdatedAt: time.Now().UTC(),
	}
	_, err := p.collection.ReplaceOne(ctx, bson.M{"_id": id}, place)
	if err != nil {
		return models.PlaceEmployee{}, fmt.Errorf("failed to update place: %w", err)
	}

	return place, nil
}

func (p *placesEmployees) Delete(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id}
	var deletedPlace models.PlaceEmployee
	err := p.collection.FindOneAndDelete(ctx, filter).Decode(&deletedPlace)
	if err != nil {
		fmt.Errorf("failed to delete place: %w", err)
	}
	return nil
}
