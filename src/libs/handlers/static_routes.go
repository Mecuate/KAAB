package handlers

import "kaab/src/models"

var RoutesDirectory = models.StaticPaths{
	LOGIN:     "/user-portal/login",
	LOGOUT:    "/user-portal/logout",
	VERSION:   "/core-engine/version",
	ANALYTICS: "/stats",
}
