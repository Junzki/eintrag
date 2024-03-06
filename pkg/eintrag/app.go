package eintrag

import (
	"database/sql"
	"eintrag/internal/eintrag"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Config struct {
	Listen                   string `json:"listen"`
	Port                     uint16 `json:"port"`
	SigningKey               string `json:"secret_key"`
	DatabaseType             string `json:"database_type" default:"postgres"`
	DatabaseConnectionString string `json:"database_connection_string"`
}

func NewConfig(configFile *string) Config {
	cfg := Config{
		Listen:     "0.0.0.0",
		Port:       8080,
		SigningKey: "eintrag-key",
	}

	if nil != configFile && "" != *configFile {
		content, err := os.ReadFile(*configFile)
		if nil != err {
			panic(err)
		}

		err = json.Unmarshal(content, &cfg)
		if nil != err {
			panic(err)
		}
	}

	return cfg
}

type App struct {
	engine *gin.Engine
	DbConn *sql.DB
	Logger *logrus.Logger
	config Config
}

func NewApp(configFile *string) App {
	app := App{
		nil,
		eintrag.DbConn,
		eintrag.LOG,
		NewConfig(configFile),
	}

	eintrag.InitDbConn(app.config.DatabaseConnectionString)

	return app
}

func (app *App) Init() {
	app.engine = gin.Default()

	app.engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})
	app.engine.POST("/login", Login)

}

func (app *App) Run() {
	_ = app.engine.Run()
}
