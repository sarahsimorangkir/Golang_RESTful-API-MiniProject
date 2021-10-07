package models

import (
	"echo-rest/db"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

type Admin struct {
	IdAdm   int    `json : "idadm`
	NameAdm string `json : "nameadm" validate :"required"`
	Jk      string `json : "jk" validate :"required"`
	Telp    string `json : "telp" validate :"required"`
}

func TakeAdmin() (Response, error) {
	var obj Admin
	var arrobj []Admin
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM admin"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdAdm, &obj.NameAdm, &obj.Jk, &obj.Telp)
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

func AllAdmin(nameadm string, jk string, telp string) (Response, error) {
	var res Response

	v := validator.New()

	adm := Admin{
		NameAdm: nameadm,
		Jk:      jk,
		Telp:    telp,
	}

	err := v.Struct(adm)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT admin (nameadm, jk, telp) VALUES (?,?,?,?)"
	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := statement.Exec(nameadm, jk, telp)
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

func UpdateAdmin(idadm int, nameadm string, jk string, telp string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE admin SET nameadm = ?,  jk = ?, telp = ?   WHERE idadm = ?"

	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := statement.Exec(nameadm, jk, telp, idadm)
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

func DeleteAdmin(idadm int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM admin WHERE idadm = ? "

	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := statement.Exec(idadm)
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
