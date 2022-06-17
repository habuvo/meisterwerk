package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"meisterwerk/entities"
	"meisterwerk/handlers"
)

const (
	list   = `SELECT id,title,start_time,end_time,address,status FROM events WHERE 1=1`
	get    = `SELECT id,title,start_time,end_time,address,status FROM events WHERE id=$1`
	upsert = `INSERT INTO events (id,title,start_time,end_time,address,status) VALUES ($1,$2,$3,$4,$5,$6) ON CONFLICT (id)
		DO
			UPDATE 
			SET 
				title = EXCLUDED.title,
				start_time = EXCLUDED.start_time,
				end_time = EXCLUDED.end_time,
				address = EXCLUDED.address,
				status = EXCLUDED.status
	`
	delete = `DELETE FROM events WHERE id=$1`
)

var _ handlers.Storage = Repository{}

func (r Repository) Get(id string) (ev entities.Event, err error) {
	query := get

	var rows *sql.Rows
	rows, err = r.db.Query(query, id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		err = fmt.Errorf("no events found with ID %s", id)
		return
	case err != nil:
		return
	}

	err = rows.Scan(&ev.ID, &ev.Title, &ev.Start, &ev.End, &ev.Address, &ev.Status)
	if err != nil {
		return
	}

	return
}

func (r Repository) List(from, to string) (ev []entities.Event, err error) {
	query := list

	var args []interface{}

	if len(from) > 0 {
		args = append(args, from)
		query = fmt.Sprintf("%s AND start_time >= $%d", query, len(args))
	}
	if len(to) > 0 {
		args = append(args, to)
		query = fmt.Sprintf("%s AND start_time <= $%d", query, len(args))
	}

	var rows *sql.Rows
	rows, err = r.db.Query(query, args...)
	if err != nil {
		return
	}

	for rows.Next() {
		e := entities.Event{}
		err = rows.Scan(&e.ID, &e.Title, &e.Start, &e.End, &e.Address, &e.Status)
		if err != nil {
			return
		}

		ev = append(ev, e)
	}

	return
}

func (r Repository) Delete(id string) error {
	_, err := r.db.Exec(delete, id)
	return err
}

func (r Repository) Upsert(te entities.TransportEvent) error {
	_, err := r.db.Exec(upsert, te.ID, te.Title, te.Start, te.End, te.Address, te.Status)
	return err
}
