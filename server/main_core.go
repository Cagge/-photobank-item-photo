package main

import (
	"gitlab.sima-land.ru/dev-dep/dev/packages/go-app-constructor/http"
	"gitlab.sima-land.ru/dev-dep/dev/packages/go-app-constructor/k8s"
	"gitlab.sima-land.ru/dev-dep/dev/packages/go-app-constructor/log"
	"gitlab.sima-land.ru/dev-dep/dev/packages/go-app-constructor/pprof"
	"gitlab.sima-land.ru/dev-dep/dev/packages/go-app-constructor/sentry"
	"gitlab.sima-land.ru/dev-dep/dev/packages/go-app-constructor/serve"
)

func registerCore() {
	app.
		Register(log.NewService()).
		Register(serve.NewService()).
		Register(http.NewService()).
		Register(pprof.NewService()).
		Register(sentry.NewService()).
		Register(k8s.NewService())
}
