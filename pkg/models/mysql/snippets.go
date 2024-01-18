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

func (m *SnippetModel) Store(title, content, expires string) (int, error) {
	// wrtie the sql statement we want to execute
	// we are using the ? as a placeholder
	stmt := `INSERT INTO snippets (title, content,created, expires)
VALUES(?,?, UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// use the exec() method on the embedded connection pool to execute the sql statement. The first parameter is the SQL statement , followed by the title, content and expiry values for the place placeholder parameters. This method returns a sql.Result object, which contains some basic information about what happened when the statement was executed.

	result, err := m.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	// use the LastInsertId() method on the result object to get the ID of our newly inserted record in the snippets table.

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	// the ID returned has the type int64, so we convert it to an int type before returning.
	return int(id), nil
}

// return a specific resource based on id
func (m *SnippetModel) Show(id int) (*models.Snippet, error) {
	return nil, nil
}

// return a collection of resources
func (m *SnippetModel) Index() ([]*models.Snippet, error) {
	return nil, nil
}
