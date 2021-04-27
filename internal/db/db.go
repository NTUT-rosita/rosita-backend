// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.countStmt, err = db.PrepareContext(ctx, count); err != nil {
		return nil, fmt.Errorf("error preparing query Count: %w", err)
	}
	if q.getDistinctAcademicYearStmt, err = db.PrepareContext(ctx, getDistinctAcademicYear); err != nil {
		return nil, fmt.Errorf("error preparing query GetDistinctAcademicYear: %w", err)
	}
	if q.getDistinctCollegeTypeStmt, err = db.PrepareContext(ctx, getDistinctCollegeType); err != nil {
		return nil, fmt.Errorf("error preparing query GetDistinctCollegeType: %w", err)
	}
	if q.getDistinctDepartmentNameStmt, err = db.PrepareContext(ctx, getDistinctDepartmentName); err != nil {
		return nil, fmt.Errorf("error preparing query GetDistinctDepartmentName: %w", err)
	}
	if q.getDistinctSchoolSystemStmt, err = db.PrepareContext(ctx, getDistinctSchoolSystem); err != nil {
		return nil, fmt.Errorf("error preparing query GetDistinctSchoolSystem: %w", err)
	}
	if q.getDistinctSemesterStmt, err = db.PrepareContext(ctx, getDistinctSemester); err != nil {
		return nil, fmt.Errorf("error preparing query GetDistinctSemester: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.countStmt != nil {
		if cerr := q.countStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countStmt: %w", cerr)
		}
	}
	if q.getDistinctAcademicYearStmt != nil {
		if cerr := q.getDistinctAcademicYearStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDistinctAcademicYearStmt: %w", cerr)
		}
	}
	if q.getDistinctCollegeTypeStmt != nil {
		if cerr := q.getDistinctCollegeTypeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDistinctCollegeTypeStmt: %w", cerr)
		}
	}
	if q.getDistinctDepartmentNameStmt != nil {
		if cerr := q.getDistinctDepartmentNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDistinctDepartmentNameStmt: %w", cerr)
		}
	}
	if q.getDistinctSchoolSystemStmt != nil {
		if cerr := q.getDistinctSchoolSystemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDistinctSchoolSystemStmt: %w", cerr)
		}
	}
	if q.getDistinctSemesterStmt != nil {
		if cerr := q.getDistinctSemesterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDistinctSemesterStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                            DBTX
	tx                            *sql.Tx
	countStmt                     *sql.Stmt
	getDistinctAcademicYearStmt   *sql.Stmt
	getDistinctCollegeTypeStmt    *sql.Stmt
	getDistinctDepartmentNameStmt *sql.Stmt
	getDistinctSchoolSystemStmt   *sql.Stmt
	getDistinctSemesterStmt       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                            tx,
		tx:                            tx,
		countStmt:                     q.countStmt,
		getDistinctAcademicYearStmt:   q.getDistinctAcademicYearStmt,
		getDistinctCollegeTypeStmt:    q.getDistinctCollegeTypeStmt,
		getDistinctDepartmentNameStmt: q.getDistinctDepartmentNameStmt,
		getDistinctSchoolSystemStmt:   q.getDistinctSchoolSystemStmt,
		getDistinctSemesterStmt:       q.getDistinctSemesterStmt,
	}
}
