package model

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	Topic string
	What  string
}

func ReadSkill() (sk Skill) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What have you learnt today?")
	fmt.Print("Topic: ")
	fmt.Scan(&sk.Topic)

	fmt.Print("What: ")
	text, _ := reader.ReadString('\n')
	sk.What = strings.TrimSpace(text)
	return
}
