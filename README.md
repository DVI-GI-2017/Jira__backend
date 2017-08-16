#### Task manager [Jira]

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