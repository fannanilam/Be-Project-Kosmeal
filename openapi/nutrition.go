package openapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kosmeal/model"
	"net/http"
)

func GetNutritionController(url string) ([]*model.Nutrition, error) {

	//	url := "https://nutrition-by-api-ninjas.p.rapidapi.com/v1/nutrition?query=2%20eggs%2C%20chocolate%2C%20sugar%2C%20milk%2C%20multipurpose%20flour%2C%20butter"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "nutrition-by-api-ninjas.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "02599545d4msh777c163d7b195a4p1beffdjsn31b3dbaf617b")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	fmt.Println(string(body))

	data := make([]*model.Nutrition, 0)
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}
