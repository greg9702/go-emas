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

### Documentation

Full project documentation can be found on [Wiki](https://github.com/greg9702/go-emas/wiki)

### Results

#### De Jong's function 1

![Alt Text](https://i.imgur.com/HCVL2oP.gif)

#### Rosenbrock's valley (De Jong's function 2)

![Alt Text](https://i.imgur.com/bt0CyiE.gif)

> More on [Wiki](https://github.com/greg9702/go-emas/wiki)

