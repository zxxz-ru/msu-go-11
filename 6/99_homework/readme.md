Задание №6 - написание тестов
Дана функция doSearch, которая отправляет запрос в поисковый движок и возвращает результат, немного преобразуя его.

```
func SearchServer(w http.ResponseWriter, r *http.Request) {}
```

Надо:
1. Реализовать простой SearchServer, который отвечает данными с нужными параметрами. Его использовать в httptest.NewServer. Это писать в `main_test.go` (чтобы для него не считалось покрытие). Посмотрите пример использования в `4/99_homework/crawler`
1. Протестить корректность работы всех входящих параметров функции `doSearch`
2. Обеспечить максимальное покрытие тестами для функции `doSearch`, включая случаи ошибок
3. Построить html-отчет о покрытии. То что не покрыто - объяснить почему покрыть невозможно.
4. Найти баги в `doSearch` :)

Дополнительно:
* данные для работы лежаит в файле `dataset.xml`
* вы может не ограничиваться функцией `SearchServer` для тестирования если вам надо проверить какой-то отдельны кейс. Их может быть много. B httptest.NewServer тоже может быть много.
* параметр `query` ищет по полям `Name` и `About`
* параметр `order` работает по полям `Id`, `Age`, `Name`

Дополнительное задание (без доп баллов оценки):
* `SearchServer` положить в `main.go`, обеспечить максимально возможное покрытие тестами (через `doSearch`), построить html-отчет. ТО что покрыть не получается - объяснить почему. В данном случае `SearchServer` будет представлен в еинственное экземпляре, но это не мешает делать другие функции для теста `doSearch` или менять его сигнатуру.