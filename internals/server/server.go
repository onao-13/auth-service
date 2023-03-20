package server

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
	"security-service/api"
	"security-service/internals/app/controller"
	user3 "security-service/internals/app/repository"
	user2 "security-service/internals/app/service"
	"security-service/internals/conf"
	"security-service/internals/security"
	"security-service/internals/security/jwt"
)

type Server struct {
	cfg  conf.Config
	ctx  context.Context
	pool *pgxpool.Pool
	srv  *http.Server
}

var log = logrus.New()

func New(cfg conf.Config, ctx context.Context) *Server {
	return &Server{cfg: cfg, ctx: ctx}
}

func (s *Server) Start() {
	var err error

	log.Infoln("Server starting...")

	s.pool, err = pgxpool.Connect(s.ctx, s.cfg.DbUrl())
	if err != nil {
		log.Fatalln("Error connecting to database. Error: ", err)
	}

	jwtProvider := jwt.NewProvider(s.cfg)
	bcryptEncryptor := security.NewBcryptEncryption()

	userRepository := user3.NewUserRepository(s.pool, s.ctx)

	authService := user2.NewAuthService(userRepository, bcryptEncryptor, jwtProvider)

	authController := controller.NewAuthController(authService)

	route := api.CreateRouter(authController)

	log.Infoln("Server started")

	s.srv = &http.Server{
		Addr:    ":" + s.cfg.Port,
		Handler: route,
	}

	s.srv.ListenAndServe()
}
