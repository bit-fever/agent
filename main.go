//=============================================================================
/*
Copyright © 2025 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package main

import (
	"crypto/tls"
	"github.com/bit-fever/agent/pkg/app"
	"github.com/bit-fever/agent/pkg/core"
	"github.com/bit-fever/agent/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
)

//=============================================================================

func main() {
	cfg := readConfig()
	file := initLogs(cfg)
	defer file.Close()

	core.StartPeriodicScan(cfg)
	router := registerServices()
	runHttpServer(router, cfg)
}

//=============================================================================

func readConfig() *app.Config {
	viper.SetConfigName("agent")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/bit-fever/")
	viper.AddConfigPath("$HOME/.bit-fever/agent")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	var cfg app.Config

	err = viper.Unmarshal(&cfg)

	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}

//=============================================================================

func initLogs(cfg *app.Config) *os.File {

	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lmicroseconds | log.Lshortfile)

	f, err := os.OpenFile(cfg.General.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	gin.DefaultWriter = wrt

	return f
}

//=============================================================================

func registerServices() *gin.Engine {

	log.Println("Registering services...")
	router := gin.Default()
	service.Init(router)

	return router
}

//=============================================================================

func runHttpServer(router *gin.Engine, cfg *app.Config) {
	log.Println("Starting HTTPS server...")

	server := &http.Server{
		Addr:      cfg.General.BindAddress,
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
		Handler: router,
	}

	log.Println("Running")
	err := server.ListenAndServeTLS("config/agent.crt", "config/agent.key")

	if err != nil {
		log.Fatal(err)
	}
}

//=============================================================================
