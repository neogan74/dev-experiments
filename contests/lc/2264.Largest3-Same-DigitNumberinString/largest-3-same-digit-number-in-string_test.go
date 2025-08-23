package largest3samedigitnumberinstring

import "testing"

func Test_largestGoodInteger(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				num: "6777133339",
			},
			want: "777",
		},
		{
			name: "Test 2",
			args: args{
				num: "2300019",
			},
			want: "000",
		},
		{
			name: "Test 3",
			args: args{
				num: "42352338",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestGoodInteger(tt.args.num); got != tt.want {
				t.Errorf("largestGoodInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
