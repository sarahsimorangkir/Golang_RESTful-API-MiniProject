package models

import (
	"echo-rest/db"
	"net/http"
)

type Product struct {
	IdProduct   int    `json : "idproduct"`
	ProductName string `json : "productname"`
	Category     string `json : "category"`
	Image   string `json : "image"`
	Tutorial   string `json : "tutorial"`

}

func TakeProduct() (Response, error) {
	var obj Product
	var arrobj []Product
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM product"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdProduct, &obj.ProductName, &obj.Category, &obj.Image, &obj.Tutorial)
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

func AllProduct(productname string, category string, image string, tutorial) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT product (productname, category, image, tutorial) VALUES (?,?,?,?,?)"
	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := statement.Exec(productname, category, image, tutorial)
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

func UpdateProduct(idproduct int, productname string, category string, image string, tutorial string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE product SET productname = ?,  category = ?, image = ?   WHERE idproduct = ?"

	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := statement.Exec(productname, category, image, tutorial, idproduct)
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

func DeleteProduct(idproduct int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM product WHERE idproduct = ? "

	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := statement.Exec(idproduct)
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
