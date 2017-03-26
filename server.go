package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/golang/glog"
)

func main() {
	// For glog
	flag.Parse()

	togglToken := os.Getenv("TOGGL_TOKEN")
	if togglToken == "" {
		glog.Error("TOGGL_TOKEN is not set.")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{'version':'v1'}")
	})

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		project := r.URL.Query().Get("project")
		switch project {
		case "engineering":
			StartTimeEntry(togglToken, 13515980)
		case "english":
			StartTimeEntry(togglToken, 29746715)
		case "output":
			StartTimeEntry(togglToken, 24747282)
		case "management":
			StartTimeEntry(togglToken, 24738197)
		case "private":
			StartTimeEntry(togglToken, 31003297)
		default:
			glog.Warning(project + " is not supported")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{'action':'Start'}")
	})

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		StopLatestTimeEntry(togglToken)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{'action':'stop'}")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		glog.Error(fmt.Sprintf("Listen Server is failed: %v", err))
		os.Exit(1)
	}
}
