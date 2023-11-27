package app

import (
	"fmt"

	"github.com/RB-PRO/sibintek/internal/app/httpas"
	"github.com/RB-PRO/sibintek/internal/app/jsonas"
)

// Запуск самого приложения
func (a *App) Run() error {

	a.llog.Info("Начало работы приложения")

	// Анонимная функция сохранения результатов по расчётам файла с массивом
	funcSumIntDone := func(arr []int) {
		a.llog.Info("Успешно прочитаны данные: %s", a.flages.dataInput)
		Sum := sum(arr) // Рассчёт суммы чисел в слайсе
		SaveStr := fmt.Sprintf("Сумма всех чисел в массиве: %d\n", Sum)
		fmt.Print(SaveStr)
		a.llog.Info(SaveStr)
		jsonas.Save(a.flages.outputFile, SaveStr)
		a.llog.Info("Успешно сохранили данные в файл: %s", a.flages.outputFile)
	}

	// Прочитать файл в слайс
	if a.flages.inputMode {
		DataJson, Err := jsonas.UnmarshalInts(a.flages.dataInput)
		if Err != nil {
			a.llog.Warn("Не получилось распарсить данные %s: %v", a.flages.dataInput, Err)
		} else {
			funcSumIntDone(DataJson)
		}
	} else {
		DataJsonFile, ErrJson := jsonas.ReadJsonInts(a.flages.dataInput)
		if ErrJson != nil { // Если есть ошибка при чтении файла
			a.llog.Warn("Не получилось получить данные из файла %s: %v", a.flages.dataInput, ErrJson)
		} else { // Если нет ошибки при чтении файла
			funcSumIntDone(DataJsonFile)
		}
	}

	// Выполнить http-запрос
	httpCore := httpas.New()
	_, ErrReq := httpCore.TestRequest(a.cf.URL)
	if ErrReq != nil {
		a.llog.Err("Выполнение запроса: %v", ErrReq)
	}
	a.llog.Info("Успешно выполнен запрос по ссылке %s", a.cf.URL)

	return nil
}

// Рассчёт суммы элементов слайса
func sum(arr []int) (sum int) {
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}
