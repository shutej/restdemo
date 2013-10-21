package main

// This package is a very basic illustration for how to create an
// Oauth2-protected REST API in Go.  It doesn't even stop and ask you if you
// want to authorize, it just auto-authorizes client1 to prove that all the
// pieces work.

import (
	"github.com/ant0ine/go-json-rest"
	"github.com/shutej/goauth2"
	"github.com/shutej/goauth2/authcache"
	"github.com/shutej/goauth2/authhandler"
	"log"
	"net/http"
	"time"
)

// User is an example model object served by the REST API.
type User struct {
	Id   string
	Name string
}

// GetUserV1 serves GET requests for users.  A real implementation would
// probably talk to a persistence layer.
func GetUserV1(w *rest.ResponseWriter, req *rest.Request) {
	user := User{
		Id:   req.PathParam("id"),
		Name: "User " + req.PathParam("id"),
	}
	w.WriteJson(&user)
}

func main() {
	// A basic auth cache.  A real implementation would probably use an
	// out-of-process auth cache, like the Redis-based one packaged in the
	// goauth2 distribution.
	authCache := authcache.NewBasicAuthCache()

	// A dummy auth handler.  A real implementation would probably talk to a
	// persistence layer.
	authHandler := authhandler.NewWhiteList("client1")

	server := goauth2.NewServer(authCache, authHandler)

	mux := http.NewServeMux()

	// Serves authorization requests from the /authorize URL path.  A real
	// implementation would probably override this to ask if we should allow the
	// application to use your credentials.
	mux.Handle("/authorize", server.MasterHandler())

	// Serves static files from the local "static" directory under the
	// /static/... URL path.
	const static = "/static/"
	mux.Handle(
		static, http.StripPrefix(static, http.FileServer(http.Dir("static"))))

	// Serves a versioned REST API which only authorized clients can access.
	apiV1 := &rest.ResourceHandler{}
	apiV1.SetRoutes(
		rest.Route{"GET", "/v1/user/:id", GetUserV1},
	)
	mux.Handle("/v1/", server.TokenVerifier(apiV1))

	// Create the HTTP server.  We hard-code port 8000.  Our jQuery has also
	// hard-coded the same port.  A real implementation would probably get the
	// host and port from a flag, configuration file, or environment variable.
	httpd := &http.Server{
		Addr:           ":8000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start the server.  We hard-code our cert and key.  A real implementation
	// must override these with values produced by a real certificate authority.
	log.Fatal(httpd.ListenAndServeTLS("certs/cert.pem", "certs/key.pem"))
}
