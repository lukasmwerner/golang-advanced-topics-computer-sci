# Week 8

Goal: To complete chapter 7 in the GOPL book and to start thinking on what my project should be.

Status: I completed chapter 7 which was about interfaces and consuming interfaces. The coolest program I wrote was an improvement on the surface program from chapter 3. 
```go
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"

	"gopl.io/ch7/eval"
)

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func main() {
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})

}

// The following code has only been marginaly modified from ch3 to allow an io.Writer and a function parameter
const (
	width, height = 600, 300
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func surface(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill white; stroke-width: 0.7;'"+" width='%d' heigh='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' />\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int, f func(x, y float64) float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := width/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

```

What it does is spawn an http server which adds a function handler to the `/plot` route. It uses a math parser written earlier in the chapter which is recursive in nature. The function passed into `?expr=` gets passed as the function to the surface program from the chapter 3 program "surface" which plots a 2d function in a "3d" surface using xml and an isometric perspective. The function `surface` was a modified version of the original `main` function. It had an `io.Writer` and a `f func(x, y float64) float64` parameter added to allow it to "print" to the `http.ResponseWriter` which implements the `io.Writer` which is why I can use `io.Writer` instead of a `http.ResponseWriter` as the parameter type, the second parameter is the parsed function from earlier in the chapter which is a math parser.

Here are some example functions to try out!
 * `pow(2,sin(y))*pow(2,sin(x))/12`
 * `sin(x*y/10)/10`
 * `sin(-x)*pow(1.5,-r)`