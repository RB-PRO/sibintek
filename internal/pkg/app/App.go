// Само рабочее приложение
package app

import (
	"github.com/RB-PRO/sibintek/internal/app/loges"
)

type App struct {
	folderLogs   string
	url          string
	nameJsonFile string
	llog         *loges.Core
}

func New() (*App, error) {

	// Корневая папка для файла логов
	FolderLogs := "logs"

	// Ссылка для запроса
	Url := "https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/27185784/detail?languageId=-1&appId=1"

	// Название файла json
	NameJsonFile := "data.json"

	// Запуск пакета логгирования
	llog, ErrLog := loges.New(FolderLogs)
	if ErrLog != nil {
		panic(ErrLog)
	}
	return &App{
		folderLogs:   FolderLogs,
		url:          Url,
		nameJsonFile: NameJsonFile,
		llog:         llog,
	}, nil
}
