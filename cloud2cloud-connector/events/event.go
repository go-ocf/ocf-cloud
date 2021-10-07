package events

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/plgd-dev/go-coap/v2/message"

	"github.com/plgd-dev/kit/v2/codec/cbor"
	"github.com/plgd-dev/kit/v2/codec/json"
)

const CorrelationIDKey = "Correlation-ID"
const SubscriptionIDKey = "Subscription-ID"
const ContentTypeKey = "Content-Type"
const AcceptKey = "Accept"
const EventTypeKey = "Event-Type"
const SequenceNumberKey = "Sequence-Number"
const EventTimestampKey = "Event-Timestamp"
const EventSignatureKey = "Event-Signature"
const AcceptEncodingKey = "Accept-Encoding"
const ContentEncodingKey = "Content-Encoding"

var ContentType_JSON = message.AppJSON.String()
var ContentType_VNDOCFCBOR = message.AppOcfCbor.String()

type EventHeader struct {
	CorrelationID   string
	ID              string
	ContentType     string
	EventType       EventType
	SequenceNumber  uint64
	EventTimestamp  time.Time
	EventSignature  string
	AcceptEncoding  []string
	ContentEncoding string
}

func ParseEventHeader(r *http.Request) (h EventHeader, _ error) {
	correlationID := r.Header.Get(CorrelationIDKey)
	subscriptionID := r.Header.Get(SubscriptionIDKey)
	if subscriptionID == "" {
		return h, fmt.Errorf("invalid " + SubscriptionIDKey)
	}
	eventType := EventType(r.Header.Get(EventTypeKey))
	switch eventType {
	case EventType_ResourceChanged,
		EventType_ResourcesPublished, EventType_ResourcesUnpublished,
		EventType_DevicesOnline, EventType_DevicesOffline, EventType_DevicesRegistered, EventType_DevicesUnregistered,
		EventType_SubscriptionCanceled:
	default:
		return h, fmt.Errorf("invalid "+EventTypeKey+"(%v)", eventType)
	}

	contentType := r.Header.Get(ContentTypeKey)
	switch contentType {
	case "":
		switch eventType {
		case EventType_SubscriptionCanceled:
		default:
			return h, fmt.Errorf("invalid " + ContentTypeKey)
		}
	case ContentType_JSON:
	case ContentType_VNDOCFCBOR:
	default:
		return h, fmt.Errorf("invalid "+ContentTypeKey+"(%v)", contentType)
	}

	seqNum := r.Header.Get(SequenceNumberKey)
	if seqNum == "" {
		return h, fmt.Errorf("invalid " + SequenceNumberKey)
	}
	sequenceNumber, err := strconv.ParseUint(seqNum, 10, 64)
	if err != nil {
		return h, fmt.Errorf("invalid "+SequenceNumberKey+"(%v): %w", seqNum, err)
	}

	evTimestamp := r.Header.Get(EventTimestampKey)
	if evTimestamp == "" {
		return h, fmt.Errorf("invalid " + EventTimestampKey)
	}
	eventTimestamp, err := strconv.ParseInt(evTimestamp, 10, 64)
	if err != nil {
		return h, fmt.Errorf("invalid "+EventTimestampKey+"(%v): %w", evTimestamp, err)
	}
	eventSignature := r.Header.Get(EventSignatureKey)
	if eventSignature == "" {
		return h, fmt.Errorf("invalid " + EventSignatureKey)
	}

	contentEncoding := r.Header.Get(ContentEncodingKey)

	var acceptEncoding []string
	v := r.Header.Get(AcceptEncodingKey)
	if r.Method == "POST" && v != "" {
		acceptEncoding = strings.Split(v, ",")
		if len(acceptEncoding) != 1 {
			return h, fmt.Errorf("invalid "+AcceptEncodingKey+"(%+v): more than 1", acceptEncoding)
		}
	}

	return EventHeader{
		CorrelationID:   correlationID,
		ID:              subscriptionID,
		ContentType:     contentType,
		EventType:       eventType,
		SequenceNumber:  sequenceNumber,
		EventTimestamp:  time.Unix(eventTimestamp, 0),
		EventSignature:  eventSignature,
		ContentEncoding: contentEncoding,
		AcceptEncoding:  acceptEncoding,
	}, nil
}

func (h EventHeader) GetContentDecoder() (func(w []byte, v interface{}) error, error) {
	var decoder func(w []byte, v interface{}) error
	switch h.ContentType {
	case ContentType_JSON:
		decoder = json.Decode
	case ContentType_VNDOCFCBOR:
		decoder = cbor.Decode
	}
	if decoder == nil {
		return nil, fmt.Errorf("%v decoder not found", h.ContentType)
	}

	return decoder, nil
}

func CalculateEventSignature(secret, contentType string, eventType EventType, subscriptionID string, seqNum uint64, timeStamp time.Time, body []byte) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(contentType))
	hash.Write([]byte(":"))
	hash.Write([]byte(eventType))
	hash.Write([]byte(":"))
	hash.Write([]byte(subscriptionID))
	hash.Write([]byte(":"))
	hash.Write([]byte(strconv.FormatUint(seqNum, 10)))
	hash.Write([]byte(":"))
	hash.Write([]byte(strconv.FormatInt(timeStamp.Unix(), 10)))
	hash.Write([]byte(":"))
	hash.Write([]byte(body))
	return hex.EncodeToString(hash.Sum(nil))
}
