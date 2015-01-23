# DNSimple Maintenance application

A super simple application designed to return a maintenance message when the API/App are offline for maintenance.


## Usage

The app is designed to be deployed to Heroku. Lower the hostname TTL and change the IP to Heroku to enable the app.


## Example

```
$ PORT=5000 go run main.go

  2015/01/21 15:09:26 Listening on 5000...
```

will result in

```
$ curl -i -H 'Host: api.example' localhost:5000/foo
  HTTP/1.1 503 Service Unavailable
  Content-Type: application/json; charset=UTF-8
  Date: Fri, 23 Jan 2015 19:50:05 GMT
  Content-Length: 123

  {"message":"Scheduled Maintenance for Jan 24, 04:00 - 05:00 UTC","href":"http://dnsimplestatus.com/incidents/j4l3lshmxmjg"}


$curl -i -H 'Host: example' localhost:5000/foo

  HTTP/1.1 503 Service Unavailable
  Content-Type: application/html; charset=UTF-8
  Date: Fri, 23 Jan 2015 19:50:12 GMT
  Content-Length: 168

  <h1>Scheduled Maintenance for Jan 24, 04:00 - 05:00 UTC</h1><p>Follow the updates at the <a href='http://dnsimplestatus.com/incidents/j4l3lshmxmjg'>status site</a></p>
```
