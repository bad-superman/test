package dao

import (
	"context"
	"time"

	// "github.com/InfluxCommunity/influxdb3-go/influx"

	"github.com/bad-superman/test/conf"
	"github.com/bad-superman/test/logging"
	influx "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var (
	// local
	INFLUXDB_TOKEN_Local = "tk2UFXxqDNYYvPFnWWoGLQMR9vHGnu5UqvVudyXN8xqYP8XTM46p12OHMUtXLjLcHxQAYx4Zh3Xr6TpQ5lRiNA=="
	INFLUXDB_URL_Local   = "http://127.0.0.1:8086"
	// 类似db
	INFLUXDB_ORG_Local    = "dc"
	INFLUXDB_BUCKET_Local = "test"
)

// influx cloud 客户端
type InfluxDBV2 struct {
	c      *conf.Config
	client influx.Client
}

func NewInfluxDBV2(c *conf.Config) *InfluxDBV2 {
	client := influx.NewClient(c.InfluxConfig.URL, c.InfluxConfig.Token)
	return &InfluxDBV2{
		c:      c,
		client: client,
	}
}

// 添加记录
// measurement 类似sql中的表名
func (i *InfluxDBV2) WritePoint(measurement string, fields map[string]interface{}, tags map[string]string, ts time.Time) error {
	point := write.NewPoint(measurement, tags, fields, time.Now())
	return i.WritePoints([]*write.Point{point})
}

// 批量添加记录
// measurement 类似sql中的表名
func (i *InfluxDBV2) WritePoints(points []*write.Point) error {
	writeAPI := i.client.WriteAPIBlocking(i.c.InfluxConfig.Org, i.c.InfluxConfig.Bucket)
	if err := writeAPI.WritePoint(context.Background(), points...); err != nil {
		logging.Errorf("InfluxDB|WritePoint error,err:%v", err)
		return err
	}
	return nil
}
