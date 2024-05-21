package users

import "github.com/google/wire"

var UserSet = wire.NewSet(
	InitController,
	wire.Bind(new(IController), new(Controller)),

	InitService,
	wire.Bind(new(IService), new(Service)),

	InitRepo,
	wire.Bind(new(IRepo), new(Repo)),

	InitModel,
	wire.Bind(new(IModel), new(Model)),
)
