package main

import (
	"log"
	"net/http"
	"tm/pkg/contracts"
	"tm/pkg/handlerutil"
)

func (s *Server) reportPeriod(w http.ResponseWriter, r *http.Request) {
	reqBody, ok := handlerutil.BodyJson[contracts.ReportRequest](w, r)
	if !ok {
		return
	}

	res, err := s.tm.ReportPeriod(
		r.Context(),
		reqBody.StartDate,
		reqBody.EndDate,
		reqBody.U100,
	)

	if err != nil {
		log.Println(err.Error())
		s.handleErr(err, w)
		return
	}

	handlerutil.Json(w, res)
	return
}
