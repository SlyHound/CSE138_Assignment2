# CSE138_Assignment2

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

## Team Contributions
- Oleksiy Omelchenko: worked on Part 1 to ensure that GET, PUT, and DELETE requests worked fine for requests sent to a single Docker container. He also provided the Dockerfile needed for the Go files to function correctly in an isolated environment.
- Jackie: worked on Part 2
- Zach: worked on building tests
