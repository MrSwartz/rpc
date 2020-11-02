package std

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

var separator = func() string {
	var sep []byte
	var ch = byte('=')

	for i := 0; i < (2<<6 - 1); i++ {
		sep = append(sep, ch)
	}
	return string(sep)
}

func PrintErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v\n", err)
	}
}

func PrintLogs(err error, body *http.Request, args interface{}, reply interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error  : %v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Value : %v\n", reply)
	fmt.Fprintf(os.Stdout, "Args  : %v\n", args)
	fmt.Fprintf(os.Stdout, "Body  : %v\n", body.Body)
	fmt.Fprintf(os.Stdout, "%v\n", separator())
}

func ApiLog(err error, out *sql.Rows, val interface{}, args interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error        : %v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Rows        : %v\n", out)
	fmt.Fprintf(os.Stdout, "Args        : %v\n", args)
	fmt.Fprintf(os.Stdout, "APIResponse : %v\n", val)
	fmt.Fprintf(os.Stdout, "%v\n", separator())
}
