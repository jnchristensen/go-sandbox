package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func incrementCookie(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visitCount")
	if err == http.ErrNoCookie {
		// initialize cookie
		cookie = &http.Cookie{
			Name:  "visits",
			Value: "0",
		}
	}
	count, err := strconv.Atoi(cookie.Value)

	if err != nil {
		log.Fatal(err)
	}

	count++

	http.SetCookie(res, &http.Cookie{
		Name:  "visitCount",
		Value: strconv.Itoa(count),
	})

	io.WriteString(res, "Website visits: "+strconv.Itoa(count))
}

func main() {
	http.HandleFunc("/visit", incrementCookie)
	http.ListenAndServe(":8080", nil)
}
