create table if not exists students
(
    id integer PRIMARY KEY AUTOINCREMENT,
    nombres varchar(100),
    apellidos varchar(100),
    nacimiento datetime,
    sexo char(1),
    calle varchar(100),
    numeroext varchar(50),
    numeroint varchar(50),
    colonia varchar(50),
    municipio varchar(50),
    estado varchar(50),
    pais varchar(50),
    cp varchar(6),
    telcelular varchar(50),
    telcasa varchar(50),
    email varchar(255),
    fechainicio datetime,
    observaciones text,
    activo char(1),
    created_at datetime,
    updated_at datetime
);
