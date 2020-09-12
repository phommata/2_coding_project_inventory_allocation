package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/oklog/oklog/pkg/group"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type AppConfig struct {
	Addr     AddressConfig
}

type AddressConfig struct {
	HTTP  string `env:"HTTP_ADDR" envDefault:":80"`
}

func main() {
	config, err := setUpConfig()
	if err != nil {
		os.Exit(1)
	}

	rtr := mux.NewRouter()
	setupBasicRoutes(rtr)
	svr := http.Server{
		Addr:    config.Addr.HTTP,
		Handler: rtr,
	}

	var g group.Group
	{
		g.Add(func() error {
			return svr.ListenAndServe()
		}, func(_ error) {
			_ = svr.Close()
		})
	}
	{
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			}
		}, func(_ error) {
		})
	}
	fmt.Println(g.Run())
}

func setUpConfig() (AppConfig, error) {
	cfg := AppConfig{}
	var err error
	if err = env.Parse(&cfg.Addr); err != nil {
		return cfg, err
	}

	return cfg, nil
}

//As dependencies are added, they will be added to this function signature as arguments.
//They will then be passed in to the `make{}Handler` function.
func setupBasicRoutes(r *mux.Router) {

}