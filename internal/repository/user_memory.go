package repository

import (
	"context"
	"sync"

	"github.com/alexaaaant/user-service/internal/domain"
)

type UserMemoryRepo struct {
	mu    sync.Mutex
	users map[int64]*domain.User
	next  int64
}

func NewUserMemoryRepo() *UserMemoryRepo {
	return &UserMemoryRepo{
		users: make(map[int64]*domain.User),
		next:  1,
	}
}

func (r *UserMemoryRepo) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.next
	r.next++
	r.users[user.ID] = user
	return nil
}

func (r *UserMemoryRepo) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	u, ok := r.users[id]
	if !ok {
		return nil, nil
	}
	return u, nil
}
