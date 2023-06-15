tm-simulator
========

tm-simulator is a tool to reproduce some scenarios and evaluate the schedulers' efficiency.

## Build
1. [Go](https://golang.org/) Version 1.9 or later
2. In the root directory of the [TM project](https://github.com/gottingen/tm), use the `make simulator` command to compile and generate `bin/tm-simulator`


## Usage

This section describes how to use the simulator.

### Flags description

```
-tm string
      Specify a TM address (if this parameter is not set, it will start a TM server from the simulator inside)
-config string
      Specify a configuration file for the TM simulator
-case string
      Specify the case which the simulator is going to run
-serverLogLevel string
      Specify the TM server log level (default: "fatal")
-simLogLevel string
      Specify the simulator log level (default: "fatal")
```

Run all cases:

    ./tm-simulator

Run a specific case with an internal TM:

    ./tm-simulator -case="casename"

Run a specific case with an external TM:

    ./tm-simulator -tm="http://127.0.0.1:2379" -case="casename"
