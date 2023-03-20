package main

import (
	"golang.org/x/net/context"
	"security-service/internals/conf"
	"security-service/internals/server"
)

func main() {
	cfg := conf.UploadDev()
	ctx := context.Background()

	srv := server.New(cfg, ctx)

	srv.Start()
}
