package main

import (
	"fmt"
	"os"
	"time"

	"github.com/davidae/proto-definition/proto"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LOL struct {
	From        string  `parquet:"name=from, type=BYTE_ARRAY"`
	LastUpdated int64   `parquet:"name=last_updated, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	To          *string `parquet:"name=to, type=BYTE_ARRAY"`
}

func main() {
	from := "David"
	to := "Lia"
	greet := proto.Greeting{
		From:                from,
		To:                  &to,
		LastUpdated:         timestamppb.New(time.Now()),
		PublisherCategories: []int32{1, 2, 3},
	}

	greetSimple := proto.GreetingSimple{
		From:                from,
		To:                  &to,
		LastUpdated:         time.Now().Unix(),
		PublisherCategories: []int32{1, 2, 3},
	}

	// lol := LOL{
	// 	From: "david", To: &to, LastUpdated: time.Now().Unix(),
	// }

	fmt.Println(greet.String())

	var err error
	w, err := os.Create("output/flat.parquet")
	if err != nil {
		panic(err)
	}

	//write
	pw, err := writer.NewParquetWriterFromWriter(w, new(proto.GreetingSimple), 4)
	if err != nil {
		panic(err)
	}
	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY
	if err = pw.Write(greetSimple); err != nil {
		panic(err)
	}

	if err := pw.WriteStop(); err != nil {
		panic(err)
	}

	if err := w.Close(); err != nil {
		panic(err)
	}
}
