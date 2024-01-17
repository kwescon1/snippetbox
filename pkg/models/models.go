// defining the top level data types that the database model will use and return
// this file is acting like a file in laravel i will name snippet model

package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
