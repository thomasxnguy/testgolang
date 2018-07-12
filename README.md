Todo Finder Tool
===

todofinder is a tool helping to search for comment in packages containing a specific key pattern such as "TODO", "FIXME" etc.. and displays them 
to the user in console.
It also provide a server mode where it acts as a API server


```
$ go run todofinder.go -package fmt -pattern TODO
/usr/local/go/src/fmt/scan.go : 740 :
TODO： accept N and Ni independently?
```

Installation
---

```
% go get -u github.com/..
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

#### Run Server Mode
```
$ go run todofinder/cmd/todofinder.go server -config ../conf/todofinder.yaml

```

By default, application will run on port 8080

##### Example API request

```
$ todofinder server -config ../conf/todofinder.yaml &
$ curl -XGET localhost:8080/search?package=fmt&pattern=TODO .
{
}

```