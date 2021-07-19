package main

import (
	"fmt"
	"time"

	go_proto "github.com/davidae/proto-definition/proto/go"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	var (
		from          = "David"
		street        = "Fakestreet 123"
		postCode      = uint32(123)
		firstName     = "Ola"
		LastName      = "Nordmann"
		useForBilling = false
	)

	greetValueWithPerson := go_proto.GreetingValue{
		From:                from,
		To:                  "",
		LastUpdated:         timestamppb.New(time.Now()),
		PublisherCategories: []int32{1, 2, 3},
		Address: &go_proto.AddressValue{
			Street:        street,
			PostCode:      postCode,
			Person:        &go_proto.PersonValue{FirstName: firstName, LastName: LastName},
			UseForBilling: useForBilling,
		},
	}

	greetValueWithoutPerson := go_proto.GreetingValue{
		From:                from,
		To:                  "",
		LastUpdated:         timestamppb.New(time.Now()),
		PublisherCategories: []int32{1, 2, 3},
		Address: &go_proto.AddressValue{
			Street:        street,
			PostCode:      postCode,
			UseForBilling: useForBilling,
		},
	}

	outValueWithPerson, err := proto.Marshal(&greetValueWithPerson)
	if err != nil {
		panic(err)
	}

	var valueToPointerWithPerson go_proto.Greeting
	if err := proto.Unmarshal(outValueWithPerson, &valueToPointerWithPerson); err != nil {
		panic(err)
	}

	fmt.Printf("valueToPointerWithPerson: %#v\n", valueToPointerWithPerson.String())
	fmt.Printf("is To nil? %t, is UseForBilling nil? %t\n", valueToPointerWithPerson.To == nil, valueToPointerWithPerson.Address.UseForBilling == nil)

	outValueWithoutPerson, err := proto.Marshal(&greetValueWithoutPerson)
	if err != nil {
		panic(err)
	}

	var valueToPointerWithoutPerson go_proto.Greeting
	if err := valueToPointerWithoutPerson.UnmarshalVT(outValueWithoutPerson); err != nil {
		panic(err)
	}

	fmt.Printf("valueToPointerWithoutPerson: %#v\n", valueToPointerWithoutPerson.String())
	fmt.Printf("is To nil? %t, is UseForBilling nil? %t\n", valueToPointerWithPerson.To == nil, valueToPointerWithPerson.Address.UseForBilling == nil)
	fmt.Printf("is person nil? %t\n", valueToPointerWithoutPerson.Address.Person == nil)
}
