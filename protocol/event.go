package protocol

type Event int

// TODO - Check specification for all events
const (
	ANNOUNCE Event = iota
	STARTED
	STOPPED
	COMPLETE
	PAUSED
)

var eventMapping = map[string]Event{
	"":         ANNOUNCE,
	"started":  STARTED,
	"stopped":  STOPPED,
	"complete": COMPLETE,
	"paused":   PAUSED,
}

var eventToString = map[Event]string{
	ANNOUNCE: "Announce",
	STARTED:  "Started",
	STOPPED:  "Stopped",
	COMPLETE: "Complete",
	PAUSED:   "Paused",
}

func (e Event) String() string {
	return eventToString[e]
}

func EventFromString(raw string) Event {
	val, ok := eventMapping[raw]
	if !ok {
		return ANNOUNCE
	}
	return val
}
