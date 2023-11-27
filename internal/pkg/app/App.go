// Само рабочее приложение
package app

import (
	"fmt"

	"github.com/RB-PRO/sibintek/internal/app/jsonas"
	"github.com/RB-PRO/sibintek/internal/app/loges"
)

// Структура рабочего приложения
type App struct {
	llog   *loges.Core
	flages Flags
	cf     jsonas.Config
}

func New() (*App, error) {

	// Парсинг командной строки
	flags := flugger()

	// Запуск пакета логгирования
	llog, ErrLog := loges.New(flags.logsFile)
	if ErrLog != nil {
		fmt.Printf("Не удалось подключить логгер по причине: %v\n", ErrLog)
	}
	llog.Info(`Настройка флагов: %+v`, flags)
	llog.Info(`Запустил логгер в папке "%s"`, flags.logsFile)

	// Чтение конфига
	cf, ErrConfig := jsonas.ReadConfig(flags.configFile)
	if ErrConfig != nil || cf.URL == "" {
		llog.Warn(`Не смог прочитать конфиг из файла "%s"`, flags.configFile)
		cf.URL = "https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/27185784/detail?languageId=-1&appId=1"
		llog.Info(`Использую для работы URL "%s"`, cf.URL)
	} else {
		llog.Info(`Прочитал конфиг из файла "%s"`, flags.configFile)
	}

	return &App{
		cf:     cf,
		llog:   llog,
		flages: flags,
	}, nil
}
