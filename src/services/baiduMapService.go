package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BaiduMapService struct {
	APIKey string
}

type LocationResponse struct {
	Status  string `json:"status"`
	Result  struct {
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
	} `json:"result"`
}

func (b *BaiduMapService) GetLocation(address string) (float64, float64, error) {
	url := fmt.Sprintf("http://api.map.baidu.com/geocoding/v3/?address=%s&output=json&ak=%s", address, b.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("failed to get location: %s", resp.Status)
	}

	var locationResponse LocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationResponse); err != nil {
		return 0, 0, err
	}

	if locationResponse.Status != "0" {
		return 0, 0, fmt.Errorf("error from Baidu API: %s", locationResponse.Status)
	}

	return locationResponse.Result.Location.Lng, locationResponse.Result.Location.Lat, nil
}

func (b *BaiduMapService) GetRoute(origin string, destination string) (string, error) {
	// Implementation for getting route from Baidu Map API
	return "", nil
}