# CSE138_Assignment2

In order to run this program, you must do some setup with Docker before you can do anything else

First, create a subnet `mynet` with IP Range 10.10.0.0/16 like this:

`$ docker network create --subnet=10.10.0.0/16 mynet`

Then build the docker image like so:

`$ docker build -t assignment2-img .`

Now we can run the main and forwarding instances of our program like so:

Main Instance: 
`docker run -p 8086:8085 --net=mynet --ip=10.10.0.2 --name="main-container" assignment2-img`

Forwarding Instance:
`docker run -p 8087:8085 --net=mynet --ip=10.10.0.3 --name="forwarding-container" -e FORWARDING_ADDRESS=10.10.0.2:8085 assignment2-img`

Then you should be able to send GET, PUT, or DELETE requests to either the main or forwarding instances.

In order to run the unit tests, run `go test -v`
## Acknowledgements
- Zach and Jackie went to Patrick’s office hours to ask about how to determine which container we are in, environment variables, and if each container is running.
- Oleksiy went to Abhay’s office hours to figure out how to forward requests to the forwarding container from the main container. He also went to Patrick’s section on Monday where he discussed various aspects of the assignment such as sending requests, how to know whether the forwarding container is up using environment variables, and forwarding the request to the main container when the forwarding container is up.

## Citations
- [Go by Example: Environment Variables](https://gobyexample.com/environment-variables)
- [Go Web Framework: Gin](https://github.com/gin-gonic/gin)
- [Go Package: http](https://golang.org/pkg/net/http/)
- [How To Correctly Serialize JSON String In Golang](https://goinbigdata.com/how-to-correctly-serialize-json-string-in-golang/)
- [Go Package: OS](https://golang.org/pkg/os/)
- [Go Package: JSON](https://golang.org/pkg/encoding/json/)
- [Go Package: IO](https://golang.org/pkg/io/ioutil/)
- [Go Handler Interface](https://divyanshushekhar.com/golang-responsewriter-request/)
- [Go HTTP Response](https://medium.com/@vivek_syngh/http-response-in-golang-4ca1b3688d6)
- [apitest Documentation](https://apitest.dev/)
- [Testing APIs in Golang using apitest](https://dev.to/eminetto/testing-apis-in-golang-using-apitest-1860)

## Team Contributions
- Oleksiy Omelchenko: worked on Part 1 to ensure that GET, PUT, and DELETE requests worked fine for requests sent to a single Docker container. He also provided the Dockerfile needed for the Go files to function correctly in an isolated environment.
- Jackie: ensured that part 2 was working where requests were either sent to the forwarding container and finally to the main container or were sent straight to the main container.
- Zach: built robust tests to ensure that our program works in situations which weren't previously tested by either part 1 or part 2 of the testing script.

