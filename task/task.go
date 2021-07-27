package task

import (
	"encoding/json"
	"errors"
	"log"
)

//GetShortURL processes the get request for short URL for given long URL
func GetShortURL(shortURL string) {
	log.Println("get the short URL for given URL")
}

//CreateShortURL processes the request to create short URL from long URL
func CreateShortURL(data []byte) error {
	log.Println("create the short URL for the given URL")

	var longURL string
	err := json.Unmarshal(data, &longURL)
	if err != nil {
		log.Println("Error while unmarshelling the request payload. Error : ", err.Error())
		return errors.New("unmarshal error in the given URL")
	}

	return nil
}
