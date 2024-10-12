package main

import (
	"fmt"
	"os"

	"github.com/rchirinos11/golearn/model"
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
	case "add":
		service.AddSkill()
	case "list":
		service.PrintSkills()
	case "delete_all":
		service.DeleteAll()
	default:
		fmt.Printf("%s is not a valid argument\n", args[1])
		os.Exit(2)
	}
}
