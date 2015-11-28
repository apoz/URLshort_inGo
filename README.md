# URLshort_inGO

URL shortener exercise written in Go

The goal of this repo is me to practice http development and play a bit with go


# To test the POST with the CURL
curl -X POST -d "{ \"longURL\": \"http://andres-pozo.com/asdfghjkl\"}" http://localhost:8080/new


# API for the URLShortener (and examples)

## How to create a new shortURL
To create a new short URL, a POST http operation has to be sent to the server with the following JSON body:

```
pozo@POZOBOE$curl -v -X POST -d "{ \"longurl\": \"http://this-is-my-long-url.com/asdfghjkl
\" }" http://localhost:8080/new
* Adding handle: conn: 0x1dc3050
* Adding handle: send: 0
* Adding handle: recv: 0
* Curl_addHandleToPipeline: length: 1
* - Conn 0 (0x1dc3050) send_pipe: 1, recv_pipe: 0
* About to connect() to localhost port 8080 (#0)
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> POST /new HTTP/1.1
> User-Agent: curl/7.30.0
> Host: localhost:8080
> Accept: */*
> Content-Length: 57
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 57 out of 57 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Wed, 25 Nov 2015 19:57:54 GMT
< Content-Length: 72
<
{"ShortID":"bRche","LongURL":"http://this-is-my-long-url.com/asdfghjkl"}* Connec
tion #0 to host localhost left intact
´´´

The JSON response of the server will contain the path of the short URL in a param called ShortID.

## When visting a shortURL
When visiting a short URL, the server sends a 301 HTTP response with the appropriate Location header containing the longURL.
´´´´
pozo@POZOBOE
$ curl -v -X GET http://localhost:8080/bRche
* Adding handle: conn: 0x6b2ef0
* Adding handle: send: 0
* Adding handle: recv: 0
* Curl_addHandleToPipeline: length: 1
* - Conn 0 (0x6b2ef0) send_pipe: 1, recv_pipe: 0
* About to connect() to localhost port 8080 (#0)
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET /bRche HTTP/1.1
> User-Agent: curl/7.30.0
> Host: localhost:8080
> Accept: */*
>
< HTTP/1.1 301 Moved Permanently
< Location: http://this-is-my-long-url.com/asdfghjkl
< Date: Wed, 25 Nov 2015 19:58:31 GMT
< Content-Length: 0
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
```


## Statistics of visits for a shortURL
The server keeps the statistics for the number of times each link is visited. The stats can be queries including a '+' sign at the end of the short URL.
´´´´
pozo@POZOBOE $ curl -v -X GET http://localhost:8080/bRche+
* Adding handle: conn: 0x6b2ef0
* Adding handle: send: 0
* Adding handle: recv: 0
* Curl_addHandleToPipeline: length: 1
* - Conn 0 (0x6b2ef0) send_pipe: 1, recv_pipe: 0
* About to connect() to localhost port 8080 (#0)
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET /bRche HTTP/1.1
> User-Agent: curl/7.30.0
> Host: localhost:8080
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Wed, 25 Nov 2015 19:58:31 GMT
< Content-Length: 38
< Content-Type: application/json
<
<{"ShortID":"bRche","NumberOfVisits":34566}
* Connection #0 to host localhost left intact
```


# TODO
* Refactor the code
* Better logging


# Ref links

http://dougblack.io/words/a-restful-micro-framework-in-go.html

http://thenewstack.io/make-a-restful-json-api-go/


