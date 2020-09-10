//+build wireinject

package generator

import (
	"github.com/google/wire"
	"github.com/hirasawayuki/go-oauth-server/controller"
	"github.com/hirasawayuki/go-oauth-server/handler"
	"github.com/hirasawayuki/go-oauth-server/registry"
	"github.com/hirasawayuki/go-oauth-server/repository"
	"github.com/hirasawayuki/go-oauth-server/usecase"
)

func InitializeAppHandler() (handler.AppHandler, func(), error) {
	wire.Build(
		registry.NewStore,
		wire.Bind(new(registry.IStores), new(registry.Store)),
		wire.Bind(new(repository.IClientRepository), new(*repository.ClientRepository)),
		wire.Struct(new(repository.ClientRepository), "*"),
		wire.Bind(new(handler.IAuthorizeHandler), new(*handler.AuthorizeHandler)),
		wire.Struct(new(handler.AuthorizeHandler), "*"),
		wire.Bind(new(controller.IClientController), new(*controller.ClientController)),
		wire.Struct(new(controller.ClientController), "*"),
		wire.Bind(new(usecase.IClientGetUseCase), new(*usecase.ClientGetUseCase)),
		wire.Struct(new(usecase.ClientGetUseCase), "*"),
		wire.Struct(new(handler.AppHandler), "*"),
	)
	return handler.AppHandler{}, nil, nil
}
