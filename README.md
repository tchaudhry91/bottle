# Bottle

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=tchaudhry91_bottle&metric=alert_status)](https://sonarcloud.io/dashboard?id=tchaudhry91_bottle)
![CI](https://github.com/tchaudhry91/bottle/workflows/CI/badge.svg)

Store and re-use data from unix pipes.
Use these utilities to pipe data into a common store across the network.

```bash
# Machine 1: Pipe data into the bottle
echo "Wow" | bottle-cli -m fill -i sample

# Crate stores the data

# Machine 2: Pipe data back
bottle-cli -m drain -i sample
Wow

```

This project contains two parts. A server, called `crate`, and a client, called `bottle-cli`. 

## Crate

This is the server that holds all bottles stored. The crate CLI operates with the following parameters

```bash
crate -h
Usage of crate:
  -db-file string
        Local file to store bottles in (default "crate.db")
  -grpc-addr string
        address to bind on (default "127.0.0.1:14999")
```

## Bottle-CLI

This is cli utility to interact with the crate. It uses gRPC behind the scenes. But does not enforce TLS yet! So do not use this on public networks as of now.
```
bottle-cli -h                                                                             130 ↵
NAME:
   bottle - Store data from pipes

USAGE:
   bottle-cli [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --server value, -s value  Address of the remote server [$BOTTLE_SERVER]
   --mode value, -m value    Mode to operate in, fill/drain/pour (default: "list") [$BOTTLE_MODE]
   --id value, -i value      ID of the bottle to operate on (default: "clip") [$BOTTLE_ID]
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)
```