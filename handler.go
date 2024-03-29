package main

import (
	"net/http"

	"github.com/TTraveller7/invokerlib/pkg/api"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	api.MonitorHandle(w, r)
}
