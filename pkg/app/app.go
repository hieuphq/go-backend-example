package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/dwarvesf/gerr"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hieuphq/backend-example/pkg/config"
	"github.com/hieuphq/backend-example/pkg/handler"
	"github.com/hieuphq/backend-example/pkg/middleware"
	"github.com/hieuphq/backend-example/pkg/validator"
	"github.com/hieuphq/backend-example/translation"
)

// App api app instance
type App struct {
	cfg config.Config
	l   gerr.Log
	th  translation.Helper
}

// LoadApp load config and init app
func LoadApp() *App {
	cls := config.DefaultConfigLoaders()
	cfg := config.LoadConfig(cls)
	l := gerr.NewSimpleLog()
	th := translation.NewTranslatorHelper()

	return &App{
		cfg: cfg,
		l:   l,
		th:  th,
	}
}

// Run api app
func (a App) Run() {
	router := a.setupRouter()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.cfg.Port),
		Handler: router,
	}

	go func() {
		// service connections
		a.l.Info("listening on ", a.cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		quit <- os.Interrupt
	}()

	select {
	case <-quit:

		a.l.Info("Shutdown Server ...")
		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.GetShutdownTimeout())
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			a.l.Error("Server Shutdown:", err)
		}
		a.l.Info("Server exiting")
	}
}

func (a App) setupRouter() *gin.Engine {
	r := gin.New()
	binding.Validator = validator.NewStructValidator(a.th)

	r.Use(middleware.NewLogDataMiddleware(a.cfg.ServiceName, a.cfg.Env))
	r.Use(cors.New(
		cors.Config{
			AllowOrigins: a.cfg.GetCORS(),
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
			AllowHeaders: []string{"Origin", "Host",
				"Content-Type", "Content-Length",
				"Accept-Encoding", "Accept-Language", "Accept",
				"X-CSRF-Token", "Authorization", "X-Requested-With", "X-Access-Token"},
			ExposeHeaders:    []string{"MeAllowMethodsntent-Length"},
			AllowCredentials: true,
		},
	))

	h := handler.NewHandler(a.cfg, a.l, a.th)

	// handlers
	r.GET("/healthz", h.Healthz)
	r.POST("/signup", h.Signup)
	return r
}
