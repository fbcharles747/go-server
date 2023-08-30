/*
 * PetSitter API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"JobApplicationsIdPut",
		strings.ToUpper("Put"),
		"/job-applications/{id}",
		JobApplicationsIdPut,
	},

	Route{
		"JobsGet",
		strings.ToUpper("Get"),
		"/jobs",
		JobsGet,
	},

	Route{
		"JobsIdDelete",
		strings.ToUpper("Delete"),
		"/jobs/{id}",
		JobsIdDelete,
	},

	Route{
		"JobsIdGet",
		strings.ToUpper("Get"),
		"/jobs/{id}",
		JobsIdGet,
	},

	Route{
		"JobsIdJobApplicationsGet",
		strings.ToUpper("Get"),
		"/jobs/{id}/job-applications",
		JobsIdJobApplicationsGet,
	},

	Route{
		"JobsIdJobApplicationsPost",
		strings.ToUpper("Post"),
		"/jobs/{id}/job-applications",
		JobsIdJobApplicationsPost,
	},

	Route{
		"JobsIdPut",
		strings.ToUpper("Put"),
		"/jobs/{id}",
		JobsIdPut,
	},

	Route{
		"JobsPost",
		strings.ToUpper("Post"),
		"/jobs",
		JobsPost,
	},

	Route{
		"RegisterUser",
		strings.ToUpper("Post"),
		"/users",
		RegisterUser,
	},

	Route{
		"UsersIdDelete",
		strings.ToUpper("Delete"),
		"/users/{id}",
		UsersIdDelete,
	},

	Route{
		"UsersIdGet",
		strings.ToUpper("Get"),
		"/users/{id}",
		UsersIdGet,
	},

	Route{
		"UsersIdJobsGet",
		strings.ToUpper("Get"),
		"/users/{id}/jobs",
		UsersIdJobsGet,
	},

	Route{
		"UsersIdPut",
		strings.ToUpper("Put"),
		"/users/{id}",
		UsersIdPut,
	},
}