package controller

import (
	"fmt"
	"net/http"
)

// HelloAction — простой контроллер, возвращающий "Hello World"
func (c *TodoController) HelloAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Todo Container!")
}
