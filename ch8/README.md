# Chapter8: Testing your application

Test utilities provided by golang seems primitive, but are essential tools required for implement automatic testing. This chapter introduces **testing** (built-in test package), **check, Ginkgo** (popular packages).

## 8.1: Go and Test
**testing** is a built-in test package, provides basic part of auto-testing. **net/http/httptest** is a test package for web applications.

Package testing used with the `go test` command. This command calls Go source file named `hoge_test.go`, filename often corresponds to the target script. In the test file, function must be implementd in below format.

``` go
func TestXxx(*testing.T) {...}
```

`go test` command executes those (prefixed) functions.

## 8.2: Unittest in Go
Normally, unit (program module) is provided data and returns output. Unittest is executed connecting multiple test as **testsuite**.

Function `TestXxx` receives `t`, pointer to struct `testing.T`. This can be used for notify faults or errors. `testing.T` provides some useful functions;

- Log: similar to `fmt.Println`, record text to error log
- Logf: similar to `fmt.Printf`, reshape by a format, record the generated text to error log
- Fail: create the fault record, but also allow continuing execution
- FailNow: create the fault record and stop execution
- Convenience function combining them also exists

Function `Error` calls `Log` then calls `Fail`.

Function `Skip` provided by `testing.T` can be used to skip time-consuming testcase on thebasic check (not only skip annoying fault record).

Function `t.Parallel()` executes multiple testcases concurrently. This is usefull to save time when testcases have no dependencies. `$ go test -v -short -parallel n` (`-v -short` are another options).

**Benchmark test** is another type of test provides by testing package. Benchmark test in Go simply iterates target function `b.N` times. `b.N` varies depending on the necessity. `$ go test -v -cover -short -bench .` (`-v -cover -short` is another options). Target benchmark file can be specified via `-bench` parameter, `.` treat all benchmark files.

## 8.3: HTTP test in Go
There are many ways to test web application. This book focuses on unittest of handlers in Go. Package `testing/httptest` plays the role in Go. `httptest` provides utility to simulate web server. This allows us to obtain HTTP response by sending HTTP request with `net/http` package.

Package `testing` provides a function `TestMain` can be used for pre/postprocess. Like;

``` go
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}
```

`setUp` and `tearDown` are functions to execute pre/postprocess required for all testcases, called only once. Function `m.Run` calls individual testcase functions. `Run` returns terminal code then put it on `os.Exit`.

Server test implemented in `http_test1` initialize global variables in `setUp` function. This shrinks individual testcase functions, makes change of initialization easy. **However**, this testcase interacts with the actual database. This is a kind of dependency and testcase is not fully independent.

## 8.4: Test double and dependency injection
**Test double** effectively makes unittest independent. Test double is used to preserve actual object, struct, and function in the simulation. For example in testing mailer, sending actual mails is bothering. We can avoid it by creating test double resembling mailer behavior. This also prevents testcase from depending on actual database conditions.

Test double work as a substitute for the actual function or struct. Using test double requires to design a software assuming existance of test double. For the service in ch7, dependency on the database exists in deep layer in the code, this make difficult to create test double.

**Dependency injection** is a popular design pattern when assuming test double. This separates dependencies in a software by creating *dependency source*. The dependency source is implemented as interface type in Go.

In this chapter, the element of dependency injection is not to supress to use `sql.Db` but to avoid a direct dependency for it. We inject `sql.DB` into struct `Post` as a *dependency source* so that the code can be tested by test double (no need to interact with actual database).

- Go interface example: [https://gobyexample.com/interfaces](https://gobyexample.com/interfaces)

## 8.5: Third-party Go test libraries
- gocheck: simple and cooperative
- Ginkgo: enables behavior-driven development in Go, complicated

## 8.6: Summary
- Go provides a `testing` package as a built-in test tool and for unittest, launched with `go test`
- `testing` provides basic test and benchmark utilities
- Unittest in Go web application can be done by `testing/httptest`
- Test double ensures independencies in test cases.
- *Dependency injection* is a design pattern to implement test double
- Go also provides third-party test library, Gocheck expand basic go test utility, Ginkgo implements tests for BDD (behavior-driven development)
