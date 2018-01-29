# gophan

gophan is Go commandline wrapper for [`phantomjs`][phantomjs]

[phantomjs]: http://phantomjs.org/

# Install

```
go get github.com/hlts2/gophan
```

## CLI Usage

```
$ gophan --help
Usage:
  gophan [command]

Available Commands:
  help        Help about any command
  run         Run binary of phantomjs internally

Flags:
  -h, --help   help for gophan

```

```
gophan run [URL or HTML] <option>

Available Options:
  -o, --out   Set the output file of the capture result (available extensions are png, jpg, pdf etc)
              $ gophan run [URL or HTML] -o output/capture.png

  -s, --set   Set custom javascript for phantomjs
              $ gophan run [URL or HTML] -s custom/load.js

  -q, --query Set CSS selector
              $ gophan run [URL or HTML] -q "#main"
```

## Example

```
go run ./example/main.go
gophan run http://localhost/ -o images/capture.png
```
![capture](images/capture.png)
