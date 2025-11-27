package order

import (
	"fmt"
	"sync"
	"time"
)

func StartDeliveryWorker(ch chan Order, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case order, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("\nOrder ID %v in worke -", order.ID)
			time.Sleep(3 * time.Second)
			fmt.Printf("\n - Order ID %v delivered", order.ID)

		case <-time.After(1 * time.Second):
			fmt.Println("\nДолго ждём, подождём ещё...")
		}

	}
}

func SafeFuncSaveOrder(order *Order, storage StorageIntf) error { // чисто эксперемент
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Перехвачена паника:", r)
		}
	}()
	storage.Save(order)
	return nil
}

func TESTSafeGetByID(ById int, storage OrderStorageMock) {
	order, _, err := storage.GetByID(ById)
	if err != nil {
		fmt.Println(fmt.Errorf("TESTSafeGetByID-Не удалось найти заказ  %+v: %w", order, err))

	} else {
		fmt.Printf("Корректный ID - %v", order)
	}
}
