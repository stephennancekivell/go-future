# go-future

A concurrency library for go.

Also see the introductory blog post [Simplifying go concurrency with Future](https://stephenn.com/2022/05/simplifying-go-concurrency-with-future.html)

## Install

```
go get github.com/stephennancekivell/go-future
```

## Usage

Create a future with a function that you want to run in a go routine, get use `.Get()` to get the value, waiting for it if its not ready.

```go
f := future.New(func() string {
    return "value"
})

value := f.Get()
```

If you have a list of futures you can transform them into a single Future with `Sequence` or `Traverse`

```go
var someFutures []Future[T] // eg
f := future.Sequence(someFutures)

valuesSlice := f.Get()
```

If you have two futures and want the fastest result you can use can use `Race`.

```go
var fa Future[T]
var fb Future[T]

f := future.Race(fa, fb)

fastestResult := f.Get()
```
