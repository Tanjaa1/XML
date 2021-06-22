package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"user-service/dto"
	tracer "user-service/tracer"
)

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	span := tracer.StartSpanFromContext(ctx, "renderJSON")
	defer span.Finish()

	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func decodeBodyRegisteredUser(ctx context.Context, r io.Reader) (*dto.RequestRegisteredUser, error) {
	span := tracer.StartSpanFromContext(ctx, "decodeBody")
	defer span.Finish()

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var rt dto.RequestRegisteredUser
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}
