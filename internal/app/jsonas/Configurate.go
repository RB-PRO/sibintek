package jsonas

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	URL string
}

// Считать слайс []int из файла FileName
//
// В аргументе передайте путь до файла с необходимой информацией.
//
// Пример файла:
//
//	`[
//		1,
//		2,
//		3
//	]`
func ReadConfig(FileName string) (cf Config, Err error) {
	if FileName == "" {
		return Config{}, fmt.Errorf("file name is zero length")
	}
	dataBytes, Err := os.ReadFile(FileName)
	if Err != nil {
		return Config{}, fmt.Errorf("os.ReadFile: %v", Err)
	}
	Err = json.Unmarshal(dataBytes, &cf)
	if Err != nil {
		return Config{}, fmt.Errorf("json.Unmarshal: %v", Err)
	}
	return cf, nil
}
