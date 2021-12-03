export class EstudianteObs {
    id: number;
    studentId: number;
    fecha: string;
    observacion: string;

    constructor(id: number,
        studentId: number,
        fecha: string,
        observacion: string
    ) {
        this.id = id;
        this.studentId = studentId;
        this.fecha = fecha;
        this.observacion = observacion;
    }

    getClone(): EstudianteObs {
        return new EstudianteObs(
            this.id,
            this.studentId,
            this.fecha,
            this.observacion);
    }

    static cloneIt(src: EstudianteObs): EstudianteObs {
        return new EstudianteObs(
            src.id,
            src.studentId,
            src.fecha,
            src.observacion);
    }

    /** Regresa una instancia con los valores de sus campos vacios */
    static newEmpty(): EstudianteObs {
        return new EstudianteObs(
            0,
            0,
            '2000-01-01',
            ''
        );
    }

    /** Regresa una copia de la instancia proporcionada, 
     * si la instancia es undefined regresa una instancia con los valores de sus campos vacios */
    static guaranteedClone(src: EstudianteObs | undefined): EstudianteObs {
        if (src === undefined) {
            return this.newEmpty();
        }
        return EstudianteObs.cloneIt(src);
    }

}