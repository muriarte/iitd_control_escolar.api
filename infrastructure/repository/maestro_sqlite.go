package repository

import (
	"database/sql"
	"fmt"

	"iitd_control_escolar.api/entity"
)

const MaestroFieldList = `nombres, apellidos, nacimiento, sexo, calle, numeroext, numeroint, colonia, 
municipio, estado, pais, cp, telcelular, telcasa, email, fechainicio, observaciones, activo`

//MaestroSQLite mysql repo
type MaestroSQLite struct {
	db *sql.DB
}

//NewMaestroSQLite create new repository
func NewMaestroSQLite(db *sql.DB) *MaestroSQLite {
	return &MaestroSQLite{
		db: db,
	}
}

//Create a maestro
func (r *MaestroSQLite) Create(e *entity.Maestro) (int, error) {
	var sqlStr = fmt.Sprintf("insert into maestros (%s) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", MaestroFieldList)
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

//Get a maestro
func (r *MaestroSQLite) Get(id int) (*entity.Maestro, error) {
	var sqlStr = fmt.Sprintf("select id, %s from maestros where id = ?", MaestroFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var b entity.Maestro
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	if rows.Next() {
		err = maestroFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		return &b, nil
	}
	return nil, nil
}

//Update a maestro
func (r *MaestroSQLite) Update(e *entity.Maestro) error {
	//	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`update maestros set nombres=?, apellidos=?, nacimiento=?, sexo=?, calle=?, 
	numeroext=?, numeroint=?, colonia=?, municipio=?, estado=?, pais=?, cp=?, telcelular=?, 
	telcasa=?, email=?, fechainicio=?, observaciones=?, activo=? where id = ?`,
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
		//		e.UpdatedAt.Format("2006-01-02"),
		e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search maestros
func (r *MaestroSQLite) Search(query string) ([]*entity.Maestro, error) {
	var sqlStr = fmt.Sprintf("select id,%s from maestros where concat(nombres, ' ',apellidos) like ?", MaestroFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var maestros []*entity.Maestro
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.Maestro
		err = maestroFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		maestros = append(maestros, &b)
	}

	return maestros, nil
}

//List maestros
func (r *MaestroSQLite) List() ([]*entity.Maestro, error) {
	var sqlStr = fmt.Sprintf("select id, %s from maestros", MaestroFieldList)
	stmt, err := r.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	var maestros []*entity.Maestro
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = stmt.Close()
	}()

	for rows.Next() {
		var b entity.Maestro
		err = maestroFullRowScan(rows, &b)
		if err != nil {
			return nil, err
		}
		maestros = append(maestros, &b)
	}
	return maestros, nil
}

//Delete a maestro
func (r *MaestroSQLite) Delete(id int) error {
	_, err := r.db.Exec("delete from maestros where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func maestroFullRowScan(rows *sql.Rows, b *entity.Maestro) error {
	err := rows.Scan(&b.ID, &b.Nombres, &b.Apellidos, &b.Nacimiento, &b.Sexo, &b.Calle, &b.NumeroExt,
		&b.NumeroInt, &b.Colonia, &b.Municipio, &b.Estado, &b.Pais, &b.CP, &b.TelCelular, &b.TelCasa,
		&b.Email, &b.FechaInicio, &b.Observaciones, &b.Activo)
	return err
}
