# BookStore-Management-Api-using-SQL-database
This is a Go implementation of CRUD API Which uses mysql docker container to make connection to the database and store/retrieve book details.

## How to Run the project in Local
To run the project in Local, 
- First Run the docker container for Mysql as described below.
- Then in another terminal run: `go run ./cmd/main.go`
- Use PostMan to test the APIs and response as on browser it will download the response in a file.

## Docker container for mysql
here user is rooot
password is set as my-secret-pw

 ```bash
    # To start the container:
    docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d -p 33060:3306 mysql:latest
    # To run sql commands inside mysql container:
    docker exec -it -u root some-mysql mysql -u root -p

```

To kill the remove container name

```bash
docker ps -all
docker kill some-mysql
docker rm  some-mysql
```

## Mysql commands to Show/Create Databases

```sql
SHOW DATABASES;
CREATE DATABASE BookStoreDB;
+--------------------+
| Database           |
+--------------------+
| BookStoreDB        |
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+

USE BookStoreDB;
SHOW TABLES;
+-----------------------+
| Tables_in_BookStoreDB |
+-----------------------+
| books                 |
+-----------------------+
```