package Controllers

import (
	"belajar-golang/latihan-2/gorm-gin-1/ApiHelpers"
	"belajar-golang/latihan-2/gorm-gin-1/Models"

	"github.com/gin-gonic/gin"
)

func GetAllTodos(context *gin.Context) {
	var todos []Models.Todos
	err := Models.GetAllTodos(&todos)

	if err != nil {
		ApiHelpers.RespondJSON(context, 404, todos)
	} else {
		ApiHelpers.RespondJSON(context, 200, todos)
	}
}

func AddNewTodo(context *gin.Context) {
	var todo Models.Todos
	context.BindJSON(&todo)
	err := Models.AddNewTodo(&todo)

	if err != nil {
		ApiHelpers.RespondJSON(context, 404, todo)
	} else {
		ApiHelpers.RespondJSON(context, 200, todo)
	}
}

func GetTodo(context *gin.Context) {
	var todo Models.Todos
	id := context.Params.ByName("id")
	err := Models.GetTodo(&todo, id)

	if err != nil {
		ApiHelpers.RespondJSON(context, 404, todo)
	} else {
		ApiHelpers.RespondJSON(context, 200, todo)
	}
}

func UpdateTodo(context *gin.Context) {
	var todo Models.Todos
	id := context.Params.ByName("id")
	err := Models.GetTodo(&todo, id)

	if err != nil {
		ApiHelpers.RespondJSON(context, 404, todo)
	}

	context.BindJSON(&todo)
	err = Models.UpdateTodo(&todo, id)

	if err != nil {
		ApiHelpers.RespondJSON(context, 404, todo)
	} else {
		ApiHelpers.RespondJSON(context, 200, todo)
	}
}

func DeleteTodo(context *gin.Context) {
	var todo Models.Todos
	id := context.Params.ByName("id")
	err := Models.DeleteTodo(&todo, id)

	if err != nil {
		ApiHelpers.RespondJSON(context, 404, todo)
	} else {
		ApiHelpers.RespondJSON(context, 200, todo)
	}
}
