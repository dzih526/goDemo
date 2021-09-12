package errorHandlling

import (
	"database/sql"
	"fmt"
)
import "github.com/pkg/errors"

var (
	db sql.DB
)

func query(sentence string) (string, error)  {
	//	...
	//	handle sql.ErrNoRows
	var result int
	err := db.QueryRow(sentence).Scan(&result)
	if err == sql.ErrNoRows{
		return "",errors.Wrap(err,"No Rows for first time scan")
	}
	// ...
}

func main() {
	var sentense string
	result, err := query(sentense)
	if err!=nil {
		fmt.Println("Original error: %T %v+\n",errors.Cause(err),errors.Cause(err))
	}
}