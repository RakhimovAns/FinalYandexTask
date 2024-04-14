# Calculus

Финальный проект Яндекс Лицея

### В случае ошибки или при любых других вопросах пишите сюда [@Rakhimov_Ans](https://t.me/Rakhimov_Ans)

## Установка
___

#### Docker
Для начала на вашем компьютере должен быть установлен [Docker](https://docker.com). Как его установить написано в [документации](https://docs.docker.com/get-docker/).

#### Клонирование репозитория

Далее необходимо клонировать репозиторий с кодом

```bash
git clone https://github.com/RakhimovAns/FinalYandexTask.git
```

#### Build

После этого нужно создать билд с помощью Docker Compose

```bash
docker-compose build
```

P.S Этот процесс может занять довольно много времени, так, что заварите чашечку кофе ☕
___
## Запуск

После этого можно запускать проект

Для того чтобы запустить все сервисы:

```bash
docker-compose up 
```
P.S Этот процесс может занять немного времени, ждем до появления слов 

**gRPC server listening on port 50051**

___
### После запуска

#### Главная страница

После перехода на [`localhost:8080`](http://localhost:8080) открывается главная страница с полем для ввода и выражениями.

<img src="static/main.png">

___
Для начала нужно зарегистрироваться **(если не войдете в аккаунт калькулятор не будет работать)**, нажав на кнопку Register. После он перейдет на страничку регистрации.

<img src="static/register.png">

___

Далее войдите в аккаунт

<img src="static/login.png">

____

После того как зашли, можете пользоваться калькулятором.

Пример заполнения:

<img src="static/example.png">

Нажимаем на кнопку submit, и появится мониторинг выражении

<img src="static/monitoring.png">

____

## Покрытие тестами

Чтобы проверить покрытие тестами, перейдите в initializers/database.go и в ConnectToDB Изменить
```go
	dsn = "host=postgres user=postgres password=postgres dbname=yandex port=5432 sslmode=disable"
```
на 
```go
	dsn = "host=localhost user=postgres password=postgres dbname=yandex port=5432 sslmode=disable"
```

P.S. Извините, что так вышло, не нашел способа исправить проблему, если знаете, как подскажите).
## Итог
[пишите](https://t.me/Rakhimov_Ans). Не судите строго, всем удачи.
____
by Rakhimov Ansar.