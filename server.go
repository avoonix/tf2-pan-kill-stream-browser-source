package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const html = `
<html>
	<head>
		<style>
			body {
				font-size: 50px;
				font-family: sans-serif;
			}
		</style>
	</head>
	<body>
		<div id='counter'>PAN_KILLS</div>
		<script>
			setInterval(function() {
				window.location.reload();
			}, 3000)
		</script>
	</body>
</html>
`

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatalf("Invalid number of arguments - usage: %s <steam username regex> <path to console.log>", os.Args[0])
	}
	fmt.Println(args)
	aORb := regexp.MustCompile(args[0] + " killed (.*) with fryingpan.")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		dat, err := ioutil.ReadFile(args[1])
		if err != nil {
			log.Println(err)
		}
		matches := aORb.FindAllStringIndex(string(dat), -1)
		fmt.Fprintf(w, strings.Replace(html, "PAN_KILLS", fmt.Sprint(len(matches)), -1))
	})
	fmt.Printf("Starting server at port 6969\n")
	if err := http.ListenAndServe("127.0.0.1:6969", nil); err != nil {
		log.Fatal(err)
	}
}
