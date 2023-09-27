package server

import "kaab/src/models"

var DirectoryTable = models.StaticPaths{

	LOGIN:     "/user-portal/login",
	LOGOUT:    "/user-portal/logout",
	VERSION:   "/core-engine/version",
	ANALYTICS: "/stats",
}
