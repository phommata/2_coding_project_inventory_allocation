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
	HTTP  string `env:"HTTP_ADDR" envDefault:":8080"`
}

func main() {
	config, err := setUpConfig()
	if err != nil {
		os.Exit(1)
	}

	rtr := mux.NewRouter()
	svr := http.Server{
		Addr:    config.Addr.HTTP,
		Handler: rtr,
	}

	//Set Up Inventory Service and Inventory Routes
	inventoryService := inventory.NewElavonService(logger, client)
	inventoryEndpoints := inventory.MakeEndpoints(inventoryService)
	setInventoryRoutes(rtr, inventoryEndpoints)

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

func setElavonRoutes(rtr *mux.Router, endpoints inventory.Endpoints) {
	rtr.Methods(http.MethodPost).Path("/inventory/allocate").inventoryAllocateHandler(endpoints.InventoryAllocateEndpoint)
}