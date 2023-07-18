# DPLL netlink client in golang

This repository contains a usage example for generic Netlink CLI interacting
with Linux DPLL API.
. 
Intel Ice driver for E810 cards relies on Linux DPLL driver for 
DPLL device queries, configuration and notifications through generic netlink.

The testbench contains kernel module built around Linux DPLL API, and a copy
of user space Python CLI for interacting with the module through generic netlink.

## How to use this?
You need a linux kernel that supports DPLL driver, and either of the two:

1. E-810 Intel NIC with ICE driver that supports netlink. At the moment of writing, this isn't in the upstream yet.
2. [dpll testbench](https://github.com/vitus133/dpll-testbench) 

## Usage
For options, run 

```bash
make help
```
For linting, you need `golangci-lint` binary in your `bin` directory 

You need to be root to use netlink dpll subsystem. 

```bash
[vitaly@rhel9 go-dpll]$ sudo ./dpll-cli 
2023/07/18 17:15:54 {0 dpll_testbench automatic  holdover 5799633565433967664 pps}
2023/07/18 17:15:59 {0 dpll_testbench automatic  unlocked 5799633565433967664 pps}
2023/07/18 17:16:04 {0 dpll_testbench automatic  locked 5799633565433967664 pps}
```
