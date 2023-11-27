package jsonas

import (
	"os"
)

// Простое сохранение результата в файл
func Save(FileName string, str string) error {
	return os.WriteFile(FileName, []byte(str), 0666)
}
