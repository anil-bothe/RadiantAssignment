## Radiant Cloud Technlogy - Assignment

1. go get -u github.com/gorilla/mux
2. go get -u gorm.io/gorm
3. go get -u gorm.io/driver/mysql
<!-- we can use (gofiber) also -->

## DB instruction
1. Add Access Host -> public ip
2. Check in phpAdmin

## Authors Routing

1. GET http://127.0.0.1:8000/authors            ==> (to get all authors) 
2. POST http://127.0.0.1:8000/authors/create    ==> (pass json)
3. POST http://127.0.0.1:8000/authors/2         ==> (author details by id)
4. PUT http://127.0.0.1:8000/authors/2/edit     ==> (update author by id)
5. DELETE http://127.0.0.1:8000/authors/2/delete ==> (delete author by id)

## Books Routing

1. GET http://127.0.0.1:8000/books            ==> (to get all books)     
2. POST http://127.0.0.1:8000/books/create    ==> (pass json)
3. POST http://127.0.0.1:8000/books/2         ==> (book details by id)
4. PUT http://127.0.0.1:8000/books/2/edit     ==> (update books by id)
5. DELETE http://127.0.0.1:8000/books/2/delete ==> (delete books by id)

## Users Routing

1. GET http://127.0.0.1:8000/users            ==> (to get all users)     
2. POST http://127.0.0.1:8000/users/create    ==> (pass json)
3. POST http://127.0.0.1:8000/users/2         ==> (book details by id)
4. PUT http://127.0.0.1:8000/users/2/edit     ==> (update users by id)
5. DELETE http://127.0.0.1:8000/users/2/delete ==> (delete users by id)


## extra
 - env set JWT_SECRETE="my_secrete"
 