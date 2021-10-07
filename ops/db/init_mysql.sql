create
database if not exists iitd;
use iitd;

create table if not exists students
(
    id int,
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
    observaciones mediumtext,
    activo char(1),
    created_at datetime,
    updated_at datetime,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

create table if not exists maestros
(
    id int,
    nombre varchar(100),
    observaciones mediumtext,
    activo char(1),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;


create table if not exists materias
(
    id int,
    nombre varchar(100),
    observaciones mediumtext,
    activo char(1),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

create table if not exists studentmaterias
(
    id int,
    studentid int,
    materiaid int,
    inicio datetime,
    fin datetime,
    observaciones mediumtext,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

create table if not exists observaciones
(
    id int,
    studentid int,
    fecha datetime,
    observacion mediumtext,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
