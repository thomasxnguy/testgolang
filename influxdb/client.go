package main

import (
	"log"
	"time"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
)

const (
	DB          = "demo"
	username    = "admin"
	password    = "admin"
	measurement = "measure_test"
)

func main() {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	//Create 10 points
	writePoints(c, 10)

	//Display points
	printPoint(c)

	//Clean up
	tearDownMeasurement(c)
}



func writePoints(clnt client.Client, nb int) {
	sampleSize := nb

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  DB,
		Precision: "ns",
	})
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < sampleSize; i++ {

		fields := map[string]interface{}{
			"up":   i,
			"down": 100.0 - i,
		}

		pt, err := client.NewPoint(
			measurement,
			nil,
			fields,
			time.Now(),
		)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	if err := clnt.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func printPoint(clnt client.Client) error {

	points, err := queryDB(clnt, fmt.Sprintf("SELECT * FROM %s", measurement))

	if err != nil {
		return err
	}

	for i, row := range points[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		up := row[1]
		down := row[2]
		log.Printf("[%2d] %s: up=%v down=%v\n", i, t.Format(time.Stamp), up, down)
	}
	return nil
}

func tearDownMeasurement(clnt client.Client) error {
	_, err := queryDB(clnt, fmt.Sprintf("DROP MEASUREMENT %s", measurement))
	return err
}

func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: DB,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
