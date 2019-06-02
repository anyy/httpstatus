## Installation

```
go get github.com/gazelle0130/httpstatus
```

## Usage

You can use httpstatus to display http status code.

```
# show all http status codes
httpstatus all
100 : Continue
101 : Switching Protocols
102 : Processing
200 : OK
201 : Created
202 : Accepted
203 : Non-Authoritative Information
204 : No Content
...
```

```
# find status by a code
httpstatus find 200
200 : OK
```
