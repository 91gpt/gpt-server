/**
 * @Author Nil
 * @Description api/chat/stream.go
 * @Date 2023/4/26 13:38
 **/

package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Stream(ctx *gin.Context) {
	str := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}
	flusher, ok := ctx.Writer.(http.Flusher)
	if !ok {
		http.Error(ctx.Writer, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Important to make it work in browsers
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")

	for _, i := range str {
		//time.Sleep(time.Second)
		log.Printf("data: %s\n", i)
		fmt.Fprintf(ctx.Writer, "data: %s\n", i)
		flusher.Flush()
	}
}

func writeOutput(w http.ResponseWriter, input string) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Important to make it work in browsers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	log.Printf("data: %s\n", input)
	fmt.Fprintf(w, "data: %s\n", input)
	flusher.Flush()
}
