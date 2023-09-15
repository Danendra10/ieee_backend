package seeders

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/danendra10/ieee_backend/database"
	"github.com/danendra10/ieee_backend/models"
)

var compounds []string = []string{
	"CO",
	"CO2",
	"NH3",
	"NH4",
	"O3",
}

var lat_lng [][]float64 = [][]float64{
	{-7.285985, 112.796380},
	{-7.281926, 112.798379},
	{-7.280807, 112.797866},
	{-7.278739, 112.797780},
	{-7.277882, 112.796707},
	{-7.276931, 112.795976},
	{-7.280481, 112.791538},
	{-7.282029, 112.792527},
	{-7.283395, 112.793060},
	{-7.287203, 112.793361},
	{-7.288559, 112.792769},
	{-7.285991, 112.795503},
	{-7.287443, 112.795745},
	{-7.290511, 112.793816},
}

func RandomizeCompounds(arr []string) (res string) {
	return arr[rand.Intn(len(arr)-1)]
}

func RandomizeLatLng(arr [][]float64) (lat, lng float64) {
	randomIndex := rand.Intn(len(arr))
	return arr[randomIndex][0], arr[randomIndex][1]
}

func GenerateCreatedAt(day time.Time) time.Time {
	dayStart := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	dayEnd := dayStart.Add(24 * time.Hour)

	diff := dayEnd.Sub(dayStart)
	randomSeconds := rand.Int63n(int64(diff.Seconds()))

	created_at := dayStart.Add(time.Duration(randomSeconds) * time.Second)

	return created_at
}
func SeedAirPolutions() error {
	for i := 0; i < 100; i++ {
		component := RandomizeCompounds(compounds)
		lat, lng := RandomizeLatLng(lat_lng)

		var days_ago int
		if i < 10 {
			days_ago = 10
		} else if i < 20 {
			days_ago = 9
		} else if i < 30 {
			days_ago = 8
		} else if i < 40 {
			days_ago = 7
		} else if i < 50 {
			days_ago = 6
		} else if i < 60 {
			days_ago = 5
		} else if i < 70 {
			days_ago = 4
		} else if i < 80 {
			days_ago = 3
		} else if i < 90 {
			days_ago = 2
		} else {
			days_ago = 1
		}

		created_at := GenerateCreatedAt(time.Now().AddDate(0, 0, -days_ago))

		airPolution := models.AirPolution{
			Lat:       lat,
			Lng:       lng,
			Compound:  component,
			CreatedAt: created_at,
		}

		if err := database.DB.Create(&airPolution).Error; err != nil {
			return fmt.Errorf("error creating air pollution: %v", err)
		}

		if i%10 == 0 {
			fmt.Println("Seeded", i, "Data")
		}
	}

	return nil
}
