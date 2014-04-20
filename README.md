martini-influxdb
================

Martini logger to influxDB

Use it in place of the Martini logger:

```golang
// initialize influxdb
conf := &influxdb.ClientConfig{
	Host:     "lecaire.nobugware.com:8086",
	Username: "root",
	Password: "totoin",
	Database: "four",
}
client, err := influxdb.NewClient(conf)
if err != nil {
	log.Fatal(err)
}
m.Use(influxlogger.Logger(client))
```
