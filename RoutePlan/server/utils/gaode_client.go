package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	commonpb "route_plan/server/proto/common"
)

var (
	GaodeProdURL = "https://restapi.amap.com/"
	GaodeAppKey  = "3aeb5e3659d0f53264f1f07521c417de"
)

type GaodeHttpClientInterface interface {
	Search(gpsInfo *commonpb.GpsInfo, keywords string, limit int) ([]Location, error)
}

type GaodeHttpClient struct {
	client *http.Client
	url    string
	key    string
}

func NewGaodeHttpClient(httpClient *http.Client) (*GaodeHttpClient, error) {
	if httpClient == nil {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		httpClient = &http.Client{
			Timeout:   time.Duration(1) * time.Second,
			Transport: tr,
		}
	}

	return &GaodeHttpClient{
		client: httpClient,
		url:    GaodeProdURL,
		key:    GaodeAppKey,
	}, nil
}

type BusinessExtension struct {
	Cost   interface{} `json:"cost"`
	Rating interface{} `json:"rating"`
}

type Poi struct {
	Name         string            `json:"name"`
	Id           string            `json:"id"`
	Location     string            `json:"location"`
	Type         string            `json:"type"`
	ProvinceName string            `json:"pname"`
	CityName     string            `json:"cityname"`
	DistrictName string            `json:"adname"`
	Address      string            `json:"address"`
	PostalCode   string            `json:"pcode"`
	CityCode     string            `json:"citycode"`
	DistrictCode string            `json:"adcode"`
	Distance     string            `json:"distance"`
	BusinessExt  BusinessExtension `json:"biz_ext"`
}

type SearchResponse struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Pois     []Poi  `json:"pois"`
	Count    string `json:"count"`
}

func (c *GaodeHttpClient) Search(gpsInfo *commonpb.GpsInfo, keywords string, limit int) ([]Location, error) {
	requestUrl := fmt.Sprintf("%sv3/place/around?location=%f,%f&keywords=%s&key=%s&offset=%d&page=1&extensions=all", c.url, gpsInfo.GetLongitude(), gpsInfo.GetLatitude(), keywords, c.key, limit)
	resp, err := c.client.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResp SearchResponse
	if err := json.Unmarshal(respByte, &searchResp); err != nil {
		return nil, err
	}

	if searchResp.Info != "OK" {
		return nil, fmt.Errorf("failed to call Gaode search API, error: %s", searchResp.Info)
	}

	var locations []Location
	for _, poi := range searchResp.Pois {
		location, err := FromPoi(poi)
		if err != nil {
			return nil, fmt.Errorf("failed to convert Gaode poi to Location, error: %s", err)
		}

		locations = append(locations, location)
	}

	return locations, nil
}
