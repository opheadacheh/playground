package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	commonpb "route_plan/server/proto/common"
	servicepb "route_plan/server/proto/service"
)

func testPlanRouteHttp() {
	jsonStr := []byte(`{"gps_info":{"latitude": 22.570134, "longitude": 113.944782}, "keywords": ["火锅", "按摩"]}`)
	url := "https://www.go-route-plan.top/v1/example/echo"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func testPlanRouteGrpc() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := servicepb.NewRoutePlanClient(conn)
	ctx := context.Background()

	// how to perform healthcheck request manually:
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	resp, err := c.PlanRoute(ctx, &servicepb.PlanRouteRequest{
		GpsInfo: &commonpb.GpsInfo{
			Latitude:  22.570134,
			Longitude: 113.944782,
		},
		Keywords: []string{
			"火锅",
			"按摩",
		},
	})
	if err != nil {
		log.Fatalf("PlanRoute failed %+v", err)
	}

	log.Printf("response: %v\n", resp)
}

func main() {
	testPlanRouteGrpc()
	// testPlanRouteHttp()
}
