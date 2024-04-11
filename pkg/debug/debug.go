package debug

import "fmt"

type Debugger struct {
	Name  string
	Level string
}

func (d Debugger) Log(level string, text string) {
	if d.Level == level {
		fmt.Println("["+d.Name+"]", text)
	}
}
