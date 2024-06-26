# Containers
### Описание задачи

Есть n контейнеров и какое-то количество шаров разного цвета, количество разных цветов тоже n. Шары в случайном порядке распределены по контейнерам. Допустима лишь одна операция с шарами: два шара в разных контейнерах можно поменять местами. Нужно ответить на вопрос: можно ли используя эту единственную операцию отсортировать шары так, чтобы выполнялись такие условия:
    • в каждом контейнере лежат лишь шары одинакового цвета
    • каждый цвет лежит лишь в одном контейнере

### Решение задачи

При операции замены шаров количество шаров в контейнере не меняется, значит, для того чтобы заполнить весь контейнер шарами одного цвета, во всех контейнерах должно быть столько же шаров сколько помещается в одном контейнере, поэтому достаточно проверить, что для каждого контейнера хватает шаров для его заполнения.

### Запуск

Программу можно запустить с помощью Golnag, а также с помощью Docker контейнера.

Для запуска программы с помощью Golang и интерактивным вводом данных через stdin:

```
go run main.go
```

Для запуска программы с интерактивным вводом данных через stdin в Docker контейнере:

```
make start
```

Для запуска заранее написанных тестов в Docker контейнере:

```
make tests
```

### Ограничения

1<= n <= 100 - количество контейнеров и цветов 0 <= количество шаров одного цвета в одном контейнере <= 1000000000

### Формат входа
Первая строчка содержит n. Остальные n строчек содержат n чисел, разделённых пробелами. Каждая строчка описывает содержимое контейнера. Каждое из чисел обозначает количество шаров i-го цвета в этом контейнере, где i - порядковый номер числа в строчке

### Выход

Написать "yes", если сортировку можно сделать, или "no", если нельзя.
