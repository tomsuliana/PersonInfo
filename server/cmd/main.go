package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"server/server/config"
	personDel "server/server/internal/Person/delivery"
	personRep "server/server/internal/Person/repository/postgres"
	personUsecase "server/server/internal/Person/usecase"
	"server/server/internal/middleware"
)

const PORT = ":8080"

var (
	host     = "localhost"
	port     = 5432
	user     = "uliana"
	password = "uliana"
	dbname   = "personinfo"

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
)

//GetPostgres gets postgres connection
func GetPostgres(psqlInfo string) (*sql.DB, error) {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}

func main() {
	router := mux.NewRouter()
	db, err := GetPostgres(psqlInfo)
	if err != nil {
		fmt.Println(err, " ", psqlInfo)
		log.Fatalf("cant connect to postgres")
		return
	}
	defer db.Close()

	baseLogger, err := config.Cfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer errorLogger.Sync()

	logger := middleware.NewACLog(baseLogger.Sugar(), errorLogger.Sugar())

	personRepo := personRep.NewPersonRepo(db)
	personUC := personUsecase.NewPersonUsecase(personRepo)
	personHandler := personDel.NewPersonHandler(personUC, logger)

	personHandler.RegisterHandler(router)

	router.Use(middleware.PanicMiddleware)
	router.Use(logger.ACLogMiddleware)

	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	fmt.Println("Server start at port", PORT[1:])
	err = server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}

}
