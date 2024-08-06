package database

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/marveloper-8/real-time-analytics-dashboard-monorepo/pkg/utils"
)

type InfluxDB struct {
	Client	influxdb2.Client
	Org			string
	Bucket	string
	Logger	*utils.Logger
}

func NewInfluxDB(url, token, org, bucket string, logger *utils.Logger) (*InfluxDB, error)  {
	client := influxdb2.NewClient(url, token)

	_, err := client.Ready(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to InfluxDB: %w", err)
	}

	return &InfluxDB{
		Client: client,
		Org: org,
		Bucket: bucket,
		Logger: logger,
	}, nil
}

func (db *InfluxDB) Close()  {
	db.Client.Close()
}

func (db *InfluxDB) WritePoint(measurement string, tags map[string]string, fields map[string]interface{}) error  {
	writeAPI := db.Client.WriteAPI(db.Org, db.Bucket)
	p := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	writeAPI.WritePoint(p)
	writeAPI.Flush()
	return nil
}

func (db *InfluxDB) Query(query string) ([]map[string]interface{}, error) {
    queryAPI := db.Client.QueryAPI(db.Org)
    result, err := queryAPI.Query(context.Background(), query)
    if err != nil {
        return nil, fmt.Errorf("query error: %w", err)
    }

    var data []map[string]interface{}
    for result.Next() {
        values := result.Record().Values()
        data = append(data, values)
    }

    if result.Err() != nil {
        return nil, fmt.Errorf("query result error: %w", result.Err())
    }

    return data, nil
}