package main

import (
	"flag"
	"os"

	"github.com/OlegLaban/weather_service/internal/handlers"
	"github.com/OlegLaban/weather_service/internal/pkg/app/rest"
	"github.com/OlegLaban/weather_service/pkg/logger/factory"
)

func main() {
	configPath := flag.String("config", "./config/rest_config.yaml", "path to application config")
	f, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	h := handlers.New()
	c, err := rest.NewFromYaml(f)
	if err != nil {
		panic(err)
	}
	lf := factory.New(factory.Config{Env: "local"})
	l := lf.NewLogger("main")
	app := rest.New(h, l, c)
	app.Run()
}
