package main

// TODO: Fix import pathes
import (
	"github.com/kyohei0423/dev-templates/go/echo/app/infra"
	"github.com/kyohei0423/dev-templates/go/echo/app/infra/db"
)

func main() {
	db := db.NewDBconnection()
	defer db.Close()
	infra.Router()
}
