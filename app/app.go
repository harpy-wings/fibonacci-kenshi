package app

import (
	"fmt"

	"github.com/harpy-wings/fibonacci-kenshi/pkg/config"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	s    *echo.Echo
	port int
}

func (a *App) ListenAndServe() error {
	return a.s.Start(fmt.Sprintf(":%d", a.port))
}

func New() (*App, error) {
	var (
	// v *viper.Viper Using Singleton pattern for Viper and Logger
	// logger *log.Logger Using Singleton pattern for Viper and Logger
	)
	app := new(App)
	app.s = echo.New()

	{
		err := config.InitConfig(viper.GetViper())
		if err != nil {
			return nil, err
		}
	}
	{
		// TODO move to logger package
		// TODO Read from config
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.WarnLevel)
	}
	// Init Instrumentations and etc.
	app.port = viper.GetInt("port")
	return app, nil
}
