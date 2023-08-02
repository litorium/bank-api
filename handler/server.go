package handler

import (
	"bank-api/manager"
	"bank-api/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager

	srv  *gin.Engine
}

func (s *server) Run() {
	// session
	store := cookie.NewStore([]byte("secret"))

	s.srv.Use(middleware.LoggerMiddleware())
	s.srv.Use(sessions.Sessions("session", store))

	// handler
	NewUserHandler(s.srv, s.usecaseManager.GetUserUsecase())
	NewLoginHandler(s.srv, s.usecaseManager.GetLoginUsecase())
	NewMerchantHandler(s.srv, s.usecaseManager.GetMerchantUsecase())
	NewPaymentHandler(s.srv, s.usecaseManager.GetPaymentUsecase())

	s.srv.Run()
}

func NewServer() Server {

	repo := manager.NewRepoManager()
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()
	return &server{
		usecaseManager: usecase,
		srv:            srv,
	}
}

