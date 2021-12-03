package repository

import (
	"database/sql"
	"fmt"

	"iitd_control_escolar.api/entity"
)

const StudentObsFieldList = `studentid, fecha, observacion`

//StudentObsSQLite mysql repo
type StudentObsSQLite struct {
	db *sql.DB
}

//NewStudentObsSQLite create new repository
func NewStudentObsSQLite(db *sql.DB) *StudentObsSQLite {
	return &StudentObsSQLite{
		db: db,
	}
}

//Observaciones an observacion
func (r *StudentObsSQLite) Create(e *entity.StudentObs) (int, error) {
	var sqlStr = fmt.Sprintf("insert into observaciones (%s) values(?,?,?)", StudentObsFieldList)
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
func (r *StudentObsSQLite) Get(id int) (*entity.StudentObs, error) {
	var sqlStr = fmt.Sprintf("select id, %s from observaciones where id = ?", StudentObsFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var b entity.StudentObs
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
func (r *StudentObsSQLite) Update(e *entity.StudentObs) error {
	//	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`update observaciones set studentId=?, fecha=?, observacion=? where id = ?`,
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
func (r *StudentObsSQLite) Search(studentId int) ([]*entity.StudentObs, error) {
	query := ""
	connector := "where "
	if studentId > 0 {
		query = fmt.Sprintf("%s%sstudentId = %d", query, connector, studentId)
		connector = " and "
	}
	var sqlStr = fmt.Sprintf("select id,%s from observaciones %s order by studentid, fecha", StudentObsFieldList, query)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var observaciones []*entity.StudentObs
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.StudentObs
		err = observacionesFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		observaciones = append(observaciones, &b)
	}

	return observaciones, nil
}

//List observaciones
func (r *StudentObsSQLite) List() ([]*entity.StudentObs, error) {
	var sqlStr = fmt.Sprintf("select id, %s from observaciones order by studentid, fecha", StudentObsFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var observaciones []*entity.StudentObs
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.StudentObs
		err = observacionesFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		observaciones = append(observaciones, &b)
	}
	return observaciones, nil
}

//Delete an observacion
func (r *StudentObsSQLite) Delete(id int) error {
	_, err := r.db.Exec("delete from observaciones where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func observacionesFullRowScan(rows *sql.Rows, b *entity.StudentObs) error {
	err := rows.Scan(&b.ID, &b.StudentId, &b.Fecha, &b.Observacion)
	// x := b.Fecha.Format("2006-01-02")
	// fmt.Println(x)
	return err
}
