package main

import (
	"Go-Project-With-Postgres/handler"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/yookoala/realpath"
)

//go:embed assets
var assets embed.FS

func main() {
	// config
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)

	config.SetConfigFile("env/config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}
	switch config.GetString("runtime.loglevel") {
	case "trace":
		log.Println(logrus.TraceLevel)
	case "debug":
		log.Println(logrus.DebugLevel)
	default:
		log.Println(logrus.InfoLevel)
	}
	log.Println("starting web service")
	// env
	env := config.GetString("runtime.environment")
	// logrus
	logrus := logrus.NewEntry(logrus.StandardLogger())
	// decoder
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	// assets
	asst, err := fs.Sub(assets, "assets")
	if err != nil {
		log.Fatal(err)
	}

	if env == "development" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		assetPath, err := realpath.Realpath(filepath.Join(wd, "assets"))
		if err != nil {
			log.Fatal(err)
		}
		asst = afero.NewIOFS(afero.NewBasePathFs(afero.NewOsFs(), assetPath))
	}
	// initialize Route
	r, err := handler.New(env, config, logrus, asst, decoder)
	if err != nil {
		log.Fatal(err)
	}
	ser := config.GetString("server.serverPort")
	if err := http.ListenAndServe(ser, r); err != nil {
		log.Fatal(err)
	}

}
