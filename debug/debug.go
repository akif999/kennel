package debug

import (
	"fmt"
	"os"
)

func Init() error {
	err := os.Remove("./debug.txt")
	if err != nil {
		return err
	}
	return nil
}

func Fprint(dInfo interface{}) error {
	content := fmt.Sprintf("%v\n", dInfo)
	f, err := os.OpenFile("./debug.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write([]byte(content)); err != nil {
		return err
	}

	return nil
}
