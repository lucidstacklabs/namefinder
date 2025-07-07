package dns

import (
	"context"
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
	"strings"
	"time"
)

type Handler struct {
	recordService *RecordService
}

func NewHandler(recordService *RecordService) *Handler {
	return &Handler{
		recordService: recordService,
	}
}

func (h *Handler) Handle(w dns.ResponseWriter, r *dns.Msg) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true

	for _, q := range r.Question {
		log.Printf("query: %s %s", q.Name, dns.TypeToString[q.Qtype])

		recordType := dns.TypeToString[q.Qtype]
		records, err := h.recordService.Query(ctx, q.Name, recordType)

		if err != nil {
			log.Printf("error querying records: %v", err)
			m.Rcode = dns.RcodeServerFailure
			continue
		}

		if len(records) == 0 {
			log.Printf("no records found for %s %s", q.Name, recordType)
			m.Rcode = dns.RcodeNameError
			continue
		}

		for _, record := range records {
			rr := h.createResourceRecord(record, q.Qtype)
			if rr != nil {
				m.Answer = append(m.Answer, rr)
			}
		}
	}

	err := w.WriteMsg(m)

	if err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func (h *Handler) createResourceRecord(record *Record, qtype uint16) dns.RR {
	recordName := strings.TrimSuffix(record.Name, ".")
	recordName = fmt.Sprintf("%s.", recordName)

	header := dns.RR_Header{
		Name:   recordName,
		Rrtype: qtype,
		Class:  dns.ClassINET,
		Ttl:    record.TTL,
	}

	switch qtype {
	case dns.TypeA:
		if ip := net.ParseIP(record.Value); ip != nil {
			return &dns.A{
				Hdr: header,
				A:   ip,
			}
		}

	case dns.TypeAAAA:
		if ip := net.ParseIP(record.Value); ip != nil {
			return &dns.AAAA{
				Hdr:  header,
				AAAA: ip,
			}
		}

	case dns.TypeCNAME:
		recordValue := strings.TrimSuffix(record.Value, ".")
		recordValue = fmt.Sprintf("%s.", recordValue)
		return &dns.CNAME{
			Hdr:    header,
			Target: recordValue,
		}

	case dns.TypeMX:
		parts := strings.Fields(record.Value)
		if len(parts) >= 2 {
			var priority uint16
			_, err := fmt.Sscanf(parts[0], "%d", &priority)

			if err != nil {
				break
			}

			mx := strings.TrimSuffix(parts[1], ".")
			mx = fmt.Sprintf("%s.", mx)
			return &dns.MX{
				Hdr:        header,
				Preference: priority,
				Mx:         mx,
			}
		}

	case dns.TypeTXT:
		return &dns.TXT{
			Hdr: header,
			Txt: []string{record.Value},
		}

	case dns.TypeNS:
		recordValue := strings.TrimSuffix(record.Value, ".")
		recordValue = fmt.Sprintf("%s.", recordValue)
		return &dns.NS{
			Hdr: header,
			Ns:  recordValue,
		}

	case dns.TypeSOA:
		parts := strings.Fields(record.Value)
		if len(parts) >= 7 {
			var serial, refresh, retry, expire, minimum uint32
			_, err := fmt.Sscanf(parts[2], "%d", &serial)
			if err != nil {
				break
			}

			_, err = fmt.Sscanf(parts[3], "%d", &refresh)
			if err != nil {
				break
			}

			_, err = fmt.Sscanf(parts[4], "%d", &retry)
			if err != nil {
				break
			}

			_, err = fmt.Sscanf(parts[5], "%d", &expire)
			if err != nil {
				break
			}

			_, err = fmt.Sscanf(parts[6], "%d", &minimum)
			if err != nil {
				break
			}

			return &dns.SOA{
				Hdr:     header,
				Ns:      parts[0],
				Mbox:    parts[1],
				Serial:  serial,
				Refresh: refresh,
				Retry:   retry,
				Expire:  expire,
				Minttl:  minimum,
			}
		}
	}

	return nil
}
