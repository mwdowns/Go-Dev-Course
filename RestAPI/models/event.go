package models

import (
	"encoding/json"
	"mwdowns/rest-api/db"
	"time"
)

type Event struct {
	id          int8
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
	Uuid        string
}

const eventsTableName = "events"

type result []any

var Objects = map[string]func() interface{}{
	"events": func() interface{} { return &Event{} },
}

func (e Event) Save() (string, error) {
	client, err := db.Client()
	if err != nil {
		return "", err
	}
	data, _, err := client.From(eventsTableName).
		Insert(e.inputs(), false, "", "", "exact").
		Execute()
	if err != nil {
		return "", err
	}
	var r result
	err = json.Unmarshal(data, &r)
	if err != nil {
		return "", err
	}
	return r.buildEvent(r[0].(map[string]interface{})).Uuid, err
}

func (e Event) Update() (string, error) {
	client, err := db.Client()
	if err != nil {
		return "", err
	}
	_, _, err = client.From(eventsTableName).Update(e.inputs(), "", "exact").Eq("uuid", e.Uuid).Execute()
	if err != nil {
		return "", err
	}
	return e.Uuid, err
}

func (r result) buildEvent(m map[string]interface{}) Event {
	t, _ := time.Parse(time.RFC3339Nano, m["date_time"].(string))
	e := Event{
		Name:        m["name"].(string),
		Description: m["description"].(string),
		Location:    m["location"].(string),
		DateTime:    t,
		UserID:      int(m["user_id"].(float64)),
		Uuid:        m["uuid"].(string),
	}
	return e
}

func GetEvents() ([]Event, error) {
	var events []Event
	var r result
	client, err := db.Client()
	if err != nil {
		return events, err
	}
	data, _, err := client.From(eventsTableName).
		Select("*", "exact", false).
		Execute()
	if err != nil {
		return events, err
	}
	err = json.Unmarshal(data, &r)

	for i, _ := range r {
		m := r[i].(map[string]interface{})
		e := r.buildEvent(m)
		events = append(events, e)
	}

	return events, nil
}

func GetEvent(id string) (Event, error) {
	var e Event
	client, err := db.Client()
	if err != nil {
		return e, err
	}
	data, _, err := client.From(eventsTableName).
		Select("*", "1", false).
		Eq("uuid", id).
		Execute()
	if err != nil {
		return e, err
	}
	var r result
	err = json.Unmarshal(data, &r)
	if err != nil {
		return e, err
	}
	return r.buildEvent(r[0].(map[string]interface{})), nil
}

func (e Event) inputs() map[string]interface{} {
	return map[string]interface{}{
		"name":        e.Name,
		"description": e.Description,
		"location":    e.Location,
		"date_time":   e.DateTime,
		"user_id":     e.UserID,
	}
}
