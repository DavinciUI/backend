package repository

import (
	"github.com/DavinciUI/backend/code"
	"github.com/bluele/gcache"
	"time"
)

import entityModule "github.com/DavinciUI/backend/entity"

func NewCachedRepository() Repository {
	return CachedRepository{
		gcache.New(200).
			LRU().
			Expiration(time.Hour * 24).
			Build(),
	}
}


type Repository interface {

	Find(code code.Code) (entityModule.Entity, error)
	FindAll() []entityModule.Entity
	Delete(code code.Code)
	Save(entity entityModule.Entity) error

}

type CachedRepository struct {
	cache gcache.Cache //not exported
}

func (repository CachedRepository) Find(code code.Code) (entityModule.Entity, error) {
	entity, err := repository.cache.Get(code)

	if err != nil {
		return nil, err
	}

	return entity.(entityModule.Entity), nil
}

func (repository CachedRepository) FindAll() (entities []entityModule.Entity) {
	internalCache := repository.cache.GetALL(true)
	entities = make([]entityModule.Entity, len(internalCache))

	i := 0
	for _, entity := range internalCache {
		entities[i] = entity.(entityModule.Entity)
		i++
	}

	return
}

func (repository CachedRepository) Delete(code code.Code) {
	repository.cache.Remove(code)
}

func (repository CachedRepository) Save(entity entityModule.Entity) error {
	return repository.cache.Set(entity.GetCode(), entity)
}
