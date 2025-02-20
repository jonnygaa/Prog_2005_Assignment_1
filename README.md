# 01-REST-diag

Simple module demonstrating the use of a WebServer and the extraction of HTTP request information.

## Deploying the demo

* Open in IDE of choice and run main.go
	* If not using an IDE, compile the source (Linux: `go build -o diag .` in the source directory; Windows: `go build -o diag.exe .`), and run the executable binary (Linux: `./diag`, Windows: `diag.exe`) afterwards. Remember to add execution permissions if you use Linux (e.g., `chmod 740 diag`).
* Per default the server will be available under http://localhost:8080, with the actual production endpoint under http://localhost:8080/diag.

## Using the demo

Use a HTTP client of choice (e.g., Browser, Postman) to point requests to the server URL. The service then extracts central header and body information from the HTTP Request and returns those to the client where it is displayed in the HTTP response body.

Using Postman (or another configurable HTTP client), explore the variation of headers (e.g., content-type), methods (e.g., POST), and possible content. For further information about different http headers, you can start off using https://en.wikipedia.org/wiki/List_of_HTTP_header_fields.

## Exploration

Explore different requests to explore the invocation of different handlers based on path patterns matching. 
* Examples (to explore redirection):
  * http://localhost:8080/diag/first
  * http://localhost:8080/diag/first/
  * http://localhost:8080/diag/first/someElement/third
  * http://localhost:8080/diag/first/someEleMent/third

Try the first example with the method POST (as opposed to GET) and observe the redirection to the specific handler. Review how the patterns are specified and vary those to observe the effects. 

Details about the routing patterns can be found here: https://go.dev/blog/routing-enhancements

Also review the http package documentation.
