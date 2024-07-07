package main

import (
	"air-drop/cmd/internal/config"
	"air-drop/cmd/internal/handler"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/middleware"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"path/filepath"
)

var configFile1, _ = filepath.Abs("cmd/etc/mm-test.yaml")
var configFile = flag.String("f", configFile1, "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.Timeout = 0

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, func(w http.ResponseWriter) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}, "*"))

	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	server.Use(middleware.ActiveUserMiddleware)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
