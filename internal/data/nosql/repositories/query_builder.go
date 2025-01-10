package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Filter interface {
	ById(id uuid.UUID) *QueryBuilder
	ByPlaceId(placeId uuid.UUID) *QueryBuilder
	ByEmployeeId(employeeId uuid.UUID) *QueryBuilder
	ByUserId(userId uuid.UUID) *QueryBuilder
	ByRole(role string) *QueryBuilder
	ByCreatedAt(start, end time.Time) *QueryBuilder
	ByUpdatedAt(start, end time.Time) *QueryBuilder

	SortBy(field string, ascending bool) *QueryBuilder

	NewLimit(limit int64) *QueryBuilder
	NewSkip(skip int64) *QueryBuilder

	Execute(ctx context.Context, result interface{}) error
}

type QueryBuilder struct {
	Collection *mongo.Collection
	Filters    bson.M
	Sort       bson.D
	Limit      int64
	Skip       int64
}

// NewQueryBuilder возвращает новый экземпляр PlaceQueryBuilder
func NewQueryBuilder(collection *mongo.Collection) *QueryBuilder {
	return &QueryBuilder{
		Collection: collection,
		Filters:    bson.M{},
		Sort:       bson.D{},
	}
}

func (qb *QueryBuilder) ById(id uuid.UUID) *QueryBuilder {
	qb.Filters["_id"] = id
	return qb
}

func (qb *QueryBuilder) ByPlaceId(placeId uuid.UUID) *QueryBuilder {
	qb.Filters["place_id"] = placeId
	return qb
}

func (qb *QueryBuilder) ByEmployeeId(employeeId uuid.UUID) *QueryBuilder {
	qb.Filters["employee_id"] = employeeId
	return qb
}

func (qb *QueryBuilder) ByUserId(userId uuid.UUID) *QueryBuilder {
	qb.Filters["user_id"] = userId
	return qb
}

func (qb *QueryBuilder) ByRole(role string) *QueryBuilder {
	qb.Filters["role"] = role
	return qb
}

func (qb *QueryBuilder) ByCreatedAt(start, end time.Time) *QueryBuilder {
	qb.Filters["created_at"] = bson.M{
		"$gte": start,
		"$lte": end,
	}
	return qb
}

func (qb *QueryBuilder) ByUpdatedAt(start, end time.Time) *QueryBuilder {
	qb.Filters["updated_at"] = bson.M{
		"$gte": start,
		"$lte": end,
	}
	return qb
}

func (qb *QueryBuilder) SortBy(field string, ascending bool) *QueryBuilder {
	direction := 1
	if !ascending {
		direction = -1
	}
	qb.Sort = append(qb.Sort, bson.E{Key: field, Value: direction})
	return qb
}

func (qb *QueryBuilder) NewLimit(limit int64) *QueryBuilder {
	qb.Limit = limit
	return qb
}

func (qb *QueryBuilder) NewSkip(skip int64) *QueryBuilder {
	qb.Skip = skip
	return qb
}

func (qb *QueryBuilder) Execute(ctx context.Context, result interface{}) error {
	options := options.Find()
	if len(qb.Sort) > 0 {
		options.SetSort(qb.Sort)
	}
	if qb.Limit > 0 {
		options.SetLimit(qb.Limit)
	}
	if qb.Skip > 0 {
		options.SetSkip(qb.Skip)
	}

	cursor, err := qb.Collection.Find(ctx, qb.Filters, options)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, result)
}
