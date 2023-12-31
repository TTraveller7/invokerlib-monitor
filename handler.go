package main

import (
	"net/http"

	"github.com/TTraveller7/invokerlib"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	invokerlib.MonitorHandle(w, r)
}
