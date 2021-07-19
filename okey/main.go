package main

import (
	"fmt"
	"time"

	go_proto "github.com/davidae/proto-definition/proto/go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	from := "David"
	to := ""
	greet := go_proto.Greeting{
		From:                from,
		To:                  &to,
		LastUpdated:         timestamppb.New(time.Now()),
		PublisherCategories: []int32{1, 2, 3},
	}

	// greetSimple := go_proto.GreetingSimple{
	// 	From:                from,
	// 	To:                  &to,
	// 	LastUpdated:         time.Now().Unix(),
	// 	PublisherCategories: []int32{1, 2, 3},
	// }

	data, err := greet.MarshalVT()
	if err != nil {
		panic(err)
	}

	greeterVT := &go_proto.Greeting{}
	if err := greeterVT.UnmarshalVT(data); err != nil {
		panic(err)
	}
	fmt.Printf("getting with UnmarshalVT: %s\n", greeterVT.String())

	greeterBasic := &go_proto.Greeting{}
	if err := proto.Unmarshal(data, greeterBasic); err != nil {
		panic(err)
	}
	fmt.Printf("getting with basic: %s\n", greeterVT.String())

	// wrap := gogo_proto.GreetingWrapp{
	// 	Baz: &gogo_proto.GreetingWrapp_BazValue{BazValue: 1},
	// }
	// fmt.Printf("wrap: %#v\n", wrap)

	reflect := greet.ProtoReflect()
	fmt.Println("protoreflect start")
	reflect.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		fmt.Printf("jsonName: %s, cardinality: %s, presence: %t\n", fd.JSONName(), fd.Cardinality().GoString(), fd.HasPresence())
		fmt.Printf("  isValid? %t, value: %#v\n", v.IsValid(), v.Interface())
		return true
	})
	fmt.Println("protoreflect end")
}
