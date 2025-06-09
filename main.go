package main

import (
	"coinpe/controllers"
	"coinpe/database"
	"coinpe/models"
	"coinpe/pkg/config"
	"coinpe/pkg/graceful"
	"coinpe/pkg/logger"
	"coinpe/pkg/validator"
	"coinpe/routers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	gormlogger "gorm.io/gorm/logger"
)

func main() {
	// Get Config from env
	cfg := &config.Config{}
	err := cfg.Load(cfg)

	if err != nil {
		logger.Fatalf("error while loading config: %s", err)
	}

	if cfg.Debug {
		logger.SetLogLevel(logrus.DebugLevel)
	}

	dbLogConfig := database.DBLogConfig{
		DefaultLogLevel:    gormlogger.Info,
		MigrationsLogLevel: gormlogger.Silent,
	}

	// Get DB connection
	db, err := database.New(
		cfg.MainDatabase,
		database.WithLogConfig(dbLogConfig),
		database.WithMigrations(models.GetMigrationModel()),
	)
	if err != nil {
		logger.Fatalf("unable to get database connection, error: %s", err)
	}

	models.AddSystemData(db, cfg.Environment)

	app := config.App{
		Config: *cfg,
		DB:     db,
	}

	ctrl := controllers.BaseController{
		DB:     app.DB,
		Config: app.Config,
	}

	validate, trans, err := validator.InitValidator()
	if err != nil {
		logger.Fatal("Unable to init validator ", err)
	}

	//adding remaining values to the controller
	ctrl.Translator = &trans
	ctrl.Validator = validate

	router := gin.New()
	if app.Config.VPCProxyCIDR != "" {
		router.SetTrustedProxies([]string{app.Config.VPCProxyCIDR})
	}
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	app.Router = router

	// Register routes
	routers.RegisterRoutes(app, ctrl)

	// Setup gin router and listening for requests
	listenAddr := ":" + app.Config.Server.Port
	server := &http.Server{
		Addr:    listenAddr,
		Handler: app.Router,
	}

	graceful := graceful.Graceful{
		HTTPServer:      server,
		ShutdownTimeout: time.Duration(5 * time.Second),
		State:           &graceful.ServerState{},
	}

	// You can generate ASCI art here
	// https://patorjk.com/software/taag/#p=display&f=Doom&t=COINPE
	banner := `
	
 ▗▄▄▖ ▗▄▖ ▗▄▄▄▖▗▖  ▗▖▗▄▄▖ ▗▄▄▄▖
▐▌   ▐▌ ▐▌  █  ▐▛▚▖▐▌▐▌ ▐▌▐▌   
▐▌   ▐▌ ▐▌  █  ▐▌ ▝▜▌▐▛▀▘ ▐▛▀▀▘
▝▚▄▄▖▝▚▄▞▘▗▄█▄▖▐▌  ▐▌▐▌   ▐▙▄▄▖

`
	graceful.ListenAndServe(banner)

}
