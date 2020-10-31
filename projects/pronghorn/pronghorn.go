package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/russross/blackfriday"
)

var port = flag.String("Port", "8080", "Sets the port that the server runs on")
var host = flag.String("Host", "0.0.0.0", "Set the listening host device")
var mdCSS = flag.String("CSS", "/._builtins/modest.css", "Sets the css file for the markdown rendered files")

func main() {
	// TODO
	flag.Parse()
	pronghorn := Pronghorn{
		Port:  *port,
		Host:  *host,
		MDCSS: *mdCSS,
	}
	fmt.Printf("Listening on %s\n", pronghorn.Addr())
	log.Fatal(http.ListenAndServe(pronghorn.Addr(), pronghorn))
}

// Pronghorn is the basic server type which contains the server config
type Pronghorn struct {
	Port  string
	Host  string
	MDCSS string
}

// Addr returns the address string based on the Port and Host
func (ph Pronghorn) Addr() string {
	return fmt.Sprintf("%s:%s", ph.Host, ph.Port)
}

// ServeHTTP is the http handler for the server
func (ph Pronghorn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serverHeader(w)
	base := path.Base(r.URL.Path)
	relativePath := "." + r.URL.Path

	if strings.HasPrefix(r.URL.Path, "/._builtins") {
		if strings.HasSuffix(base, "retro.css") {
			writeContentType(w, ".css")
			fmt.Fprint(w, MDCSS_Retro)
			return
		} else if strings.HasSuffix(base, "modest.css") {
			writeContentType(w, ".css")
			fmt.Fprint(w, MDCSS_Modest)
			return
		}
	}

	if strings.HasSuffix(base, ".md") {

		file, err := ioutil.ReadFile(relativePath)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "pronghorn error on %v: %v", relativePath, err)
			return
		}

		var htmlRenderer = blackfriday.HtmlRenderer(blackfriday.HTML_COMPLETE_PAGE, base, ph.MDCSS)
		output := blackfriday.Markdown(file, htmlRenderer,
			blackfriday.EXTENSION_AUTOLINK+
				blackfriday.EXTENSION_AUTO_HEADER_IDS+
				blackfriday.EXTENSION_BACKSLASH_LINE_BREAK+
				blackfriday.EXTENSION_DEFINITION_LISTS+
				blackfriday.EXTENSION_FENCED_CODE+
				blackfriday.EXTENSION_FOOTNOTES+
				blackfriday.EXTENSION_HARD_LINE_BREAK+
				blackfriday.EXTENSION_HEADER_IDS+
				blackfriday.EXTENSION_JOIN_LINES+
				blackfriday.EXTENSION_LAX_HTML_BLOCKS+
				blackfriday.EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK+
				blackfriday.EXTENSION_NO_INTRA_EMPHASIS+
				blackfriday.EXTENSION_SPACE_HEADERS+
				blackfriday.EXTENSION_STRIKETHROUGH+
				blackfriday.EXTENSION_TABLES+
				blackfriday.EXTENSION_TITLEBLOCK,
		)
		w.Write(output)
		writeContentType(w, ".html")
		return

	}
	writeContentType(w, base)
	body, err := ioutil.ReadFile(relativePath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "pronghorn error on %v: %v", relativePath, err)
		return
	}
	w.Write(body)
}

func writeContentType(w http.ResponseWriter, base string) {
	if strings.HasSuffix(base, ".css") {
		w.Header().Set("Content-Type", "text/css")
	}
	if strings.HasSuffix(base, ".js") {
		w.Header().Set("Content-Type", "text/javascript")
	}
	if strings.HasSuffix(base, ".html") {
		w.Header().Set("Content-Type", "text/html")
	}
}

func serverHeader(w http.ResponseWriter) {
	w.Header().Set("Server", "pronghorn")
}

var MDCSS_Retro = `pre,
code {
  font-family: Menlo, Monaco, "Courier New", monospace;
}

pre {
  padding: 0.5rem;
  line-height: 1.25;
  overflow-x: scroll;
}

@media print {
  *,
  *:before,
  *:after {
    background: transparent !important;
    color: #000 !important;
    box-shadow: none !important;
    text-shadow: none !important;
  }

  a,
  a:visited {
    text-decoration: underline;
  }

  a[href]:after {
    content: " (" attr(href) ")";
  }

  abbr[title]:after {
    content: " (" attr(title) ")";
  }

  a[href^="#"]:after,
  a[href^="javascript:"]:after {
    content: "";
  }

  pre,
  blockquote {
    border: 1px solid #999;
    page-break-inside: avoid;
  }

  thead {
    display: table-header-group;
  }

  tr,
  img {
    page-break-inside: avoid;
  }

  img {
    max-width: 100% !important;
  }

  p,
  h2,
  h3 {
    orphans: 3;
    widows: 3;
  }

  h2,
  h3 {
    page-break-after: avoid;
  }
}

a,
a:visited {
  color: #01ff70;
}

a:hover,
a:focus,
a:active {
  color: #2ecc40;
}

.retro-no-decoration {
  text-decoration: none;
}

html {
  font-size: 12px;
}

@media screen and (min-width: 32rem) and (max-width: 48rem) {
  html {
    font-size: 15px;
  }
}

@media screen and (min-width: 48rem) {
  html {
    font-size: 16px;
  }
}

body {
  line-height: 1.85;
}

p,
.retro-p {
  font-size: 1rem;
  margin-bottom: 1.3rem;
}

h1,
.retro-h1,
h2,
.retro-h2,
h3,
.retro-h3,
h4,
.retro-h4 {
  margin: 1.414rem 0 0.5rem;
  font-weight: inherit;
  line-height: 1.42;
}

h1,
.retro-h1 {
  margin-top: 0;
  font-size: 3.998rem;
}

h2,
.retro-h2 {
  font-size: 2.827rem;
}

h3,
.retro-h3 {
  font-size: 1.999rem;
}

h4,
.retro-h4 {
  font-size: 1.414rem;
}

h5,
.retro-h5 {
  font-size: 1.121rem;
}

h6,
.retro-h6 {
  font-size: 0.88rem;
}

small,
.retro-small {
  font-size: 0.707em;
}

/* https://github.com/mrmrs/fluidity */

img,
canvas,
iframe,
video,
svg,
select,
textarea {
  max-width: 100%;
}

html,
body {
  background-color: #222;
  min-height: 100%;
}

html {
  font-size: 18px;
}

body {
  color: #fafafa;
  font-family: "Courier New";
  line-height: 1.45;
  margin: 6rem auto 1rem;
  max-width: 48rem;
  padding: 0.25rem;
}

pre {
  background-color: #333;
}

blockquote {
  border-left: 3px solid #01ff70;
  padding-left: 1rem;
}
`

var MDCSS_Modest = `@media print {
	*,
	*:before,
	*:after {
	  background: transparent !important;
	  color: #000 !important;
	  box-shadow: none !important;
	  text-shadow: none !important;
	}
  
	a,
	a:visited {
	  text-decoration: underline;
	}
  
	a[href]:after {
	  content: " (" attr(href) ")";
	}
  
	abbr[title]:after {
	  content: " (" attr(title) ")";
	}
  
	a[href^="#"]:after,
	a[href^="javascript:"]:after {
	  content: "";
	}
  
	pre,
	blockquote {
	  border: 1px solid #999;
	  page-break-inside: avoid;
	}
  
	thead {
	  display: table-header-group;
	}
  
	tr,
	img {
	  page-break-inside: avoid;
	}
  
	img {
	  max-width: 100% !important;
	}
  
	p,
	h2,
	h3 {
	  orphans: 3;
	  widows: 3;
	}
  
	h2,
	h3 {
	  page-break-after: avoid;
	}
  }
  
  pre,
  code {
	font-family: Menlo, Monaco, "Courier New", monospace;
  }
  
  pre {
	padding: 0.5rem;
	line-height: 1.25;
	overflow-x: scroll;
  }
  
  a,
  a:visited {
	color: #3498db;
  }
  
  a:hover,
  a:focus,
  a:active {
	color: #2980b9;
  }
  
  .modest-no-decoration {
	text-decoration: none;
  }
  
  html {
	font-size: 12px;
  }
  
  @media screen and (min-width: 32rem) and (max-width: 48rem) {
	html {
	  font-size: 15px;
	}
  }
  
  @media screen and (min-width: 48rem) {
	html {
	  font-size: 16px;
	}
  }
  
  body {
	line-height: 1.85;
  }
  
  p,
  .modest-p {
	font-size: 1rem;
	margin-bottom: 1.3rem;
  }
  
  h1,
  .modest-h1,
  h2,
  .modest-h2,
  h3,
  .modest-h3,
  h4,
  .modest-h4 {
	margin: 1.414rem 0 0.5rem;
	font-weight: inherit;
	line-height: 1.42;
  }
  
  h1,
  .modest-h1 {
	margin-top: 0;
	font-size: 3.998rem;
  }
  
  h2,
  .modest-h2 {
	font-size: 2.827rem;
  }
  
  h3,
  .modest-h3 {
	font-size: 1.999rem;
  }
  
  h4,
  .modest-h4 {
	font-size: 1.414rem;
  }
  
  h5,
  .modest-h5 {
	font-size: 1.121rem;
  }
  
  h6,
  .modest-h6 {
	font-size: 0.88rem;
  }
  
  small,
  .modest-small {
	font-size: 0.707em;
  }
  
  /* https://github.com/mrmrs/fluidity */
  
  img,
  canvas,
  iframe,
  video,
  svg,
  select,
  textarea {
	max-width: 100%;
  }
  
  @import url(http://fonts.googleapis.com/css?family=Open+Sans+Condensed:300,300italic,700);
  
  @import url(http://fonts.googleapis.com/css?family=Arimo:700,700italic);
  
  html {
	font-size: 18px;
	max-width: 100%;
  }
  
  body {
	color: #444;
	font-family: "Open Sans Condensed", sans-serif;
	font-weight: 300;
	margin: 0 auto;
	max-width: 48rem;
	line-height: 1.45;
	padding: 0.25rem;
  }
  
  h1,
  h2,
  h3,
  h4,
  h5,
  h6 {
	font-family: Arimo, Helvetica, sans-serif;
  }
  
  h1,
  h2,
  h3 {
	border-bottom: 2px solid #fafafa;
	margin-bottom: 1.15rem;
	padding-bottom: 0.5rem;
	text-align: center;
  }
  
  blockquote {
	border-left: 8px solid #fafafa;
	padding: 1rem;
  }
  
  pre,
  code {
	background-color: #fafafa;
  }
  `
