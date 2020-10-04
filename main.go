package main

import (
	"MealMap/Config/Route"
	"fmt"
	"net/http"
)

func main()  {
	fmt.Print("asdfsdf")
	routeTable := Route.SetupRoute()

	// Run
	s := &http.Server{
		Handler:        routeTable,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}