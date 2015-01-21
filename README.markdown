# DNSimple API Maintenance application

A super simple application designed to return a maintenance message when the API is offline for maintenance.


## Usage

The app is designed to be deployed to Heroku. Lower the api hostname TTL and change the IP to Heroku to enable the app.


## Example

```
$ PORT=5000 go run main.go

  2015/01/21 15:09:26 Listening on 5000...
```

Any HTTP request will then return

```
$ curl -i localhost:5000/api/v1/domains

  HTTP/1.1 503 Service Unavailable
  Content-Type: application/json; charset=UTF-8
  Date: Wed, 21 Jan 2015 14:09:47 GMT
  Content-Length: 123

  {"message":"Scheduled Maintenance for Jan 24, 04:00 - 05:00 UTC","href":"http://dnsimplestatus.com/incidents/j4l3lshmxmjg"}
```

