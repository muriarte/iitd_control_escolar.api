package student

import (
	"iitd_control_escolar.api/entity"
	"iitd_control_escolar.api/infrastructure/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newFixtureStudent() *entity.Student {
	return &entity.Student{
		ID:        entity.NewID(),
		Email:     "ozzy@metalgods.net",
		Nombres:   "Ozzy",
		Apellidos: "Osbourne",
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := repository.NewStudentInmem()
	m := NewService(repo)
	u := newFixtureStudent()
	_, err := m.CreateStudent(u.Nombres, u.Apellidos, u.Nacimiento, u.Sexo, u.Calle, u.NumeroExt, u.NumeroInt,
		u.Colonia, u.Municipio, u.Estado, u.Pais, u.CP, u.TelCelular, u.TelCasa, u.Email, u.FechaInicio,
		u.Observaciones, u.Activo)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
	assert.True(t, u.UpdatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := repository.NewStudentInmem()
	m := NewService(repo)
	u1 := newFixtureStudent()
	u2 := newFixtureStudent()
	u2.Nombres = "Lemmy"

	uID, _ := m.CreateStudent(u1.Nombres, u1.Apellidos, u1.Nacimiento, u1.Sexo, u1.Calle, u1.NumeroExt, u1.NumeroInt,
		u1.Colonia, u1.Municipio, u1.Estado, u1.Pais, u1.CP, u1.TelCelular, u1.TelCasa, u1.Email, u1.FechaInicio,
		u1.Observaciones, u1.Activo)

	_, _ = m.CreateStudent(u2.Nombres, u2.Apellidos, u2.Nacimiento, u2.Sexo, u2.Calle, u2.NumeroExt, u2.NumeroInt,
		u2.Colonia, u2.Municipio, u2.Estado, u2.Pais, u2.CP, u2.TelCelular, u2.TelCasa, u2.Email, u2.FechaInicio,
		u2.Observaciones, u2.Activo)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchStudents("ozzy")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Osbourne", c[0].Apellidos)

		c, err = m.SearchStudents("dio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListStudents()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetStudent(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.Nombres, saved.Nombres)
	})
}

func Test_Update(t *testing.T) {
	repo := repository.NewStudentInmem()
	m := NewService(repo)
	u := newFixtureStudent()
	id, err := m.CreateStudent(u.Nombres, u.Apellidos, u.Nacimiento, u.Sexo, u.Calle, u.NumeroExt, u.NumeroInt,
		u.Colonia, u.Municipio, u.Estado, u.Pais, u.CP, u.TelCelular, u.TelCasa, u.Email, u.FechaInicio,
		u.Observaciones, u.Activo)
	assert.Nil(t, err)
	saved, _ := m.GetStudent(id)
	saved.Nombres = "Dio"
	//saved.Books = append(saved.Books, entity.NewID())
	assert.Nil(t, m.UpdateStudent(saved))
	updated, err := m.GetStudent(id)
	assert.Nil(t, err)
	assert.Equal(t, "Dio", updated.Nombres)
	assert.False(t, updated.UpdatedAt.IsZero())
	//assert.Equal(t, 1, len(updated.Books))
}

func TestDelete(t *testing.T) {
	repo := repository.NewStudentInmem()
	m := NewService(repo)
	u1 := newFixtureStudent()
	u2 := newFixtureStudent()
	u2ID, _ := m.CreateStudent(u2.Nombres, u2.Apellidos, u2.Nacimiento, u2.Sexo, u2.Calle, u2.NumeroExt, u2.NumeroInt,
		u2.Colonia, u2.Municipio, u2.Estado, u2.Pais, u2.CP, u2.TelCelular, u2.TelCasa, u2.Email, u2.FechaInicio,
		u2.Observaciones, u2.Activo)

	err := m.DeleteStudent(u1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteStudent(u2ID)
	assert.Nil(t, err)
	_, err = m.GetStudent(u2ID)
	assert.Equal(t, entity.ErrNotFound, err)

	//Check that if the items contains subitems (i.e. customer has orders or invoices) it cannot be deleted
	//u3 := newFixtureStudent()
	//id, _ := m.CreateStudent(u3.Nombres, u3.Apellidos, u3.Nacimiento, u3.Sexo, u3.Calle, u3.NumeroExt, u3.NumeroInt,
	//	u3.Colonia, u3.Municipio, u3.Estado, u3.Pais, u3.CP, u3.TelCelular, u3.TelCasa, u3.Email, u3.FechaInicio,
	//	u3.Observaciones, u3.Activo)
	//saved, _ := m.GetStudent(id)
	//saved.Books = []entity.ID{entity.NewID()}
	//_ = m.UpdateStudent(saved)
	//err = m.DeleteStudent(id)
	//assert.Equal(t, entity.ErrCannotBeDeleted, err)
}
