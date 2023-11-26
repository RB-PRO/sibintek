package httpas

import (
	"fmt"
	"net/http"
)

// Функция выполнения тестового [запроса]
//
// # Укажите в качестве аргумента ссылку, по которой необходимо выполнить запрос
//
// [запроса]: https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/27185784/detail?languageId=-1&appId=1
func (core *Core) TestRequest(url string) (bool, error) {

	// Создаёми запрос
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return false, ErrNewRequest
	}
	req.Header.Add("authority", "www.massimodutti.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("referer", "https://www.massimodutti.com/")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36")

	// Выполняем запрос
	res, ErrDo := core.Do(req)
	if ErrDo != nil {
		return false, fmt.Errorf("Do: %v", ErrDo)
	}
	defer res.Body.Close()

	// Проверка статуса ответа и вывод соответствующего результата
	if res.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, fmt.Errorf("status code is %v", res.StatusCode)
}
