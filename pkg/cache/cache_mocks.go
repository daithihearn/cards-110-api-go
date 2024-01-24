package cache

import "time"

type MockCache[T any] struct {
	Cache[T]
	MockSetErr    *[]error
	MockGetResult *[]T
	MockGetExists *[]bool
	MockGetErr    *[]error
	MockDeleteErr *[]error
}

func (m *MockCache[T]) Set(key string, value T, expiration time.Duration) error {
	// Get the first element of the error array and remove it from the array, return nil if the array is empty
	var err error
	if len(*m.MockSetErr) > 0 {
		err = (*m.MockSetErr)[0]
		*m.MockSetErr = (*m.MockSetErr)[1:]
	} else {
		err = nil
	}

	return err
}

func (m *MockCache[T]) Get(key string) (T, bool, error) {
	// Get the first element of the result array and remove it from the array, return nil if the array is empty
	var result T
	if len(*m.MockGetResult) > 0 {
		result = (*m.MockGetResult)[0]
		*m.MockGetResult = (*m.MockGetResult)[1:]
	}

	// Get the first element of the exists array and remove it from the array, return nil if the array is empty
	var exists bool
	if len(*m.MockGetExists) > 0 {
		exists = (*m.MockGetExists)[0]
		*m.MockGetExists = (*m.MockGetExists)[1:]
	}

	// Get the first element of the error array and remove it from the array, return nil if the array is empty
	var err error
	if len(*m.MockGetErr) > 0 {
		err = (*m.MockGetErr)[0]
		*m.MockGetErr = (*m.MockGetErr)[1:]
	} else {
		err = nil
	}

	return result, exists, err
}

func (m *MockCache[T]) Delete(key string) error {
	// Get the first element of the error array and remove it from the array, return nil if the array is empty
	var err error
	if len(*m.MockDeleteErr) > 0 {
		err = (*m.MockDeleteErr)[0]
		*m.MockDeleteErr = (*m.MockDeleteErr)[1:]
	} else {
		err = nil
	}

	return err
}
