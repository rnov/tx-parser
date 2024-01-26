package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	jr "tx-parser/internal/parser/handler/request"
	"tx-parser/internal/parser/service"
)

// ParserHandler is a handler that exposes tx parser functionality thorough REST API
type ParserHandler struct {
	parser service.Parser
}

func NewParserHandler(parser service.Parser) *ParserHandler {
	return &ParserHandler{
		parser: parser,
	}
}

func (h *ParserHandler) SubscribeAddress(w http.ResponseWriter, r *http.Request) {
	req := &jr.AddAddressReq{}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Address == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok := h.parser.Subscribe(req.Address)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *ParserHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	resp := h.parser.GetTransactions(address)
	rBody := &jr.GetTxsResp{
		resp,
	}
	body, jsonErr := json.Marshal(rBody)
	if jsonErr != nil {
		// note should log error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(body)
}

func (h *ParserHandler) GetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	resp := h.parser.GetCurrentBlock()
	rBody := &jr.GetBlockResp{
		Decimal: resp,
	}
	body, jsonErr := json.Marshal(rBody)
	if jsonErr != nil {
		// note should log error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(body)
}
