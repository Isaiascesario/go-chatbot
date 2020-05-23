package dao

import (
	"fmt"
)

// GetItem !
func GetItem(nome string) (int, error) {
	var query = "select id as name from items where name->>'pt_BR' ilike $1  || '%'  order by name->>'pt_BR' asc limit 1"
	var rows, err = db.Query(query, nome)
	if err != nil {
		return 0, fmt.Errorf("db.Query: %s", err)
	}
	defer rows.Close()
	var um int
	for rows.Next() {
		err = rows.Scan(&um)
		if err != nil {
			return 0, fmt.Errorf("rows.Next: %s", err)
		}
		break
	}
	return um, nil
}

// GetCreature !
func GetCreature(nome string) (int, error) {
	var query = "select id as name from creatures  where name->>'pt_BR' ilike $1  || '%' order by name->>'pt_BR' asc limit 1"
	var rows, err = db.Query(query, nome)
	if err != nil {
		return 0, fmt.Errorf("db.Query: %s", err)
	}
	defer rows.Close()
	var um int
	for rows.Next() {
		err = rows.Scan(&um)
		if err != nil {
			return 0, fmt.Errorf("rows.Next: %s", err)
		}
		break
	}
	return um, nil
}
