// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// Connected clients list
type ConnectedClients struct {
	Usename Name           // usename
	Client  sql.NullString // client
	State   sql.NullString // state
	Count   sql.NullInt64  // count
}

// GetConnectedClients runs a custom query, returning results as ConnectedClients.
func GetConnectedClients(db XODB) ([]*ConnectedClients, error) {
	var err error

	// sql query
	var sqlstr = `SELECT usename, ` +
		`CASE WHEN client_hostname IS NULL THEN client_addr::text ELSE client_hostname END AS client, ` +
		`state, count(*) ` +
		`FROM pg_stat_activity ` +
		`WHERE state IS NOT NULL ` +
		`GROUP BY 1,2,3 ` +
		`ORDER BY 4 desc,3`

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ConnectedClients{}
	for q.Next() {
		cc := ConnectedClients{}

		// scan
		err = q.Scan(&cc.Usename, &cc.Client, &cc.State, &cc.Count)
		if err != nil {
			return nil, err
		}

		res = append(res, &cc)
	}

	return res, nil
}