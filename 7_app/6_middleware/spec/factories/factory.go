package factories

import "learning/app/models"

type Factory struct {
	Store models.Store
}

func New(store models.Store) *Factory {
	return &Factory{
		Store: store,
	}
}
