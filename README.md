# HTTP-Monitoring
implementation of  a service with Golang programming language to monitor
HTTP endpoints so that in some configurable periods (e.g., the 30s, 1m, 5m) this
service sends HTTP requests to the endpoint and logs the response status code
and registers an alert for that URL.
