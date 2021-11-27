export class EstudianteMateria {
    id: number;
    studentId: number;
    materiaId: number;
    materiaNombre: string;
    inicio: string;
    fin: string;
    observaciones: string;

    constructor(id: number,
        studentId: number,
        materiaId: number,
        materiaNombre: string,
        inicio: string,
        fin: string,
        observaciones: string
    ) {
        this.id = id;
        this.studentId = studentId;
        this.materiaId = materiaId;
        this.materiaNombre = materiaNombre;
        this.inicio = inicio;
        this.fin = fin;
        this.observaciones = observaciones;
    }

    getClone(): EstudianteMateria {
        return new EstudianteMateria(
            this.id,
            this.studentId,
            this.materiaId,
            this.materiaNombre,
            this.inicio,
            this.fin,
            this.observaciones);
    }

    static cloneIt(src: EstudianteMateria): EstudianteMateria {
        return new EstudianteMateria(
            src.id,
            src.studentId,
            src.materiaId,
            src.materiaNombre,
            src.inicio,
            src.fin,
            src.observaciones);
    }

    static newEmpty(): EstudianteMateria {
        return new EstudianteMateria(
            0,
            0,
            0,
            '',
            '2000-01-01',
            '2000-01-01',
            ''
        );
    }

    static guaranteedClone(src: EstudianteMateria | undefined): EstudianteMateria {
        if (src === undefined) {
            return this.newEmpty();
        }
        return EstudianteMateria.cloneIt(src);
    }

}