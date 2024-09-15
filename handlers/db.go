package handlers

import "short_url/database"

var (
	db = database.GetDBConnection()
)