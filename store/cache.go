package store

import "okki.hu/garric/ppnext/model"

// Cache implements Repository.
// Instead of providing a stand-alone implementation, Cache is wrapping another Repository, and provides
// caching functionality based on the model.Room name.
type Cache struct {
	repo  Repository
	cache map[string]*model.Room
}

// NewCache returns a new Cache, that wraps the given Repository
func NewCache(r Repository) *Cache {
	return &Cache{
		repo:  r,
		cache: make(map[string]*model.Room),
	}
}

// Load returns the cached instance of a model.Room with the given name. If the cache does not
// contain the requested object yet, it is loaded from the Repository, and then added to the cache.
func (c *Cache) Load(name string) (*model.Room, error) {
	room, exists := c.cache[name]
	if exists {
		return room, nil
	}
	room, error := c.repo.Load(name)
	if error == nil {
		c.cache[room.Name] = room
	}
	return room, error
}

// Save persists the given model.Room object, and invalidates any instance of the same
// object stored in the cache.
func (c *Cache) Save(room *model.Room) error {
	c.Invalidate(room.Name)
	return c.repo.Save(room)
}

// Delete completely removes a model.Room from the underlying Repository, and cache.
func (c *Cache) Delete(name string) error {
	c.Invalidate(name)
	return c.repo.Delete(name)
}

func (c *Cache) Invalidate(name string) {
	delete(c.cache, name)
}

// Exists returns if a given user exists in any room.
// Cache will first check the user in the cached rooms, this can
// result in the user being found quickly without hitting the repository.
// If the user is not found in the cache, we must check the repository as well.
func (c *Cache) Exists(user string) (bool, error) {
	for _, room := range c.cache {
		for u := range room.Votes {
			if u == user {
				return true, nil
			}
		}
	}
	return c.repo.Exists(user)
}