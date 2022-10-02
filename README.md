# Go Course Exercises

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
#### ARDANLABS GO TRAINING
- https://github.com/ardanlabs/gotraining
- [topics](https://github.com/ardanlabs/gotraining/tree/master/topics/go)
- [race conditions](https://github.com/ardanlabs/gotraining/blob/master/topics/go/concurrency/data_race/README.md)  

Commands:
- go run -race main.go  
vs  
- go run main.go  