package main

import (
	"fmt"
	"logiTrack/config"
	"logiTrack/internal/database"
	"logiTrack/internal/httpapi"
	"logiTrack/internal/logger"
	"logiTrack/internal/order"
	"net/http"
)

func main() {

	dbConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		logger.Log.WithError(err).Fatal("Критическая ошибка загрузки конфигурации БД")
	}

	db, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		logger.Log.WithError(err).Fatal("Критическая ошибка инициализации БД")
	}
	defer db.Close()

	//storage := *order.NewOrderStorageMock()
	storage := order.NewPostgresOrderStorage(db.DB)
	hand := httpapi.NewOrderHandler(storage)

	//TODO Убрать подальше
	//Order1 := order.Order{
	//	ID:           1,
	//	CustomerName: "Ivan",
	//	Status:       "Delivered",
	//}
	//Order2 := order.Order{
	//	ID:           2,
	//	CustomerName: "Aleksandr",
	//	Status:       "Cancelled",
	//}
	//Order3 := order.Order{
	//	ID:           3,
	//	CustomerName: "Maria",
	//	Status:       "Pending",
	//}
	//Order4 := order.Order{
	//	ID:           4,
	//	CustomerName: "Norm",
	//	Status:       "Shipped",
	//}
	//Order5 := order.Order{
	//	ID:           5,
	//	CustomerName: "Elena",
	//	Status:       "Processing",
	//}
	//Order6 := order.Order{
	//	ID:           6,
	//	CustomerName: "Dmitry",
	//	Status:       "Delivered",
	//}
	//Order7 := order.Order{
	//	ID:           7,
	//	CustomerName: "Olga",
	//	Status:       "Returned",
	//}
	//Order8 := order.Order{
	//	ID:           8,
	//	CustomerName: "Sergey",
	//	Status:       "Awaiting Payment",
	//}
	//Order9 := order.Order{
	//	ID:           9,
	//	CustomerName: "Tatiana",
	//	Status:       "Cancelled",
	//}
	//
	//storage.Save(&Order1)
	//order.SafeFuncSaveOrder(&Order3, &storage) // попытка записать опасный слайс
	//storage.Save(&Order2)
	//order.SafeFuncSaveOrder(&Order4, &storage) // попытка записать норм слайс в безопасной функции
	//order.SafeFuncSaveOrder(&Order5, &storage)
	//order.SafeFuncSaveOrder(&Order6, &storage)
	//order.SafeFuncSaveOrder(&Order7, &storage)
	//order.SafeFuncSaveOrder(&Order8, &storage)
	//order.SafeFuncSaveOrder(&Order9, &storage)

	mux := http.NewServeMux()
	hand.RegisterRoutes(mux)

	//wrappedMux := httpapi.LoggingMiddleware(mux)
	chainedHandler := httpapi.RecoverMiddleware(
		httpapi.LoggingMiddleware(mux),
	)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", chainedHandler)

}
