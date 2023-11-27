// Пакет jsonas используется для взаимодействия с json-файлами
package jsonas

import (
	"encoding/json"
	"fmt"
	"os"
)

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
func ReadJsonInts(FileName string) (data []int, Err error) {
	if FileName == "" {
		return nil, fmt.Errorf("file name is zero length")
	}
	dataBytes, Err := os.ReadFile(FileName)
	if Err != nil {
		return nil, fmt.Errorf("os.ReadFile: %v", Err)
	}
	Err = json.Unmarshal(dataBytes, &data)
	if Err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", Err)
	}
	return data, nil
}

// Парсинг строки в []int
func UnmarshalInts(input string) (data []int, Err error) {
	if Err = json.Unmarshal([]byte(input), &data); Err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", Err)
	}
	return data, nil
}
