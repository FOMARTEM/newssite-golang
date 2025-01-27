# Примеры curl запросов для api

# Работа с пользователем

# Создание пользователя
curl -X POST  http://127.0.0.1:8081/signup \
 -H "Content-Type: application/json" \
  -d ' { "name" : "test", "email" : "test@mail.com", "password" : "12345678"}'

# Авторизация
curl -X POST  http://127.0.0.1:8081/login \
 -H "Content-Type: application/json" \
 -d '{ "email" : "test@mail.com", "password" : "12345678"}'

# Получение профиля
curl -X GET http://127.0.0.1:8081/profile \
 -H "Authorization: Bearer  "

# Обновление профиля
# Поменять можно имя и пароль, в будущем добавится повторный ввод пароля
curl -X PUT http://127.0.0.1:8081/profile \
 -H "Authorization: Bearer  " \
 -H "Content-Type: application/json" \
 -d '{ "name" : "test1", "email" : "", "password" : "12345678"}'

# В будущем обновление профиля будет изменено, а так же будет добавлено обновление прав администратора

# Работа с постами
# В будущем будут добавлены права что бы только определённые пользователи могли создавать/редактировать/менять посты

# Создание поста
curl -X POST  http://127.0.0.1:8081/post \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc5NzMwODcsImlkIjoxfQ.zNjNwktAg27l9l8wzEHvlGblZVQM0fE2xBcPnuPvZhg" \
 -H "Content-Type: application/json" \
 -d '{ "title" : "kdsfgdklfg", "body" : "1"}'

# Получение всех постов 
curl -X GET  http://127.0.0.1:8081/posts

# Получение поста по id 
# Вместо id вставить id поста полученного при создании поста или получении всех постов
curl -X GET http://127.0.0.1:8081/post/1 \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc5NzE4NjIsImlkIjoxfQ.2KbpylDiPGE_OlmAW9JoN-PwKjuE3hFFFw22ybLt-LM"

# Обновление поста по id 
# Вместо id вставить id поста полученного при создании поста или получении всех постов
curl -X PUT http://127.0.0.1:8081/post/1 \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc5NzE4NjIsImlkIjoxfQ.2KbpylDiPGE_OlmAW9JoN-PwKjuE3hFFFw22ybLt-LM"\
 -H "Content-Type: application/json" \
 -d '{ "title" : "teeeeest", "body" : "wertyuiopasdfghjkl;zxcvbnm,qweoxdfdfcgedvgvfedgavgjdvsvfdgbfhdsfjdfgdsgfdsfkhgdhfgdshf"}'

# Удаление поста по id 
# Вместо id вставить id поста полученного при создании поста или получении всех постов
curl -X DELETE http://127.0.0.1:8081/post/1 \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc5NzE4NjIsImlkIjoxfQ.2KbpylDiPGE_OlmAW9JoN-PwKjuE3hFFFw22ybLt-LM"

# Работа с комментариями
# Предполагается что front-end будет запрашивать отдельно посты, отдельно комментарии к нему для оптимизации

# Создание комментария
curl -X POST http://127.0.0.1:8081/comment \
  -H "Authorization: Bearer  " \
  -H "Content-Type: application/json" \
  -d '{"comment" : "", "postid" : }'

# Получение комментариев к посту
# Вместо id вставить id поста полученного при создании поста или получении всех постов
curl -X GET http://127.0.0.1:8081/comments/id \
  -H "Authorization: Bearer  "

# Удаление комментария по id
curl -X DELETE http://127.0.0.1:8081/comment/id \
  -H "Authorization: Bearer  " 

# Удаление комментариев у поста по post_id
# Вместо id вставить id поста полученного при создании поста или получении всех постов
curl -X DELETE http://127.0.0.1:8081/comments/id \
  -H "Authorization: Bearer  " 

