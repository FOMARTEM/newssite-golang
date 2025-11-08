# Примеры curl запросов для api
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjYyMDQzNDYsImlkIjoxM30.4tBJJCs-oQtqv2wQUp1KvYAR7Ip1YC1hiWaZrZQLzQ8" \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjYyMDgyNzcsImlkIjo3fQ.tCruIZWraEiUTe62B2Y9VDLkjvlRtjStKrnaQZndA8E" \

# Работа с пользователем

# Создание пользователя
curl -X POST  http://127.0.0.1:8081/signup \
 -H "Content-Type: application/json" \
  -d ' { "name" : "Artem", "email" : "1@gmail.com", "password" : "12345678"}'

# Авторизация
curl -X POST  http://127.0.0.1:8081/login \
 -H "Content-Type: application/json" \
 -d '{ "email" : "admin@mail.com", "password" : "123456789"}'

# Получение профиля
curl -X GET http://127.0.0.1:8081/profile \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjYyMDQzNDYsImlkIjoxM30.4tBJJCs-oQtqv2wQUp1KvYAR7Ip1YC1hiWaZrZQLzQ8"

# Обновление профиля
# Поменять можно имя и пароль, в будущем добавится повторный ввод пароля
curl -X PUT http://127.0.0.1:8081/profile \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjYyMDQzNDYsImlkIjoxM30.4tBJJCs-oQtqv2wQUp1KvYAR7Ip1YC1hiWaZrZQLzQ8" \
 -H "Content-Type: application/json" \
 -d '{ "name" : "Shahov", "email" : "1@gmail.com", "password" : "123456789"}'

# В будущем обновление профиля будет изменено, а так же будет добавлено обновление прав администратора

# Изменение прав пользователя
curl -X PUT http://127.0.0.1:8081/rules \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjYyMDgyNzcsImlkIjo3fQ.tCruIZWraEiUTe62B2Y9VDLkjvlRtjStKrnaQZndA8E" \
 -H "Content-Type: application/json" \
 -d '{ "email" : "1@gmail.com", "adminrole": 7}'

# Работа с постами
# В будущем будут добавлены права что бы только определённые пользователи могли создавать/редактировать/менять посты

# Создание поста
curl -X POST  http://127.0.0.1:8081/post \
 -H "Authorization: Bearer " \
 -H "Content-Type: application/json" \
 -d '{ "title" : "", "body" : ""}'

# Получение всех постов 
curl -X GET  http://127.0.0.1:8081/posts

# Получение поста по id 
# Вместо id вставить id поста полученного при создании поста или получении всех постов
curl -X GET http://127.0.0.1:8081/post/id \
 -H "Authorization: Bearer "

# Обновление поста по id 
# Вместо id вставить id поста полученного при создании поста или получении всех постов
curl -X PUT http://127.0.0.1:8081/post/id \
 -H "Authorization: Bearer "\
 -H "Content-Type: application/json" \
 -d '{ "title" : "", "body" : ""}'

# Удаление поста по id 
# Вместо id вставить id поста полученного при создании поста или получении всех постов
curl -X DELETE http://127.0.0.1:8081/post/id \
 -H "Authorization: Bearer "

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

