package memdb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/deeper-x/quotaweb/settings"
	"github.com/redis/go-redis/v9"
)

// Quoter quota interface
type Quoter interface {
	Get() (int, error)
	Set() (bool, error)
	Ping() string
	isAllowed(int) bool
}

// Quota is the quota manager
type Quota struct {
	client *redis.Client
	name   string
	value  int
	ctx    context.Context
}

// NewQuota returns new quota
func NewQuota() Quoter {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	return &Quota{
		client: client,
		name:   settings.Name,
		value:  0,
		ctx:    context.Background(),
	}
}

// Ping test method
func (q *Quota) Ping() string {
	return "Pong"
}

// Get returns current value
func (q *Quota) Get() (int, error) {
	resInt := 0
	res, err := q.client.Get(q.ctx, settings.Name).Result()
	if err != nil {
		res = "0"
	}

	resInt, err = strconv.Atoi(res)
	if err != nil {
		return resInt, err
	}

	return resInt, nil
}

// isAllowed tells is current connection numbers is < maxAllowed
func (q *Quota) isAllowed(cur int) bool {
	return cur < settings.MaxAllowed
}

// Set increment value, after reading current value and checking if increment is allowed
func (q *Quota) Set() (bool, error) {
	res := false
	cur, err := q.client.Get(q.ctx, q.name).Result()
	if err != nil {
		cur = "0"
		fmt.Println("First insert")
	}

	curInt, err := strconv.Atoi(cur)
	if err != nil {
		return res, err
	}

	if q.isAllowed(curInt) {
		next := curInt + 1
		err = q.client.Set(q.ctx, q.name, next, 0).Err()
		if err != nil {
			return res, err
		}

		q.value = next
		res = true
		return res, nil
	}

	return res, nil
}
