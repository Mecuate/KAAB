package db

const (
	NODES             = "nodes"
	ENDPOINTS         = "endpoints"
	ACCOUNTS          = "accounts"
	MEDIA             = "media"
	SCHEMAS           = "schemas"
	FILES             = "files"
	DATA_ENTRY_EVENTS = "data_entry_events"
	KNOWN_HOST        = "known_host"
	PASSWORDS         = "passwords"
	STATS             = "stats"
	USERS             = "users"
	INSTANCE_INFO     = "instanceInfo"
	ACTIVE            = "ACTIVE"
	SUSPENDED         = "SUSPENDED"
	DELETED           = "DELETED"
)

type STU struct {
	elements []string
}

func (s STU) Contains(target string) bool {
	for _, elem := range s.elements {
		if elem == target {
			return true
		}
	}
	return false
}

func (s STU) IsActive(target string) bool {
	return target == s.elements[0]
}

func (s STU) Activate() string {
	return ACTIVE
}

func (s STU) Suspend() string {
	return SUSPENDED
}

func (s STU) Delete() string {
	return DELETED
}

var STATUS = STU{[]string{ACTIVE, SUSPENDED, DELETED}}
