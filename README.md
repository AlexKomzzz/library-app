https://github.com/AlexKomzzz/library-app.git

Задание: Спроектировать базу данных, в которой содержится авторы
книг и сами книги. Необходимо написать сервис который будет по 
автору искать книги, а по книге искать её авторов.
Требования к сервису: 
 Сервис должен принимать запрос по GRPC.
 Должна быть использована база данных MySQL
 Код сервиса должен быть хорошо откомментирован
 Код должен быть покрыт unit тестами
 В сервисе должен лежать Dockerfile, для запуска базы данных с
тестовыми данными
 Должна быть написана документация, как запустить сервис. 
Плюсом будет если в документации будут указания на команды, 
для запуска сервиса и его окружения, через Makefile 
 код должен быть выложен на github.


### Создание protoc
    make protoc
    
либо

    protoc -I=grpc_api/proto          \
            --go_out=. --go-grpc_out=.\
            grpc_api/proto/library.proto

### Запрос клиента должен состоять из:
1. ключевого слова ("authors" - если ищем авторов по книге, или "book" - если ищем книгу по авторам)
2. после ключевого слова следуют аргументы в виде названия книги (через пробелы), либо фамилии авторов

Пример запроса: authors Harry Potter
    result: Rowling



Docker MySQL
    docker run --name mysql -dp 3306:3306 -e MYSQL_ROOT_PASSWORD='qwerty' -v tom:/var/lib/mysql --network mynet mysql
    docker run --name mysql -dp 3306:3306 -e MYSQL_ROOT_PASSWORD='qwerty' -v /home/alex/GoProject/library-app/docker/volumes/tom:/var/lib/mysql mysql

    при запуске контейнера указать путь к volume 
    docker run --name mysql -dp 3306:3306 -e MYSQL_ROOT_PASSWORD='qwerty' -v $'your_path'/library-app/docker/volumes/tom:/var/lib/mysql mysql



зайти в оболочку

    docker exec -it mysql /bin/bash


    mysql -uroot -pqwerty  

создание базы данных:
    $ CREATE DATABASE library;
    $ use library;

создание таблицы:
    $ create table if not exists books 
        ( 
            id INT NOT NULL AUTO_INCREMENT, 
            title VARCHAR(255),
            PRIMARY KEY (id)
        );

    $ create table if not exists authors
        ( 
            id INT NOT NULL AUTO_INCREMENT, 
            surname VARCHAR(255),
            PRIMARY KEY (id)
        );

    $ create table if not exists authors_books
        ( 
            id_book INT, 
            id_author INT,
            FOREIGN KEY (id_book) REFERENCES books(id),
            FOREIGN KEY (id_author) REFERENCES authors(id)
        );

    $ DESC authors_books  - просмотр состава таблицы

    $ INSERT INTO books (title) VALUES ('harry potter');
    $ INSERT INTO authors (surname) VALUES ('Rowling');
    $ INSERT INTO authors_books (id_book, id_author) VALUES (1, 1);

    SELECT authors.surname FROM books JOIN authors_books ON books.id = authors_books.id_book JOIN authors ON authors_books.id_author = authors.id WHERE books.title = 'hobbit';


Содержание БД:
3 ТАБЛИЦЫ - книги, авторы и автор-книга
КНИГИ - 'harry potter', 'fantstic animals', 'hobbit'
АВТОРЫ - 'Rowling', 'Tolkin'


Запуск клиента:

    $ go run ./cmd/client/main.go author Tolkin
    $ go run ./cmd/client/main.go book hobbit


Создать docker образ:

    $ docker build -f ./dockerfile.multi -t lib .

Создать и запустить контейнер:

    $ docker run -it -dp 8080:8080 --network mynet --rm --name library lib

Запустить клиента из конетейнера:

    $ docker exec -it library /client


## Docker compose

Сборка и запуск:

    $ docker compose up --build -d