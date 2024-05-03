package httpListener

import (
	"context"
	"fmt"
	"net/http"

	catHandler "projectsphere/cats-social/internal/cat/handler"
	catRepository "projectsphere/cats-social/internal/cat/repository"
	catService "projectsphere/cats-social/internal/cat/service"
	matchHandler "projectsphere/cats-social/internal/match/handler"
	matchRepository "projectsphere/cats-social/internal/match/repository"
	matchService "projectsphere/cats-social/internal/match/service"
	userHandler "projectsphere/cats-social/internal/user/handler"
	userRepository "projectsphere/cats-social/internal/user/repository"
	userService "projectsphere/cats-social/internal/user/service"
	"projectsphere/cats-social/pkg/database"
	"projectsphere/cats-social/pkg/middleware/auth"
	"projectsphere/cats-social/pkg/utils/config"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type HttpImpl struct {
	HttpRouter *HttpRouterImpl
	httpServer *http.Server
}

func NewHttpProtocol(
	HttpRouter *HttpRouterImpl,
) *HttpImpl {
	return &HttpImpl{
		HttpRouter: HttpRouter,
	}
}

func (p *HttpImpl) setupRouter() *gin.Engine {
	return p.HttpRouter.Router()
}

func (p *HttpImpl) Listen() {
	app := p.setupRouter()

	serverPort := fmt.Sprintf(":%v", config.Get().Application.Port)
	p.httpServer = &http.Server{
		Addr:    serverPort,
		Handler: app,
	}

	log.Info().Msgf("Server started on Port %s ", serverPort)
	err := p.httpServer.ListenAndServe()
	if err != nil {
		log.Printf(err.Error())
	}
}

func (p *HttpImpl) Shutdown(ctx context.Context) error {
	if err := p.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

func Start() *HttpImpl {

	config := config.Get()

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgresql://%s:%s@%s/%s?%s", config.DB.Postgre.User, config.DB.Postgre.Pass, config.DB.Postgre.Host, config.DB.Postgre.Name, config.DB.Postgre.Params))
	if err != nil {
		// without db we can't do anything so should be aware if we can't connect
		panic(err.Error())
	}

	postgresConnector := database.NewPostgresConnector(context.TODO(), db)

	catRepo := catRepository.NewCatRepo(postgresConnector)
	catSvc := catService.NewCatService(catRepo)
	catHandler := catHandler.NewCatHandler(catSvc)

	matchRepo := matchRepository.NewMatchRepo(postgresConnector)
	matchSvc := matchService.NewMatchService(matchRepo, catRepo)
	matchhandler := matchHandler.NewMatchHandler(matchSvc)

	userRepo := userRepository.NewUserRepo(postgresConnector)

	jwtAuth := auth.NewJwtAuth(
		config.Auth.AccessTokenExpiredTime,
		config.Auth.SecretKey,
		userRepo.IsUserExist,
	)

	userSvc := userService.NewUserService(userRepo, config.Auth.BcryptSalt, jwtAuth)
	userHandler := userHandler.NewUserHandler(userSvc)

	httpHandlerImpl := NewHttpHandler(
		userHandler,
		catHandler,
		matchhandler,
		jwtAuth,
	)
	httpRouterImpl := NewHttpRoute(httpHandlerImpl)
	httpImpl := NewHttpProtocol(httpRouterImpl)

	return httpImpl
}
