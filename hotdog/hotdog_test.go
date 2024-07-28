package main

import (
	"errors"
	"testing"
)

func TestCharge(t *testing.T) {
	var (
		testChargeStruct = []struct {
			inputCard  BetterCreditCard
			amount     int
			outputCard BetterCreditCard
			err        CreditError
		}{
			{
				BetterCreditCard{1000},
				500,
				BetterCreditCard{500},
				nil,
			},
			{
				BetterCreditCard{20},
				20,
				BetterCreditCard{0},
				nil,
			},
			{
				BetterCreditCard{150},
				1000,
				BetterCreditCard{150},
				NOT_ENOUGH_CREDIT,
			},
		}
	)

	for _, test := range testChargeStruct {
		t.Run("", func(t *testing.T) {
			output, err := Charge(test.inputCard, test.amount)
			if output != test.outputCard || !errors.Is(err, test.err) {
				t.Errorf("expected %v but got %v\n expected %v but got %v\\n\"", test.outputCard, output, test.err, err)
			}
		})
	}
}
