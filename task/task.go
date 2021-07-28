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
	GetShortURL(shortURL string) (string, error)
	CreateShortURL(data []byte) (string, error)
}

type task struct{}

var (
	Action TASK
)

func init() {
	Action = task{}
}

//GetShortURL processes the get request for short URL for given long URL
func (t task) GetShortURL(longURL string) (string, error) {
	log.Println("get the short URL for given URL")
	protocol, urlWithoutProtocol := seperateURLComponents(longURL)

	entry := Cache.get(urlWithoutProtocol)
	if entry == "" {
		return "", errors.New("short URL for entered URL is not present. You can create it using create api")
	}

	shortURL := fmt.Sprintf("%s://%s", protocol, entry)

	return shortURL, nil
}

//CreateShortURL processes the request to create short URL from long URL
func (t task) CreateShortURL(data []byte) (string, error) {
	log.Println("create the short URL for the given URL")

	var longURL string
	err := json.Unmarshal(data, &longURL)
	if err != nil {
		log.Println("Error while unmarshelling the request payload. Error : ", err.Error())
		return "", errors.New("unmarshal error in the given URL")
	}

	protocol, urlWithoutProtocol := seperateURLComponents(longURL)

	//check if the URL is already present or not
	var hash string
	hash = Cache.get(urlWithoutProtocol)
	if hash == "" {
		//generate hash and store it
		hash = fmt.Sprintf("%x", createHash(urlWithoutProtocol))
		Cache.insert(urlWithoutProtocol, hash)
	}

	shortURL := fmt.Sprintf("%s://%s", protocol, hash)

	log.Println(longURL)

	log.Printf("%s", shortURL)

	return shortURL, nil
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
