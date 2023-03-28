package utility

import (
	"app/domain/entity"
	"app/infra/googlemap"

	// . "app/utility/logger"
	"regexp"
	"strings"
)

func FormatAddressFromGoogleMap(addressComps []googlemap.AddressComponents) *entity.Address {
	var addr entity.Address

	// address_componentsをループして、Stripe APIのAddressにフォーマットする
	var line1Head string
	var line1Arr [3]string
	for _, comp := range addressComps {
		switch comp.Types[0] {
		case "street_number", "sublocality_level_2":
			line1Head = comp.LongName
		case "route", "sublocality_level_3":
			line1Arr[0] = comp.LongName
		case "sublocality_level_4":
			line1Arr[1] = comp.LongName
		case "premise":
			numRegex := regexp.MustCompile(`^[0-9０-９]+$`)
			if numRegex.MatchString(comp.LongName) {
				line1Arr[2] = comp.LongName
			} else {
				addr.Line2 = &comp.LongName
			}
		case "":
		case "locality":
			addr.City = &comp.LongName
		case "administrative_area_level_1":
			addr.State = &comp.ShortName
		case "postal_code":
			addr.PostalCode = &comp.LongName
		case "country":
			addr.Country = &comp.LongName
			addr.CountryCode = &comp.ShortName
		}
	}
	addr.Line1 = &line1Head
	for i, s := range line1Arr {
		if s == "" {
			continue
		}
		if i > 0 {
			*addr.Line1 += "-"
		}
		*addr.Line1 += s
	}

	// addr.Line1とaddr.Line2を分割する
	if idx := strings.Index(*addr.Line1, ","); idx != -1 {
		*addr.Line2 = (*addr.Line1)[idx+2:]
		*addr.Line1 = (*addr.Line1)[:idx]
	}

	return &addr
}
