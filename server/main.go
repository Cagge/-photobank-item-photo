package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	constructorApp "gitlab.sima-land.ru/dev-dep/dev/packages/go-app-constructor/app"
	"gitlab.sima-land.ru/dev-dep/dev/packages/go-app-constructor/cmd"
	"gitlab.sima-land.ru/dev-dep/dev/packages/go-app-constructor/sentry"
	"gopkg.in/gographics/imagick.v2/imagick"
	"net/http"
	. "photobank-item-photo/app/pkg/downloader/models"
	. "photobank-item-photo/app/pkg/resize/models"
	"photobank-item-photo/app/resize"
	"time"
)

var BuildVersion = "development"
var BuildName = "consumer-photobank-item-photo"

var app constructorApp.Container

func main() {
	t0 := time.Now()
	defer sentry.Recover()
	configuration, err := constructorApp.
		NewConfiguration().
		WithName(BuildName).
		WithVersion(BuildVersion).
		Read()
	if err != nil {
		panic(err)
	}

	cmdSvc := cmd.NewService()
	app = constructorApp.
		NewContainer(configuration).
		Register(cmdSvc)
	registerCore()

	if err := app.Invoke(cmdSvc); err != nil {
		panic(err)
	}
	imagick.Initialize()
	defer imagick.Terminate()
	cfgResize := ConfigResize{}
	err = configuration.Unmarshal(&cfgResize)
	if err != nil {
		log.Error(err)
	}
	cfgDataPhoto := ConfigDataPhoto{}
	err = configuration.Unmarshal(&cfgDataPhoto)
	if err != nil {
		log.Error(err)
	}
	out, _, err := resize.ControlResize(cfgDataPhoto, cfgResize, http.DefaultClient)
	if err != nil {
		log.Error(err)
	}
	fmt.Println("Create", out, "img")
	if err != nil {
		log.Fatal(err)
	}

	t1 := time.Now()
	fmt.Printf("Elapsed time: %v", t1.Sub(t0))
}
