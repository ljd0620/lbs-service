package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GaodeMapService struct {
	APIKey string
}

type LocationResponse struct {
	Status string `json:"status"`
	Result struct {
		Location string `json:"location"`
	} `json:"result"`
}

func (g *GaodeMapService) GetLocation(address string) (string, error) {
	url := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?key=%s&address=%s", g.APIKey, address)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get location: %s", resp.Status)
	}

	var locationResponse LocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationResponse); err != nil {
		return "", err
	}

	if locationResponse.Status != "1" {
		return "", fmt.Errorf("error from Gaode API: %s", locationResponse.Status)
	}

	return locationResponse.Result.Location, nil
}

func (g *GaodeMapService) GetRoute(origin, destination string) (string, error) {
	url := fmt.Sprintf("https://restapi.amap.com/v3/direction/driving?key=%s&origin=%s&destination=%s", g.APIKey, origin, destination)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get route: %s", resp.Status)
	}

	var routeResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&routeResponse); err != nil {
		return "", err
	}

	if routeResponse["status"] != "1" {
		return "", fmt.Errorf("error from Gaode API: %v", routeResponse["info"])
	}

	// Assuming the route information is in routeResponse["route"]
	routeInfo, _ := json.Marshal(routeResponse["route"])
	return string(routeInfo), nil
}