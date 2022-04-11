package util

var SipCode = map[string]int{
	"NORMAL_CLEARING":                200,
	"ORIGINATOR_CANCEL":              487,
	"NO_ANSWER":                      480,
	"NETWORK_OUT_OF_ORDER":           502,
	"NO_ROUTE_DESTINATION":           404,
	"NO_ROUTE_TRANSIT_NET":           404,
	"NO_USER_RESPONSE":               408,
	"NORMAL_CIRCUIT_CONGESTION":      503,
	"NORMAL_TEMPORARY_FAILURE":       503,
	"NORMAL_UNSPECIFIED":             480,
	"NUMBER_CHANGED":                 410,
	"OUTGOING_CALL_BARRED":           403,
	"RECOVERY_ON_TIMER_EXPIRE":       504,
	"REDIRECTION_TO_NEW_DESTINATION": 410,
	"REQUESTED_CHAN_UNAVAIL":         503,
	"SERVICE_NOT_IMPLEMENTED":        501,
	"SUBSCRIBER_ABSENT":              480,
	"SWITCH_CONGESTION":              503,
	"UNALLOCATED_NUMBER":             404,
	"USER_BUSY":                      486,
}

var SipStatus = map[int]string{
	200: "NORMAL_CLEARING",
	487: "ORIGINATOR_CANCEL",
	480: "NO_ANSWER",
	502: "NETWORK_OUT_OF_ORDER",
	404: "NO_ROUTE_DESTINATION",
	408: "NO_USER_RESPONSE",
	503: "NORMAL_CIRCUIT_CONGESTION",
	410: "NUMBER_CHANGED",
	403: "OUTGOING_CALL_BARRED",
	504: "RECOVERY_ON_TIMER_EXPIRE",
	501: "SERVICE_NOT_IMPLEMENTED",
	486: "USER_BUSY",
}

var DirectionStatus = map[int]string{
	1: "inbound",
	2: "local",
	3: "outbound",
}
