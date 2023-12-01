package handle

import (
	"fmt"
	"net/http"
)

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, WorldGO!")
}
