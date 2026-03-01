package models

import (
	"mwdowns/rest-api/DB"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	UserID      int64     `json:"user_id"`
	Uuid        uuid.UUID `json:"uuid"`
}

const saveEventQuery = "INSERT INTO events(name, description, location, dateTime, uuid, user_id) VALUES (?, ?, ?, ?, ?, ?)"
const updateEventQuery = "UPDATE events SET name = ?, description = ?, location = ?, dateTime = ? WHERE uuid = ?"
const deleteEventQuery = "DELETE FROM events WHERE uuid = ?"
const getAllEventsQuery = "SELECT name, description, location, dateTime, uuid, user_id FROM events"
const getEventQuery = "SELECT name, description, location, dateTime, uuid, user_id FROM events WHERE uuid = ?"

func (e Event) Save() (string, error) {
	eventUuid := uuid.New()
	e.Uuid = eventUuid
	stmt, err := db.DB.Prepare(saveEventQuery)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Uuid, e.UserID)
	if err != nil {
		return "", err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return "", err
	}

	return eventUuid.String(), nil
}

func (e Event) Update() (string, error) {
	stmt, err := db.DB.Prepare(updateEventQuery)

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Uuid)
	if err != nil {
		return "", err
	}
	return e.Uuid.String(), err
}

func (e Event) Delete() error {
	stmt, err := db.DB.Prepare(deleteEventQuery)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Uuid)
	return err
}

func GetEvents() ([]Event, error) {
	rows, err := db.DB.Query(getAllEventsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Name, &event.Description, &event.Location, &event.DateTime, &event.Uuid, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEvent(id uuid.UUID) (*Event, error) {
	row := db.DB.QueryRow(getEventQuery, id)
	var event Event
	err := row.Scan(&event.Name, &event.Description, &event.Location, &event.DateTime, &event.Uuid, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
