<h1>Тестове завдання для Global Logic</h1>

# Сутність завдання
Завдання можна переглянути [тут](TASK.md).


### Швидкий огляд
Цей сервер було згенеровано проєктом [swagger-codegen](https://github.com/swagger-api/swagger-codegen). Swagger-документацію можна подивитися [тут](https://github.com/AndriiPopovych/gses/blob/main/gses2swagger.yaml).
Використовуючи [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) з віддаленого сервера можна легко створити заглушку сервера.

- Версія API: 1.0.0
- Версія збірки: 2022-07-24T15:47:24.214Z


### Запуск сервера
Для запуску сервера слідуйте цим простим настановам:

- Перейменуйте `.env.example` на `.env` та встановіть туди свої змінні.
- Пропишіть в консолі команду ```go run main.go```.
- Для швидкої перевірки також можна використати Postman замість редактора Swagger.


### Запуск сервера з Docker

- Установіть Docker;
- Пропишіть в консолі команду `docker build -t yourProjectName .`;
- Перевірте наявність Docker's image шляхом команди `docker images`;
- Запустіть Docker за допомогою команди `docker `


### Підключені бібліотеки
  - `github.com/go-mail/mail` - для надсилання електронних листів через SMTP сервер;
  - `github.com/gorilla/mux`  - для створення HTTP запитів;
  - `github.com/joho/godotenv`- для простого отримання доступу до `.env` файлу, де зберігається ключі доступу до API тощо.