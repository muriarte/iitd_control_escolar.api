package repository

import (
	"database/sql"
	"fmt"

	"time"

	"iitd_control_escolar.api/entity"
)

//const StudentFieldList =  `nombres, apellidos, nacimiento, sexo, calle, numeroext, numeroint, colonia,
//municipio, estado, pais, cp, telcelular, telcasa, email, fechainicio, observaciones, activo, created_at`

//StudentMySQL mysql repo
type StudentMySQL struct {
	db *sql.DB
}

//NewStudentMySQL create new repository
func NewStudentMySQL(db *sql.DB) *StudentMySQL {
	return &StudentMySQL{
		db: db,
	}
}

//Create a student
func (r *StudentMySQL) Create(e *entity.Student) (int, error) {
	var sqlStr = fmt.Sprintf("insert into students (%s) values(?,?,?,?,?)", StudentFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.Nombres,
		e.Apellidos,
		e.Nacimiento,
		e.Sexo,
		e.Calle,
		e.NumeroExt,
		e.NumeroInt,
		e.Colonia,
		e.Municipio,
		e.Estado,
		e.Pais,
		e.CP,
		e.TelCelular,
		e.TelCasa,
		e.Email,
		e.FechaInicio,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

//Get a student
func (r *StudentMySQL) Get(id int) (*entity.Student, error) {
	var sqlStr = fmt.Sprintf("select id, %s from students where id = ?", StudentFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var b entity.Student
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		err = studentFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
	}
	return &b, nil
}

//Update a student
func (r *StudentMySQL) Update(e *entity.Student) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update students set email = ?, firstname = ?, lastname = ?, age = ?, updated_at = ? where id = ?",
		e.Nombres,
		e.Apellidos,
		e.Nacimiento,
		e.Sexo,
		e.Calle,
		e.NumeroExt,
		e.NumeroInt,
		e.Colonia,
		e.Municipio,
		e.Estado,
		e.Pais,
		e.CP,
		e.TelCelular,
		e.TelCasa,
		e.Email,
		e.FechaInicio,
		e.UpdatedAt.Format("2006-01-02"),
		e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search students
func (r *StudentMySQL) Search(query string) ([]*entity.Student, error) {
	var sqlStr = fmt.Sprintf("select %s from students where concat(firstName, ' ',lastname) like ?", StudentFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var students []*entity.Student
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.Student
		err = studentFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		students = append(students, &b)
	}

	return students, nil
}

//List students
func (r *StudentMySQL) List() ([]*entity.Student, error) {
	var sqlStr = fmt.Sprintf("select id, %s from students", StudentFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var students []*entity.Student
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.Student
		err = studentFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		students = append(students, &b)
	}
	return students, nil
}

//Delete a student
func (r *StudentMySQL) Delete(id int) error {
	_, err := r.db.Exec("delete from students where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
