package repository

import (
	"database/sql"
	"fmt"
	"iitd_control_escolar.api/entity"
	"time"
)

const StudentFieldList = `nombres, apellidos, nacimiento, sexo, calle, numeroext, numeroint, colonia, 
municipio, estado, pais, cp, telcelular, telcasa, email, fechainicio, observaciones, activo, created_at`

//StudentSQLite mysql repo
type StudentSQLite struct {
	db *sql.DB
}

//NewStudentSQLite create new repository
func NewStudentSQLite(db *sql.DB) *StudentSQLite {
	return &StudentSQLite{
		db: db,
	}
}

//Create a student
func (r *StudentSQLite) Create(e *entity.Student) (entity.ID, error) {
	var sqlStr = fmt.Sprintf("insert into students (%s) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", StudentFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return e.ID, err
	}
	rsp, err := stmt.Exec(
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
		e.Observaciones,
		e.Activo,
		time.Now().Format("2006-01-02"),
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
	return entity.ID(lastID), nil
}

//Get a student
func (r *StudentSQLite) Get(id entity.ID) (*entity.Student, error) {
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
	for rows.Next() {
		err = studentFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
	}
	return &b, nil
}

//Update a student
func (r *StudentSQLite) Update(e *entity.Student) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`update students set nombres=?, apellidos=?, nacimiento=?, sexo=?, calle=?, 
                    numeroext=?, numeroint=?, colonia=?, municipio=?, estado=?, pais=?, cp=?, telcelular=?, 
                    telcasa=?, email=?, fechainicio=?, observaciones=?, activo=?, created_at=? where id = ?`,
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
		e.Observaciones,
		e.Activo,
		e.UpdatedAt.Format("2006-01-02"),
		e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search students
func (r *StudentSQLite) Search(query string) ([]*entity.Student, error) {
	var sqlStr = fmt.Sprintf("select %s from students where (firstName || ' ' || lastname) like ?", StudentFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var students []*entity.Student
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
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
func (r *StudentSQLite) List() ([]*entity.Student, error) {
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
func (r *StudentSQLite) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from students where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func studentFullRowScan(rows *sql.Rows, b *entity.Student) error {
	err := rows.Scan(&b.ID, &b.Nombres, &b.Apellidos, &b.Nacimiento, &b.Sexo, &b.Calle, &b.NumeroExt,
		&b.NumeroInt, &b.Colonia, &b.Municipio, &b.Estado, &b.Pais, &b.CP, &b.TelCelular, &b.TelCasa,
		&b.Email, &b.FechaInicio, &b.Observaciones, &b.Activo, &b.CreatedAt)
	return err
}
