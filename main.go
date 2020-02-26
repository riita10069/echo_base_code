package main

import (
	"github.com/riita10069/echo_base_code/router"
)

func main() {
	r := router.New()
  r.Logger.Fatal(r.Start(":8000"))
}
