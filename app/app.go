package app

import (
	"fmt"

	"github.com/harpy-wings/fibonacci-kenshi/internal/controllers"
	"github.com/harpy-wings/fibonacci-kenshi/pkg/config"
	"github.com/harpy-wings/fibonacci-kenshi/pkg/constants"
	"github.com/harpy-wings/fibonacci-kenshi/pkg/fibonacci"
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
		FB  fibonacci.IFibonacci
		err error
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

	// create Modules
	{
		var ops = []fibonacci.Option{fibonacci.OptionWithCaching(true)}
		if v := viper.GetInt(constants.MaxBitSize); v != 0 {
			ops = append(ops, fibonacci.OptionWithMaxBitSize(v))
		}
		FB, err = fibonacci.New(ops...)
		if err != nil {
			return nil, err
		}
	}
	{
		//register Controllers
		C := controllers.NewDefaultController(FB)
		err = C.Register(app.s)
		if err != nil {
			return nil, err
		}
	}

	return app, nil
}
