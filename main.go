package main

import (
	"GoLangCourseExercisesMore/action"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math"
	"os"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
)

// GO suffers no fools!!!

// crypto bcrypt https://pkg.go.dev/golang.org/x/crypto/bcrypt
// run the following to install bcrypt:
// ---> go get golang.org/x/crypto/bcrypt
// run the following to update bcrypt:
// ---> go get -u golang.org/x/crypto/bcrypt

// Test run command:
// go test

// Test coverage command:
// go test -cover
// go test -coverprofile c.out
// go tool cover -html=c.out

// Benchmark command (ninja_level_13/01/example_01):
// go test -bench .

/*
	Concurrency
	===========
	https://go.dev/doc/effective_go#concurrency

	Goroutines
	They're called goroutines because the existing terms—threads, coroutines, processes, and so on—convey inaccurate connotations.
	A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space.
	It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap,
	and grow by allocating (and freeing) heap storage as required.

	Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue
	to run. Their design hides many of the complexities of thread creation and management.

	Prefix a function or method call with the go keyword to run the call in a new goroutine. When the call completes, the
	goroutine exits, silently. (The effect is similar to the Unix shell's & notation for running a command in the background.)


	USEFUL LINKS
	=============
	https://talks.golang.org/
	https://talks.golang.org/2012/concurrency.slide#25

	https://go.dev/blog/pipelines

	Go Concurrency Patterns: Context
	https://go.dev/blog/context


	Go statements
	=============
	https://go.dev/ref/spec

	A "go" statement starts the execution of a function call as an independent concurrent thread of control, or goroutine,
	within the same address space.

	GoStmt = "go" Expression .
	he expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted
	as for expression statements.

	The function value and parameters are evaluated as usual in the calling goroutine, but unlike with a regular call,
	program execution does not wait for the invoked function to complete. Instead, the function begins executing independently
	in a new goroutine. When the function terminates, its goroutine also terminates. If the function has any return values,
	they are discarded when the function completes.  <----- !!!!!
*/

var a int = 8

const (
	length     = 10
	weight int = 100
)

type personExample struct {
	first string
	last  string
	age   int
}

//var array [10]int

func main() {
	// first comment in test_01 branch
	// second comment in test_02 branch
	// b_03
	// b_04

	//fmt.Printf("%+v\n", a)
	//fmt.Println(a)
	//arrayActions()
	// printNumberInDifferentFormats(255)
	// printNumberInDifferentFormats(256)
	// printNumberInDifferentFormats(512)

	//fmt.Printf("%#U\n", 90)
	//level4Ex1()
	//level4Ex2()
	//level4Ex3()
	//level4Ex4()
	//level4Ex5()
	//level4Ex6() - Incorrect solution by lecturer!!!
	//level4Ex7()
	//level4Ex8()
	//level4Ex9()
	//level4Ex10()

	//level5Ex1()
	//level5Ex2()
	//level5Ex3()
	//level5Ex4()

	//Level6Ex1()
	//Level6Ex2()
	//Level6Ex3()
	//Level6Ex4()
	//Level6Ex5()
	//Level6Ex6()
	//Level6Ex7()
	//Level6Ex8()
	//Level6Ex9()
	//Level6Ex10()
	//testManyCores()

	//Level7Ex1()
	//Level7Ex2()

	//sortFunc()

	//testBcrypt()

	//Level8Ex1()
	//Level8Ex2()
	//Level8Ex3()
	//Level8Ex4()
	//Level8Ex5()

	//getSystemData()

	//Level9Ex1() - ??? WHY the exercise is not working if not implemented in main - ???
	/*
		// ====================================
		// =========== Level9Ex1 ==============
		// Go routines
		fmt.Println("Go routines - Stated")
		fmt.Println("Begin CPUs: ", runtime.NumCPU())
		fmt.Println("Begin Go routines: ", runtime.NumGoroutine())

		var waitGroup sync.WaitGroup
		waitGroup.Add(2)

		//go action1(waitGroup) - not working ??? - DEAD LOCK - WHY???
		go func() {
			fmt.Println("Action - 1")
			waitGroup.Done()
		}()
		//go action2(waitGroup) - not working ??? - DEAD LOCK - WHY???
		go func() {
			fmt.Println("Action - 1")
			waitGroup.Done()
		}()

		fmt.Println("Mid CPUs: ", runtime.NumCPU())
		fmt.Println("Mid Go routines: ", runtime.NumGoroutine())

		waitGroup.Wait()

		fmt.Println("End CPUs: ", runtime.NumCPU())
		fmt.Println("End Go routines: ", runtime.NumGoroutine())
		fmt.Println("Go routines - Finished")

		// Go routines
		// =========== Level9Ex1 ==============
		// ====================================
	*/

	//Level9Ex2()

	/*
		//Level9Ex5()
		//Level9Ex4()
		//Level9Ex3()
		// ====================================
		// =========== Level9Ex5 ==============
		// =========== Level9Ex4 ==============
		// =========== Level9Ex3 ==============
		// Go routines
		// run as (to see race condition):
		// 			go run -race test.go
		// OUTPUT:
		// CPUs: 10
		// Goroutines: 1
		// Count:  99
		// Found 1 data race(s)
		// exit status 66

		var count int64 = 0
		fmt.Println("CPUs:", runtime.NumCPU())
		fmt.Println("Goroutines:", runtime.NumGoroutine())

		var wg sync.WaitGroup
		type Mutex struct {
		}
		//var m sync.Mutex

		wg.Add(100)
		for i := 0; i < 100; i++ {
			go func() {
				//runtime.Gosched() // https://pkg.go.dev/runtime#Gosched

				atomic.AddInt64(&count, 1)
				fmt.Printf("count: %d\n", atomic.LoadInt64(&c))
				//m.Lock()
				//count++
				//m.Unlock()
				//fmt.Printf("count: %d\n", count)
				wg.Done()
			}()
		}
		fmt.Println("Goroutines:", runtime.NumGoroutine())
		wg.Wait()

		fmt.Println("CPUs:", runtime.NumCPU())
		fmt.Println("Goroutines:", runtime.NumGoroutine())
		fmt.Println("Count: ", count)

		// Go routines
		// =========== Level9Ex3 ==============
		// =========== Level9Ex4 ==============
		// =========== Level9Ex5 ==============
		// ====================================
	*/

	//Level10Ex1()
	//Level10Ex2()
	//Level10Ex3()
	//Level10Ex4()
	//Level10Ex5()
	//Level10Ex6()

	// Level10Ex7()
	// Level10Ex7TryToEnhanceAndStartEachLoopFromNextDecade()

	// https://go.dev/blog/defer-panic-and-recover

	// Level11Ex1()
	// Level11Ex2()
	// Level11Ex3()
	// Level11Ex5()
	// Level11Ex5()

	// ninja_level_13
	// example_01.MainFunc()

	// ninja_level_13
	// example_02:
	//s := "one two three, four , five one two two two two three"
	//res := strings.Fields(s)
	//m := make(map[string]int)
	//for _, v := range res {
	//	m[v]++
	//}
	//for k, v := range m {
	//	fmt.Printf("Value %9s found: %2d times\n", k, v)
	//}
	//fmt.Println("Finished.")

	//res, _ := test(1, 2, 3, 4)
	//fmt.Println(res)

	//fmt.Println("Hello, Test")
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println("inside loop")
	//	}
	//}
	//for i := 0; i < 10; i++ {
	//	fmt.Println("inside loop")
	//}
	//time.Sleep(2 * time.Second)

	type User struct {
		Name        string `json:"full_name"`
		Age         int    `json:"age,omitempty"`
		Active      bool   //`json:"-"`
		LastLoginAt string
	}

	type Data struct {
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	type Issue struct {
		UserId1 Data `json:"userId1"`
	}

	//type Val struct {
	//	IssuesData Issue `json:"issuesData"`
	//}

	type Val struct {
		IssuesData map[string]Data `json:"issuesData"`
	}

	//	val := `{
	//  "issuesData":{
	//    "userId1": {
	//      "id": 1,
	//      "title": "delectus aut autem",
	//      "completed": false
	//    }
	//  }
	//}`
}

func test(arg ...interface{}) (int, error) {
	count := 0
	for i, v := range arg {
		fmt.Printf("%d - %d\n", i, v)
		val, err := v.(int)
		if !err {
			return 0, fmt.Errorf("")
		}
		count = count + val
	}
	return count, nil
}

func Level11Ex5() int {
	// Just a link to testing resources:
	// https://pkg.go.dev/testing
	// https://www.golang-book.com/books/intro/12

	a := 10
	b := 15
	res := action.AddInt(a, b)
	fmt.Printf("%d + %d = %d", a, b, res)
	return res
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, fmt.Errorf("unable to handle negative number: %v", f) - my previous solution
		// e := errors.New("unable to handle negative number")
		e := fmt.Errorf("unable to handle negative number: %v", f)
		return 0, sqrtError{"11.111 N", "45.444 W", e}
	}
	return 42, nil
}

func Level11Ex4() {
	res, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
		panic(err)
		return
	}

	fmt.Println("SQRT = ", res)
}

type customErr struct {
	errorMessage string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("custom error: %v", ce.errorMessage)
}

func Level11Ex3() {
	c1 := customErr{"Generated error by code"}
	fooLevel11Ex3(c1)
}

func fooLevel11Ex3(e error) {
	fmt.Println("Foo func: ", e.(customErr).errorMessage) // assertion of type
}

type personLevel11Ex2 struct {
	First   string
	Last    string
	Sayings []string
}

func Level11Ex2() {
	p1 := personLevel11Ex2{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Println("failed toJSON method: %w", err)
		return
	}

	fmt.Println(string(bs))
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("failed Marshal: %w", err)
	}
	return bs, nil
}

type personLevel11Ex1 struct {
	First   string
	Last    string
	Sayings []string
}

func Level11Ex1() {
	p1 := personLevel11Ex1{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Errorf("marshal failed: %w", err)
		log.Fatalln("marshal failed: %w", err)
	}
	fmt.Println(string(bs))
}

func Level10Ex7() {
	c := make(chan int)
	//q := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 100; i++ {
				c <- i
			}
		}()
	}

	for n := 0; n < 100; n++ {
		fmt.Println(<-c)
	}

	fmt.Println("Finished...")
}

func Level10Ex7TryToEnhanceAndStartEachLoopFromNextDecade() {
	var countAtomic int64 = 0
	c := make(chan int)
	//q := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := int(countAtomic); i < int(atomic.AddInt64(&countAtomic, 10)); i++ {
				c <- i
			}
			atomic.AddInt64(&countAtomic, 10)
		}()
	}

	array := make([]int, 100)
	//array = array[100:100]
	for n := 0; n < 100; n++ {
		val := <-c
		fmt.Println(n, val)

		array[n] = val
	}

	fmt.Println("Sorted array")
	sort.Ints(array)
	for i := 0; i < 100; i++ {
		fmt.Println(i, array[i])
	}

	fmt.Println("Finished...")
}

func Level10Ex6() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Finished...")
}

func Level10Ex5() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func gen10Ex4(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive10Ex4(c <-chan int, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func Level10Ex4() {
	q := make(chan int)
	c := gen10Ex4(q)

	receive10Ex4(c, q)

	fmt.Println("about to exit")
}

func gen10Ex3() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive10Ex3(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func Level10Ex3() {
	c := gen10Ex3()
	receive10Ex3(c)

	fmt.Println("about to exit")
}

func Level10Ex2() {
	// ====================================
	// Channels
	// Level10Ex2
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
	// Level10Ex2
	// ====================================
}

func Level10Ex1() {
	// ====================================
	// Channels
	// Level10Ex1

	// solution 1:
	//c := make(chan int)

	//go func() {
	//	c <- 42
	//}()

	// solution 2:
	c := make(chan int, 1)
	c <- 42

	fmt.Println(<-c)

	// Level10Ex1
	// ====================================
}

type person92 struct {
	name string
}

type human92 interface {
	Speak()
}

func SaySomething(h human92) {
	fmt.Println("Person - SaySomething implementation from person: ")
	h.Speak()
}

func (t *person92) Speak() {
	fmt.Println("Person - Speak implementation from person")
}

func Level9Ex2() {
	p := person92{name: "James Bond"}

	SaySomething(&p)
}

func action1(waitGroup sync.WaitGroup) {
	fmt.Println("Action - 1")
	waitGroup.Done()
}

func action2(waitGroup sync.WaitGroup) {
	fmt.Println("Action - 2")
	waitGroup.Done()
}

func Level9Ex1() {
	fmt.Println("Level 9 Ex 1")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go action1(waitGroup)
	go action2(waitGroup)

	waitGroup.Wait()

	fmt.Println("Level 9 Ex 1 - Finished")
}

func getSystemData() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

type user85 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type By_Age []user85

func (a By_Age) Len() int           { return len(a) }
func (a By_Age) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a By_Age) Less(i, j int) bool { return a[i].Age < a[j].Age }

type By_Last []user85

func (n By_Last) Len() int           { return len(n) }
func (n By_Last) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n By_Last) Less(i, j int) bool { return n[i].Last < n[j].Last }

func Level8Ex5() {
	u1 := user85{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user85{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user85{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user85{u1, u2, u3}

	fmt.Println(users)

	fmt.Println("Sort by Age:")
	sort.Sort(By_Age(users))
	for _, u := range users {
		fmt.Printf("Age: %d\n", u.Age)
		fmt.Printf("\t%s %s\n", u.First, u.Last)
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Printf("\t\t%s\n", s)
		}
	}

	fmt.Println("Sort by Last:")
	sort.Sort(By_Last(users))
	for _, u := range users {
		fmt.Printf("Age: %d\n", u.Age)
		fmt.Printf("\t%s %s\n", u.First, u.Last)
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Printf("\t\t%s\n", s)
		}
	}
}

func Level8Ex4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	//sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}

type user83 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func Level8Ex3() {
	u1 := user83{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user83{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user83{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user83{u1, u2, u3}

	fmt.Println(users)

	fmt.Println("---")
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Somthing went wrong!")
	}
}

type user struct {
	First string
	Age   int
}

func Level8Ex1() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	v, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Unable to convert to JSON")
		return
	}
	fmt.Println("JSON: ", string(v))
}

type Person82 struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func Level8Ex2() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	//res := Person82{}
	var res []Person82
	fmt.Println(res)
	err := json.Unmarshal([]byte(s), &res)
	if err != nil {
		fmt.Println("Unable to parse JSON")
		return
	}
	fmt.Println(res)
	for i, v := range res {
		fmt.Println("Value", i, v)
	}
}

func testBcrypt() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	newPass := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(newPass))
	if err != nil {
		fmt.Println("Incorrect pass")
		return
	}
	fmt.Println("You're logged in!")
}

type PersonSort struct {
	Name string
	Age  int
}

type ByAge []PersonSort

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []PersonSort

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name }

func sortFunc() {
	p1 := PersonSort{Name: "JamesBond", Age: 32}
	p2 := PersonSort{Name: "Moneypenny", Age: 27}
	p3 := PersonSort{Name: "Q", Age: 64}
	p4 := PersonSort{Name: "M", Age: 56}

	people := []PersonSort{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people)) // sort by Age
	fmt.Println(people)
	sort.Sort(ByName(people)) // sort by Name
	fmt.Println(people)
}

func Level7Ex1() {
	x := 100
	fmt.Println("x: ", x)
	fmt.Println("Address of x: ", &x)
}

type Person7 struct {
	first string
	last  string
	age   int
}

func changeMe(p *Person7) {
	(*p).first = "New_" + (*p).first
	// p.first = "New_" + (*p).first // will work also
}

func Level7Ex2() {
	p := Person7{
		first: "James",
		last:  "Bond",
		age:   42,
	}
	fmt.Println("Person: ", p)
	changeMe(&p)
	fmt.Println("Person: ", p)
}

func worker(n int) {
	for i := 0; i < 100000000; i++ {
		fmt.Printf("Worker %d [%d]\n", n, i)
	}
}

func testManyCores() {
	go worker(1)
	go worker(2)
	go worker(3)
	go worker(4)
	go worker(5)
	go worker(6)
	go worker(7)
	go worker(8)
	go worker(9)
	go worker(10)
}

func enclosure() func() int {
	internalVar := 100
	return func() int {
		internalVar++
		return internalVar
	}
}

func Level6Ex10() {
	f := enclosure()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

// ----- Level6Ex9() ------

func callBackFunc() {
	fmt.Println("From callBackFunc func")
}

func funcWithCallBackArg(f func()) {
	f()
}

func Level6Ex9() {
	funcWithCallBackArg(callBackFunc)
}

// ----- Level6Ex8() ------

func returnsFunc() func() {
	return func() {
		fmt.Println("From returned func")
	}
}

func Level6Ex8() {
	rf := returnsFunc()
	rf()
}

func Level6Ex7() {
	f := func() {
		fmt.Println("unonimous func")
	}
	f()
}

func Level6Ex6() {
	func() {
		fmt.Println("one more unonimous func")
	}()
}

// ----- Level6Ex5() ------

type Shape interface {
	area() float64
}

type Square struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.width * s.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s Shape) {
	fmt.Printf("Area: %f\n", s.area())
}

func infoWithInterface(i interface{}) {
	switch v := i.(type) {
	case Circle:
		fmt.Printf("Circle area: %f\n", v.area())
	case Square:
		fmt.Printf("Square area: %f\n", v.area())
	default:
		fmt.Println("Unknown shape")
	}
}

func Level6Ex5() {
	square := Square{
		height: 10.0,
		width:  10.0,
	}

	circle := Circle{
		radius: 10.0,
	}

	info(circle)
	info(square)

	infoWithInterface(circle)
	infoWithInterface(square)
}

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) speak() {
	fmt.Printf("%s %s (%d) is speaking now", p.first, p.last, p.age)
}

func Level6Ex4() {
	person1 := Person{
		"James",
		"Bond",
		43,
	}

	person1.speak()
}

func deferredFunction() {
	fmt.Println("The deferred function")
}

func Level6Ex3() {
	fmt.Println("Before deferred function")
	defer deferredFunction()
	fmt.Println("After deferred function")
}

func Level6Ex2() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Sum of int: ", foo62(values...))
	fmt.Println("Sum of int: ", bar62(values))
}

func foo62(array ...int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return sum
}

func bar62(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func foo61() int {
	x := 0
	return x
}

func bar61() (int, string) {
	x := 100
	name := "Bar"
	return x, name
}

func Level6Ex1() {
	fmt.Println("foo61", foo61())
	res, nameRes := bar61()
	fmt.Println("bar61", res, nameRes)
}

func level5Ex4() {
	p := struct {
		firstName      string
		hobbies        map[string]string
		iceCreamFlavor []string
	}{
		firstName: "James",
		hobbies: map[string]string{
			"Tennis":   "First",
			"Football": "Last",
		},
		iceCreamFlavor: []string{"Strawberry", "Lemon", "Vanilla"},
	}
	fmt.Println(p)
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicleType vehicle
	fourWheel   bool
}

type sedan struct {
	vehicleType vehicle
	luxury      bool
}

func level5Ex3() {
	t := truck{
		vehicleType: vehicle{doors: 2, color: "black"},
		fourWheel:   false,
	}
	s := sedan{
		vehicleType: vehicle{doors: 4, color: "red"},
		luxury:      true,
	}
	fmt.Printf("Vehicle type: %T\n", t)
	fmt.Printf("\t doors: %v, color: %v\n", t.vehicleType.doors, t.vehicleType.color)
	fmt.Printf("\t fourWheel: %v\n", t.fourWheel)

	fmt.Printf("Vehicle type: %T\n", s)
	fmt.Printf("\t doors: %v, color: %v\n", s.vehicleType.doors, s.vehicleType.color)
	fmt.Printf("\t fourWheel: %v\n", s.luxury)

	fmt.Println("")
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.vehicleType.doors)
	fmt.Println(s.vehicleType.color)
}

type person struct {
	firstName      string
	lastName       string
	iceCreamFlavor []string
}

func level5Ex2() {
	m := map[string]person{}
	person1 := person{firstName: "James", lastName: "Bond", iceCreamFlavor: []string{"strawberry", "vanilla", "grapefruit"}}
	person2 := person{firstName: "Super", lastName: "Man", iceCreamFlavor: []string{"raspberry", "pistachio", "lemon"}}

	m[person1.lastName] = person1
	m[person2.lastName] = person2

	for _, v := range m {
		fmt.Printf("Person: %v %v\n", v.firstName, v.lastName)
		for _, val := range v.iceCreamFlavor {
			fmt.Printf("\tFlavor: %v\n", val)
		}
	}

}

func level5Ex1() {
	person1 := person{firstName: "James", lastName: "Bond", iceCreamFlavor: []string{"strawberry", "vanilla", "grapefruit"}}
	person2 := person{firstName: "Super", lastName: "Man", iceCreamFlavor: []string{"raspberry", "pistachio", "lemon"}}

	fmt.Printf("Person: %v %v\n", person1.firstName, person1.lastName)
	for _, v := range person1.iceCreamFlavor {
		fmt.Printf("\tFlavor: %v\n", v)
	}
	fmt.Println("============")
	fmt.Printf("Person: %v %v\n", person2.firstName, person2.lastName)
	for _, v := range person2.iceCreamFlavor {
		fmt.Printf("\tFlavor: %v\n", v)
	}

}

func level4Ex10() {
	a := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	a["Todd McLeod"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	for key, v := range a {
		fmt.Printf("For: %v\n", key)
		for i, favorite := range v {
			fmt.Printf("\tFavorite (%v): %v\n", i+1, favorite)
		}
	}

	fmt.Printf("After deletion of %v\n", "moneypenny_miss")
	delete(a, "moneypenny_miss")
	for key, v := range a {
		fmt.Printf("For: %v\n", key)
		for i, favorite := range v {
			fmt.Printf("\tFavorite (%v): %v\n", i+1, favorite)
		}
	}
}

func level4Ex9() {
	a := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	a["Todd McLeod"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	for key, v := range a {
		fmt.Printf("For: %v\n", key)
		for i, favorite := range v {
			fmt.Printf("\tFavorite (%v): %v\n", i+1, favorite)
		}
	}
}

func level4Ex8() {
	a := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}
	for key, v := range a {
		for _, favorite := range v {
			fmt.Printf("Key: %v, Favorite: %v\n", key, favorite)
		}
	}
}

func level4Ex7() {
	a := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	// a[0] = []string{"James", "Bond", "Shaken, not stirred"}
	// a[1] = []string{"Miss", "Moneypenny", "Helloooooo, James."}
	for i, v := range a {
		fmt.Printf("Record %v: %v\n", i, v)
		for j, v1 := range v {
			fmt.Printf("\tRecord %v, Data %v: %v\n", i, j, v1)
		}
	}
}

func level4Ex6() {
	a := make([]string, 0, 50)
	a = append(a, "Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming")
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	for i := 0; i < len(a); i++ {
		fmt.Printf("Index: %v, state: %v\n", i, a[i])
	}
}

func level4Ex5() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	a = append(a[:3], a[6:]...)
	fmt.Println(a)
}

func level4Ex4() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	a = append(a, 52)
	fmt.Println(a)

	a = append(a, 53, 54, 55)
	fmt.Println(a)

	y := []int{56, 57, 58, 59, 60}
	a = append(a, y...)
	fmt.Println(a)
}

func level4Ex3() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s1 := a[:5]
	s2 := a[5:]
	s3 := a[2:7]
	s4 := a[1:6]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

func level4Ex2() {
	a := []int{1, 2, 3, 4, 5}
	s := a[:]
	for i, v := range s {
		fmt.Printf("%v : %v of type %T\n", i, v, v)
	}
	fmt.Printf("Array type: %T", s)
}

func level4Ex1() {
	a := [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("%v : %v of type %T\n", i, v, v)
	}
	fmt.Printf("Array type: %T", a)
}

func anotherExercise() {
	aa := 42
	printNumberInDifferentFormats(aa)
	bb := aa << 1
	printNumberInDifferentFormats(bb)
}

// Level 2: ex #1
// binary - decimal - hex
func printNumberInDifferentFormats(num int) {
	fmt.Printf("Binary: %10b, Decimal: %d, Hex: %#5X - with 0X\n", num, num, num)
	fmt.Printf("Binary: %10b, Decimal: %d, Hex: %#5x - with 0x\n", num, num, num)
	fmt.Printf("Binary: %10b, Decimal: %d, Hex: %5x - without ox\n", num, num, num)
	fmt.Println("------------")
}

func arrayActions() {
	array := [9]int{}

	printArray(array[:])

	for i := 0; i < len(array); i++ {
		array[i] = (i + 1) * 10
	}

	printArray(array[:])
}

func printArray(array []int) {
	fmt.Println("==> printArray")
	fmt.Printf("Array length is %d\n", len(array))
	for i, v := range array {
		fmt.Printf("array[%d] = %d\n", i, v)
	}
	fmt.Println("<== printArray")
}
