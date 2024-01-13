package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type MockCollection[T any] struct {
	Collection[T]
	MockFindOneResult *[]T
	MockFindOneExists *[]bool
	MockFindOneErr    *[]error
	MockFindResult    *[][]T
	MockFindErr       *[]error
	MockUpsertErr     *[]error
	MockUpdateOneErr  *[]error
}

func (m *MockCollection[T]) FindOne(ctx context.Context, filter bson.M) (T, bool, error) {
	// Get the first element of the result array and remove it from the array, return nil if the array is empty
	var result T
	if len(*m.MockFindOneResult) > 0 {
		result = (*m.MockFindOneResult)[0]
		*m.MockFindOneResult = (*m.MockFindOneResult)[1:]
	}

	// Get the first element of the exists array and remove it from the array, return nil if the array is empty
	var exists bool
	if len(*m.MockFindOneExists) > 0 {
		exists = (*m.MockFindOneExists)[0]
		*m.MockFindOneExists = (*m.MockFindOneExists)[1:]
	}

	// Get the first element of the error array and remove it from the array, return nil if the array is empty
	var err error
	if len(*m.MockFindOneErr) > 0 {
		err = (*m.MockFindOneErr)[0]
		*m.MockFindOneErr = (*m.MockFindOneErr)[1:]
	} else {
		err = nil
	}

	return result, exists, err
}

func (m *MockCollection[T]) Find(ctx context.Context, filter bson.M) ([]T, error) {
	// Get the first element of the result array and remove it from the array, return nil if the array is empty
	var result []T
	if len(*m.MockFindResult) > 0 {
		result = (*m.MockFindResult)[0]
		*m.MockFindResult = (*m.MockFindResult)[1:]
	}

	// Get the first element of the error array and remove it from the array, return nil if the array is empty
	var err error
	if len(*m.MockFindErr) > 0 {
		err = (*m.MockFindErr)[0]
		*m.MockFindErr = (*m.MockFindErr)[1:]
	} else {
		err = nil
	}

	return result, err
}

func (m *MockCollection[T]) Upsert(ctx context.Context, t T, id string) error {
	// Get the first element of the error array and remove it from the array, return nil if the array is empty
	var err error
	if len(*m.MockUpsertErr) > 0 {
		err = (*m.MockUpsertErr)[0]
		*m.MockUpsertErr = (*m.MockUpsertErr)[1:]
	} else {
		err = nil
	}

	return err
}

func (m *MockCollection[T]) UpdateOne(ctx context.Context, t T, id string) error {
	// Get the first element of the error array and remove it from the array, return nil if the array is empty
	var err error
	if len(*m.MockUpdateOneErr) > 0 {
		err = (*m.MockUpdateOneErr)[0]
		*m.MockUpdateOneErr = (*m.MockUpdateOneErr)[1:]
	} else {
		err = nil
	}

	return err
}

func (m *MockCollection[T]) FindOneAndUpdate(ctx context.Context, filter bson.M, update bson.M) (T, error) {
	var result T
	return result, nil
}

func (m *MockCollection[T]) FindOneAndReplace(ctx context.Context, filter bson.M, replacement T) (T, error) {
	var result T
	return result, nil
}
