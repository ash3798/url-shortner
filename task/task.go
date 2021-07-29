package task

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

type TASK interface {
	GetShortURL(urlData []byte) ([]byte, error)
	CreateShortURL(data []byte) ([]byte, error)
}

type task struct{}

var (
	Action TASK
)

func init() {
	Action = task{}
}

type URLDetails struct {
	Url      string `json:"url"`
	ShortURL string `json:"short_url"`
}

//GetShortURL processes the get request for short URL for given long URL
func (t task) GetShortURL(urlData []byte) ([]byte, error) {

	var urlInfo URLDetails
	err := json.Unmarshal(urlData, &urlInfo)
	if err != nil {
		log.Println("Error while unmarshelling the request payload. Error : ", err.Error())
		return []byte(""), errors.New("unmarshal error in the given URL")
	}

	protocol, urlWithoutProtocol := seperateURLComponents(urlInfo.Url)

	entry := Cache.get(urlWithoutProtocol)
	if entry == "" {
		return []byte(""), errors.New("short URL for entered URL is not present. You can create it using create api")
	}

	urlInfo.ShortURL = fmt.Sprintf("%s://%s", protocol, entry)

	result, err := json.Marshal(urlInfo)
	if err != nil {
		log.Println("Error while marshelling the request response. Error : ", err.Error())
		return []byte(""), errors.New("error while marshelling the result")
	}
	return result, nil
}

//CreateShortURL processes the request to create short URL from long URL
func (t task) CreateShortURL(data []byte) ([]byte, error) {

	var urlInfo URLDetails
	err := json.Unmarshal(data, &urlInfo)
	if err != nil {
		log.Println("Error while unmarshelling the request payload. Error : ", err.Error())
		return []byte(""), errors.New("unmarshal error in the given URL")
	}

	protocol, urlWithoutProtocol := seperateURLComponents(urlInfo.Url)

	//check if the URL is already present or not
	var hash string
	hash = Cache.get(urlWithoutProtocol)
	if hash == "" {
		//generate hash and store it
		hash = fmt.Sprintf("%x", createHash(urlWithoutProtocol))
		Cache.insert(urlWithoutProtocol, hash)
	}

	urlInfo.ShortURL = fmt.Sprintf("%s://%s", protocol, hash)

	result, err := json.Marshal(urlInfo)
	if err != nil {
		log.Println("Error while marshelling the request response. Error : ", err.Error())
		return []byte(""), errors.New("error while marshelling the result")
	}
	return result, nil
}

func createHash(longURL string) []byte {
	shortURL := sha256.Sum256([]byte(longURL))
	return shortURL[25:]
}

func seperateURLComponents(URL string) (string, string) {
	comp := strings.Split(URL, "://")

	var protocol string
	var urlWithoutProtocol string

	//seperate protocol and rest of the URL
	if len(comp) == 2 {
		protocol = comp[0]
		urlWithoutProtocol = comp[1]
	} else {
		protocol = ""
		urlWithoutProtocol = comp[0]
	}

	return protocol, urlWithoutProtocol
}
