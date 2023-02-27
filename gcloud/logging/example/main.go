package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	lg "github.com/TeamMomentum/go-pkg/gcloud/logging"
)

func main() {
	lg.SetFlags(1)
	lg.SetPrefix("main: ")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger := lg.New(os.Stderr, "", 0)
		logger.SetFlags(0)
		logger.SetEntryHTTPRequest(r)
		logger.SetEntryLabels(map[string]string{
			"key1": "value1",
			"key2": "value2",
		})
		logger.SetEntryLabelFunc("key2", func(v string) string {
			_, _, line, _ := runtime.Caller(4)
			return fmt.Sprintf("this is %s at %d", v, line)
		})

		logger.Debugf("now: %s", time.Now().Format(time.DateTime))

		logger.Println("finished func")
	})

	lg.SetEntryLabel("KEY1", "VALUE1")
	lg.Printf("port=%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		lg.Errorf("http error: %s", err)
	}
}
