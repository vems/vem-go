package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	add "github.com/vems/pb/add"

	"golang.org/x/net/context"
)

type SumReply struct {
	Error string `json:"error,omitempty"`
	Data  int64  `json:"data,omitempty"`
}

func sumRequestHandler(c client, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	firstValue := r.URL.Query().Get("first")
	secondValue := r.URL.Query().Get("second")

	if firstValue == "" || secondValue == "" {
		sumErrorReply := SumReply{
			Error: "Please specify first / second in the query params",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(sumErrorReply)
		return
	}

	addCh := getSum(c, ctx, firstValue, secondValue)

	addReply := <-addCh
	if err := addReply.err; err != nil {
		sumErrorReply := SumReply{
			Error: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(sumErrorReply)
		return
	}

	p := addReply.v
	sumReply := SumReply{
		Data: p,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sumReply)
}

type sumResults struct {
	v   int64
	err error
}

func getSum(c client, ctx context.Context, first string, second string) chan sumResults {
	ch := make(chan sumResults, 1)

	go func() {
		firstInt, err := strconv.ParseInt(first, 10, 64)
		if err != nil {
			ch <- sumResults{err: err}
			return
		}

		secondInt, err := strconv.ParseInt(second, 10, 64)
		if err != nil {
			ch <- sumResults{err: err}
			return
		}

		res, err := c.Sum(ctx, &add.SumRequest{
			A: firstInt,
			B: secondInt,
		})
		ch <- sumResults{res.V, err}
	}()

	return ch
}
