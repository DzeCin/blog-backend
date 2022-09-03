/*
 * Blog API
 *
 * This is a blog API
 *
 * API version: 1.0.0
 * Contact: dzenancindrak@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package blog

import (
	"context"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name           string
	Method         string
	Pattern        string
	EnabledAuth    bool
	ctxHandlerFunc ctxHandler
}

type Routes []Route

type ctxHandler func(w http.ResponseWriter, r *http.Request, ctx *context.Context)

func ContextHandler(handler ctxHandler, ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, ctx)
	}
}

func NewRouter(ctx *context.Context) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var contHandler ctxHandler
		var handler http.Handler
		contHandler = route.ctxHandlerFunc
		handler = http.HandlerFunc(ContextHandler(contHandler, ctx))
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"HealthCheck",
		strings.ToUpper("Get"),
		"/",
		false,
		HealthCheck,
	},

	Route{
		"AddPost",
		strings.ToUpper("Post"),
		"/posts",
		false,
		AddPost,
	},

	Route{
		"DeletePost",
		strings.ToUpper("Delete"),
		"/posts/{postId}",
		false,
		DeletePost,
	},

	Route{
		"GetPost",
		strings.ToUpper("Get"),
		"/posts/{postId}",
		false,
		GetPost,
	},

	Route{
		"GetPosts",
		strings.ToUpper("Get"),
		"/posts",
		false,
		GetPosts,
	},

	Route{
		"UpdatePost",
		strings.ToUpper("Patch"),
		"/posts/{postId}",
		false,
		UpdatePost,
	},
}
