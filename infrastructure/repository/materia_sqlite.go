package repository

import (
	"database/sql"
	"fmt"

	"iitd_control_escolar.api/entity"
)

const MateriaFieldList = `nombre, observaciones, activo`

//MateriaSQLite mysql repo
type MateriaSQLite struct {
	db *sql.DB
}

//NewMateriaSQLite create new repository
func NewMateriaSQLite(db *sql.DB) *MateriaSQLite {
	return &MateriaSQLite{
		db: db,
	}
}

//Create a materia
func (r *MateriaSQLite) Create(e *entity.Materia) (int, error) {
	var sqlStr = fmt.Sprintf("insert into materias (%s) values(?,?,?)", MateriaFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return e.ID, err
	}
	rsp, err := stmt.Exec(
		e.Nombre,
		e.Observaciones,
		e.Activo,
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

//Get a materia
func (r *MateriaSQLite) Get(id int) (*entity.Materia, error) {
	var sqlStr = fmt.Sprintf("select id, %s from materias where id = ?", MateriaFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var b entity.Materia
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	if rows.Next() {
		err = materiaFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		return &b, nil
	}
	return nil, nil
}

//Update a materia
func (r *MateriaSQLite) Update(e *entity.Materia) error {
	//	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`update materias set nombre=?, observaciones=?, activo=? where id = ?`,
		e.Nombre,
		e.Observaciones,
		e.Activo,
		//		e.UpdatedAt.Format("2006-01-02"),
		e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search materias
func (r *MateriaSQLite) Search(query string) ([]*entity.Materia, error) {
	var sqlStr = fmt.Sprintf("select id,%s from materias where nombre like ?", MateriaFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var materias []*entity.Materia
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.Materia
		err = materiaFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		materias = append(materias, &b)
	}

	return materias, nil
}

//List materias
func (r *MateriaSQLite) List() ([]*entity.Materia, error) {
	var sqlStr = fmt.Sprintf("select id, %s from materias", MateriaFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var materias []*entity.Materia
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.Materia
		err = materiaFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		materias = append(materias, &b)
	}
	return materias, nil
}

//Delete a materia
func (r *MateriaSQLite) Delete(id int) error {
	_, err := r.db.Exec("delete from materias where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func materiaFullRowScan(rows *sql.Rows, b *entity.Materia) error {
	err := rows.Scan(&b.ID, &b.Nombre, &b.Observaciones, &b.Activo)
	return err
}
