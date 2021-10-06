package models

import (
	"echo-rest/db"
	"net/http"
)

type Event struct {
	IdEvent          int    `json : "idevent"`
	EventName        string `json : "eventname"`
	EventSchedule    string `json : "eventschedule"`
	EventDescription string `json : "eventdescription"`
	Image            string `json : "image"`
}

func TakeEvent() (Response, error) {
	var obj Event
	var arrobj []Event
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM event"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdEvent, &obj.EventName, &obj.EventSchedule, &obj.EventDescription, &obj.Image)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "succes"
	res.Data = arrobj

	return res, nil
}

func AllEvent(eventname string, eventschedule string, eventdescription string, image string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT event (eventname, eventschedule, eventdescription, image) VALUES (?,?,?,?,?)"
	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := statement.Exec(eventname, eventschedule, eventdescription, image)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertId,
	}

	return res, nil

}

func UpdateEvent(idevent int, eventname string, eventschedule string, eventdescription string, image string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE event SET eventname = ?,  eventschedule = ?, eventdescription = ?, image = ?   WHERE idevent = ?"

	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := statement.Exec(eventname, eventschedule, eventdescription, image)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteEvent(idevent int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM event WHERE idevent = ? "

	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := statement.Exec(idevent)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
