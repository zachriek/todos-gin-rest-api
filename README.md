# Simple Todo List RESTful API
It is a just simple RESTful API with Go using:
1. **Gin Framework**
2. **Gorm** 

## Installation & Run
```bash
# Download this project
$ go get github.com/zachriek/todos-gin-rest-api

# Download Gin Framework
$ go get github.com/gin-gonic/gin

# Download GORM
$ go get github.com/jinzhu/gorm
```

Setting DB in main.go
```go
Config.DB, err = gorm.Open("mysql", "DBUSERNAME:DBPASSWORD@tcp(127.0.0.1:3306)/DBNAME?charset=utf8&parseTime=True&loc=Local")
```

## Structure
```
├── Models
│   ├── Todos.go // Todos models
|	├── Scheme.go // Todos struct and tabel
├── Config
│   └── Database.go // Global DB
├── Controllers
│   └── Todos.go // Todos Controller
├── ApiHelpers
│   └── Response.go // response function
├── Routers
|   └── Routers.go // Routers
└── main.go
```

## API

#### /todos
* `GET` : Get all todos
* `POST` : Create a new todo

#### /todos/:id
* `GET` : Get a todo
* `PUT` : Update a todo
* `DELETE` : Delete a todo

#Post Params
```
{
	"name": "Todo 1",
	"completed": false
}
```
