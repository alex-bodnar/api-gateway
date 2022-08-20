package users_redis

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"

	"authenticator/internal/api/domain/users"
	"authenticator/internal/api/repository"
	"authenticator/pkg/errs"
)

var _ repository.Users = &Repository{}

// Repository implements repository.FilmCache
type Repository struct {
	redisCache *cache.Cache
	timeLive   time.Duration
}

// NewRepository constructor.
func NewRepository(timeLive time.Duration, rdb *redis.Client) *Repository {
	return &Repository{
		redisCache: cache.New(&cache.Options{Redis: rdb}),
		timeLive:   timeLive,
	}
}

// SetByName add new film to cache.
func (r Repository) Save(ctx context.Context, userData users.User) error {
	item := cache.Item{
		Ctx:   ctx,
		Key:   userData.Name,
		Value: userData,
		TTL:   r.timeLive,
	}

	if err := r.redisCache.Set(&item); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	return nil
}

// GetByName - get user by name.
func (r Repository) GetByName(ctx context.Context, name string) (users.User, error) {
	result := users.User{}
	if err := r.redisCache.Get(ctx, name, &result); err != nil {
		if err == cache.ErrCacheMiss {
			return users.User{}, errs.NotFound{What: "user"}
		}

		return users.User{}, errs.Internal{Cause: err.Error()}
	}

	return result, nil
}
