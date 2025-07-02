package main

import (
	"log"
	"net/http"
	"tm/pkg/contracts"
	"tm/pkg/handlerutil"
)

func (s *Server) reportTimeline(w http.ResponseWriter, r *http.Request) {
	reqBody, ok := handlerutil.BodyJson[contracts.TimelineRequest](w, r)
	if !ok {
		return
	}

	res, err := s.tm.ReportTimeline(
		r.Context(),
		reqBody.StartDate,
		reqBody.EndDate,
	)

	if err != nil {
		log.Println(err.Error())
    s.handleErr(err, w)
		return
	}

	handlerutil.Json(w, res)
	return

}
