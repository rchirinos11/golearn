package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rchirinos11/golearn/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Skill{})

	parseArgs(db)
}

func parseArgs(db *gorm.DB) {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("I have no args!")
		return
	}

	switch args[1] {
	case "add":
		addSkill(db)
	case "list":
		listSkills(db)
	default:
		fmt.Println("I don't know what you want me to do!")
	}
}

func listSkills(db *gorm.DB) {
	var skills []model.Skill
	db.Find(&skills)
	fmt.Println("In total, you have learnt:")
	fmt.Println("===========================================")
	fmt.Printf("%-15s %s\n", "What", "When")
	for _, skill := range skills {
		fmt.Printf("%-15s %s\n", skill.What, skill.CreatedAt.Format(time.UnixDate))
	}
}

func addSkill(db *gorm.DB) {
	var newSkill model.Skill
	fmt.Println("What have you learnt today?")
	fmt.Scan(&newSkill.What)
	fmt.Println("You have learnt:", newSkill.What)

	db.Create(&newSkill)
}
