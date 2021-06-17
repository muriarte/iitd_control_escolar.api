package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/rs/cors"
	"iitd_control_escolar.api/infrastructure/repository"
	"iitd_control_escolar.api/usecase/student"

	//"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"iitd_control_escolar.api/api/handler"
	"iitd_control_escolar.api/config"

	//"github.com/eminetto/clean-architecture-go-v2/pkg/metric"
	_ "github.com/mattn/go-sqlite3"
)

// Puntero a la estructura DB, nos permite manejar la
// base de datos
var db *sql.DB

func main() {

	db = GetSqliteConnection()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	// Build http handlers dependencies
	studentRepo := repository.NewStudentSQLite(db) //repository.NewStudentMySQL(db)
	studentService := student.NewService(studentRepo)

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	//handlers

	// Build http handlers

	//student
	handler.MakeStudentHandlers(r, studentService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		//Handler:      context.ClearHandler(http.DefaultServeMux),
		Handler:  context.ClearHandler(c.Handler(r)),
		ErrorLog: logger,
	}

	// Desplegamos directorio actual
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error al obtener directorio actual %s", err.Error())
	}
	fmt.Printf("Directorio actual:%s\n", pwd)

	// Realizamos test de acceso a la base de datos
	sts, err := studentService.ListStudents()
	if err != nil {
		log.Fatalf("Error al hacer test de acceso a la base de datos %s", err.Error())
	}
	fmt.Printf("Estudiantes en base de datos: %d\n", len(sts))

	fmt.Printf("iitd Api Server listening on %s\n", ":"+strconv.Itoa(config.API_PORT))
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetSqliteConnection() *sql.DB {
	// Para evitar realizar una nueva conexión en cada llamada a
	// la función GetConnection.
	if db != nil {
		return db
	}
	// Declaramos la variable err para poder usar el operador
	// de asignación “=” en lugar que el de asignación corta,
	// para evitar que cree una nueva variable db en este scope y
	// en su lugar que inicialice la variable db que declaramos a
	// nivel de paquete.
	var err error
	// Conexión a la base de datos
	db, err = sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func GetMysqlConnection() *sql.DB {
	// Para evitar realizar una nueva conexión en cada llamada a
	// la función GetConnection.
	if db != nil {
		return db
	}
	// Declaramos la variable err para poder usar el operador
	// de asignación “=” en lugar que el de asignación corta,
	// para evitar que cree una nueva variable db en este scope y
	// en su lugar que inicialice la variable db que declaramos a
	// nivel de paquete.
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	var err error
	// Conexión a la base de datos
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
