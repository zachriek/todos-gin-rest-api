package Models

import "belajar-golang/latihan-2/gorm-gin-1/Config"

func GetAllTodos(todo *[]Todos) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func AddNewTodo(todo *Todos) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodo(todo *Todos, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *Todos, id string) (err error) {
	Config.DB.Save(todo)
	return nil
}

func DeleteTodo(todo *Todos, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
