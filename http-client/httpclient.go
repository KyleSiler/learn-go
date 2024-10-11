package main

import (
	"fmt"
	"io"
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
		fmt.Println(err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fileName := "data/" + location + "-" + time.Now().Format("20060102") + ".json"
	os.WriteFile(fileName, body, 0666)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go callAPI("roseville", "https://www.autonationchryslerdodgejeepramroseville.com/apis/widget/INVENTORY_LISTING_DEFAULT_AUTO_NEW:inventory-data-bus1/getInventory?bodyStyle=Truck%20Crew%20Cab&compositeType=new&make=Jeep&model=Gladiator&year=2024&pageSize=50", &wg)
	wg.Wait()
}
