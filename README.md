TodoFinder 
===

todofinder is a tool that search for comment containing a specific key pattern such as "TODO" or "FIXME" in golang packages, and display results to the end-user with essential informations (files name, lines, comments) 

The tool has two execution mode :
- search mode : results are display in console stdout. It is the mode by default 
- server mode : service run as an API server on port 8080 by default


```
$ go run todofinder/cmd/todofinder.go -package fmt -pattern TODO 
/usr/local/go/src/fmt/scan.go : 740 :
TODO： accept N and Ni independently?
```

Installation
---

The project use "dep", a prototype dependency management tool for Go.
For more information, please visit https://github.com/golang/dep
```
% go get -u github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder
% dep ensure
```

Usage
---

```
$ todofinder -h
usage: todofinder <command>
The commands are:
   search - run in command line  (*default)
   server - run in server mode

search [-package input_package_name] [-pattern input_search_pattern]
  -package string
       package to search
  -pattern string
       patttern to search
```

### Command Line mode

```
$ go run todofinder/cmd/todofinder.go -h
```
or
```
$ go run todofinder/cmd/todofinder.go search -h
```

```
Usage of search:
  -package string
       package to search
  -pattern string
       patttern to search
```

### Server mode

```
$ go run todofinder/cmd/todofinder.go server -h
```

```
Usage of server:
  -config string
    	configuration file path
```

#### Run in server mode
```
$ go run todofinder/cmd/todofinder.go server -config ../conf/todofinder.yaml

```
The application runs as an HTTP server at port 8080 (default). It provides the following RESTful endpoints:

* `GET /search`: search a specific pattern in all .go file from a package

##### Example API request

```
$ todofinder server -config ../conf/todofinder.yaml &amp;
$ curl -XGET 'localhost:8080/search?package=fmt&amp;pattern=TODO' .
{
    "result": [
        {
            "file": "/usr/local/Cellar/go/1.6.3/libexec/src/fmt/format.go",
            "pos": 332,
            "com": "TODO: Avoid buffer by pre-padding.\n"
        },
        {
            "file": "/usr/local/Cellar/go/1.6.3/libexec/src/fmt/scan.go",
            "pos": 747,
            "com": "TODO: accept N and Ni independently?\n"
        }
    ]
}

```

##### API Error codes


| Error Code | Description |
| --- | --- |
| NOT_FOUND | Resource was not found |
| METHOD_NOT_ALLOWED| Call endpoint using a not supported method |
| INTERNAL_SERVER_ERROR| An issue occurred server side |
| UNAUTHORIZED| User is not authorized to call this resource |
| BAD_PARAMETER| Call endpoint using a bad or missing parameter |
| PACKAGE_NOT_FOUND| Cannot find the package to search on |
| NO_SOURCE| The package does not contain valid source |
| SOURCE_NOT_READABLE| The package .go source files are not readable |
