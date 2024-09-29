package app

import (
	"context"
	"ddd-template/configs"
	"ddd-template/infrastructures/gorm/userdao"
	"ddd-template/infrastructures/user"
	"ddd-template/interfaces/httphandlers"
	"ddd-template/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

type App struct {
	Config  *configs.Config
	Handler http.Handler
}

func Setup() (*App, error) {
	ctx := context.Background()

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Setup user
	userDao := userdao.NewUserDao(db)
	userRepository := user.NewRepository(userDao)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userHandler := httphandlers.NewUserHandler(userUseCase)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", userHandler.GetById)
	})

	// Load configs
	cfg := configs.New()
	if err := cfg.Load(ctx); err != nil {
		return nil, err
	}
	return &App{
		Config:  cfg,
		Handler: r,
	}, nil
}
