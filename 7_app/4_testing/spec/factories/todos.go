package factories

import (
	"context"
	"fmt"
	"learning/app/models"

	"github.com/jackc/pgx/v5/pgtype"
)

func (f *Factory) CreateTodo(context context.Context, overrides map[string]interface{}) (models.Todo, error) {
	params := models.CreateTodoParams{
		Title:       "Default Todo Title",
		Description: pgtype.Text{String: "Default Todo Description", Valid: true},
		Completed:   pgtype.Bool{Bool: false, Valid: true},
	}

	if title, ok := overrides["title"]; ok {
		params.Title = title.(string)
	}
	if description, ok := overrides["description"]; ok {
		params.Description.String = description.(string)
	}
	if completed, ok := overrides["completed"]; ok {
		params.Completed.Bool = completed.(bool)
	}

	todo, err := f.Store.CreateTodo(context, params)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (f *Factory) CreateTodos(ctx context.Context, count int, prepareParams func(map[string]interface{}, int)) ([]models.Todo, error) {
	var todos []models.Todo
	for i := 0; i < count; i++ {
		todoParams := map[string]interface{}{
			"title":       fmt.Sprintf("Todo %d", i+1),
			"description": fmt.Sprintf("Description for Todo %d", i+1),
			"completed":   false,
		}
		prepareParams(todoParams, i)
		todo, err := f.CreateTodo(ctx, todoParams)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
