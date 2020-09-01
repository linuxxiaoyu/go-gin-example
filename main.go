// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/linuxxiaoyu/go-gin-example
// @license.name MIT
// @license.url https://github.com/linuxxiaoyu/go-gin-example/blob/master/LICENSE
// @basePath /api/v1
package main

import (
	"fmt"
	"net/http"

	"github.com/linuxxiaoyu/go-gin-example/pkg/setting"
	"github.com/linuxxiaoyu/go-gin-example/routers"
)

func main() {
	router := routers.Init()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
