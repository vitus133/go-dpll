# DPLL netlink client in golang

This is a CLI for interacting with Linux netlink DPLL API.

## Prerequisites
You need a linux kernel that supports DPLL driver, and either of the two:

1. E-810 Intel NIC with ICE driver that supports netlink. At the moment of writing, this isn't in the upstream yet.
2. [dpll testbench](https://github.com/vitus133/dpll-testbench) 

## Build

```bash
make build
```

## Usage
For options, run 

```bash
./dpll-cli
```
For linting, you need `golangci-lint` binary in your `bin` directory 

You need to be root to use netlink dpll subsystem. 

## Examples
### Dump devices

```bash
[core@cnfde22 go-dpll]$ sudo ./dpll-cli dumpDevices
{"timestamp":"2024-03-04T10:42:16.525504569Z","id":0,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"locked-ho-acquired","clockId":"0x507c6fffff1fb4e8","type":"eec","temp":0}
{"timestamp":"2024-03-04T10:42:16.525504569Z","id":1,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"locked-ho-acquired","clockId":"0x507c6fffff1fb4e8","type":"pps","temp":0}
{"timestamp":"2024-03-04T10:42:16.525504569Z","id":2,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"locked-ho-acquired","clockId":"0x507c6fffff1fb484","type":"eec","temp":0}
{"timestamp":"2024-03-04T10:42:16.525504569Z","id":3,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"locked-ho-acquired","clockId":"0x507c6fffff1fb484","type":"pps","temp":0}
```
### Dump pins
```bash
[core@cnfde22 go-dpll]$ sudo ./dpll-cli dumpPins
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":0,"clockId":"0x507c6fffff1fb4e8","boardLabel":"CVL-SDP22","panelLabel":"","packageLabel":"","type":"int-oscillator","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":5,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":1,"clockId":"0x507c6fffff1fb4e8","boardLabel":"CVL-SDP20","panelLabel":"","packageLabel":"","type":"int-oscillator","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":4,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":2,"clockId":"0x507c6fffff1fb4e8","boardLabel":"C827_0-RCLKA","panelLabel":"","packageLabel":"","type":"mux","frequency":1953125,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":8,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":3,"clockId":"0x507c6fffff1fb4e8","boardLabel":"C827_0-RCLKB","panelLabel":"","packageLabel":"","type":"mux","frequency":1953125,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":9,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":4,"clockId":"0x507c6fffff1fb4e8","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":3,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":5,"clockId":"0x507c6fffff1fb4e8","boardLabel":"SMA2/U.FL2","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":2,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":6,"clockId":"0x507c6fffff1fb4e8","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":-855.06},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":7,"clockId":"0x507c6fffff1fb4e8","boardLabel":"REF-SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":8,"clockId":"0x507c6fffff1fb4e8","boardLabel":"REF-SMA2/U.FL2","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":9,"clockId":"0x507c6fffff1fb4e8","boardLabel":"PHY-CLK","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":156250000,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":10,"clockId":"0x507c6fffff1fb4e8","boardLabel":"MAC-CLK","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":156250000,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":11,"clockId":"0x507c6fffff1fb4e8","boardLabel":"CVL-SDP21","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":12,"clockId":"0x507c6fffff1fb4e8","boardLabel":"CVL-SDP23","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":13,"clockId":"0x507c6fffff1fb4e8","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":3,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":14,"clockId":"0x507c6fffff1fb4e8","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":3,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":15,"clockId":"0x507c6fffff1fb4e8","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":3,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":16,"clockId":"0x507c6fffff1fb4e8","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":3,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":17,"clockId":"0x507c6fffff1fb484","boardLabel":"CVL-SDP22","panelLabel":"","packageLabel":"","type":"int-oscillator","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":5,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":18,"clockId":"0x507c6fffff1fb484","boardLabel":"CVL-SDP20","panelLabel":"","packageLabel":"","type":"int-oscillator","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":4,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":19,"clockId":"0x507c6fffff1fb484","boardLabel":"C827_0-RCLKA","panelLabel":"","packageLabel":"","type":"mux","frequency":1953125,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":8,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":20,"clockId":"0x507c6fffff1fb484","boardLabel":"C827_0-RCLKB","panelLabel":"","packageLabel":"","type":"mux","frequency":1953125,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":9,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":470.53},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":22,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA2/U.FL2","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":2,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":23,"clockId":"0x507c6fffff1fb484","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":0,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":24,"clockId":"0x507c6fffff1fb484","boardLabel":"REF-SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":25,"clockId":"0x507c6fffff1fb484","boardLabel":"REF-SMA2/U.FL2","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":26,"clockId":"0x507c6fffff1fb484","boardLabel":"PHY-CLK","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":156250000,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":27,"clockId":"0x507c6fffff1fb484","boardLabel":"MAC-CLK","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":156250000,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":28,"clockId":"0x507c6fffff1fb484","boardLabel":"CVL-SDP21","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":29,"clockId":"0x507c6fffff1fb484","boardLabel":"CVL-SDP23","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":30,"clockId":"0x507c6fffff1fb484","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":20,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":31,"clockId":"0x507c6fffff1fb484","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":20,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":32,"clockId":"0x507c6fffff1fb484","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":20,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:42:56.593649599Z","id":33,"clockId":"0x507c6fffff1fb484","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":20,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
```

## Monitor
This will monitor and print both devices and pins notifications

```bash
[core@cnfde22 go-dpll]$ sudo ./dpll-cli monitor
{"timestamp":"2024-03-04T10:43:57.513510052Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":629.21},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:43:57.56025115Z","id":6,"clockId":"0x507c6fffff1fb4e8","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":2370.9},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:43:57.596982527Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":629.21},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:43:57.647676234Z","id":6,"clockId":"0x507c6fffff1fb4e8","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":2370.9},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:43:58.485730147Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":1067.07},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:43:58.522891803Z","id":6,"clockId":"0x507c6fffff1fb4e8","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":2724.32},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T10:43:58.585007033Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":1067.07},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
```
## Output format

Currently only JSON is supported. 
It can be used for filtering with `jq`, for example:

```bash
[core@cnfde22 go-dpll]$ sudo ./dpll-cli dumpPins | jq -c '. | select(.type=="mux") | .boardLabel'
"C827_0-RCLKA"
"C827_0-RCLKB"
"C827_0-RCLKA"
"C827_0-RCLKB"
```