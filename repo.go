package main

import "fmt"

var currentBeerId int
var currentTodoId int

var todos Todos
var beers Beers

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
	RepoCreateBeer(Beer{Name: "Stella", Abv:"4.5%", DidWeDrinkIt:true})
	RepoCreateBeer(Beer{Name: "Amstel Light"})
}

func RepoFindBeer(id int) Beer {
	for _, t := range beers {
		if t.Id == id {
			return t
		}
	}
	// return empty Beer if not found
	return Beer{}
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

func RepoCreateBeer(t Beer) Beer {
	currentBeerId += 1
	t.Id = currentBeerId
	beers = append(beers, t)
	return t
}


//this is bad, I don't think it passes race condtions
func RepoCreateTodo(t Todo) Todo {
	currentTodoId += 1
	t.Id = currentTodoId
	todos = append(todos, t)
	return t
}


func RepoDestroyBeer(id int) error {
	for i, t := range beers {
		if t.Id == id {
			beers = append(beers[:i], beers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Beer with id of %d to delete", id)
}


func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
