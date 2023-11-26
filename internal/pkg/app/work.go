package app

import (
	"fmt"

	"github.com/RB-PRO/sibintek/internal/app/httpas"
	"github.com/RB-PRO/sibintek/internal/app/jsonas"
)

// Запуск самого приложения
func (a *App) Run() error {

	a.llog.Info("Начало работы приложения")
	a.llog.Info(`Запустил логгер в папке "%s"`, a.folderLogs)

	// Прочитать файл в слайс
	DataJsonFile, ErrJson := jsonas.ReadJsonInts(a.nameJsonFile)
	if ErrJson != nil { // Если есть ошибка при чтении файла
		a.llog.Warn("Не получилось получить данные из файла %s: %v", a.nameJsonFile, ErrJson)
	} else { // Если нет ошибки при чтении файла
		a.llog.Info("Успешно прочитан файл %s", a.nameJsonFile)

		Sum := sum(DataJsonFile) // Рассчёт суммы чисел в слайсе
		fmt.Printf("Сумма всех чисел в массиве: %d\n", Sum)
		a.llog.Info("Сумма всех чисел в массиве: %d\n", Sum)
	}

	// Выполнить http-запрос
	httpCore := httpas.New()
	_, ErrReq := httpCore.TestRequest(a.url)
	if ErrReq != nil {
		a.llog.Err("Выполнение запроса: %v", ErrReq)
	}
	a.llog.Info("Успешно выполнен запрос по ссылке %s", a.url)

	return nil
}

// Рассчёт суммы элементов слайса
func sum(arr []int) (sum int) {
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}
