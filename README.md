# GoQuiz

[![Golang](https://www.clipartmax.com/png/middle/111-1112331_how-couldnt-you-love-him-gopher-mic-drop.png)](https://golang.org/doc/code.html)

Urlshort is a program that will look at the path of any incoming web request and redirect the user a new page if exists, like a bit.ly.

## Install

To run this program, you will need to:

Example:

```sh
export GOPATH=$(go env GOPATH)
go get github.com/vinixavier/urlshort
cd $GOPATH/vinixavier/urlshort
go install
```

## Using

The program have following help:

```sh
Usage of urlshort:
    -file string
        Specify a YAML or JSON file with path to URL.
```

See examples of YAML files in the `data/*.yaml` directory.

## TODO

- Include new handler: JSON
- Include new handler: BoltDB