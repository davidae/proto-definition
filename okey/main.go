package main

import (
	"fmt"
	"time"

	greeting_proto "github.com/davidae/proto-definition/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	from := "David"
	to := "Lia"
	greet := greeting_proto.Greeting{
		From:                from,
		To:                  &to,
		LastUpdated:         timestamppb.New(time.Now()),
		PublisherCategories: []int32{1, 2, 3},
	}

	// greetSimple := greeting_proto.GreetingSimple{
	// 	From:                from,
	// 	To:                  &to,
	// 	LastUpdated:         time.Now().Unix(),
	// 	PublisherCategories: []int32{1, 2, 3},
	// }

	data, err := greet.MarshalVT()
	if err != nil {
		panic(err)
	}

	greeterVT := &greeting_proto.Greeting{}
	if err := greeterVT.UnmarshalVT(data); err != nil {
		panic(err)
	}
	fmt.Printf("getting with UnmarshalVT: %s\n", greeterVT.String())

	greeterBasic := &greeting_proto.Greeting{}
	if err := proto.Unmarshal(data, greeterBasic); err != nil {
		panic(err)
	}
	fmt.Printf("getting with basic: %s\n", greeterVT.String())
}
