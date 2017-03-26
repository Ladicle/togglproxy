package main

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/vzvu3k6k/go-toggl/toggl"
)

func StartTimeEntry(token string, projectID int) {
	c := toggl.NewClient(token)
	if _, err := c.TimeEntries.Start(&toggl.TimeEntry{
		ProjectID:   projectID,
		CreatedWith: "toggl2gcal",
	}); err != nil {
		glog.Error(fmt.Sprintf("Can not start %d time entries: %v", projectID, err))
	}
}

func StopLatestTimeEntry(token string) {
	c := toggl.NewClient(token)
	entry, err := c.TimeEntries.Current()
	if err != nil {
		glog.Warning(fmt.Sprintf("Can not get current time entries: %v", err))
	}
	if _, err = c.TimeEntries.Stop(entry.ID); err != nil {
		glog.Error(fmt.Sprintf("Can not stop time entries: %v", err))
	}
	// TODO: google calendarに登録
}
