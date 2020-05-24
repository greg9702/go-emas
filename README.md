# __go-emas__
Evolutionary Multi-Agent System implemented in GO

### Requirements
```
golang >= 1.13
```
> All required golang dependencies will automatically download at the first start-up

### How to run

To run an application, run in project directory:
```
go run cmd/main.go -logFile <PATH> -logLevel <VALUE>
```
> Where PATH is path to a log file and VALUE is an integer

To run tests, run in project directory:
```
 go test test/* -v
```
