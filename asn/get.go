package asn

import (
	"encoding/csv"
	"errors"
	"marcopolo/utils"
	"os"
	"sort"
	"strings"
)

func (asn Asn) GetIPs(asnPath string) (IpCollection, error) {
	for i := range asn.PrioritiesNames {
		asn.PrioritiesNames[i] = utils.RemoveSpace(strings.ToLower(asn.PrioritiesNames[i]))
	}
	for i := range asn.ForbiddenNames {
		asn.ForbiddenNames[i] = utils.RemoveSpace(strings.ToLower(asn.ForbiddenNames[i]))
	}
	f, err := os.Open(asnPath)
	if err != nil {
		return IpCollection{}, err
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return IpCollection{}, err
	}
	var ipCollection IpCollection
	for _, record := range records[1:] {
		for j := range record {
			record[j] = utils.RemoveSpace(record[j])
		}
		ipStartS := record[0]
		ipEndS := record[1]
		currentASNID := strings.ToLower(record[2])
		currentASNName := strings.ToLower(record[3])
		if strings.Contains(ipStartS, ":") { //ONLY IPV4 Ips
			continue
		}
		if asn.IsForbidden(currentASNName) {
			continue
		}
		priorityIndex, ispriority, priorityName := asn.IsPriority(currentASNName)
		ipStart := utils.IpToUint32(ipStartS)
		ipEnd := utils.IpToUint32(ipEndS)
		countN := ipEnd - ipStart + 1
		ipFull := IpRange{
			Start:             ipStart,
			End:               ipEnd,
			AsnNameInputPrior: priorityName,
			AsnID:             currentASNID,
			AsnName:           currentASNName,
			Quantity:          countN,
			Priority:          priorityIndex,
		}
		if !ispriority {
			ipCollection.Remaining.Quantity += countN
			ipCollection.Remaining.IPs = append(ipCollection.Remaining.IPs, ipFull)
			continue
		}
		ipCollection.Priorities.Quantity += ipFull.Quantity
		ipCollection.Priorities.IPs = append(ipCollection.Priorities.IPs, ipFull)
		var found bool
		for _, asID := range ipCollection.Priorities.AsnIDs {
			if asID == currentASNID {
				found = true
				break
			}
		}
		if !found {
			ipCollection.Priorities.AsnIDs = append(ipCollection.Priorities.AsnIDs, currentASNID)
		}
	}
	sort.Slice(ipCollection.Priorities.IPs, func(i, j int) bool {
		return ipCollection.Priorities.IPs[i].Priority < ipCollection.Priorities.IPs[j].Priority
	})
	for _, specialName := range asn.PrioritiesNames {
		var found bool
		for _, specialData := range ipCollection.Priorities.IPs {
			if specialData.AsnNameInputPrior == specialName {
				found = true
				break
			}
		}
		if !found {
			return IpCollection{}, errors.New("ASN not found or repeated, please recheck your ASN priority List " + specialName)
		}
	}
	return ipCollection, nil
}
