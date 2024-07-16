package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	l "service-hf-voucher-p5/external/logger"
	ps "service-hf-voucher-p5/external/strings"
	"service-hf-voucher-p5/internal/core/application"
	"service-hf-voucher-p5/internal/core/domain/entity"
	"service-hf-voucher-p5/internal/core/domain/entity/dto"
	"strconv"
	"time"
)

type Handler interface {
	HandlerVoucher(rw http.ResponseWriter, req *http.Request)
	HealthCheck(rw http.ResponseWriter, req *http.Request)
}

type handler struct {
	app application.Application
}

func NewHandler(app application.Application) Handler {
	return handler{app: app}
}

func (h handler) HandlerVoucher(rw http.ResponseWriter, req *http.Request) {

	var routes = map[string]http.HandlerFunc{
		"get hermes_foods/voucher":      h.getVoucherByID,
		"post hermes_foods/voucher":     h.saveVoucher,
		"put hermes_foods/voucher/{id}": h.updateVoucherByID,
	}

	handler, err := router(req.Method, req.URL.Path, routes)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h handler) HealthCheck(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "OK"}`))
}

func (h *handler) getVoucherByID(rw http.ResponseWriter, req *http.Request) {
	msgID := l.MessageID(req.Header.Get(l.MessageIDKey))
	id := getID("voucher", req.URL.Path)

	v, err := h.app.GetVoucherByID(msgID, id)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
		return
	}

	if v == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "voucher not found"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(ps.MarshalString(v)))

}

func (h *handler) saveVoucher(rw http.ResponseWriter, req *http.Request) {

	msgID := l.MessageID(req.Header.Get(l.MessageIDKey))
	var buff bytes.Buffer

	var reqVoucher dto.RequestVoucher

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqVoucher); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	percentage, err := strconv.ParseInt(reqVoucher.Percentage, 10, 64)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error parsing percentage: %v"}`, err)
		return
	}

	voucher := entity.Voucher{
		Code:       reqVoucher.Code,
		Percentage: percentage,
	}

	if len(reqVoucher.ExpiresAt) > 0 {
		voucher.ExpiresAt.Value = new(time.Time)
		if err := voucher.ExpiresAt.SetTimeFromString(reqVoucher.ExpiresAt); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
			return
		}
	}

	reqVoucher.ExpiresAt = voucher.ExpiresAt.Format()

	v, err := h.app.SaveVoucher(msgID, reqVoucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(v)))

}

func (h *handler) updateVoucherByID(rw http.ResponseWriter, req *http.Request) {
	msgID := l.MessageID(req.Header.Get(l.MessageIDKey))
	id := getID("voucher", req.URL.Path)

	var buff bytes.Buffer

	var reqVoucher dto.RequestVoucher

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqVoucher); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	v, err := h.app.UpdateVoucherByID(msgID, id, reqVoucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to update voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(v)))
}
