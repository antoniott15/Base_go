package main

import(

	"github.com/sirupsen/logrus"
	"basego/database"
	"basego"

)



func main() {
	mongo, err := database.New("mongodb://localhost:8000")
	if err != nil {
		logrus.Error(err)
	}
	api := basego.NewAPI(mongo,"/api", ":4200")

	if err := api.Launch(); err != nil{
		logrus.Error(err)
	}

	if err := api.ConnectWithGRPC(); err != nil {
		logrus.Error(err)
	}

}