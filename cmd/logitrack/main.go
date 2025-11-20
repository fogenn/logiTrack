package main

import (
	"fmt"
	"logiTrack/internal/httpapi"
	"logiTrack/internal/order"
	"net/http"
)

func main() {

	OrderStorageMock := *order.NewOrderStorageMock()
	hand := httpapi.NewOrderHandler(&OrderStorageMock)

	Order1 := order.Order{
		ID:           1,
		CustomerName: "Ivan",
		Status:       "Delivered",
	}
	Order2 := order.Order{
		ID:           2,
		CustomerName: "Aleksandr",
		Status:       "Cancelled",
	}
	Order3 := order.Order{
		ID:           3,
		CustomerName: "Maria",
		Status:       "Pending",
	}
	Order4 := order.Order{
		ID:           4,
		CustomerName: "Norm",
		Status:       "Shipped",
	}
	Order5 := order.Order{
		ID:           5,
		CustomerName: "Elena",
		Status:       "Processing",
	}
	Order6 := order.Order{
		ID:           6,
		CustomerName: "Dmitry",
		Status:       "Delivered",
	}
	Order7 := order.Order{
		ID:           7,
		CustomerName: "Olga",
		Status:       "Returned",
	}
	Order8 := order.Order{
		ID:           8,
		CustomerName: "Sergey",
		Status:       "Awaiting Payment",
	}
	Order9 := order.Order{
		ID:           9,
		CustomerName: "Tatiana",
		Status:       "Cancelled",
	}

	OrderStorageMock.Save(&Order1)
	order.SafeFuncSaveOrder(&Order3, &OrderStorageMock) // попытка записать опасный слайс
	OrderStorageMock.Save(&Order2)
	order.SafeFuncSaveOrder(&Order4, &OrderStorageMock) // попытка записать норм слайс в безопасной функции
	order.SafeFuncSaveOrder(&Order5, &OrderStorageMock)
	order.SafeFuncSaveOrder(&Order6, &OrderStorageMock)
	order.SafeFuncSaveOrder(&Order7, &OrderStorageMock)
	order.SafeFuncSaveOrder(&Order8, &OrderStorageMock)
	order.SafeFuncSaveOrder(&Order9, &OrderStorageMock)

	mux := http.NewServeMux()
	hand.RegisterRoutes(mux)

	//wrappedMux := httpapi.LoggingMiddleware(mux)
	chainedHandler := httpapi.RecoverMiddleware(
		httpapi.LoggingMiddleware(mux),
	)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", chainedHandler)

}
