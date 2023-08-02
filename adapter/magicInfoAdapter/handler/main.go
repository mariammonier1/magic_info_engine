package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	magicInfoOutputPort "magic_info_engine/port/outputPort/magicInfoOutPutPort"
	"net/http"
)

type magicInfoInit struct {
	 magicInfoOutputPort.IMagicInfoDevicesHandlerPort
}

func NewMagicInfoConnection() magicInfoOutputPort.IMagicInfoHandlerPort {

	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"username": "admin",
		"password": "admin2016",
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("http://192.168.1.9:7001/MagicInfo/auth", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return nil

	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)

		return nil
	}

	type loginResonse struct {
		Token string `json:"token"`
	}
	var response loginResonse

	json.Unmarshal(body, &response)

	return &magicInfoInit{
		&DeviceOperations{token: response.Token, client: &http.Client{}},
	}
}
