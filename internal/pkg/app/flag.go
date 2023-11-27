package app

import (
	"flag"
)

// Структура конфигурационных данных, которая должна быть олучена из "флагов"
// аргумента запуска программы
type Flags struct {
	// Переключатель определения в ходного типа данных
	//	inputMode=false - из файла
	//	inputMode=true - из стандартного ввода
	inputMode bool

	// Входные данные, в зависимости от inputModе:
	//	inputMode=false - Название файла
	//	inputMode=true - Входная строка
	dataInput string

	// Название выходного файла, куда будут сохраняться результаты
	outputFile string

	// Название конфигурационного файла, где лежат данные о ссылке и тд
	configFile string

	// папка для логов
	logsFile string
}

// Обработка флагов путём обработки аргументов командной строки
//
//	--file [файл] - название файла с массивом цифр
//
// или
//
//	--stdin [Данные] -  стандартный ввод
//
//	--output [файл] - файл для вывода данных
//
//	--config [файл] - файл конфига
func flugger() (f Flags) {
	flag.StringVar(&f.outputFile, "output", "", "название выходного файла")
	flag.StringVar(&f.logsFile, "log", "", "название папки с логами")
	flag.StringVar(&f.configFile, "config", "", "название файла с конфигом самого приложения")
	flag.StringVar(&f.dataInput, "file", "", "название файла с массивом цифр")
	stdin := flag.String("stdin", "", "значение для стандартноо ввода stdin")
	flag.Parse()
	if *stdin != "" {
		f.inputMode = true
		f.dataInput = *stdin
	}
	return f
}
