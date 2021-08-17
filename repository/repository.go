package repository

import (
	"github.com/DavinciUI/backend/code"
	"time"
)
import "github.com/bluele/gcache"

func NewCachedRepository() Repository {
	return CachedRepository{
		gcache.New(200).
			LRU().
			Expiration(time.Hour * 24).
			Build(),
	}
}

type Entity interface {
	GetCode() code.Code
}

type Repository interface {

	Find(code code.Code) (Entity, error)
	FindAll() []Entity
	Delete(code code.Code)
	Save(entity Entity) error

}

type CachedRepository struct {
	cache gcache.Cache //not exported
}

func (repository CachedRepository) Find(code code.Code) (Entity, error) {
	entity, err := repository.cache.Get(code)

	return entity.(Entity), err
}

func (repository CachedRepository) FindAll() (entities []Entity) {
	internalCache := repository.cache.GetALL(true)
	entities = make([]Entity, len(internalCache))

	i := 0
	for _, entity := range internalCache {
		entities[i] = entity.(Entity)
		i++
	}

	return
}

func (repository CachedRepository) Delete(code code.Code) {
	repository.cache.Remove(code)
}

func (repository CachedRepository) Save(entity Entity) error {
	return repository.cache.Set(entity.GetCode(), entity)
}
