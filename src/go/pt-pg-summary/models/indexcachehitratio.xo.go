// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// Table cache hit ratio
type IndexCacheHitRatio struct {
	Name  string          // name
	Ratio sql.NullFloat64 // ratio
}

// GetIndexCacheHitRatio runs a custom query, returning results as IndexCacheHitRatio.
func GetIndexCacheHitRatio(db XODB) (*IndexCacheHitRatio, error) {
	var err error

	// sql query
	const sqlstr = `SELECT 'index hit rate' AS name, ` +
		`CASE WHEN sum(idx_blks_hit) IS NULL ` +
		`THEN 0 ` +
		`ELSE (sum(idx_blks_hit)) / sum(idx_blks_hit + idx_blks_read) ` +
		`END AS ratio ` +
		`FROM pg_statio_user_indexes ` +
		`WHERE (idx_blks_hit + idx_blks_read) > 0`

	// run query
	XOLog(sqlstr)
	var ichr IndexCacheHitRatio
	err = db.QueryRow(sqlstr).Scan(&ichr.Name, &ichr.Ratio)
	if err != nil {
		return nil, err
	}

	return &ichr, nil
}
