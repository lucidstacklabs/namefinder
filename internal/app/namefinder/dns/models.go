package dns

import (
	"fmt"
	"github.com/miekg/dns"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Record struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	NamespaceID string             `bson:"namespace_id" json:"namespace_id"`
	Name        string             `bson:"name" json:"name"`
	Type        RecordType         `bson:"type" json:"type"`
	Value       string             `bson:"value" json:"value"`
	TTL         uint32             `bson:"ttl" json:"ttl"`
	Class       RecordClass        `bson:"class" json:"class"`
	CreatorType ActorType          `bson:"actor_type" json:"actor_type"`
	CreatorID   string             `json:"creator_id" bson:"creator_id" json:"creator_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type ActorType string

const (
	ActorTypeAdmin  ActorType = "admin"
	ActorTypeApiKey ActorType = "apikey"
)

type RecordClass string

const (
	RecordClassInternet RecordClass = "IN"
)

func GetRecordClass(class dns.Class) (RecordClass, error) {
	switch class {
	case dns.ClassINET:
		return RecordClassInternet, nil
	default:
		return "", fmt.Errorf("record class %v is not supported", class)
	}
}

type RecordType string

const (
	RecordTypeA     RecordType = "A"
	RecordTypeAAAA  RecordType = "AAAA"
	RecordTypeCNAME RecordType = "CNAME"
	RecordTypeMX    RecordType = "MX"
	RecordTypeTXT   RecordType = "TXT"
	RecordTypeSOA   RecordType = "SOA"
	RecordTypeNS    RecordType = "NS"
)

func GetRecordType(dnsType dns.Type) (RecordType, error) {
	switch dnsType {
	case dns.Type(dns.TypeA):
		return RecordTypeA, nil
	case dns.Type(dns.TypeAAAA):
		return RecordTypeAAAA, nil
	case dns.Type(dns.TypeCNAME):
		return RecordTypeCNAME, nil
	case dns.Type(dns.TypeMX):
		return RecordTypeMX, nil
	case dns.Type(dns.TypeTXT):
		return RecordTypeTXT, nil
	case dns.Type(dns.TypeSOA):
		return RecordTypeSOA, nil
	case dns.Type(dns.TypeNS):
		return RecordTypeNS, nil
	default:
		return "", fmt.Errorf("dns type %v is not supported", dnsType)
	}
}
