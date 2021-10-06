package config

import (
	"context"
	"kosmeal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBLog *mongo.Client

func InitDB() {
	var err error
	//                              username:password@tcp(host:port)/nama_database
	DB, err = gorm.Open(mysql.Open("root:fanna123n@tcp(127.0.0.1:3306)/kosmeal"))
	if err != nil {
		panic(err)
	}
}

func InitLog() {
	var err error
	DBLog, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/kosmeal"))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err = DBLog.Connect(ctx)
	if err != nil {
		panic(err)
	}

	DBLog.ListDatabaseNames(ctx, bson.M{})
}

func InitMigration() {
	DB.AutoMigrate(
		&model.User{},
		&model.Admin{},
		&model.FoodCategory{},
		&model.Recipe{},
		&model.Comment{},
	)
}
