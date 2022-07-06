package main

import (
	"github.com/GustavoFreitas22/the_guardian/config"
	"github.com/GustavoFreitas22/the_guardian/routes"
)

func main() {
	config.DbConnection()
	routes.HandleRoutes()
}
