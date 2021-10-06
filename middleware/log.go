package middleware

import (
	"kosmeal/config"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	coll := config.DBLog.Database("kosmeal").Collection("logs")

	return func(c echo.Context) error {

		log := bson.M{
			"time":   time.Now(),
			"method": c.Request().Method,
			"path":   c.Path(),
		}
		coll.InsertOne(c.Request().Context(), log)

		return next(c)
	}
}

//	"fmt"
//	"time"

//	"github.com/labstack/echo/v4"
//	"go.mongodb.org/mongo-driver/bson"

// func Log(next echo.HandlerFunc) echo.HandlerFunc {
//	coll := config.DBLog.Database("kosmeal").Collection("logs")

//	return func(c echo.Context) error {
//		log := bson.M{
//			"time":   time.Now(),
//			"method": c.Request().Method,
//			"path":   c.Path(),
//		}
//		coll.InsertOne(c.Request().Context(), log)
//		fmt.Println(log)
//		return next(c)
//	}
//}
