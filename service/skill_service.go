package service

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rchirinos11/golearn/model"
	"gorm.io/gorm"
)

type Service interface {
	AddSkill()
	PrintSkills()
	DeleteAll() error
	Edit(args []string)
	PrintByDate()
	DeleteOne(args []string)
}

type SkillService struct {
	DB *gorm.DB
}

func (sr *SkillService) AddSkill() {
	newSkill := model.ReadSkill()
	sr.DB.Create(&newSkill)
	fmt.Println("Saved")
}

func (sr *SkillService) PrintSkills() {
	var skills []model.Skill
	sr.DB.Find(&skills)
	fmt.Println("In total, you have learnt:")
	printSeparator()
	fmt.Printf("%-5s %-12s %-40s %s\n", "Index", "Topic", "What", "When")
	printSeparator()

	for _, skill := range skills {
		fmt.Printf("%-5d %-12s %-40s %s\n", skill.ID, skill.Topic, skill.What, skill.CreatedAt.Format(time.UnixDate))
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
		os.Remove("../test.db")
	}
}

func (sr *SkillService) DeleteOne(args []string) {
	modifierArgError(args)
	sr.DB.Delete(&model.Skill{}, args[2])
}

func (sr *SkillService) Edit(args []string) {
	modifierArgError(args)
	id, err := strconv.ParseUint(args[2], 10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(4)
	}
	updated := model.ReadSkill()
	updated.ID = uint(id)
	sr.DB.Save(&updated)
}

func (sr *SkillService) PrintByDate() {
}

func modifierArgError(args []string) {
	if len(args) < 3 {
		fmt.Println("Please provide an id for this option")
		fmt.Println("Format: golearn <option> <id>")
		os.Exit(3)
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
