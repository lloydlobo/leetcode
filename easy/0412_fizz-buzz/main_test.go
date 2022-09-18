package main

import (
	"log"
	"os"
	"os/exec"
	"testing" // Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the "go test" command, which automates execution of any function of the form

	"github.com/fatih/color"
	"github.com/google/go-cmp/cmp"
)

// TestFizzBuzz
// $ go test -v
// $ go test -json
//
// fizzBuzz: func fizzBuzz(n int) []string
//
// https://pkg.go.dev/github.com/google/go-cmp/cmp
// Be careful when using cmp.Equal as it may lead to a panic condition
// It is intended to only be used in tests, as performance is not a goal and it may panic if it cannot compare the values.
// Its propensity towards panicking means that its unsuitable for production environments where a spurious panic may be fatal.
func TestFizzBuzz(t *testing.T) {
	n := 16
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16"}
	got := fizzBuzz(n)
	log.Println(cmp.Equal(got, want))
} // if got == want { t.Errorf("FizzBuzz(%v) = %d; want %d", n, got, want) }

func TestFizzBuzzAppend(t *testing.T) {
	n := 16
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16"}
	got := FizzBuzzAppend(n)
	log.Println(cmp.Equal(got, want))
}

// BenchmarkFizzBuzz
//
// n := 16
// $ go test -bench=.
// 2022/09/18 15:13:25 true
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkFizzBuzz-12              968226              1243 ns/op
// PASS
// ok      github.com/lloydlobo/leetcode   2.194s
func BenchmarkFizzBuzz(b *testing.B) {
	n := 10000000
	for i := 0; i < b.N; i++ {
		fizzBuzz(n)
	}
}

// BenchmarkFizzBuzzAppend
//
// 2022/09/18 16:12:23 true
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkFizzBuzz-12                   1        1458140285 ns/op
// BenchmarkFizzBuzzAppend-12             6         184199832 ns/op
func BenchmarkFizzBuzzAppend(b *testing.B) {
	n := 10000000
	for i := 0; i < b.N; i++ {
		FizzBuzzAppend(n)
	}
}

// n := 10000000
// 2022/09/18 15:51:53 true
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkFizzBuzz-12                   1        1154388657 ns/op
// PASS

// F is a type passed to fuzz tests\.
func TestLog(t *testing.T) {
	log.Println(cmp.Equal(1, 1))
}

// editFile, error := os.Open(filename)
func logToFile(filename string) {
	if filename == "" {
		filename = "test.log"
	}
	file, error := os.Create(filename)
	log.SetFlags(log.Ldate | log.Lshortfile)
	if error != nil {
		log.Fatalf("%s not created, %v", filename, error)
	}
	log.SetOutput(file)
	pwd := osExecLocation() // $ pwd # current directory
	log.Printf("Creating %v! in %s", file, pwd)
	log.Printf("Printing benchmark %v! in %s", osExecLogBenchmark(), pwd)
	file.Close()
}

func osExecLocation() string {
	out, err := exec.Command("pwd").Output() // Output runs the command and returns its standard output.
	if err != nil {
		log.Fatal(err)
	}
	output := string(out)
	color.Set(color.FgYellow, color.Bold)
	log.Println(string(out))
	color.Unset()
	return output
}

func osExecLogBenchmark() string {
	cmd := "go test -bench=."
	cmd = "go test -bench=~/.ghq/github.com/lloydlobo/leetcode/easy/0412_fizz-buzz/main_test.go"
	cmd = "go test -v"
	out, err := exec.Command(cmd).Output() // Output runs the command and returns its standard output.
	if err != nil {
		log.Fatal(err)
	}
	output := string(out)
	return output
}

// func Open(name string) (*File, error) {
// 	return OpenFile(name, O_RDONLY, 0)
// }
// func Chdir(dir string) error {

/*
https://stackoverflow.com/questions/24534072/how-to-compare-if-two-structs-slices-or-maps-are-equal
m1 := map[string]int{
    "a": 1,
    "b": 2,
}
m2 := map[string]int{
    "a": 1,
    "b": 2,
}
fmt.Println(cmp.Equal(m1, m2)) // will result in true

Important Note:

Be careful when using cmp.Equal as it may lead to a panic condition

    It is intended to only be used in tests, as performance is not a goal and it may panic if it cannot compare the values. Its propensity towards panicking means that its unsuitable for production environments where a spurious panic may be fatal.


*/

/*

func fizzBuzz(n int) []string {
	var output []string

	for i := 1; i < n+1; i++ {
		output = append(output, fmt.Sprint(i))

		if (i)%15 == 0 {
			output[i-1] = "FizzBuzz"
		} else if (i)%3 == 0 {
			output[i-1] = "Fizz"
		} else if (i)%5 == 0 {
			output[i-1] = "Buzz"
		}
	}
	return output
}

func main() {
	fmt.Println("Hello, world")
	n := 16
	fmt.Printf("fizzBuzz: %v\n", fizzBuzz(n))
}

*/

/*

func getTestSearchArgs() ([]int, int, int) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	want := 4

	return nums, target, want
}

// TestSearchForLoop is for testing binary search algorithm
// to find index of required target int.
//
// PASS
// ok      binary-search   0.001s
//
// --- FAIL: TestSearchForLoop (0.00s)
// main_test.go:36: Search([-1 0 3 5 8 12], 9) = -1; want 4
// FAIL
// exit status 1
// FAIL    binary-search   0.001s
func TestSearchForLoop(t *testing.T) {
	nums, target, want := getTestSearchArgs()

	got := SearchForLoop(nums, target)
	if got != want {
		t.Errorf("Search(%d, %v) = %v; want %v", nums, target, got, want)
	}
}

// To run all benchmarks, use -bench=., as shown below:
//
//	go test -bench=.
//
// goos: linux
// goarch: amd64
// pkg: binary-search
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkSearchForLoop-12              426310108                2.815 ns/op
// BenchmarkRandInt-12             95099158                12.29 ns/op
// BenchmarkPrimeNumbers-12          471997              2384 ns/op
// BenchmarkBigLen-12                     1        1810085852 ns/op
// PASS
// ok      binary-search   5.634s
func BenchmarkSearchForLoop(b *testing.B) {
	nums, target, _ := getTestSearchArgs()

	for i := 0; i < b.N; i++ {
		SearchForLoop(nums, target)
	}
}

// # Benchmarks Results
//
// To run all benchmarks, use -bench=., as shown below:
//
// $ go test -bench=.
//
// goos: linux
// goarch: amd64
// pkg: binary-search
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkRandInt-12             88516600                12.22 ns/op
// BenchmarkPrimeNumbers-12          470682              2391 ns/op
// BenchmarkBigLen-12                     1        1866591573 ns/op
// PASS
// ok      binary-search   4.119s

// # Benchmarks
//
// Functions of the form
//
//	func BenchmarkXxx(*testing.B)
//
// are considered benchmarks, and are executed by the "go test" command when
// its -bench flag is provided. Benchmarks are run sequentially.
//
// For a description of the testing flags, see
// https://golang.org/cmd/go/#hdr-Testing_flags.
//
// To run all benchmarks, use -bench=., as shown below:
//
// $ go test -bench=.
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

// To run all benchmarks, use -bench=., as shown below:
//
// $ go test -bench=.
//
// SIMPLE BENCHMARKING SUITE
// https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
func primeNumbers(max int) []int {
	var primes []int
	for i := 2; i < max; i++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

// To run all benchmarks, use -bench=., as shown below:
//
// $ go test -bench=.
//
// b.N specifies the number of iterations; the value is not fixed, but dynamically allocated, ensuring that the benchmark runs for at least one second by default.
func BenchmarkPrimeNumbers(b *testing.B) {
	num := 121
	for i := 0; i < b.N; i++ {
		primeNumbers(num)
	}
}

// The benchmark function must run the target code b.N times.
// During benchmark execution, b.N is adjusted until the benchmark function lasts
// long enough to be timed reliably. The output
//
//	BenchmarkRandInt-8   	68453040	        17.8 ns/op
//
// means that the loop ran 68453040 times at a speed of 17.8 ns per loop.
//
// If a benchmark needs some expensive setup before running, the timer
// may be reset:
func BenchmarkBigLen(b *testing.B) {
	// big := NewBig(40)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// big.Len()
		// fmt.Printf("big: %v\n", big)
		NewBig(43)
	}
}

func NewBig(n int) int {
	output := fib(n)
	return output
}

func fib(n int) int {
	if n < 2 {
		return 1
	}
	output := fib(n-1) + fib(n-2)
	return output
}

// SIMPLE TESTING SUITE
//
// Abs returns the absolute value of x.
func Abs(i int) int {
	output := math.Abs(float64(i))

	return int(output)
}

// TestAbs for Warmup testing
// T is a type passed to Test functions to manage test state and support formatted test logs.
// FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf
func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}


*/
/*
// TestHelloName calls the Hello function, passing a name value with which the function should be able to return a valid response message.
func TestHelloName(t *testing.T) {
	name := "Fred"
	want := regexp.MustCompile(`\b` + name + `\b`) // MustCompile is like Compile but panics if the expression cannot be parsed\.
	msg, err := Hello("Fred")                      // Run Hello Command of package greetings
	if !want.MatchString(msg) || err != nil {
		t.Fatalf( // var t *testing.T & Fatalf is equivalent to Logf followed by FailNow.
			`Hello("Fred") = %q, %v, want match for %#q, nil`,
			msg, err, want,
		)
	}
}

// TestHelloEmpty calls the Hello function with an empty string. This test is designed to confirm that your error handling works.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`,
			msg, err,
		)
	}
}

// To implement the test run:
// $ go test
// PASS
// ok      example.com/greetings   0.001s


*/
