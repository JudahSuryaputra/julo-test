package http

import (
	"julo-case-study/http/Auth"
	"julo-case-study/http/User"
	"julo-case-study/http/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/gocraft/dbr"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(dbConn *dbr.Connection) {
	a.Router = mux.NewRouter().StrictSlash(true)

	a.authRoutes(dbConn)
	a.userRoutes(dbConn)
}

func (a *App) authRoutes(dbConn *dbr.Connection) {
	a.Router.Use(middlewares.SetContentTypeMiddleware)

	initAccount := Auth.InitAccount{DBConn: dbConn}
	a.Router.Handle("/api/v1/init", initAccount).Methods(http.MethodPost)
}

func (a *App) userRoutes(dbConn *dbr.Connection) {
	user := a.Router.PathPrefix("/api/v1").Subrouter()
	user.Use(middlewares.CommonAuthJwtVerify)

	enableWallet := User.EnableWallet{DBConn: dbConn}
	getBalance := User.GetBalance{DBConn: dbConn}
	user.Handle("/wallet", enableWallet).Methods(http.MethodPost)
	user.Handle("/wallet", getBalance).Methods(http.MethodGet)
}

func (a *App) RunServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("PORT")
	}
	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"})

	log.Printf("\nServer starting on Port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CombinedLoggingHandler(os.Stderr, handlers.CORS(headersOK, originsOK, methodsOK)(a.Router))))
}
