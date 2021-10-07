package repository

import (
	"database/sql"
	"fmt"

	"iitd_control_escolar.api/entity"
)

const StudentMateriasFieldList = `studentId, materiaId, inicio, fin, observaciones`

//StudentMateriassQLite mysql repo
type StudentMateriaSQLite struct {
	db *sql.DB
}

//NewStudentMateriassQLite create new repository
func NewStudentMateriaSQLite(db *sql.DB) *StudentMateriaSQLite {
	return &StudentMateriaSQLite{
		db: db,
	}
}

//Create a studentmaterias
func (r *StudentMateriaSQLite) Create(e *entity.StudentMateria) (int, error) {
	var sqlStr = fmt.Sprintf("insert into studentmaterias (%s) values(?,?,?,?,?)", StudentMateriasFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return e.ID, err
	}
	rsp, err := stmt.Exec(
		e.StudentId,
		e.MateriaId,
		e.Inicio,
		e.Fin,
		e.Observaciones,
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

//Get a studentmaterias
func (r *StudentMateriaSQLite) Get(id int) (*entity.StudentMateria, error) {
	var sqlStr = fmt.Sprintf("select id, %s from studentmaterias where id = ?", StudentMateriasFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var b entity.StudentMateria
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	if rows.Next() {
		err = studentMateriasFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		return &b, nil
	}
	return nil, nil
}

//Update a studentmaterias
func (r *StudentMateriaSQLite) Update(e *entity.StudentMateria) error {
	//	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`update studentmaterias set studentId=?, materiaId=?, inicio=?, fin=?, observaciones=? where id = ?`,
		e.StudentId,
		e.MateriaId,
		e.Inicio,
		e.Fin,
		e.Observaciones,
		//		e.UpdatedAt.Format("2006-01-02"),
		e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search studentmateriass
func (r *StudentMateriaSQLite) Search(studentId int, materiaId int) ([]*entity.StudentMateria, error) {
	query := ""
	connector := "where "
	if studentId > 0 {
		query = fmt.Sprintf("%s%sstudentId = %d", query, connector, studentId)
		connector = " and "
	}
	if materiaId > 0 {
		query = fmt.Sprintf("%s%smateriaId = %d", query, connector, materiaId)
	}
	var sqlStr = fmt.Sprintf("select id,%s from studentmaterias %s", StudentMateriasFieldList, query)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var studentMaterias []*entity.StudentMateria
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.StudentMateria
		err = studentMateriasFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		studentMaterias = append(studentMaterias, &b)
	}

	return studentMaterias, nil
}

//List studentmateriass
func (r *StudentMateriaSQLite) List() ([]*entity.StudentMateria, error) {
	var sqlStr = fmt.Sprintf("select id, %s from studentmaterias", StudentMateriasFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var studentmateriass []*entity.StudentMateria
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.StudentMateria
		err = studentMateriasFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		studentmateriass = append(studentmateriass, &b)
	}
	return studentmateriass, nil
}

//Delete a studentmaterias
func (r *StudentMateriaSQLite) Delete(id int) error {
	_, err := r.db.Exec("delete from studentmaterias where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func studentMateriasFullRowScan(rows *sql.Rows, b *entity.StudentMateria) error {
	err := rows.Scan(&b.ID, &b.StudentId, &b.MateriaId, &b.Inicio, &b.Fin, &b.Observaciones)
	return err
}
