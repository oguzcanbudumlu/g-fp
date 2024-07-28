package main

import (
	"errors"
	"fmt"
)

// Bad
const (
	BADHOTDOG_PRICE = 4
)

type BadCreditCard struct {
	credit int
}
type BadHotdog struct {
}

func (c *BadCreditCard) charge(amount int) {
	if amount <= c.credit {
		c.credit -= amount
	} else {
		panic("no more credit")
	}
}

func orderBadHotdog(c *BadCreditCard) BadHotdog {
	c.charge(BADHOTDOG_PRICE)
	return BadHotdog{}
}

// Better
type BetterCreditCard struct {
	credit int
}

type BetterHotdog struct {
	price int
}

type CreditError error

type PaymentFn func(BetterCreditCard, int) (BetterCreditCard, CreditError)

func NewBetterCreditCard(initialCredit int) BetterCreditCard {
	return BetterCreditCard{credit: initialCredit}
}

func NewBetterHotdog() BetterHotdog {
	return BetterHotdog{
		price: 4,
	}
}

func OrderBetterHotdog(c BetterCreditCard, pay PaymentFn) (BetterHotdog, func() (BetterCreditCard, CreditError)) {
	hotdog := NewBetterHotdog()
	chargeFn := func() (BetterCreditCard, CreditError) {
		return pay(c, hotdog.price)
	}
	return hotdog, chargeFn
}

var (
	NOT_ENOUGH_CREDIT CreditError = CreditError(errors.New("not enough credit"))
)

func Charge(c BetterCreditCard, amount int) (BetterCreditCard, CreditError) {
	if amount <= c.credit {
		c.credit -= amount
		return c, nil
	}
	return c, NOT_ENOUGH_CREDIT
}

func main() {
	myCard := NewBetterCreditCard(1000)
	hotdog, creditFn := OrderBetterHotdog(myCard, Charge)
	fmt.Println("%+v\n", hotdog)
	newCard, err := creditFn()
	if err != nil {
		panic("User has no credit")
	}
	myCard = newCard
	fmt.Printf("%+v\n", myCard)
}
