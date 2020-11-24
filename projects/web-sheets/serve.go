package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/tealeg/xlsx"
)

var port = flag.String("Port", "8080", "Sets the port that the server runs on")
var host = flag.String("Host", "0.0.0.0", "Set the listening host device")

func main() {
	flag.Parse()
	http.HandleFunc("/", sheetHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func writeHTMLSheet(w io.Writer, sh xlsx.Sheet) {
	fmt.Fprintln(w, "<style>table, th, td { border: 1px solid black; border-collapse: collapse; text-align: left; padding: 10px; }</style>")
	fmt.Fprintln(w, "<table>")
	for r := 0; r < sh.MaxRow; r++ {
		fmt.Fprintln(w, "<tr>")
		for c := 0; c < sh.MaxCol; c++ {
			cell, _ := sh.Cell(r, c)
			cellStyle := cell.GetStyle()
			font := cellStyle.Font
			fill := cellStyle.Fill

			fillColor := strings.TrimPrefix(fill.FgColor, "FF")
			textColor := strings.TrimPrefix(font.Color, "FF")

			style := fmt.Sprintf("font-family: '%v'; color: #%v; background: #%v;", font.Name, textColor, fillColor)
			fontJSON, _ := json.Marshal(&font)
			fillJSON, _ := json.Marshal(&fill)
			if r == 0 {
				fmt.Fprintf(w, "<th style=\"%v\" data-font='%v' data-fill='%v'>%v</th>\n", style, string(fontJSON), string(fillJSON), cell.Value)
			} else {
				fmt.Fprintf(w, "<td style=\"%v\" data-font='%v' data-fill='%v'>%v</td>\n", style, string(fontJSON), string(fillJSON), cell.Value)
			}
		}
		fmt.Fprintln(w, "</tr>")
	}
	fmt.Fprintln(w, "</table>")
}

func sheetHandler(w http.ResponseWriter, r *http.Request) {
	var wb, err = xlsx.OpenFile("." + r.URL.Path)
	if err != nil {
		fmt.Fprintln(w, "500 Internal Server Error")
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	sh := wb.Sheet["index"]
	writeHTMLSheet(w, *sh)
}
