package main

import (
	"bytes"
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

	greet := go_proto.Greeting{
		From:                from,
		To:                  nil,
		LastUpdated:         timestamppb.New(time.Now()),
		PublisherCategories: []int32{1, 2, 3},
		Address: &go_proto.AddressPointer{
			Street:        &street,
			PostCode:      &postCode,
			Person:        &go_proto.PersonPointer{FirstName: &firstName, LastName: &LastName},
			UseForBilling: &useForBilling,
		},
	}

	greetValue := go_proto.GreetingValue{
		From: from,
		// To:                  &to,
		LastUpdated:         timestamppb.New(time.Now()),
		PublisherCategories: []int32{1, 2, 3},
		Address: &go_proto.AddressValue{
			Street:        street,
			PostCode:      postCode,
			Person:        &go_proto.PersonValue{FirstName: firstName, LastName: LastName},
			UseForBilling: useForBilling,
		},
	}

	outPointer, err := proto.Marshal(&greet)
	fmt.Println(outPointer)
	fmt.Println(err)

	outValue, err := proto.Marshal(&greetValue)
	fmt.Println(outValue)
	fmt.Println(err)

	fmt.Printf("string pointer out,string value out:\n%q\n%q\n", string(outPointer), string(outValue))
	fmt.Printf("pointer out equal value out? %t\n", bytes.Compare(outPointer, outValue) == 0)
	fmt.Printf("pointer out equal value out? %t\n", string(outPointer) == string(outValue))

	var valueToPointer go_proto.Greeting
	if err := proto.Unmarshal(outValue, &valueToPointer); err != nil {
		panic(err)
	}

	fmt.Printf("valueToPointer: %#v\n", valueToPointer.String())

	// data, err := greet.MarshalVT()
	// if err != nil {
	// 	panic(err)
	// }

	// greeterVT := &go_proto.Greeting{}
	// if err := greeterVT.UnmarshalVT(data); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("getting with UnmarshalVT: %s\n", greeterVT.String())

	// greeterBasic := &go_proto.Greeting{}
	// if err := proto.Unmarshal(data, greeterBasic); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("getting with basic: %s\n", greeterVT.String())

	// // wrap := gogo_proto.GreetingWrapp{
	// // 	Baz: &gogo_proto.GreetingWrapp_BazValue{BazValue: 1},
	// // }
	// // fmt.Printf("wrap: %#v\n", wrap)

	// reflect := greet.ProtoReflect()
	// fmt.Println("protoreflect start")
	// reflect.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
	// 	fmt.Printf("jsonName: %s, cardinality: %s, presence: %t\n", fd.JSONName(), fd.Cardinality().GoString(), fd.HasPresence())
	// 	fmt.Printf("  isValid? %t, value: %#v\n", v.IsValid(), v.Interface())
	// 	return true
	// })
	// fmt.Println("protoreflect end")
}
