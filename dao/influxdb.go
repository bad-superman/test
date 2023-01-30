package dao

import (
	"context"
	"time"

	"github.com/bad-superman/test/logging"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var (
	INFLUXDB_TOKEN = "_NcniwamS9a1rqUtToaDRMuchjWwpcYKaYjk-xq3mUm2UkteFnaTuMLbaTXxN3RWDQuXLH3NLMDrCg6Jsx3CAg=="
	INFLUXDB_URL   = "https://ap-southeast-2-1.aws.cloud2.influxdata.com"
	// 类似db
	INFLUXDB_ORG    = "test"
	INFLUXDB_BUCKET = "test"
)

// influx cloud 客户端
type InfluxDB struct {
	client influxdb2.Client
}

func NewInfluxDB() *InfluxDB {
	return &InfluxDB{
		client: influxdb2.NewClient(INFLUXDB_URL, INFLUXDB_TOKEN),
	}
}

// 添加记录
// measurement 类似sql中的表名
func (i *InfluxDB) WritePoint(measurement string, fields map[string]interface{}, tags map[string]string) error {
	writeAPI := i.client.WriteAPIBlocking(INFLUXDB_ORG, INFLUXDB_BUCKET)

	point := write.NewPoint(measurement, tags, fields, time.Now())
	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		logging.Errorf("InfluxDB|WritePoint error,err:%v", err)
		return err
	}
	return nil
}
