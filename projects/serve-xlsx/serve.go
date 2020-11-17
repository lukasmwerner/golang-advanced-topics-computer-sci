package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/tealeg/xlsx"
)

var port = flag.String("Port", "8080", "Sets the port that the server runs on")
var host = flag.String("Host", "0.0.0.0", "Set the listening host device")

func main() {
	flag.Parse()
	http.HandleFunc("/", sheetHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func sheetHandler(w http.ResponseWriter, r *http.Request) {
	var wb, err = xlsx.OpenFile("." + r.URL.Path)
	if err != nil {
		fmt.Fprintln(w, "500 Internal Server Error")
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	var sh = wb.Sheet["index"]
	fmt.Fprintln(w, "<style>table, th, td {   border: 1px solid black;   border-collapse: collapse; }</style>")
	fmt.Fprintln(w, "<table>")
	for r := 0; r < sh.MaxRow; r++ {
		fmt.Fprintln(w, "<tr>")
		for c := 0; c < sh.MaxCol; c++ {
			cell, _ := sh.Cell(r, c)
			if r == 0 {
				fmt.Fprintf(w, "<th>%v</th>\n", cell.Value)
			} else {
				fmt.Fprintf(w, "<td>%v</td>\n", cell.Value)
			}
		}
		fmt.Fprintln(w, "</tr>")
	}
	fmt.Fprintln(w, "</table>")
}
