package dao

import "fmt"

func getItems(nome String) (int, error) {
	var query = "select id as name from items  where name->>'pt_BR' ilike %$1% limit 1"
	var rows, err = db.Query(query, nome)
	if err != nil {
		return fundamental.Info{}, fmt.Errorf("db.Query: %s", err)
	}
	defer rows.Close()
	var um int
	for rows.Next() {
		err = rows.Scan(&um)
		if err != nil {
			return fundamental.Info{}, fmt.Errorf("rows.Next: %s", err)
		}
		break
	}
	return um, nil

}
func getCreatures(nome String) (int, error) {
	var query = "select id as name from creatures  where name->>'pt_BR' ilike %$1% limit 1"
	var rows, err = db.Query(query, nome)
	if err != nil {
		return fundamental.Info{}, fmt.Errorf("db.Query: %s", err)
	}
	defer rows.Close()
	var um int
	for rows.Next() {
		err = rows.Scan(&um)
		if err != nil {
			return fundamental.Info{}, fmt.Errorf("rows.Next: %s", err)
		}
		break
	}
	return um, nil
}
