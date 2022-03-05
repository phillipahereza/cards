package cards

import "testing"

func TestIsCardNumberValid(t *testing.T) {
	type args struct {
		cardNumber string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty string for card number", args: args{cardNumber: ""}, want: false},
		{name: "card number with very few digits", args: args{cardNumber: "5237 43423"}, want: false},
		{name: "card number with very many digits", args: args{cardNumber: "52374342343232323232323232323233232332323"}, want: false},
		{name: "card number with non integer chars", args: args{cardNumber: "S237 251b 2477 B1ZZ"}, want: false},
		{name: "card number with spaces", args: args{cardNumber: "52 37 251 6 2 477 8 1 3 3"}, want: true},
		{name: "invalid Visa card number", args: args{cardNumber: "4485256453420222"}, want: false},
		{name: "valid Visa card number", args: args{cardNumber: "4024007183198980776"}, want: true},
		{name: "valid Amex card number", args: args{cardNumber: "376683179465306"}, want: true},
		{name: "valid Maestro card number", args: args{cardNumber: "0604928919175879"}, want: true},
		{name: "invalid Maestro card number", args: args{cardNumber: "0604928919175889"}, want: false},
		{name: "invalid AMEX card number", args: args{cardNumber: "376683179465316"}, want: false},
		{name: "valid JCB card number", args: args{cardNumber: "3545209258685463"}, want: true},
		{name: "invalid JCB card number", args: args{cardNumber: "35429600947496920892"}, want: false},
		{name: "invalid Mastercard card number", args: args{cardNumber: "5362590016635917"}, want: false},
		{name: "valid Mastercard card number", args: args{cardNumber: "2720990220245697"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCardNumberValid(tt.args.cardNumber); got != tt.want {
				t.Errorf("IsCardNumberValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardScheme(t *testing.T) {
	type args struct {
		cardNumber string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "valid amex starting with 34", args: args{cardNumber: "347491610166640"}, want: amex},
		{name: "valid amex starting with 37", args: args{cardNumber: "375983475514786"}, want: amex},
		{name: "invalid amex starting with 34", args: args{cardNumber: "34749161016640"}, want: unknown},
		{name: "invalid amex starting with 37", args: args{cardNumber: "37598345514786"}, want: unknown},
		{name: "valid JCB", args: args{cardNumber: "3540390463054243"}, want: jcb},
		{name: "invalid JCB", args: args{cardNumber: "1140390463054243"}, want: unknown},
		{name: "valid Maestro starting with 6", args: args{cardNumber: "6759649826438453"}, want: maestro},
		{name: "valid Maestro starting with 50", args: args{cardNumber: "5038433578905422"}, want: maestro},
		{name: "valid Maestro starting with 56-58", args: args{cardNumber: "58909798440638784"}, want: maestro},
		{name: "invalid Maestro", args: args{cardNumber: "589097984406387844323232"}, want: unknown},
		{name: "valid visa", args: args{cardNumber: "4719668336571185"}, want: visa},
		{name: "valid visa with 13 digit number", args: args{cardNumber: "4929211035486"}, want: visa},
		{name: "invalid visa with 12 digit number", args: args{cardNumber: "429211035486"}, want: unknown},
		{name: "valid mastercard with 51-55 range", args: args{cardNumber: "5210673913091665"}, want: mastercard},
		{name: "valid mastercard with 2221-2720 range", args: args{cardNumber: "2720990220245697"}, want: mastercard},
		{name: "invalid mastercard", args: args{cardNumber: "54936547274580171"}, want: unknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CardScheme(tt.args.cardNumber); got != tt.want {
				t.Errorf("CardScheme() = %q, want %q", got, tt.want)
			}
		})
	}
}