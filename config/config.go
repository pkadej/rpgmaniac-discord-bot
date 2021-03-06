package config

import (
	"encoding/json"
	"io/ioutil" //it will be used to help us read our config.json file.
	"log"       //used to print errors majorly.
)

var (
	Channel string
	Game    string

	config *configStruct //To store value extracted from config.json.
)

type configStruct struct {
	Channel string `json : "Channel"`
	Game    string `json : "Game"`
}

func ReadConfig() error {
	log.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json") // ioutil package's ReadFile method which we read config.json and return it's value we will then store it in file variable and if an error ocurrs it will be stored in err .

	//Handling error and printing it using fmt package's Println function and returning it .
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// Here we performing a simple task by copying value of file into config variable which we have declared above , and if there any error we are storing it in err . Unmarshal takes second arguments reference remember it .
	err = json.Unmarshal(file, &config)

	//Handling error
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// After storing value in config variable we will access it and storing it in our declared variables .
	Channel = config.Channel
	Game = config.Game

	//If there isn't any error we will return nil.
	return nil

}
