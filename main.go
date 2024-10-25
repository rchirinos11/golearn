package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rchirinos11/golearn/model"
	"github.com/rchirinos11/golearn/notify"
	"github.com/rchirinos11/golearn/service"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Skill{})

	service := service.SkillService{DB: db}
	parseArgs(&service)
}

func parseArgs(service *service.SkillService) {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No args provided")
		os.Exit(1)
	}

	switch args[1] {
	case "start":
		go start(service)
		select {}
	case "add":
		service.AddSkill()
	case "list":
		service.PrintSkills()
	case "del":
		service.DeleteOne(args)
	case "edit":
		service.Edit(args)
	case "bydate":
		service.PrintByDate()
	case "delete_all":
		service.DeleteAll()
	default:
		fmt.Printf("%s is not a valid argument\n", args[1])
		os.Exit(2)
	}
}

func start(service *service.SkillService) {
	notifyCmd := notify.InitNotifier()
	interval := time.Hour * 4
	fmt.Println("Started service, recurrence:", interval)
	printNext(interval)

	for {
		select {
		case <-time.Tick(interval):
			notifyCmd.Notify("Golearn", "Stop being lazy")
			service.AddSkill()
			printNext(interval)
		}
	}
}

func printNext(interval time.Duration) {
	fmt.Println("Next notification at", time.Now().Add(interval).Format("15:04:05"))
}
