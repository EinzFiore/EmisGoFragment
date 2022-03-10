package emisHelpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12"
	configs "gitlab.com/EinzFiore/emis-modules/configs"
)

type ApiMeta struct {
	Version  int         `json:"version"`
	Message  string      `json:"message"`
	Errors   interface{} `json:"errors"`
	Metadata interface{} `json:"metadata"`
	Results  interface{} `json:"results"`
}

func HttpClient() *http.Client {
	timeOut := configs.GetConfig().MaxRequestTimeout
	maxTimeout, err := strconv.Atoi(timeOut)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{Timeout: time.Duration(maxTimeout) * time.Second}
	return client
}

func RequestWithAuth(method string, url string, c *gin.Context) ([]byte, error) {
	endpoint := url
	client := HttpClient()

	token := c.GetHeader("Authorization")

	// hard payload, next will handle this
	payload := map[string]string{"foo": "baz"}
	jsonData, err := json.Marshal(payload)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	var result ApiMeta
	err = json.Unmarshal([]byte(string(body)), &result)
	if err != nil {
		return []byte(result.Message), err
	}

	if result.Errors != nil {
		msgError := result.Errors.(map[string]interface{})
		// sometimes emis response return errors with code and message
		_, ok := msgError["message"]
		if ok {
			return []byte(result.Message), errors.New(fmt.Sprint(msgError["message"]))
		} else {
			return []byte(result.Message), errors.New(fmt.Sprint(result.Errors))
		}

	}

	// must to change results type interface to byte
	resData, err := json.Marshal(result.Results)
	if err != nil {
		return resData, err
	}

	// this helper only return results from response emis
	return resData, nil
}

func IrisRequestHTTP(method string, url string, c iris.Context) ([]byte, error) {
	endpoint := url
	client := HttpClient()

	token := c.GetHeader("Authorization")

	// hard payload, next will handle this
	payload := map[string]string{"foo": "baz"}
	jsonData, err := json.Marshal(payload)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	var result ApiMeta
	err = json.Unmarshal([]byte(string(body)), &result)
	if err != nil {
		return []byte(result.Message), err
	}

	if result.Errors != nil {
		msgError := result.Errors.(map[string]interface{})
		// sometimes emis response return errors with code and message
		_, ok := msgError["message"]
		if ok {
			return []byte(result.Message), errors.New(fmt.Sprint(msgError["message"]))
		} else {
			return []byte(result.Message), errors.New(fmt.Sprint(result.Errors))
		}

	}

	// must to change results type interface to byte
	resData, err := json.Marshal(result.Results)
	if err != nil {
		return resData, err
	}

	// this helper only return results from response emis
	return resData, nil
}
