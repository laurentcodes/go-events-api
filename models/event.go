package models

import (
	"time"

	"example.com/events-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
    INSERT INTO events (name, description, location, date_time, user_id) 
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id
  `

	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&e.ID)

	return err
}

func (e Event) Update() error {
	query := `
    UPDATE events
    SET name = $1, description = $2, location = $3, date_time = $4
    WHERE id = $5
  `

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func GetEvent(id int64) (*Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events WHERE id = $1`

	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func Delete(event *Event) error {
	query := `DELETE FROM events WHERE id = $1`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&event.ID)

	return err
}

func (e Event) Register(user_id int64) error {
	query := `
    INSERT INTO registrations (event_id, user_id)
		VALUES ($1, $2)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, user_id)

	return err
}

func (e Event) CancelRegistration(user_id int64) error {
	query := `
		DELETE FROM registrations 
		WHERE event_id = $1 AND user_id = $2
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, user_id)

	return err
}
