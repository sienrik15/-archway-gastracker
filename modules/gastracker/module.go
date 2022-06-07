package gastracker

import (
	"github.com/giansalex/archway-gastracker/database"

	"github.com/forbole/juno/v3/modules"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
	_ modules.BlockModule   = &Module{}
)

// Module represents the x/gastracker module handler
type Module struct {
	db *database.Db
}

// NewModule allows to build a new Module instance
func NewModule(db *database.Db) *Module {
	return &Module{
		db: db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "gastracker"
}
