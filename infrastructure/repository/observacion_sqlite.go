package repository

import (
	"database/sql"
	"fmt"

	"iitd_control_escolar.api/entity"
)

const ObservacionesFieldList = `studentid, fecha, observacion`

//ObservacionSQLite mysql repo
type ObservacionSQLite struct {
	db *sql.DB
}

//NewObservacionSQLite create new repository
func NewObservacionSQLite(db *sql.DB) *ObservacionSQLite {
	return &ObservacionSQLite{
		db: db,
	}
}

//Observaciones an observacion
func (r *ObservacionSQLite) Create(e *entity.Observacion) (int, error) {
	var sqlStr = fmt.Sprintf("insert into observaciones (%s) values(?,?,?)", ObservacionesFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return e.ID, err
	}
	rsp, err := stmt.Exec(
		e.StudentId,
		e.Fecha,
		e.Observacion,
		//		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}

	lastID, err := rsp.LastInsertId()
	if err != nil {
		return e.ID, err
	}
	return int(lastID), nil
}

//Get an observacion
func (r *ObservacionSQLite) Get(id int) (*entity.Observacion, error) {
	var sqlStr = fmt.Sprintf("select id, %s from observaciones where id = ?", ObservacionesFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var b entity.Observacion
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	if rows.Next() {
		err = observacionesFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		return &b, nil
	}
	return nil, nil
}

//Update an observacion
func (r *ObservacionSQLite) Update(e *entity.Observacion) error {
	//	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`update observaciones set studentId=?, observacion=? where id = ?`,
		e.StudentId,
		e.Fecha,
		e.Observacion,
		//		e.UpdatedAt.Format("2006-01-02"),
		e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search observaciones
func (r *ObservacionSQLite) Search(studentId int) ([]*entity.Observacion, error) {
	query := ""
	connector := "where "
	if studentId > 0 {
		query = fmt.Sprintf("%s%sstudentId = %d", query, connector, studentId)
		connector = " and "
	}
	var sqlStr = fmt.Sprintf("select id,%s from observaciones %s", ObservacionesFieldList, query)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var observaciones []*entity.Observacion
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.Observacion
		err = observacionesFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		observaciones = append(observaciones, &b)
	}

	return observaciones, nil
}

//List observaciones
func (r *ObservacionSQLite) List() ([]*entity.Observacion, error) {
	var sqlStr = fmt.Sprintf("select id, %s from observaciones", ObservacionesFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var observaciones []*entity.Observacion
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.Observacion
		err = observacionesFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		observaciones = append(observaciones, &b)
	}
	return observaciones, nil
}

//Delete an observacion
func (r *ObservacionSQLite) Delete(id int) error {
	_, err := r.db.Exec("delete from observaciones where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func observacionesFullRowScan(rows *sql.Rows, b *entity.Observacion) error {
	err := rows.Scan(&b.ID, &b.StudentId, &b.Fecha, &b.Observacion)
	return err
}
