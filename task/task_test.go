package task

import (
	"encoding/json"
	"testing"

	"github.com/ash3798/url-shortner/config"
)

func TestCreateShortURL(t *testing.T) {
	if !config.InitEnv() {
		t.Errorf("failed to initalize the env variables")
	}

	sendUrlInfo := URLDetails{
		Url: "https://www.google.com/test/path",
	}

	data, err := json.Marshal(sendUrlInfo)
	if err != nil {
		t.Errorf("Error while marshelling the payload for create URL function")
	}

	//Testcase : check the result of the createShortURL api to see if it return the same long URL that was sent
	res, err := Action.CreateShortURL([]byte(data))
	if err != nil {
		t.Errorf("failed while creating the short URL for long URL")
	}

	receivedURLInfo := URLDetails{}
	err = json.Unmarshal(res, &receivedURLInfo)
	if err != nil {
		t.Errorf("error while unmarshelling the received response from create short Url request")
	}

	if receivedURLInfo.Url != sendUrlInfo.Url {
		t.Errorf("sent long URL and long URL in received result should be same. Expected : %s , Received: %s", sendUrlInfo.Url, receivedURLInfo.Url)
	}

	//Testcase : Verify if calling the createShortURL again with same URL results in same short URL
	res, err = Action.CreateShortURL([]byte(data))
	if err != nil {
		t.Errorf("failed while creating the short URL for long URL")
	}

	receivedAgainURLInfo := URLDetails{}
	err = json.Unmarshal(res, &receivedAgainURLInfo)
	if err != nil {
		t.Errorf("error while unmarshelling the received response from create short Url request")
	}

	if receivedURLInfo.ShortURL != receivedAgainURLInfo.ShortURL {
		t.Errorf("Same result should be returned if we try to create short URL again for same longURL. Expected : %s , Received: %s", receivedURLInfo.ShortURL, receivedAgainURLInfo.ShortURL)
	}

	//Testcase : Verify if createShortURL and getShortURL both return the same shortURL as result
	res, err = Action.GetShortURL(data)
	if err != nil {
		t.Errorf("failed while doing get request on the longURL")
	}

	receivedGetURLInfo := URLDetails{}
	err = json.Unmarshal(res, &receivedGetURLInfo)
	if err != nil {
		t.Errorf("error while unmarshelling the received response from create short Url request")
	}

	if receivedURLInfo.ShortURL != receivedGetURLInfo.ShortURL {
		t.Errorf("sent long URL and long URL in received result should be same. Expected : %s , Received: %s", receivedURLInfo.ShortURL, receivedGetURLInfo.ShortURL)
	}
}
