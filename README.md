# madprobectl

This project has been made to provide companies wanting simple aliveness probe to be easily installed on their systems.

## Getting Started

### Prerequisites

This project has been developed using `go version go1.13`

It uses the well known `spf13/cobra`.

Golang can be installed from the official [go website](https://golang.org/dl/).

### Installing

To get the CLI, simply do:

 `$ go get -u github.com/MadJlzz/madprobectl`

or build it directly from sources.

It's very simple to use the CLI. At any time, if you need help, you can use `-h` or `--help`
for more information about the commands you can use.

```
$ ./madprobectl --help
```

will output:

```
madprobectl is a lightweight CLI used to manage madprobe's probes state.

For example, it is very easy to create a new probe:
  madprobectl probe create -f probe.yaml

Or maybe you want to delete one?
  madprobectl probe delete simple-http-probe

For an exhaustive list of available commands, please refer to the documentation.

Usage:
  madprobectl [command]

Available Commands:
  help        Help about any command
  probe       Command used to interact with probes

Flags:
      --config string   config file (default is $HOME/.madprobectl.yaml)
  -h, --help            help for madprobectl

Use "madprobectl [command] --help" for more information about a command.
```

## Contributing

I'll be more than happy to have feedback on the way I designed this application. Things can always be done better and
I m eager to learn what could be improved!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details