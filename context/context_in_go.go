package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type key int
type userId int

const userIdKeyctx key = 1

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User-Id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Panicln(err)
	}

	ctx := context.WithValue(r.Context(), userIdKeyctx, userId(idInt))

	result := processLongTask(ctx)

	w.Write([]byte(result))
}

func processLongTask(ctx context.Context) string {
	id := ctx.Value(userIdKeyctx)

	select {
	case <-time.After(3 * time.Second):
		return fmt.Sprintf("done processing %d\n", id)
	case <-ctx.Done():
		log.Println("request canceled")
		return ""
	}
}
