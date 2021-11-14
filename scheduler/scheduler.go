package scheduler

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func mainScheduler() {
	// fmt.Println("I am running task.")
}

func Run() {
	mainScheduler()
	go func() {
		gocron.Every(10).Second().Do(mainScheduler)
		<-gocron.Start()
	}()
}
