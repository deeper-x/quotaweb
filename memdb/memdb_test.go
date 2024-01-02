package memdb

import (
	"testing"
)

type MockQuota struct {
	name  string
	value int
}

// NewQuoter interface return
func NewQuoter() MockQuota {
	return MockQuota{
		name:  "mockQuota",
		value: 0,
	}
}

func (mq *MockQuota) Ping() string {
	return "Pong"
}

// Get mocks get call
func (mq *MockQuota) Get() (int, error) {
	return 1, nil
}

// Set mock setting value call
func (mq *MockQuota) Set() error {
	mq.value++

	return nil
}

func TestPing(t *testing.T) {
	q := NewQuoter()
	got := q.Ping()
	expected := "Pong"

	if got != expected {
		t.Errorf("%s different from %s", got, expected)
	}
}

func TestGet(t *testing.T) {
	q := NewQuoter()

	_, err := q.Get()
	if err != nil {
		t.Error(err)
	}
}

func TestSet(t *testing.T) {
	q := NewQuoter()

	err := q.Set()
	if err != nil {
		t.Error(err)
	}

	if q.value != 1 {
		t.Error("value not incremented correctly")
	}
}
