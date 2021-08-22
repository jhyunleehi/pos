package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"pos/api"
	"runtime"
	"strings"
	"time"

	"pos/common/events"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
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

	logstr := viper.GetString("logging.level")
	level, err := log.ParseLevel(logstr)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	//logging
	log.SetLevel(level)
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: time.RFC3339,
		NoColors:        true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf("[%s:%d %s()] ", path.Base(f.File), f.Line, funcName)
		},
	})
	logdir := viper.GetString("logging.directory")
	log.Debugf("%v", logdir)
	finfo, err := os.Stat(logdir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(logdir, 0755)
		if err != nil {
			log.Error(err)
		}
		log.Debugf("%v", finfo)
	}
	logFile, err := os.OpenFile(logdir+"////pos_"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error(err.Error())
	}
	log.SetOutput(logFile)

	//event
	events.Setup()

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

	logDir := viper.GetString("ginfw.log.directory")
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
