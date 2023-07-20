package dao

import (
	"context"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/influx"
	"github.com/bad-superman/test/logging"
	// influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	// "github.com/influxdata/influxdb-client-go/v2/api/write"
)

var (
	INFLUXDB_TOKEN = "dzkAyPe-78NUZXmIhOL5WQP_aL2jgf-8giUPfif-hCbKczT39OMEYIJFWfgTnvdXV7POjKyjXb1_VMBqRMPJGQ=="
	INFLUXDB_URL   = "https://eu-central-1-1.aws.cloud2.influxdata.com"
	// 类似db
	INFLUXDB_ORG    = "test"
	INFLUXDB_BUCKET = "test"
)

// influx cloud 客户端
type InfluxDB struct {
	client *influx.Client
}

func NewInfluxDB() *InfluxDB {
	client, err := influx.New(influx.Configs{
		HostURL:   INFLUXDB_URL,
		AuthToken: INFLUXDB_TOKEN,
	})
	if err != nil {
		panic(err)
	}
	return &InfluxDB{
		client: client,
	}
}

// 添加记录
// measurement 类似sql中的表名
func (i *InfluxDB) WritePoint(measurement string, fields map[string]interface{}, tags map[string]string, ts time.Time) error {
	point := influx.NewPoint(measurement, tags, fields, ts)
	return i.WritePoints([]*influx.Point{point})
}

// 批量添加记录
// measurement 类似sql中的表名
func (i *InfluxDB) WritePoints(points []*influx.Point) error {
	if err := i.client.WritePoints(context.Background(), INFLUXDB_BUCKET, points...); err != nil {
		logging.Errorf("InfluxDB|WritePoint error,err:%v", err)
		return err
	}
	return nil
}
