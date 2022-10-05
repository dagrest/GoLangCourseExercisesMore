# Go Course Exercises More

### GO suffers no fools!!!

- [crypto bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)  
- run the following to install bcrypt:  
- ---> go get golang.org/x/crypto/bcrypt  
- run the following to update bcrypt:  
- ---> go get -u golang.org/x/crypto/bcrypt  

### Test run command:  
- go test  

### Test coverage command:  
- go test -cover  
- go test -coverprofile c.out  
- go tool cover -html=c.out  

### Benchmark command (ninja_level_13/01/example_01):  
- go test -bench .  

### Format command:  
- go fmt ./...  

[An Introduction to Programming in Go (Book)](https://www.golang-book.com/books/intro)  

Map in Go:  
- [Go maps in action](https://go.dev/blog/maps)  
- [Macro View of Map Internals In Go](https://www.ardanlabs.com/blog/2013/12/macro-view-of-map-internals-in-go.html)

### Concurrency

    ARDANLABS GO TRAINING
- https://github.com/ardanlabs/gotraining
- [topics](https://github.com/ardanlabs/gotraining/tree/master/topics/go)
- [race conditions](https://github.com/ardanlabs/gotraining/blob/master/topics/go/concurrency/data_race/README.md)  
- [Photos fom book](https://photos.google.com/share/AF1QipMTxq4L6HMxzow9bZLa2UU71z5R3AH-45a417xHGEtBCB7UyXZDUX4PL5KGTcheEg?key=Si1oSUJ5VU9BSWVPczdpREd5Z2N0eHBaYTZabnpn)  
Commands:
- go run -race main.go  
vs  
- go run main.go  

#### Concurrency Patterns
- [Go Blog](https://go.dev/blog/all): Search for concurrency
- [Pipeline Pattern](https://go.dev/blog/pipelines)
- [Rob Pike - Fan In example](https://go.dev/talks/2012/concurrency.slide#25)  

- FAN OUT  
  Multiple functions are reading from the same channel until that channel is closed
- FAN IN  
  A function can read from multiple inputs and proceed until all are  closed by  
  multiplexing the input channels onto a single channel that's closed when  
  all the input are closed.


    PATTERN
there's a pattern to our pipeline functions:
- stages close their outbound channels when all the send operations ae done.
- stages keep receiving values from inbound channels until those channels are closed.

[Source - pipelines](https://go.dev/blog/pipelines)

### Channels  
- [Channel types](https://go.dev/ref/spec#Channel_types)

### [Go Proverbs](https://go-proverbs.github.io/)
