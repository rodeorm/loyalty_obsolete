package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"loyalty/internal/model"
	"loyalty/internal/repo"
	"net/http"
	"time"
)

func StartClient(storage repo.Storage, accrualSystemAddress string) {
	go makeGetRequest(storage, accrualSystemAddress)
}

func makeGetRequest(storage repo.Storage, accrualSystemAddress string) {
	for {
		orders, _ := storage.SelectProcessingOrders()
		for _, order := range *orders {
			fmt.Println("Попытка обновить", order.Number)
			url := fmt.Sprintf(accrualSystemAddress + "/api/orders/" + order.Number)
			r, _ := http.Get(url)
			order := model.ExtOrder{}

			bodyBytes, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				log.Println(err)
				break
			}

			err = json.Unmarshal(bodyBytes, &order)
			if err != nil {
				log.Println(err)
				break
			}
			fmt.Println(order)
			ctx := context.TODO()
			storage.UpdateOrder(ctx, &order)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
