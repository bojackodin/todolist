package data

import "database/sql"

type Models struct {
	Tasks TaskModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Tasks: TaskModel{DB: db},
	}
}
