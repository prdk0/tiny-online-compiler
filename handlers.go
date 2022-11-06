package main

import (
	"experimental/executor"
	"html/template"
	"net/http"
)

type requestVar struct {
	Lang string
	Code string
}

type responseVar struct {
	Result string
}

var ts = template.Must(template.ParseFiles("/home/pradeek/Projects/Golang/cloud_ide/experimental/home.page.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost != r.Method {
		ts.Execute(w, nil)
		return
	}
	rs := requestVar{
		Lang: r.FormValue("select_lang"),
		Code: r.FormValue("textarea_ide"),
	}
	result := executor.ExecuteRequest(rs.Lang, "sampleProgram", rs.Code)

	resv := responseVar{
		Result: result,
	}

	ts.Execute(w, resv)
}
