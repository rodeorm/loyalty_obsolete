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
		orders, err := storage.SelectProcessingOrders()
		if err != nil {
			fmt.Println("Ошибка при получении заказов на обновление статуса", err)
			break
		}
		for _, order := range *orders {
			url := fmt.Sprintf(accrualSystemAddress + "/api/orders/" + order.Number)
			r, _ := http.Get(url)
			if r.StatusCode != 200 {
				break
			}

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
		time.Sleep(50 * time.Millisecond)
	}
}
