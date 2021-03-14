# Redis

Цель задания – разработать приложение имплементацию in-memory Redis кеша.

## Необходимый функционал:

* Клиент и сервер tcp(telnet)/REST API
* Key-value хранилище строк, списков, словарей
* Возможность установить TTL на каждый ключ
* Реализовать операторы: GET, SET, DEL, KEYS
* Реализовать покрытие несколькими тестами функционала

## Дополнительно (необязательно):

* Реализовать операторы: HGET, HSET, LGET, LSET
* Реализовать сохранение на диск
* Масштабирование (на серверной или на клиентское стороне)
* Авторизация
* Нагрузочные тесты


## REST API server:

Для запуска сервера нужно использовать команду `docker-compose up`. При запуске будут созданы хендлеры и конструкторы кэша, которые зададут начальные значения для простоты использования.

В обычном Redis реализованы все команды отдельно для каждого типа данных, в моей реализации существуют 4 разных типа хранилищ. Отдельно для строк, списков, словарей и общий для всех этих типов.

То есть, реализованые команды LSET, LGET и т.д. добавляют данные лишь в отдельное хранилище для списков. А команды SET, GET и т.д. для общего хранилища, добавление в хранилище по ключу, в котором уже есть данные происходит через команду APPEND.

## Функционал:

API принимает данные в формате JSON и возвращает строку с результатом.
В случае успешного запроса возвращается статус-200, при ошибке-400, а при использовании другого метода-405.
На каждый ключ можно установить ttl в теле запроса, если не передать значение ttl то по умолчанию поставится значение -1.
При отправке каждого запроса нужно задавать BasicAuth и передавать логин и пароль(admin,admin).
Для сохранения нужно отправить запрост /SAVE.Сохранение производится в текстовый файл в этой же дирректории.

### Операторы:

| Оператор                | Метод | Url          | Body                                                         | Пример успешного ответа                                                                                                                    |
|-----------------------|--------|--------------|--------------------------------------------------------------|-----------------------------------------------------------------------------------------|
| GET                  | GET    | /GET/{key}            | --                                                           | Resp: &{ValueD:map[key1:value1 key2:value2] ValueL:[a b] ValueS:string ttl:-1}                                                            | --                                                               |
| SET                   | PUT    | /SET         | {"key":"key1","valueS":"string","valueL":["a","b"],"valueD":{"key1":"value1","key2":"value2"}}                                                          | Resp: OK                                  |
| DEL          | DELETE    | /DEL/{key}  | --                                                           | Resp: OK                                                                            |
| KEYS               | GET | /KEYS         | --                                                           | Resp: [key3 key5]                                                                                   | --                                                               |
| SET with ttl| PUT   | /key         | {"key":"key6","valueS":"string","ttl" : 20} |     Resp: OK       |
| APPEND          | POST | /APPEND| {"key":"key6","valueL":["b","e"]}  | Resp: &{ValueD:map[key1:value1 key2:value2] ValueL:[a b b e] ValueS:string ttl:-1}                                     |


