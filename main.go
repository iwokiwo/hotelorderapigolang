package main

import (
	"fmt"
	"iwogo/Config"
	"iwogo/Models"
	"iwogo/Routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func errorWrapper(err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 	}
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})

	//----------------rabitmq-----------------
	// conn, errs := amqp.Dial(os.Getenv("RABBITMQ_HOST"))
	// errorWrapper(errs, "Failed to connect rabbitmq")
	// defer conn.Close()

	// ch, errs := conn.Channel()
	// errorWrapper(errs, "Failed to open a channel")
	// defer ch.Close()

	//---------------------------------------------

	if err != nil {
		fmt.Println("Status DB:", err)
	}

	db.AutoMigrate(
		&Models.User{},
		&Models.Category{},
		&Models.Product{},
		&Models.Unit{},
		&Models.Page{},
		&Models.Slider{},
		&Models.SliderRelation{},
		&Models.CategoryRelation{},
		&Models.Setting{},
		&Models.Branch{},
		&Models.Store{},
		&Models.Img{},
		&Models.Coupon{},
		&Models.Order{},
		&Models.OrderDetail{},
		&Models.Payment{},
		&Models.PaymentType{})

	//Routes.SubscribeMessage(ch, "go-queue_order")
	router := Routes.SetupRouter(db)
	router.Run(":" + os.Getenv("APP_PORT"))

}
