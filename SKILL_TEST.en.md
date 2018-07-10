# Background:

A development team created a command-line tool to create a list of TODOs written within
the source code.

This tool searches for comments that include the term “TODO” from the specified Go
package and displays them on the standard output.


For example, you are able to seach for TODO comments in a package like fmt below.

```
$ go run main.go fmt
/usr/local/go/src/fmt/scan.go : 740 :
TODO： accept N and Ni independently?
```

Please refer to the attached main.go file for this tool’s source code.
This main.go file will run on Go 1.7 and above.

# Task:

Your team has decided to expand upon the above tool in the following ways:


* Add the ability to change what string to search for (“TODO”, “FIXME”, etc.)
* Change the code to allow the tool to run on Go 1.6.2 and above
* Provide a server mode
* When the tool is launched in server mode, it should act as an API server
* If you send an HTTP request with the package import path, (fmt, net, http, etc.) and a
*ring to be searched for on this server, the result should be returned in JSON format
* At the very least, the response should contain the appropriate file path, line number, and
the comment


Please change main.go in order to fulfill the requirements.
Feel free to create the API response and request formats as you need to.
There is no need to limit yourself to using a single package or file.
Also, you do not need to comply with the way the attached main.go file is written; please feel
free to change and improve the way it is written.

Important! Please be aware of the following before submitting your answers:

* This tool is to be maintained by the team
* Work to actively change and improve the existing code as you see fit
* The code you submit must be verifiable (must be able to check that it works)
* The source code must be able to generate an API document
