package dns

import (
	"github.com/miekg/dns"
	"log"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(w dns.ResponseWriter, r *dns.Msg) {
	log.Print("received dns query")
}
