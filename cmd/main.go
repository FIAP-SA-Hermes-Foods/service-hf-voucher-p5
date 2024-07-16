package main

import (
	"context"
	"log"
	"net/http"
	"os"
	l "service-hf-voucher-p5/external/logger"
	voucherrpc "service-hf-voucher-p5/internal/adapters/rpc"
	"service-hf-voucher-p5/internal/core/application"
	httpH "service-hf-voucher-p5/internal/handler/http"

	"github.com/marcos-dev88/genv"
)

func init() {
	if err := genv.New(); err != nil {
		l.Errorf("", "error set envs %v", " | ", err)
	}
}

func main() {

	router := http.NewServeMux()

	ctx := context.Background()

	voucherRPC := voucherrpc.NewVoucherRPC(ctx, os.Getenv("HOST_VOUCHER"), os.Getenv("PORT_VOUCHER"))

	voucherWorkerRPC := voucherrpc.NewVoucherWorkerRPC(ctx, os.Getenv("HOST_VOUCHER"), os.Getenv("PORT_VOUCHER"))

	app := application.NewApplication(ctx, voucherRPC, voucherWorkerRPC)

	h := httpH.NewHandler(app)

	router.Handle("/hermes_foods/voucher/", http.StripPrefix("/", httpH.Middleware(h.HandlerVoucher)))
	router.Handle("/hermes_foods/voucher", http.StripPrefix("/", httpH.Middleware(h.HandlerVoucher)))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_HTTP_PORT"), router))
}
