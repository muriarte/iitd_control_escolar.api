package repository

import (
	"database/sql"
	"fmt"

	"iitd_control_escolar.api/entity"
)

const StudentMateriasFieldList = `studentId, materiaId, inicio, fin, observaciones`
const StudentMateriasFieldListPrefixed = `sm.studentId, sm.materiaId, sm.inicio, sm.fin, sm.observaciones`

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
func (r *StudentMateriaSQLite) Create(e *entity.StudentMateria) (int, string, error) {
	var sqlStr = fmt.Sprintf("insert into studentmaterias (%s) values(?,?,?,?,?)", StudentMateriasFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return e.ID, "", err
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
		return e.ID, "", err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, "", err
	}

	lastID, err := rsp.LastInsertId()
	if err != nil {
		return e.ID, "", err
	}
	materiaNombre := r.getMateriaNombre(e.MateriaId)
	return int(lastID), materiaNombre, nil
}

//Get a studentmaterias
func (r *StudentMateriaSQLite) Get(id int) (*entity.StudentMateria, error) {
	var sqlStr = fmt.Sprintf("select sm.id, m.nombre, %s from studentmaterias sm left join materias m on m.id=sm.materiaid where sm.id = ?", StudentMateriasFieldListPrefixed)
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
func (r *StudentMateriaSQLite) Update(e *entity.StudentMateria) (string, error) {
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
		return "", err
	}
	materiaNombre := r.getMateriaNombre(e.MateriaId)
	return materiaNombre, nil
}

//Search studentmateriass
func (r *StudentMateriaSQLite) Search(studentId int, materiaId int) ([]*entity.StudentMateria, error) {
	query := ""
	connector := "where "
	if studentId > 0 {
		query = fmt.Sprintf("%s%ssm.studentId = %d", query, connector, studentId)
		connector = " and "
	}
	if materiaId > 0 {
		query = fmt.Sprintf("%s%ssm.materiaId = %d", query, connector, materiaId)
	}
	var sqlStr = fmt.Sprintf("select sm.id,m.nombre,%s from studentmaterias sm left join materias m on m.id=sm.materiaid %s order by sm.studentid,sm.inicio,sm.fin", StudentMateriasFieldListPrefixed, query)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var studentMaterias []*entity.StudentMateria
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
		studentMaterias = append(studentMaterias, &b)
	}

	return studentMaterias, nil
}

//List studentmateriass
func (r *StudentMateriaSQLite) List() ([]*entity.StudentMateria, error) {
	var sqlStr = fmt.Sprintf("select sm.id, m.nombre, %s from studentmaterias sm left join materias m on m.id=sm.materiaid order by sm.studentid,sm.inicio,sm.fin", StudentMateriasFieldListPrefixed)
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
	var materiaNombre sql.NullString
	err := rows.Scan(&b.ID, &materiaNombre, &b.StudentId, &b.MateriaId, &b.Inicio, &b.Fin, &b.Observaciones)
	b.MateriaNombre = materiaNombre.String
	return err
}

func (r *StudentMateriaSQLite) getMateriaNombre(materiaId int) string {
	dalMateria := NewMateriaSQLite(r.db)
	materia, err := dalMateria.Get(materiaId)
	if err != nil {
		return "**  No encontrada **"
	}
	if materia == nil {
		return fmt.Sprintf("** Materia con id %d No encontrada **", materiaId)
	}
	return materia.Nombre
}
