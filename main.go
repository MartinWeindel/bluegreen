package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func coloring(color string, reloadSecs int) func(http.ResponseWriter, *http.Request) {
	hostname, _ := os.Hostname()
	return func(w http.ResponseWriter, req *http.Request) {
		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")
		s := fmt.Sprintf(`<html>
<head>
  <title>%s</title>
  <meta http-equiv="refresh" content="%d" />
</head>
<body style="background-color:%s;">
  <h2>%s %s</h2>
</body>
</html>`,
			color, reloadSecs, color, hostname, timestamp)
		io.WriteString(w, s)
	}
}

func main() {
	color := os.Getenv("COLOR")
	if color == "" {
		color = "yellow"
	}

	http.HandleFunc("/", coloring(color, 1))
	port := 8080
	fmt.Printf("listening on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
		os.Exit(1)
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
