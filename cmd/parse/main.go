package main

import (
	db2 "ETL/internal/db"
	"ETL/internal/models"
	"ETL/internal/static"
	"ETL/internal/web"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
)

func main() {
	db, err := db2.InitDB()
	if err != nil {
		panic(err)
	}
	err = db2.AutoMigrate(db)
	if err != nil {
		panic(err)
	}
	client := resty.New()
	url := static.FirstUrl

	for i := range static.TotalRequests {
		body, err := web.GetBody(client, url)
		if err != nil {
			panic(err)
		}

		var response models.ApiResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			panic(err)
		}

		for _, event := range response.Results {
			if err := db.Create(&event).Error; err != nil {
				log.Fatalf("failed to create event: %v", err)
			}
		}
		url = response.Next
		log.Printf("requests processed: %d\n", i)
	}
}
