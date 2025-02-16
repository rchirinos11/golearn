package cmd

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

type Config struct {
	home    string
	service *service.SkillService
}

func (c *Config) RunCli() {
	c.initService()
	c.parseArgs()
}

func (c *Config) initService() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Skill{})

	c.service = &service.SkillService{DB: db}
}

func (c *Config) parseArgs() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No args provided")
		os.Exit(1)
	}

	switch args[1] {
	case "start":
		go c.start()
		select {}
	case "add":
		c.service.AddSkill()
	case "list":
		c.service.PrintSkills()
	case "del":
		c.service.DeleteOne(args)
	case "edit":
		c.service.Edit(args)
	case "byid":
		c.service.FilterBy(args, 1)
	case "bydate":
		c.service.FilterBy(args, 2)
	case "delete_all":
		c.service.DeleteAll()
	case "filter":
		c.service.FilterBy(args, 0)
	case "test":
		c.Write()
	default:
		fmt.Printf("%s is not a valid argument\n", args[1])
		os.Exit(2)
	}
}

func (c *Config) start() {
	notifyCmd := notify.InitNotifier()
	interval := time.Hour * 4
	fmt.Println("Started service, recurrence:", interval)
	printNext(interval)

	for {
		select {
		case <-time.Tick(interval):
			notifyCmd.Notify("Golearn", "Stop being lazy")
			c.service.AddSkill()
			printNext(interval)
		}
	}
}

func printNext(interval time.Duration) {
	fmt.Println("Next notification at", time.Now().Add(interval).Format("15:04:05"))
}

func (c *Config) Write() {
	filename := c.home + "/config"
	file, err := os.Open(filename)
	if file == nil || err != nil {
		fmt.Println("File doesn't exist")
		os.Exit(1)
	}
	var content []byte
	_, err = file.Read(content)
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(2)
	}
}
