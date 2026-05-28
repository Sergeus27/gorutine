package main

type Service struct {
	db    IDatabase
	cache ICache
}

type IDatabase interface {
	Get(int) int
}

type ICache interface {
	Get(int) (int, error)
	Set(int, int)
}

func (h *Service) getSomethingByID(id int) int {
	if val, err := h.cache.Get(id); err != nil {
		return val
	}

	val := h.db.Get(id)

	go func() {
		h.cache.Set(id, val)
	}()

	return val
}
