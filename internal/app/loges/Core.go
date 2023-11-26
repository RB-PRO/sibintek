// Пакет самодельного логгера
//
// Вообще обычно такое не делают и используют что-то вроде zap, logrus, zerolog и тд
package loges

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Интерфейс логгера
type Logger interface {
	// Логировать состояние
	Info(layout string, value ...interface{})

	// Логгировать предупреждение
	Warn(layout string, value ...interface{})

	// Логгировать ошибку
	Err(layout string, value ...interface{})
}

// Ядро логгера
type Core struct {
	logInfo    *log.Logger
	logWarning *log.Logger
	logError   *log.Logger
}

// Создать объект логгирования
func New(ZeroDirection string) (*Core, error) {

	// Проверка существования папки
	if _, err := os.Stat(ZeroDirection); os.IsNotExist(err) {
		return nil, fmt.Errorf("folder '%s' does not exist", ZeroDirection)
	}

	// Путь до файла с логами
	FilePathName := ZeroDirection + "/" + time.Now().Format("2006-01-02_15-04") + ".log"
	flags := log.LstdFlags | log.Lshortfile
	FileInfo, _ := os.OpenFile(FilePathName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	FileWarinig, _ := os.OpenFile(FilePathName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	FileError, _ := os.OpenFile(FilePathName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	return &Core{
		logInfo:    log.New(FileInfo, "Info:\t", flags),
		logWarning: log.New(FileWarinig, "Warn:\t", flags),
		logError:   log.New(FileError, "Error:\t", flags),
	}, nil
}
