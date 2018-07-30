package posts

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/rohitpavaskar/training/api/models"
)

func Get() (data []map[string]string, err error) {
	// Prepare statement for reading data
	stmtOut, err := models.Db.Prepare("SELECT * FROM posts")
	if err != nil {
		return data, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query()
	if err != nil {
		log.Fatalln(err)
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return data, err
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		var row = make(map[string]string)
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			return data, err
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			row[columns[i]] = value
			// fmt.Println(columns[i], ": ", value)
		}
		data = append(data, row)
	}
	if err = rows.Err(); err != nil {
		if err != nil {
			return data, err
		}
	}

	defer rows.Close()
	return data, nil
}

func Insert(data map[string]string) (err error) {
	// Prepare statement for inserting data
	stmtIns, err := models.Db.Prepare("INSERT INTO  poswts (title,description) VALUES( ? ,?)") // ? = placeholder
	if err != nil {
		return err
	}
	defer stmtIns.Close()                                     // Close the statement when we leave main() / the program terminates
	_, err = stmtIns.Exec(data["title"], data["description"]) // Insert tuples (i, i^2)
	fmt.Println()
	if err != nil {
		return err
	}
	return nil
}

func First(id int) (data map[string]string, err error) {
	// Prepare statement for reading data
	stmtOut, err := models.Db.Prepare("SELECT * FROM posts WHERE id = " + strconv.FormatInt(int64(id), 10) + " LIMIT 1")
	if err != nil {
		return data, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query()
	if err != nil {
		return data, err
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return data, err
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var row = make(map[string]string)
	// Fetch rows
	for rows.Next() {

		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			return data, err
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			row[columns[i]] = value
			// fmt.Println(columns[i], ": ", value)
		}
		data = row
	}

	if err = rows.Err(); err != nil {
		return data, err
	}

	defer rows.Close()
	return data, nil
}

func Update(data map[string]string, id int) error {
	// Prepare statement for inserting data
	stmtIns, err := models.Db.Prepare("UPDATE  posts SET title = ? , description = ? WHERE id = ?") // ? = placeholder
	if err != nil {
		return err
	}
	defer stmtIns.Close()                                         // Close the statement when we leave main() / the program terminates
	_, err = stmtIns.Exec(data["title"], data["description"], id) // Insert tuples (i, i^2)
	if err != nil {
		return err
	}
	return nil
}

func Delete(id int) error {
	// Prepare statement for inserting data
	stmtIns, err := models.Db.Prepare("DELETE FROM posts WHERE id = ?") // ? = placeholder
	if err != nil {
		return err
	}
	defer stmtIns.Close()     // Close the statement when we leave main() / the program terminates
	_, err = stmtIns.Exec(id) // Insert tuples (i, i^2)
	if err != nil {
		return err
	}
	return nil
}
