package main

import (
	"backend/services/dbService"
	"backend/services/routerService"
)

func main() {
	dbService.Init()
	routerService.Init()
}
