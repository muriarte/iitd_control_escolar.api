package main

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"iitd_control_escolar.api/pkg/metric"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"iitd_control_escolar.api/infrastructure/repository"
	"iitd_control_escolar.api/usecase/student"

	//"github.com/prometheus/client_golang/prometheus/promhttp"

	// lightweight middleware management
	"github.com/urfave/negroni"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"iitd_control_escolar.api/api/handler"
	"iitd_control_escolar.api/api/middleware"
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

	//userRepo := repository.NewUserMySQL(db)
	//userService := user.NewService(userRepo)
	//
	//loanUseCase := loan.NewService(userService, studentService)

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}
	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(middleware.Cors),
		middleware.Metrics(metricService),
		negroni.NewLogger(),
	)

	// Build http handlers

	//student
	handler.MakeStudentHandlers(r, *n, studentService)

	////user
	//handler.MakeUserHandlers(r, *n, userService)
	//
	////loan
	//handler.MakeLoanHandlers(r, *n, studentService, userService, loanUseCase)

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	fmt.Printf("iitd Api Server listening on %s", ":"+strconv.Itoa(config.API_PORT))
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
