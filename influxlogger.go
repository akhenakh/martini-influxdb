package influxlogger

import (
	"log"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/influxdb/influxdb-go"
)

// Logger returns a middleware handler that logs the request as it goes in and the response as it goes out.
func InfluxDBLogger(client *influxdb.Client) martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
		start := time.Now()
		log.Printf("Started %s %s", req.Method, req.URL.Path)

		rw := res.(martini.ResponseWriter)
		c.Next()
		t := time.Since(start)
		log.Printf("Completed %v %s in %v\n", rw.Status(), http.StatusText(rw.Status()), t)
		if client != nil {
			s := &influxdb.Series{
				Name:    "respTime",
				Columns: []string{"method", "path", "duration"},
				Points: [][]interface{}{
					[]interface{}{req.Proto},
					[]interface{}{req.URL.RequestURI()},
					[]interface{}{time.Millisecond * t},
				},
			}
			err := client.WriteSeries([]*influxdb.Series{s})
			if err != nil {
				log.Println(err)
			}
		}
	}
}
