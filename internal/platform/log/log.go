package log

import (
	"context"
	"fmt"
	"log"

	"github.com/michelpessoa/desafioObservalityOtel/internal/platform/http"
)

func Info(ctx context.Context, msg string, tags ...string) {
	reqId := http.GetRequestID(ctx)
	formattedLog := fmt.Sprintf("[level:INFO] [x-request-id:%s] [msg:%s]", reqId, msg)

	for _, tag := range tags {
		formattedLog = fmt.Sprintf("%s [%s]", formattedLog, tag)
	}

	log.Printf(formattedLog)
}

func Error(ctx context.Context, msg string, err error, tags ...string) {
	reqId := http.GetRequestID(ctx)
	formattedLog := fmt.Sprintf("[level:ERROR] [x-request-id:%s] [msg:%s] [error:%s]", reqId, msg, err.Error())

	for _, tag := range tags {
		formattedLog = fmt.Sprintf("%s [%s]", formattedLog, tag)
	}

	log.Printf(formattedLog)
}

func Tag(k, v any) string {
	return fmt.Sprintf("%v:%v", k, v)
}
