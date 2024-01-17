// basically this file is acting like a laravel file i will name a snippet service or probally snippet controller

package mysql

import (
	"database/sql"
	"kod.net/snippetbox/pkg/models"
)

// define a snippetModel type which wraps a sql.DB connection pool

type SnippetModel struct {
	DB *sql.DB
}

// this will insert a new snippet into the database

func (m *SnippetModel) store(title, content, expires string) (int, error) {
	return 0, nil
}

// return a specific resource based on id
func (m *SnippetModel) show(id int) (*models.Snippet, error) {
	return nil, nil
}

// return a collection of resources
func (m *SnippetModel) index() ([]*models.Snippet, error) {
	return nil, nil
}
