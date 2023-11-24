package main

import (
	"training-go-hacktiv8-3/db"
	"github.com/jasonlvhit/gocron"
	_ "github.com/lib/pq"
)

func main() {
	db.Connect()
	gocron.Every(15).Seconds().Do(
		// ctrl.sensorRun
			db.Connect()
			waterValue := rand.Float64() * 100
			windValue := rand.Float64() * 100
		
			sensor := model.SensorData{}
			err := db.GetDB().Create(&sensor).Error
		
			if err != nil {
				panic(err)
			}
		
			if windValue < 6 && waterValue < 5 {
				fmt.Printf("Status Wind: aman", "Status Water: aman")
			} else if windValue >= 7 && windValue <= 15 && (waterValue >= 6 && waterValue <= 8) {
				fmt.Printf("Status Wind: siaga", "Status Water: siaga")
			} else {
				fmt.Printf("Status Wind: bahaya", "Status Water: bahaya")
			}
	)
	<-gocron.Start()
}
