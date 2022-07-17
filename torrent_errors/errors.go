package torrent_errors

type TrackerError struct {
	Code    int
	Message string
}

func (t TrackerError) Error() string {
	return t.Message
}

var (
	InvalidRequestType = TrackerError{
		Code:    100,
		Message: "Invalid request type",
	}
	MissingInfoHash = TrackerError{
		Code:    101,
		Message: "info_hash missing from request",
	}
	MissingPeerID = TrackerError{
		Code:    102,
		Message: "peer_id missing from request",
	}
	MissingPort = TrackerError{
		Code:    103,
		Message: "port missing from request",
	}
	InvalidPort = TrackerError{
		Code:    104,
		Message: "Invalid port",
	}
	InvalidAuth = TrackerError{
		Code:    150,
		Message: "Invalid info hash",
	}
	InvalidPeerID = TrackerError{
		Code:    151,
		Message: "Peer ID invalid",
	}
	InvalidNumWant = TrackerError{
		Code:    152,
		Message: "num_want invalid",
	}
	BadClient = TrackerError{
		Code:    153,
		Message: "Client not whitelisted",
	}
	InofHashNotFound = TrackerError{
		Code:    480,
		Message: "Unknown infohash",
	}
	ClientRequestTooFast = TrackerError{
		Code:    500,
		Message: "Too many requests",
	}
	GenericError = TrackerError{
		Code:    900,
		Message: "Generic",
	}
	MessageMalformedRequest = TrackerError{
		Code:    901,
		Message: "Malformed request",
	}
	QueryParseFail = TrackerError{
		Code:    902,
		Message: "Could not parse request",
	}
)
