#### Task manager [Jira]

[![Build Status](https://travis-ci.org/DVI-GI-2017/Jira__backend.svg?branch=develop)](https://travis-ci.org/DVI-GI-2017/Jira__backend)
#### Технологии
- Golang для написания логики,

#### Разработка и запуск
- `go run main.go`

#### Зависимости
Интерфейс к MongoDB: `mgo.v2` 

JSON Web Tokens : `dgrijalva/jwt-go` 

```bash
$ go get github.com/dgrijalva/jwt-go
$ go get gopkg.in/mgo.v2
```

#### Генерация ключей
```bash
$ mkdir rsa && cd rsa
$ openssl genrsa -out app.rsa 1024
$ openssl rsa -in app.rsa -pubout > app.rsa.pub
```

#### Быстрая проверка (через консоль)
```bash
for i in {1..15}; 
do echo '{"email": "test", "password": "password"}' | curl -d @- http://localhost:3000/api/v1/test; 
done
```
