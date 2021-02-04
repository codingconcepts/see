# see
see is a simple version of the watch command.

## Installation

```
$ go get -u github.com/codingconcepts/see
```

## Usage

```
$ see -h
Usage of see:
  -n uint
        number of seconds to wait between executions (default 5)
```

## Example

The following see command fetches pods from all namespaces in Kubernetes every 10 seconds: 
```
$ see -n 10 kubectl get pods --all-namespaces
```