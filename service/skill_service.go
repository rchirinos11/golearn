package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rchirinos11/golearn/model"
	"gorm.io/gorm"
)

type Service interface {
	AddSkill()
	PrintSkills()
	DeleteAll() error
	Edit()
	PrintByDate()
	DeleteOne()
}

type SkillService struct {
	DB *gorm.DB
}

func (sr *SkillService) AddSkill() {
	var newSkill model.Skill
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What have you learnt today?")
	fmt.Print("Topic: ")
	fmt.Scan(&newSkill.Topic)

	fmt.Print("What: ")
	text, _ := reader.ReadString('\n')
	newSkill.What = strings.TrimSpace(text)

	fmt.Println("Saving...")

	sr.DB.Create(&newSkill)
}

func (sr *SkillService) PrintSkills() {
	var skills []model.Skill
	sr.DB.Find(&skills)
	fmt.Println("In total, you have learnt:")
	printSeparator()
	fmt.Printf("%-5s %-12s %-40s %s\n", "Index", "Topic", "What", "When")
	printSeparator()

	for i, skill := range skills {
		fmt.Printf("%-5d %-12s %-40s %s\n", i+1, skill.Topic, skill.What, skill.CreatedAt.Format(time.UnixDate))
	}
}

func (sr *SkillService) DeleteAll() {
	fmt.Print("Are you sure? this will delete everything (y/N) ")
	if evalInput() {
		fmt.Println("Enter else")
		return
	} else {
		fmt.Println("Enter else")
	}
	fmt.Print("This is meant for testing only, I will literally delete the db file, are you really sure? (y/N) ")
	if evalInput() {
		os.Remove("../bkp.db")
	}
}

func evalInput() bool {
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)

	switch input {
	case "y":
		fallthrough
	case "yes":
		return false
	}

	fmt.Println("Exiting...")
	return true
}

func printSeparator() {
	for i := 0; i < 90; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}
