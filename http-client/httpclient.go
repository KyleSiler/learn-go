package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type CatFact struct {
	Text string
	Type string
}

func callAPI(location string, endpoint string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Println(err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	fileName := "data/" + location + "-" + time.Now().Format("20060102") + ".json"
	os.WriteFile(fileName, body, 0666)
}

func postAPI(location string, endpoint string, requestFilename string, wg *sync.WaitGroup) {
	defer wg.Done()

	requestBody, err := os.ReadFile("requests/" + requestFilename)
	if err != nil {
		log.Println(err.Error())
		return
	}
	resp, err := http.Post(endpoint, "application/json", bytes.NewReader(requestBody))
	fmt.Println(resp.Request.Header)
	if err != nil {
		log.Println(err.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fileName := "data/" + location + "-" + time.Now().Format("20060102") + ".json"
	os.WriteFile(fileName, body, 0666)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go callAPI("roseville", "https://www.autonationchryslerdodgejeepramroseville.com/apis/widget/INVENTORY_LISTING_DEFAULT_AUTO_NEW:inventory-data-bus1/getInventory?bodyStyle=Truck%20Crew%20Cab&compositeType=new&make=Jeep&model=Gladiator&year=2024&pageSize=50", &wg)
	go callAPI("sacrmento", "https://www.sacsuperstore.com/apis/widget/INVENTORY_LISTING_DEFAULT_AUTO_NEW:inventory-data-bus1/getInventory?make=Jeep&model=Gladiator&year=2024&pageSize=50", &wg)
	go callAPI("folsom", "https://www.sacsuperstore.com/apis/widget/INVENTORY_LISTING_DEFAULT_AUTO_NEW:inventory-data-bus1/getInventory?make=Jeep&model=Gladiator&year=2024&pageSize=50", &wg)
	go callAPI("reno", "https://www.lithiajeepreno.com/apis/widget/INVENTORY_LISTING_DEFAULT_AUTO_NEW:inventory-data-bus1/getInventory?year=2024&gvBodyStyle=Truck&model=Gladiator&make=Jeep&pageSize=50", &wg)
	go postAPI("elkgrove", "https://equ6hxb6wg-dsn.algolia.net/1/indexes/*/queries?x-algolia-agent=Algolia%20for%20JavaScript%20(4.9.1)%3B%20Browser%20(lite)%3B%20JS%20Helper%20(3.4.4)&x-algolia-api-key=da97ef494552f47ecc6f47068888d120&x-algolia-application-id=EQU6HXB6WG", "elkgrove-request.json", &wg)
	go postAPI("placerville", "https://2591j46p8g-dsn.algolia.net/1/indexes/*/queries?x-algolia-agent=Algolia%20for%20JavaScript%20(4.9.1)%3B%20Browser%20(lite)%3B%20JS%20Helper%20(3.4.4)&x-algolia-api-key=78311e75e16dd6273d6b00cd6c21db3c&x-algolia-application-id=2591J46P8G", "placerville-request.json", &wg)
	wg.Wait()
}
