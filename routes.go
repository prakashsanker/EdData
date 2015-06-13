package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "log"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes[]Route

var db *sql.DB
var err error

func NewRouter() *mux.Router {
	db, err = sql.Open("mysql", "psanker:123@/education_data")
	err = db.Ping()

	if err != nil {
		fmt.Println("Failed to prepare connection to database")
		log.Fatal("Error:", err.Error())
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route {
		"Index",
		"GET",
		"/",
		Index,
	},
	Route {
		"getDistrict",
		"GET",
		"/district/{districtId}",
		getDistrict,
	},
	Route {
		"getDistricts",
		"GET",
		"/districts",
		getDistricts,
	},
	Route {
		"getDistrictActivities",
		"GET",
		"/district/{districtId}/activities",
		getDistrictActivities,
	},
	Route {
		"getDistrictActivity",
		"GET",
		"/district/{districtId}/activities/{activityCode}",
		getDistrictActivity,
	},
	Route {
		"getSubActivities",
		"GET",
		"/district/{districtId}/activities/{activityCode}/subActivities",
		getSubActivities,
	},
	Route {
		"getSubActivity",
		"GET",
		"/district/{districtId}/subActivities/{subActivityCode}",
		getSubActivity,
	},
	Route {
		"getExpenses",
		"GET",
		"/district/{districtId}/subActivities/{subActivityCode/expenses",
		getExpenses
	}

}