package main

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api/write"
)

var (
	INFLUXDB_TOKEN = "_NcniwamS9a1rqUtToaDRMuchjWwpcYKaYjk-xq3mUm2UkteFnaTuMLbaTXxN3RWDQuXLH3NLMDrCg6Jsx3CAg=="
)

func main() {
	token := INFLUXDB_TOKEN
	url := "https://ap-southeast-2-1.aws.cloud2.influxdata.com"
	client := influxdb2.NewClient(url, token)

	org := "test"
	bucket := "test"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 5; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
	}
}
