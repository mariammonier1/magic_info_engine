package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"magic_info_engine/domain/domainModel"
	"net/http"
)

type DeviceOperations struct {
	token  string
	client *http.Client
}

func (thisMII *DeviceOperations) ListAllDevices() (*domainModel.MagicInfoResponseMessage, error) {

	var response *domainModel.MagicInfoResponseMessage
	req, err := http.NewRequest("GET", "http://192.168.1.9:7001/MagicInfo/restapi/v1.0/rms/devices?startIndex=0&pageSize=10", nil)

	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return response, err
	}

	req.Header.Add("api_key", thisMII.token)
	req.Header.Add("Authorization", thisMII.token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "JSESSIONID=E7F07DAAE368BE2BA30BEB61CF8B65EA")

	// Send the HTTP request
	resp, err := thisMII.client.Do(req)

	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return response, err
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return response, err
	}
	fmt.Println("=====================", body)

	json.Unmarshal(body, &response)
	fmt.Println("==============response=======", response)
	return response, err

}
