package go_simple_sql

import (
	"database/sql"
	"fmt"
)

type CONN struct {
	session  sql.DB
	err error
}

func (c *CONN) InitDB(ip, port, user, pwd, dbname, charset string) {
	url := user + ":" + pwd + "@" + "tcp(" + ip + ":" + port + ")/" + dbname + "?charset=" + charset
	db, err := sql.Open("mysql", url)
	if c.err != nil {
		fmt.Println("mysql init fail")
	} else {
		c.session = *db
		c.err = err
		fmt.Println("mysql init success")
	}
}

func (c *CONN) Query(text string) ([]map[string]string, error) {
	rows, err := c.session.Query(text)
	result := make([]map[string]string, 0)
	if err != nil {
		return result, err
	}
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		result = append(result, record)
	}
	return result, err
}

func (c *CONN) Update(text string) (int64, error) {
	tx, err := c.session.Begin()
	if err != nil {
		return 0, err
	}
	result, err := tx.Exec(text)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	rows, err := result.RowsAffected()
	return rows, err
}

func (c *CONN) Insert(text string) (int64, error) {
	tx, err := c.session.Begin()
	if err != nil {
		return 0, err
	}
	result, err := tx.Exec(text)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	id, err := result.LastInsertId()
	return id, err
}

func (c *CONN) Delete(text string) (int64, error) {
	tx, err := c.session.Begin()
	if err != nil {
		return 0, err
	}
	result, err := tx.Exec(text)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	rows, err := result.RowsAffected()
	return rows, err
}
