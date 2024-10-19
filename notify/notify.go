package notify

import (
	"fmt"
	"os"
	"os/exec"
)

type Notifier struct {
	Cmd string
}

// Basically a wrapper for libnotify. Uses notify-send command
func InitNotifier() *Notifier {
	cmd, err := exec.LookPath("notify-send")
	if err != nil {
		fmt.Println("Error: libnotify not installed or notify-send command not found")
		os.Exit(5)
	}

	return &Notifier{Cmd: cmd}
}

func (n *Notifier) Notify(title, message string) {
	exec.Command(n.Cmd, title, message, "-a", "golearn").Run()
}
