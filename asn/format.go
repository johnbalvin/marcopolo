package asn

type IpCollection struct {
	Priorities IPRangeWrapper
	Remaining  IPRangeWrapper
}
type IPRangeWrapper struct {
	Quantity uint32
	AsnIDs   []string
	IPs      []IpRange
}

type IpRange struct {
	AsnName           string
	AsnNameInputPrior string
	AsnID             string
	Start             uint32
	End               uint32
	Quantity          uint32
	Priority          int
}

type Asn struct {
	PrioritiesNames []string
	ForbiddenNames  []string
}
