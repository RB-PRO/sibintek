package httpas

import "net/http"

// Ядро HTTP-запросника
type Core struct {
	*http.Client
}

// Создать объект HTTP запросника
func New() *Core {
	client := &http.Client{}
	return &Core{client}
}
