package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TencentMapService struct {
	APIKey string
}

func NewTencentMapService(apiKey string) *TencentMapService {
	return &TencentMapService{APIKey: apiKey}
}

func (t *TencentMapService) GetLocation(address string) (interface{}, error) {
	url := fmt.Sprintf("https://apis.map.qq.com/ws/geocoder/v1/?address=%s&key=%s", address, t.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (t *TencentMapService) GetRoute(origin, destination string) (interface{}, error) {
	url := fmt.Sprintf("https://apis.map.qq.com/ws/direction/v1/driving/?from=%s&to=%s&key=%s", origin, destination, t.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}