package main

import (
	"reflect"
	"testing"
)

func TestValidCouponCodes(t *testing.T) {
	tests := []struct {
		name         string
		code         []string
		businessLine []string
		isActive     []bool
		want         []string
	}{
		{
			name:         "example1",
			code:         []string{"SAVE20", "", "PHARMA5", "SAVE@20"},
			businessLine: []string{"restaurant", "grocery", "pharmacy", "restaurant"},
			isActive:     []bool{true, true, true, true},
			want:         []string{"PHARMA5", "SAVE20"},
		},
		{
			name:         "example2",
			code:         []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"},
			businessLine: []string{"grocery", "electronics", "invalid"},
			isActive:     []bool{false, true, true},
			want:         []string{"ELECTRONICS_50"},
		},
		{
			name:         "business ordering and lexicographic",
			code:         []string{"ZPHARM", "APPLY", "B_ELEC", "AELEC", "DINNER", "BLUNCH"},
			businessLine: []string{"pharmacy", "restaurant", "electronics", "electronics", "restaurant", "grocery"},
			isActive:     []bool{true, true, true, true, true, true},
			want:         []string{"AELEC", "B_ELEC", "BLUNCH", "ZPHARM", "APPLY", "DINNER"},
		},
		{
			name:         "inactive and invalid codes filtered",
			code:         []string{"GOOD", "BAD-CHAR", "", "OK_UNDERSCORE"},
			businessLine: []string{"grocery", "grocery", "grocery", "restaurant"},
			isActive:     []bool{false, true, true, true},
			want:         []string{"OK_UNDERSCORE"},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := validateCoupons(tc.code, tc.businessLine, tc.isActive)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("validateCoupons() = %v, want %v", got, tc.want)
			}
		})
	}
}
