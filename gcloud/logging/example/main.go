package main

import (
	"fmt"
	"net/http"
	"runtime"

	lg "github.com/TeamMomentum/go-pkg/gcloud/logging"
)

func main() {
	entry := lg.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		entry.SetHTTPRequest(r)
		entry.SetLabels(map[string]string{
			"key1": "value1",
			"key2": "value2",
		})
		entry.SetLabelFunc("key2", func(v string) string {
			_, _, line, _ := runtime.Caller(3)
			return fmt.Sprintf("this is %s at %d", v, line)
		})

		entry.Debugf("received an HTTP request")

		entry.Infof("finished func")
	})

	port := "8080"
	entry.Infof("PORT=%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		entry.Errorf("serve error: %s", err)
	}
}
