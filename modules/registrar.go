package modules

import (
	"github.com/forbole/juno/v3/modules/registrar"
	"github.com/giansalex/archway-gastracker/database"
	"github.com/giansalex/archway-gastracker/modules/gastracker"

	junomod "github.com/forbole/juno/v3/modules"
)

// ModulesRegistrar represents the modules.Registrar that allows to register all custom modules
type ModulesRegistrar struct {
}

// NewModulesRegistrar allows to build a new ModulesRegistrar instance
func NewModulesRegistrar() *ModulesRegistrar {
	return &ModulesRegistrar{}
}

// BuildModules implements modules.Registrar
func (r *ModulesRegistrar) BuildModules(ctx registrar.Context) junomod.Modules {
	db := database.Cast(ctx.Database)

	return []junomod.Module{
		gastracker.NewModule(db),
	}
}
