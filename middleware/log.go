package middleware

/*	"fmt"
	"kosmeal/config"
	"kosmeal/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	coll := config.DBLog.Database("kosmeal").Collection("logs")

	return func(c echo.Context) error {
		data := new(model.User)
		if c.Request().Method != http.MethodGet {
			if err := c.Bind(&data); err != nil {
				fmt.Println(err)
			}
		}

		log := bson.M{
			"time":    time.Now(),
			"method":  c.Request().Method,
			"path":    c.Path(),
			"request": data,
		}
		response := next(c)
		log["response"] = c.Response().Status
		coll.InsertOne(c.Request().Context(), log)
		fmt.Println(log)
		return response
	}
}
*/
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
