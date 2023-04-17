# О проекте
___
Микросервис для работы с балансом пользователей. 
Позволяет производить:
* Зачисление средств
* Резервирование, списание, возврат средств при совершении заказа
* Перевод средств от пользователя к пользователю
* Получение баланса пользователя
* Получение месячного отчета с группировкой по разным продуктам 
* Детальный отчет о движении средств на балансе отдельного юзера с функциями сортировки по 
сумме и дате

Сервис предоставляет HTTP API и принимает/отдает запросы/ответы в формате JSON.

## Что под капотом?

* Роутер - [Gin Web Framework](https://github.com/gin-gonic/gin) 
* Config - [Viper](https://github.com/spf13/viper)
* БД - Postgres
* Логирование - [Logrus](https://github.com/sirupsen/logrus)

## Запуск

Проект запускается с помощью Docker-compose. 
* **Первый запуск:**

  * Собираем образы и запускаем контейнеры:
    ```
    make init
    ```
  * Создаем таблицы, для этого необходимо узнать id контейнера:
    ```
    docker ps
    ```
  * Подключаемся к контейнеру:
    ```
    docker exec -it вставьте_айди bin/bash
    ```
  * Подключаемся к postgres. В проекте используется имя по умолчанию:
    ```
    psql -U postgres
    ```
  * Копируем содержимое файла DB.sql и выполняем в терминале.
  * Выходим из контейнера.
* **Завершение работы:**
  ```
  make stop
  ```
* **Возообновление работы:**
  ```
  make start
  ```
* **Сброс БД:**
  ```
  make reset
  ```

## Запросы через Postman
* Зачисление средств (создаст пользователя если его не было):
  ```
  POST
  ```
  ```
  http://localhost:8000/balance/
  ```
  ```json
  {
    "id":1,
    "amount":100.0
  }
  ```
* Резервирование средств на заказ:
  ```
  POST
  ```
  ```
  http://localhost:8000/reserve/
  ```
  ```json
  {
    "order_id":1,
    "user_id":1,
    "product_id":1,
    "price":10.0
  }
  ```
* Списание зарезервированных средств в пользу компании, то есть заказ пришел:
  ```
  POST
  ```
  ```
  http://localhost:8000/reserve/revenue-confirm
  ```
  ```json
  {
    "order_id":1,
    "user_id":1,
    "product_id":1,
    "price":10.0
  }
  ```
* Отмена заказа:
  ```
  POST
  ```
  ```
  http://localhost:8000/reserve/revenue-deny
  ```
  ```json
  {
    "order_id":1,
    "user_id":1,
    "product_id":1,
    "price":10.0
  }
  ```
* Перевод средств от пользователя к пользователю:
  ```
  POST
  ```
  ```
  http://localhost:8000/balance/move-funds
  ```
  ``` json
  {
    "sender_id":2,
    "recipient_id":1,
    "amount":15.0
  }
  ```
* Получение баланса пользователя:
  ```
  GET
  ```
  ```
  http://localhost:8000/balance/1
  ```
* Получение месячного отчета с группировкой по разным продуктам (возвращает имя созданного файла с расширением .csv):
  ```
  POST
  ```
  ```
  http://localhost:8000/report/
  ```
  ```json
  {
  "month": "April"
  }
  ```
* Детальный отчет о движении средств на балансе отдельного юзера с функциями сортировки по
  сумме или дате:
  ```
  POST
  ```
  ```
  http://localhost:8000/operationsJournal/date / http://localhost:8000/operationsJournal/amount
  ```
  ```json
  {
    "id":1,
    "amount":0.0
  }
  ```
## Contact

Ivan Konoplich - konoplich_i@mail.ru

Project Link: [https://github.com/IvanKonoplich/Wallet-Service.git](https://github.com/IvanKonoplich/Wallet-Service.git)

