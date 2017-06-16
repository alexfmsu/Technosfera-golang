# Technosfera-golang
### Задание 1
Реализовать функции, которые нужны для прохождения тестов в main_test.go

[task1](/task1/)
##
### Задание 2
##### Функции

1. Реализовать функцию, возвращающую тип значения<br/>
   (использовать fmt, как в тестах запрещено)
2. Реализовать [извлечение корня методом Ньютона](https://ru.wikipedia.org/wiki/Алгоритм_нахождения_корня_n-ной_степени)<br/>
   (использовать пакет math запрещено)
3. Реализовать функцию [мемоизации](https://ru.wikipedia.org/wiki/Мемоизация)

[task2/funcs](/task2/funcs/)<br/><br/>
##### ООП

1. Реализовать объект календаря

[task2/calendar](/task2/calendar/)<br/><br/>
##### Творческое задание

1. Реализовать простую игру (описание в папке с игрой)

[task2/game](/task2/game-0/)

Реализация: [t.me/open_door_bot](http://t.me/open_door_bot)
##
### Задание 3

##### Балансировщик нагрузки

Есть некоторое количество серверов, нагрузка на которых распределяется методом "по кругу".<br/>
Есть балансер, который должен распределить запросы равномерно.

Суть в том, чтобы на каждый сервер попало равномерное количество нагрузки.

У балансера есть несколько функций:

Init - инициализирует собственно балансер - представьте что устанавливает соединения с указанным колчиеством серверов.<br/>
GiveStat - даёт статистику, сколько запросов пришло на каждый из серверов.<br/>
GiveNode - эта функция фвзывается, когда пришел запрос. мы получаем номер сервера, на который идти.

[task3/balancer](/task3/0_balancer/)<br/><br/>
##### Pipeline

Требуется реализовать функцию последовательно выполняющую все переданные операции 

Pipe(funcs ...job) 

Где 
type job func(in, out chan interface{})<br/><br/>

[task3/pipeline](/task3/1_pipeline/)<br/><br/>

##### Творческое задание

Дальнейшее развитие game-0
Добавляем взаимодействие нескольких игроков
Раньше был 1 игрок и игра была по сути синхронна
теперь добавляется второй ( и третий, четвёртый ) и они оба могут "вводить" команды
значит у нас появляется асинхронное взаимодействие
надо использовать каналы и горутины для обработки команд игроков в каждой комнате
Смотрите примеры с мультипрексорами для организации ввода

Большая часть команд так же относится к одному игроку, но появляется несколько новых команд:
"сказать" - транслируется всем другим игрокам в этой комнате
"сказать_игроку" - отправляет персональное сообщение другому игроку, остальные его не слышат
Если выполнить команду "осмотреться", что можно увидеть другого игрока
Это значит, что у комнаты появляется список игроков в ней, которым надо передать сообщение

Чтобы как-то отличать игроков друг от друга, теперь при инициализации надо задать имя ( функция NewPlayer )
Естественно, количество игроков не должно быть захардкожено и может меняться

У игрока появляется метод HandleInput, который принимает от него входящее сообщение. 
А так же метод HandleOutput, в который мы пишем то что игро должен видеть у себя на экране.
И метод GetOutput, который возвращает канал, принимающий сообщения

[task3/game](/task3/game-1/)

Реализация: [t.me/open_door_bot](http://t.me/open_door_bot)
##
### Задание 4

##### Crawler
Реализовать sitemap builder

Функция Crawl(url string) []string
* На вход передается web адрес корня сайта (/)
* на выход список всех страниц, на которые есть ссылки на всех страницах этого домена.

Нужно обрабатывать статусы редирект (301), и учитывать цикилческие ссылки

В данном тесте мы используем пакет net/http/httptest для эмуляции внешнего сервера

Не найденные урлы ( 404 ) учитывать не надо. Внешние урлы (хосты, отличные от текущего) к результату не относятся. Результат отсортирован в порядке появления на сайте.
в задании crawler добавил чуть описания:

Неизвестные урлы учитывать не надо. 
Внешние урлы (хосты, отличные от текущего) к результату не относятся. Результат отсортирован в порядке появления на сайте.

[task4/crawler](/task4/crawler/)<br/><br/>


##### Game

[task4/game](/task4/game-2/)

Реализация: [t.me/open_door_bot](http://t.me/open_door_bot)
##
### Задание 6
Задание №6 - написание тестов
Дана функция doSearch, которая отправляет запрос в поисковый движок и возвращает результат, немного преобразуя его.

```
func SearchServer(w http.ResponseWriter, r *http.Request) {}
```

Надо:
1. Реализовать простой SearchServer, который отвечает данными с нужными параметрами. Его использовать в httptest.NewServer. Это писать в `main_test.go` (чтобы для него не считалось покрытие). Посмотрите пример использования в `4/99_homework/crawler`
1. Протестить корректность работы всех входящих параметров функции `doSearch` через тесты, написанные в `main_test.go`. Корректность означает мы проверили тестом, что функция вернула столько данных, сколько надо, в том порядке, в котором мы запросили.
2. Обеспечить максимальное покрытие тестами для функции `doSearch`, включая случаи ошибок
3. Построить html-отчет о покрытии. То что не покрыто - объяснить почему покрыть невозможно.
4. Найти баги в `doSearch` :)

Дополнительно:
* данные для работы лежаит в файле `dataset.xml`
* вы может не ограничиваться функцией `SearchServer` если вам надо проверить какой-то отдельны кейс. Их может быть много. B httptest.NewServer тоже может быть много.
* параметр `query` ищет по полям `Name` и `About`
* параметр `order_field` работает по полям `Id`, `Age`, `Name`
* Если `query` пустой, то делаем только сортироку, т.е. возвращаем все записи
* Если `order_field` пустой, то возврщаем по `Name`, в порядке возрастания

Дополнительное задание (без доп баллов оценки):
* `SearchServer` положить в `main.go`, обеспечить максимально возможное покрытие тестами (через `doSearch`), построить html-отчет. ТО что покрыть не получается - объяснить почему. В данном случае `SearchServer` будет представлен в еинственное экземпляре, но это не мешает делать другие функции для теста `doSearch` или менять его сигнатуру.

[task6](/task6/)
##
### Задание 7
Простая orm `myorm`, реализованная через кодогенерацию.

Надо написать программу `myorm/godegen`, которая на вход получает имя файла, разбирает его, находит там структуры, у которых есть `//myorm:$tablename`, где `$tablename` - имя таблицы, с которой будем работать.

Работать надо через AST. Регулярками или иным ручным методом парсить файл нельзя.

После этого для данной структуры создаёт методы
* `FindByPK` - находит запись по первичному ключу
* `Create` - создаёт новую запись, LastInsertId кладёт в поле, помеченное как `myorm:"primary_key"`
* `Update` - обновляет запись, используя в качестве ключа поле с `myorm:"primary_key"`

Результат должен быть записан в файл `{$tablename}_myorm.go` в папке, в которой находится структура

Так же на пакет дополнительно генерируется метод `SetDB`, который принимает объект базы данных, с которой будет вестись работа, и устанавливает её в переменную уровня пакета.

В качестве примера - смотрите `main.go` + `user/user.go`. В репу надо закоммитить кодогенератор + сгенерённый код в user, чтобы можно было запустить программу и она отработала.

Кодогенератор должен быть прокрыт каким-нибудь тестами, можно через моки. Так же надо сгенерировать отчет по покрытию.

[task7/myorm](/task7/myorm)
##
### Задание 9

Интеграция с cgo:

* написать через cgo биндинги к golang для любому сишному коду/библиотеки
* написать тесты на го к этому коду - чтобы убедиться что всё работает. можно переписать тест-кейсы из сишного кода, если они там были
* убедиться что нет утечек памяти - погонять через bench
* закоммить в репу код, необходимый для компилции - полный код это это какая-то редкая либа, инструкции по установке если это публичная либа

По выбору кода для интеграции:
* допускается покрытие не всей библиотеки, если она большая, но не менее 5-и функций, которые могут быть с пользой исопльзованы. т.е. не min(a, b)
* если это внешняя (публичная) библиотека - у неё ещё не должно быть биндингов к golang
* принимается код, который вы писали в рамках курсовых или дипломных
* хорошо подойдут какие-то алгоритмы для вычисления чего-то: хеш по строке, сжатие данных, или шифрование

[task9](/task9/)
