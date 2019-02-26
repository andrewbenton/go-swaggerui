Embedded Swagger UI for Go
==========================

This library exposes a simple parameterized `http.Handler` that will display
the Swagger UI.  All resources needed to run a standalone swagger-ui are
packaged into a static filesystem and deployed with the library, which means
that your documentation can be visually served alongside your application.

Example
-------

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/andrewbenton/go-swaggerui"
)

func main() {
	swaggerJSON := "..."

	router := mux.NewRouter()
	router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(swaggerJSON))
	})
	router.PathPrefix("/swagger-ui").Handler(swaggerui.Handle(swaggerui.Config{
		AppName: "swagger-demo",
		SwaggerURL: "/swagger.json",
		RootUIPath: "/swagger-ui",
	}))

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}
}
```

Credits
-------

When looking for something to do what this library does, I found a
[post](https://www.ribice.ba/serving-swaggerui-golang/) from
[ribice](https://github.com/ribice) that helped me get started.

Implementation of the UI is thanks to [Swagger UI](https://github.com/swagger-api/swagger-ui).
