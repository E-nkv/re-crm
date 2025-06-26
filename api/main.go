package main

import (
	"log"
	"os"
	"re-crm/app"
	"re-crm/entities"
	"re-crm/repositories"
	"re-crm/services"
	"re-crm/utils"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("err loading .env")
	}
	dsn := os.Getenv("PG_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("err connecting to db: ", err)
	}

	if res := db.Exec("DROP TABLE IF EXISTS users"); res.Error != nil {
		log.Fatal("err droping table if exists")
	}
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatal("err migrating user schema")
	}
	res := db.Create([]*entities.User{
		{Nick: "admin", Pass: utils.Bcryptify("admin"), Role: "admin"},
		{Nick: "manager", Pass: utils.Bcryptify("manager"), Role: "manager"},
		{Nick: "rep", Pass: utils.Bcryptify("rep"), Role: "sales_rep"},
	})
	if res.Error != nil {
		log.Fatal("😡 err inserting 3 users", res.Error)
	}
	var us []*entities.User
	db.Model(&entities.User{}).Find(&us)
	DB = db
}

func main() {

	UserRepoPg := repositories.NewUserRepoPg(DB)
	AuthSvc := services.NewAuthService(UserRepoPg)
	app := app.NewApp(AuthSvc)
	app.Mount()
	log.Fatal(app.Run())

}
