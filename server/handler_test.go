package server

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/ash3798/url-shortner/config"
	"github.com/ash3798/url-shortner/task"
)

const (
	testLongURL  = "https://www.google.com"
	testShortURL = "http://8e0cc6d2dcf8f5"

	post = "POST"
	get  = "GET"

	postRequestPath = "/url"
	getRequestPath  = "/url"
	hostname        = "localhost"
	protocol        = "http"
)

type mockTask struct{}

var (
	mockGetShortURLFunc func() ([]byte, error)
	mockCreateShortURL  func() ([]byte, error)
)

func init() {
	//inializing mock functions with dummy ok response
	mockCreateShortURL = func() ([]byte, error) { return []byte(""), nil }
	mockGetShortURLFunc = func() ([]byte, error) { return []byte(""), nil }
}

func (t mockTask) GetShortURL(urlData []byte) ([]byte, error) {
	res, err := mockGetShortURLFunc()
	return res, err
}

func (t mockTask) CreateShortURL(data []byte) ([]byte, error) {
	res, err := mockCreateShortURL()
	return res, err
}

//TestCase : Post request handler returns the 200 StatusOK with all correct parameters
func TestPostRequestHandler(t *testing.T) {
	config.InitEnv()

	mockCreateShortURL = func() ([]byte, error) { return []byte(""), nil }
	task.Action = mockTask{}

	body := `{ "url" :"` + testLongURL + `"}`
	req, err := http.NewRequest(post, protocol+"://"+hostname+strconv.Itoa(config.Manager.ApplicationPort)+postRequestPath, bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Errorf("Error while creating request")
	}

	recorder := httptest.NewRecorder()

	HandleRequest(recorder, req)
	expectedStatusCode := 200
	if recorder.Code != expectedStatusCode {
		t.Errorf("status code returned is not as expected. Expected : %d , Actual: %d", expectedStatusCode, recorder.Code)
	}
}

//Testcase : Unmarshal error in the payload sent via user should result in 400 status code
func TestPostRequestHandlerWithUnmarshalError(t *testing.T) {
	config.InitEnv()

	mockCreateShortURL = func() ([]byte, error) { return []byte(""), errors.New("unmarshal error in payload") }
	task.Action = mockTask{}

	body := `{`
	req, err := http.NewRequest(post, protocol+"://"+hostname+strconv.Itoa(config.Manager.ApplicationPort)+postRequestPath, bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Errorf("Error while creating request")
	}

	recorder := httptest.NewRecorder()

	HandleRequest(recorder, req)

	expectedStatusCode := 400
	if recorder.Code != expectedStatusCode {
		t.Errorf("status code returned is not as expected. Expected : %d , Actual: %d", expectedStatusCode, recorder.Code)
	}
}

//TestCase : Request handler should return error if request is made using wrong method
func TestPostRequestHandlerWithWrongMethod(t *testing.T) {
	config.InitEnv()

	mockCreateShortURL = func() ([]byte, error) { return []byte(""), nil }
	task.Action = mockTask{}

	body := `{ "url" :"` + testLongURL + `"}`
	req, err := http.NewRequest("PUT", protocol+"://"+hostname+strconv.Itoa(config.Manager.ApplicationPort)+postRequestPath, bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Errorf("Error while creating request")
	}

	recorder := httptest.NewRecorder()

	HandleRequest(recorder, req)
	expectedCode := 405
	if recorder.Code != expectedCode {
		t.Errorf("status code returned is not as expected. Expected : %d , Actual: %d", expectedCode, recorder.Code)
	}
}

//TestCase : Request handler should return error if there is error in processing
func TestPostRequestHandlerWithErrorInCreatingURL(t *testing.T) {
	config.InitEnv()

	mockCreateShortURL = func() ([]byte, error) { return []byte(""), errors.New("error while creating short URL") }
	task.Action = mockTask{}

	body := `{ "url" :"` + testLongURL + `"}`
	req, err := http.NewRequest(post, protocol+"://"+hostname+strconv.Itoa(config.Manager.ApplicationPort)+postRequestPath, bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Errorf("Error while creating request")
	}

	recorder := httptest.NewRecorder()

	HandleRequest(recorder, req)
	expectedCode := 400
	if recorder.Code != expectedCode {
		t.Errorf("status code returned is not as expected. Expected : %d , Actual: %d", expectedCode, recorder.Code)
	}
}

//TESTCASE : check if the get request shows status OK with correct payload
func TestGetRequestHandler(t *testing.T) {
	config.InitEnv()

	mockGetShortURLFunc = func() ([]byte, error) { return []byte(""), nil }
	task.Action = mockTask{}

	body := `{ "url" :"` + testLongURL + `"}`
	req, err := http.NewRequest(get, protocol+"://"+hostname+strconv.Itoa(config.Manager.ApplicationPort)+getRequestPath, bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Errorf("Error while creating request")
	}

	recorder := httptest.NewRecorder()

	HandleRequest(recorder, req)
	expectedStatusCode := 200
	if recorder.Code != expectedStatusCode {
		t.Errorf("status code returned is not as expected. Expected : %d , Actual: %d", expectedStatusCode, recorder.Code)
	}
}
