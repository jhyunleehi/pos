package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"pos/api"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.SetLevel(log.TraceLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: time.RFC3339,
		NoColors:        false,
	})

	env := flag.String("env", "config", "Environment")
	viper.SetConfigName(*env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf(err.Error())
	}
	server := viper.GetString("server.ibof.ip")
	port := viper.GetInt("server.ibof.port")
	log.Printf("[%s][%d]\n", server, port)
}

func setupRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"*"}
	r.Use(cors.New(corsConfig))

	api.ApplyRoutes(r)

	return r
}

func main() {

	logDir := viper.GetString("ginfw.log.dir")
	_, err := os.Stat(logDir)
	if os.IsNotExist(err) {
		os.MkdirAll(logDir, 0755)
	}
	logFile, _ := os.OpenFile(logDir+"////ginfw_"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// HTTP Configuration
	server := &http.Server{
		Addr:           ":" + viper.GetString("ginfw.http.port"),
		Handler:        setupRouter(),
		ReadTimeout:    time.Duration(viper.GetInt("ginfw.http.read_time_out")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("ginfw.http.write_time_out")) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
