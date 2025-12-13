package main

import (
	"regexp"
	"sort"
)

// validateCoupons filters valid coupon codes and returns them sorted by business line
// (electronics, grocery, pharmacy, restaurant) and then lexicographically by code.
func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	order := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}
	validCode := regexp.MustCompile(`^[A-Za-z0-9_]+$`)

	type coupon struct {
		code string
		biz  string
	}

	var coupons []coupon
	for i := range code {
		biz := businessLine[i]
		if !isActive[i] {
			continue
		}
		if _, ok := order[biz]; !ok {
			continue
		}
		if code[i] == "" || !validCode.MatchString(code[i]) {
			continue
		}
		coupons = append(coupons, coupon{code: code[i], biz: biz})
	}

	sort.Slice(coupons, func(i, j int) bool {
		if order[coupons[i].biz] == order[coupons[j].biz] {
			return coupons[i].code < coupons[j].code
		}
		return order[coupons[i].biz] < order[coupons[j].biz]
	})

	result := make([]string, len(coupons))
	for i, c := range coupons {
		result[i] = c.code
	}
	return result
}
