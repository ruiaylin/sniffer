package logger
import (
	"fmt"
	"os"
)


type Logger struct {
	ClassName string
}

func (logger *Logger) Debug(format string, args ... interface{}) {
	prefix := fmt.Sprintf("%20v => ", logger.ClassName)
	if len(args) > 0 {
		fmt.Printf(prefix + format + "\n", args)
	} else {
		fmt.Println(prefix + format)
	}
}

func (logger *Logger) Error(format string, args ... interface{}) {
	prefix := fmt.Sprintf("%v => ", logger.ClassName)
	if len(args) > 0 {
		fmt.Printf(prefix + format + "\n", args)
	} else {
		fmt.Println(prefix + format)
	}
	os.Exit(-1)
}