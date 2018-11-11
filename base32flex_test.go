package base32flex

import (
	"encoding/base32"
	"reflect"
	"strings"
	"testing"
)

func Test_base32flex(t *testing.T) {
	tests := []struct {
		name string
		data string
	}{
		{"", "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalBytes, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(tt.data)
			if err != nil {
				t.Fatal(err)
			}

			// LowerEncoding
			lowerString := LowerEncoding.EncodeToString(originalBytes)
			if strings.ContainsAny(lowerString, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
				t.Fatal("LowerEncoding has problem.")
			}
			if strings.ContainsAny(lowerString, "01lo") {
				t.Fatal("LowerEncoding has problem.")
			}

			lowerBytes, err := LowerEncoding.DecodeString(lowerString)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(originalBytes, lowerBytes) {
				t.Fatal("LowerEncoding has problem.")
			}

			// UpperEncoding
			upperString := UpperEncoding.EncodeToString(originalBytes)
			if strings.ContainsAny(upperString, "abcdefghijklmnopqrstuvwxyz") {
				t.Fatal("UpperEncoding has problem.")
			}
			if strings.ContainsAny(upperString, "01IO") {
				t.Fatal("UpperEncoding has problem.")
			}

			upperBytes, err := UpperEncoding.DecodeString(upperString)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(originalBytes, upperBytes) {
				t.Fatal("UpperEncoding has problem.")
			}
		})
	}
}
