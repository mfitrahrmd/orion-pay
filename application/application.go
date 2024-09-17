package application

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mfitrahrmd/orion-pay/model"
	"github.com/mfitrahrmd/orion-pay/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Application struct {
	port               int
	Db                 *gorm.DB
	engine             *gin.Engine
	Router             gin.IRouter
	UserRepository     *repository.UserRepository
	TransferRepository *repository.TransferRepository
}

func NewApplication() *Application {
	return &Application{
		port: 3000,
	}
}

func (app *Application) Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	g := gin.Default()
	app.engine = g
	app.Router = g
}

func (app *Application) SetupDB() {
	var err error
	app.Db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Fatalln(err)
	}
	app.Db.AutoMigrate(&model.User{}, &model.Wallet{}, &model.Transfer{}, &model.Entry{})
	app.UserRepository = repository.NewUserRepository(app.Db)
	app.TransferRepository = repository.NewTransferRepository(app.Db)
}

func (app *Application) Run() {
	err := app.engine.Run(fmt.Sprintf(":%d", app.port))
	if err != nil {
		log.Fatalln(err)
	}
}

func (app *Application) SetPort(port int) *Application {
	app.port = port

	return app
}
