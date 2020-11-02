package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"rpc/db_server/internal/config"
	"rpc/db_server/internal/models"
)

func main() {
	rpc.Register(new(models.CompanyServer))
	rpc.Register(new(models.EmployeeServer))
	rpc.HandleHTTP()

	cnf := config.NewHttp()
	l, e := net.Listen(cnf.Proto, cnf.Port)
	fmt.Fprintf(os.Stderr, "Error : %v\n", e)
	http.Serve(l, nil)
}
