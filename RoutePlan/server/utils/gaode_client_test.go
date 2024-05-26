package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	commonpb "route_plan/server/proto/common"
)

func TestSearch(t *testing.T) {
	// Create a test server that returns a predefined JSON response.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{
            "suggestion": {
                "keywords": [],
                "cities": []
            },
            "count": "434",
            "infocode": "10000",
            "pois": [
                {
                    "parent": [],
                    "distance": "304",
                    "pcode": "440000",
                    "importance": [],
                    "biz_ext": {
                        "cost": "122.00",
                        "rating": "4.7",
                        "meal_ordering": "0"
                    },
                    "recommend": "0",
                    "type": "餐饮服务;餐饮相关场所;餐饮相关",
                    "photos": [
                        {
                            "title": [],
                            "url": "http://store.is.autonavi.com/showpic/789f077f4f2a2d36493b0a14b0f17b02"
                        },
                        {
                            "title": [],
                            "url": "http://store.is.autonavi.com/showpic/9f6bb8e747b730a12a7af9a855fa04cb"
                        },
                        {
                            "title": [],
                            "url": "http://store.is.autonavi.com/showpic/fdfabebe10aa407a9593e66455450dec"
                        }
                    ],
                    "discount_num": "0",
                    "gridcode": "3313678521",
                    "typecode": "050000",
                    "shopinfo": "1",
                    "poiweight": [],
                    "citycode": "0755",
                    "adname": "南山区",
                    "children": [],
                    "alias": [],
                    "tel": "0755-26927998",
                    "id": "B0FFL1LUAM",
                    "tag": "鸳鸯锅,涮肉,羊肉,羊火锅,羊肉火锅",
                    "event": [],
                    "entr_location": [],
                    "indoor_map": "0",
                    "email": [],
                    "timestamp": "2023-11-08 14:07:22",
                    "website": [],
                    "address": "打石一路与仙鼓路交汇处万科云城东里停车场厂入口",
                    "adcode": "440305",
                    "pname": "广东省",
                    "biz_type": "diner",
                    "cityname": "深圳市",
                    "postcode": [],
                    "match": "0",
                    "business_area": "西丽",
                    "indoor_data": {
                        "cmsid": [],
                        "truefloor": [],
                        "cpid": [],
                        "floor": []
                    },
                    "childtype": [],
                    "exit_location": [],
                    "name": "快乐小羊(万科云城)",
                    "location": "113.945474,22.572797",
                    "shopid": [],
                    "navi_poiid": [],
                    "groupbuy_num": "0"
                },
                {
                    "parent": "B0FFILB6PO",
                    "distance": "419",
                    "pcode": "440000",
                    "importance": [],
                    "biz_ext": {
                        "cost": [],
                        "rating": "4.5",
                        "meal_ordering": "0"
                    },
                    "recommend": "0",
                    "type": "餐饮服务;中餐厅;火锅店",
                    "photos": [
                        {
                            "title": [],
                            "url": "http://store.is.autonavi.com/showpic/acfaf5bd13c48115f02f0ec0a18a25e2"
                        },
                        {
                            "title": [],
                            "url": "http://store.is.autonavi.com/showpic/c904e79dcd1ac6627e5478bd4af9c1bc"
                        },
                        {
                            "title": [],
                            "url": "http://store.is.autonavi.com/showpic/7df1914642af9dec59a4801b7c426a25"
                        }
                    ],
                    "discount_num": "0",
                    "gridcode": "3313678521",
                    "typecode": "050117",
                    "shopinfo": "1",
                    "poiweight": [],
                    "citycode": "0755",
                    "adname": "南山区",
                    "children": [],
                    "alias": [],
                    "tel": "17324421208",
                    "id": "B0H224GI0H",
                    "tag": "鸳鸯锅,千层肚,菌汤,肥牛",
                    "event": [],
                    "entr_location": "113.942741,22.573330",
                    "indoor_map": "0",
                    "email": [],
                    "timestamp": "2023-11-10 12:58:36",
                    "website": [],
                    "address": "西丽街道云城万科西里商业区B1",
                    "adcode": "440305",
                    "pname": "广东省",
                    "biz_type": "diner",
                    "cityname": "深圳市",
                    "postcode": [],
                    "match": "0",
                    "business_area": "西丽",
                    "indoor_data": {
                        "cmsid": [],
                        "truefloor": [],
                        "cpid": [],
                        "floor": []
                    },
                    "childtype": "324",
                    "exit_location": [],
                    "name": "大龙燚(万科店)",
                    "location": "113.942629,22.573333",
                    "shopid": [],
                    "navi_poiid": [],
                    "groupbuy_num": "0"
                }
            ],
            "status": "1",
            "info": "OK"
        }`))
	}))
	defer ts.Close()

	h, err := NewGaodeHttpClient(ts.Client())
	if err != nil {
		t.Fatalf("Expected no error while creating client but got: %v", err)
	}
	h.url = ts.URL + "/"

	locations, err := h.Search(&commonpb.GpsInfo{Latitude: 0.1, Longitude: 0.1}, "test", 2)
	if err != nil {
		t.Fatal(err)
	}

	expectedNumOfLocations := 2
	if len(locations) != expectedNumOfLocations {
		t.Errorf("Expected %d locations, but got %d", expectedNumOfLocations, len(locations))
	}

	expectedFirstLocationName := "快乐小羊(万科云城)"
	if locations[0].Name != expectedFirstLocationName {
		t.Errorf("Expected first business name to be %s, but got %s", expectedFirstLocationName, locations[0].Name)
	}

	expectedSecondLocationName := "大龙燚(万科店)"
	if locations[1].Name != expectedSecondLocationName {
		t.Errorf("Expected first business name to be %s, but got %s", expectedSecondLocationName, locations[1].Name)
	}

	expectedFirstLocationDistance := float64(304)
	if locations[0].Distance != expectedFirstLocationDistance {
		t.Errorf("Expected first business distance to be %f, but got %f", expectedFirstLocationDistance, locations[0].Distance)
	}

	expectedFirstLocationLatitude := 22.572797
	if locations[0].Latitude != expectedFirstLocationLatitude {
		t.Errorf("Expected first business distance to be %f, but got %f", expectedFirstLocationLatitude, locations[0].Latitude)
	}

	expectedFirstLocationLongitude := 113.945474
	if locations[0].Longitude != expectedFirstLocationLongitude {
		t.Errorf("Expected first business distance to be %f, but got %f", expectedFirstLocationLongitude, locations[0].Longitude)
	}

	expectedFirstLocationCost := 122.0
	if locations[0].Cost != expectedFirstLocationCost {
		t.Errorf("Expected first business distance to be %f, but got %f", expectedFirstLocationCost, locations[0].Cost)
	}

	expectedFirstLocationRating := 4.7
	if locations[0].Rating != expectedFirstLocationRating {
		t.Errorf("Expected first business distance to be %f, but got %f", expectedFirstLocationRating, locations[0].Rating)
	}
}
