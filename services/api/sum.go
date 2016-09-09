package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	add "github.com/vems/pb/add"

	"golang.org/x/net/context"
)

func sumRequestHandler(c client, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	firstValue := r.URL.Query().Get("first")
	secondValue := r.URL.Query().Get("second")

	if firstValue == "" || secondValue == "" {
		http.Error(w, "Please specify first / second in the query params", http.StatusBadRequest)
		return
	}

	addCh := getSum(c, ctx, firstValue, secondValue)

	addReply := <-addCh
	if err := addReply.err; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p := addReply.v
	json.NewEncoder(w).Encode(p)
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
