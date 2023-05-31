package main

import (
	"apixyz/src/apixyz/database"
	"apixyz/src/apixyz/midleware"
	"apixyz/src/apixyz/util"
	"apixyz/src/config"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"os/signal"
	"time"

	ac "apixyz/src/apixyz/agent/controller"
	ar "apixyz/src/apixyz/agent/repository"
	as "apixyz/src/apixyz/agent/service"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/rs/cors"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./src/config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}
}

func main() {
	if err := database.InitAllConnection(); err != nil {
		util.Logkoe.Error(err)
	}

	defer database.CloseAllConnection()
	addr := config.DetermineListenAddress()
	host := config.Hostname()
	env := config.Environment()
	timeout := config.Timeout()

	var isDev bool
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
		util.Logkoe.Info("running in Production env")
		isDev = false
	} else {
		util.Logkoe.Info("running in Debug env")
		gin.SetMode(gin.DebugMode)
		isDev = true
	}

	secureFunc := midleware.Secure(isDev)
	router := gin.New()
	router.Use(requestid.New())
	router.Use(midleware.Recovery())
	// midleware.LoggingActivity()
	router.Use(healthcheck.Default())
	// router.Use(midleware.RequestLoggerActivity())
	router.Use(secureFunc)
	router.Use(midleware.TimeoutMiddleware(time.Duration(timeout) * time.Second))

	customerRepository := ar.CreateCustomerRepoImpl()
	customerService := as.CreateServiceImpl(customerRepository)

	ac.CreateCustomerController(router, customerService)

	c := cors.AllowAll()
	hdl := c.Handler(router)
	subhost := host[7:]

	router.Use(healthcheck.Default())
	server := &http.Server{
		Addr:         addr,
		Handler:      hdl,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {

		<-quit
		logger.Println("Server apixyz is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server apixyz: %v\n", err)
		}
		close(done)
	}()

	logger.Println("Server apixyz is ready to handle requests at", subhost, addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", addr, err)
	}

	<-done
	logger.Println("Server apixyz stopped")

}
