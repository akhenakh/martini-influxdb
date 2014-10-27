martini-influxdb
================

A simple Martini logger to influxDB

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

Then query for `code` as status code, `duration`, `url` and `method`

```sql
SELECT duration FROM resp_time WHERE code = 200
```

Note: This is using the REST api on every requests which is probably not what you want on heavy traffic but fun enough to play with InfluxDB on small project.

Note: Remember, you should forget Martini see http://blog.codegangsta.io/blog/2014/05/19/my-thoughts-on-martini/

![demo](https://github.com/akhenakh/martini-influxdb/raw/master/img/graph.png)
