package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// MyProfile
	rt.router.PUT("/login", rt.wrap(rt.doLogin))
	rt.router.POST("/signin", rt.wrap(rt.doRegister))

	return rt.router
}
