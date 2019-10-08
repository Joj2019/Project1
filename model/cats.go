package model

import (
	"errors"
	"log"

	"github.com/hysem/charts/container"
	"github.com/hysem/charts/domain"
)

const (
	catQuery string = "SELECT `category`, `count` FROM cats"
)

// GetCatsChartData retrieves the cats data from db
func GetCatsChartData() ([]domain.ChartData, error) {
	db := container.DBReader()
	result := []domain.ChartData{}
	if err := db.Select(&result, catQuery); err != nil {
		log.Println("Error while fetching cats data from db ->", err)
		return result, errors.New("Unable to complete the request")
	}
	return result, nil
}
