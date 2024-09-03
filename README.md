# Netlink DPLL test container #

This repo helps to build a container image for netlink DPLL debug

## Bulid ##

```bash
export IMG='quay.io/vgrinber/tools:dpll'
```
(replace with your repository)

```bash
podman build -t $IMG --no-cache -f Containerfile  . && podman push $IMG
```

## Hack ##
The [Containerfile](Containerfile) specifies two hacks done to the kernel tools here:

### 1. Add DPLL YNL spec

```
# Comment out this line when dpll.yaml will be upstream
COPY dpll.yaml /linux/Documentation/netlink/specs/dpll.yaml
```

### 2. Make monitoring call blocking

The upstream implementation of `ynl.py` opens the monitoring socket in non-blocking mode, which means that if there is no notification at the exact same moment you run the cli, it will exit.
If we want to wait for notifications, this hack must be done:

```
# Uncomment this line if you want cli to block while waiting on netlink notifications
RUN sed -i 's/, socket.MSG_DONTWAIT//g' lib/ynl.py 
```

## Use ##

### Use netlink cli from linux kernel tools ###

```bash
oc debug no/cnfde21.ptp.lab.eng.bos.redhat.com --image=quay.io/vgrinber/tools
Starting pod/cnfde21ptplabengbosredhatcom-debug-tl8hl ...
To use host binaries, run `chroot /host`
Pod IP: 10.16.230.5
If you don't see a command prompt, try pressing enter.

sh-5.1# python3 cli.py --spec /linux/Documentation/netlink/specs/dpll.yaml --dump device-get
[{'clock-id': 5799633565433967664,
  'id': 0,
  'lock-status': 'locked-ho-acq',
  'mode': 'automatic',
  'mode-supported': ['automatic'],
  'module-name': 'ice',
  'type': 'eec'},
 {'clock-id': 5799633565433967664,
  'id': 1,
  'lock-status': 'locked-ho-acq',
  'mode': 'automatic',
  'mode-supported': ['automatic'],
  'module-name': 'ice',
  'type': 'pps'},
 {'clock-id': 5799633565433966964,
  'id': 2,
  'lock-status': 'unlocked',
  'mode': 'automatic',
  'mode-supported': ['automatic'],
  'module-name': 'ice',
  'type': 'eec'},
 {'clock-id': 5799633565433966964,
  'id': 3,
  'lock-status': 'unlocked',
  'mode': 'automatic',
  'mode-supported': ['automatic'],
  'module-name': 'ice',
  'type': 'pps'}]

sh-5.1# python3 cli.py --spec /linux/Documentation/netlink/specs/dpll.yaml --subscribe monitor

```
The last command above will block waiting for a notification.

### Use GoLang dpll-cli ###

```bash
§ oc debug no/cnfde22.ptp.lab.eng.bos.redhat.com  --image=quay.io/vgrinber/tools:dpll
Starting pod/cnfde22ptplabengbosredhatcom-debug-762vj ...
To use host binaries, run `chroot /host`
Pod IP: 10.16.230.7
If you don't see a command prompt, try pressing enter.
sh-5.1# dpll-cli 
Interact with dpll devices via netlink

Get and monitor devices and pins.

Usage:
  dpll-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dumpDevices Dump DPLL devices
  dumpPins    Dump DPLL pins
  help        Help about any command
  monitor     Monitors DPLL events

Flags:
  -h, --help   help for dpll-cli

Use "dpll-cli [command] --help" for more information about a command.
```
#
```bash
sh-5.1§ dpll-cli dumpDevices
{"timestamp":"2024-03-04T13:18:26.12219909Z","id":0,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"locked-ho-acquired","clockId":"0x507c6fffff1fb4e8","type":"eec","temp":0}
{"timestamp":"2024-03-04T13:18:26.12219909Z","id":1,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"locked-ho-acquired","clockId":"0x507c6fffff1fb4e8","type":"pps","temp":0}
{"timestamp":"2024-03-04T13:18:26.12219909Z","id":2,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"locked-ho-acquired","clockId":"0x507c6fffff1fb484","type":"eec","temp":0}
{"timestamp":"2024-03-04T13:18:26.12219909Z","id":3,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"locked-ho-acquired","clockId":"0x507c6fffff1fb484","type":"pps","temp":0}
```
#
```bash
sh-5.1§ dpll-cli dumpPins   
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":0,"clockId":"0x507c6fffff1fb4e8","boardLabel":"CVL-SDP22","panelLabel":"","packageLabel":"","type":"int-oscillator","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":5,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":1,"clockId":"0x507c6fffff1fb4e8","boardLabel":"CVL-SDP20","panelLabel":"","packageLabel":"","type":"int-oscillator","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":4,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":2,"clockId":"0x507c6fffff1fb4e8","boardLabel":"C827_0-RCLKA","panelLabel":"","packageLabel":"","type":"mux","frequency":1953125,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":8,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":3,"clockId":"0x507c6fffff1fb4e8","boardLabel":"C827_0-RCLKB","panelLabel":"","packageLabel":"","type":"mux","frequency":1953125,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":9,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":4,"clockId":"0x507c6fffff1fb4e8","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":3,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":5,"clockId":"0x507c6fffff1fb4e8","boardLabel":"SMA2/U.FL2","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":2,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":6,"clockId":"0x507c6fffff1fb4e8","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":-1025.68},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":7,"clockId":"0x507c6fffff1fb4e8","boardLabel":"REF-SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":8,"clockId":"0x507c6fffff1fb4e8","boardLabel":"REF-SMA2/U.FL2","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":9,"clockId":"0x507c6fffff1fb4e8","boardLabel":"PHY-CLK","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":156250000,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":10,"clockId":"0x507c6fffff1fb4e8","boardLabel":"MAC-CLK","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":156250000,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":11,"clockId":"0x507c6fffff1fb4e8","boardLabel":"CVL-SDP21","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":12,"clockId":"0x507c6fffff1fb4e8","boardLabel":"CVL-SDP23","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change","pinParentDevice":{"parentId":1,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":13,"clockId":"0x507c6fffff1fb4e8","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":3,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":14,"clockId":"0x507c6fffff1fb4e8","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":3,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":15,"clockId":"0x507c6fffff1fb4e8","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":3,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":16,"clockId":"0x507c6fffff1fb4e8","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":3,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":17,"clockId":"0x507c6fffff1fb484","boardLabel":"CVL-SDP22","panelLabel":"","packageLabel":"","type":"int-oscillator","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":5,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":18,"clockId":"0x507c6fffff1fb484","boardLabel":"CVL-SDP20","panelLabel":"","packageLabel":"","type":"int-oscillator","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":4,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":19,"clockId":"0x507c6fffff1fb484","boardLabel":"C827_0-RCLKA","panelLabel":"","packageLabel":"","type":"mux","frequency":1953125,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":8,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":20,"clockId":"0x507c6fffff1fb484","boardLabel":"C827_0-RCLKB","panelLabel":"","packageLabel":"","type":"mux","frequency":1953125,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":9,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":293.35},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":22,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA2/U.FL2","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":2,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":23,"clockId":"0x507c6fffff1fb484","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":0,"state":"selectable","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":24,"clockId":"0x507c6fffff1fb484","boardLabel":"REF-SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":25,"clockId":"0x507c6fffff1fb484","boardLabel":"REF-SMA2/U.FL2","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":26,"clockId":"0x507c6fffff1fb484","boardLabel":"PHY-CLK","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":156250000,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":27,"clockId":"0x507c6fffff1fb484","boardLabel":"MAC-CLK","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":156250000,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":28,"clockId":"0x507c6fffff1fb484","boardLabel":"CVL-SDP21","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":29,"clockId":"0x507c6fffff1fb484","boardLabel":"CVL-SDP23","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change","pinParentDevice":{"parentId":3,"direction":"output","prio":0,"state":"connected","phaseOffsetPs":0},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-480307,"phaseAdjustMax":480307,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":30,"clockId":"0x507c6fffff1fb484","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":20,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":31,"clockId":"0x507c6fffff1fb484","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":20,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":32,"clockId":"0x507c6fffff1fb484","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":20,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:35.556823556Z","id":33,"clockId":"0x507c6fffff1fb484","boardLabel":"","panelLabel":"","packageLabel":"","type":"synce-eth-port","frequency":0,"frequencySupported":{"frequencyMin":0,"frequencyMax":0},"capabilities":"state-can-change","pinParentDevice":{"parentId":0,"direction":"","prio":0,"state":"","phaseOffsetPs":0},"pinParentPin":{"parentId":20,"parentState":"disconnected"},"phaseAdjustMin":0,"phaseAdjustMax":0,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
```
#
```bash
sh-5.1§ dpll-cli monitor 
{"timestamp":"2024-03-04T13:18:45.463926946Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":-140.83},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:45.501497332Z","id":6,"clockId":"0x507c6fffff1fb4e8","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":1819.96},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:45.554733964Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":-140.83},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:45.601119257Z","id":6,"clockId":"0x507c6fffff1fb4e8","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":1819.96},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:46.485397444Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":178.72},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:46.553495107Z","id":6,"clockId":"0x507c6fffff1fb4e8","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":-1552.09},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:46.591702689Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":178.72},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:46.641055384Z","id":6,"clockId":"0x507c6fffff1fb4e8","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":-1552.09},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:47.488435035Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":-179.41},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:47.550756174Z","id":6,"clockId":"0x507c6fffff1fb4e8","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":3658},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:47.607000476Z","id":21,"clockId":"0x507c6fffff1fb484","boardLabel":"SMA1","panelLabel":"","packageLabel":"","type":"ext","frequency":1,"frequencySupported":{"frequencyMin":10000000,"frequencyMax":10000000},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":3,"state":"connected","phaseOffsetPs":-179.41},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
{"timestamp":"2024-03-04T13:18:47.655550104Z","id":6,"clockId":"0x507c6fffff1fb4e8","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":1,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":3658},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":0,"fractionalFrequencyOffset":0,"moduleName":"ice"}
^C
sh-5.1§

```

## Run on any host
```bash
[core@cnfdg33 ~]$ sudo podman run --privileged --network=host --rm quay.io/vgrinber/tools:dpll  /usr/local/bin/dpll-cli dumpDevices
Trying to pull quay.io/vgrinber/tools:dpll...
Getting image source signatures
Copying blob 2460cdd52686 done  
Copying blob 975c02cd39cf done  
Copying blob 85b7afa82422 done  
Copying blob 7f74bb9cc326 done  
Copying config fc1fd6310c done  
Writing manifest to image destination
Storing signatures
{"timestamp":"2024-08-18T13:02:01.413781385Z","id":0,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"unlocked","clockId":"0x507c6fffff1fb4e8","type":"eec","temp":0}
{"timestamp":"2024-08-18T13:02:01.413781385Z","id":1,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"locked-ho-acquired","clockId":"0x507c6fffff1fb4e8","type":"pps","temp":0}
{"timestamp":"2024-08-18T13:02:01.413781385Z","id":2,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"locked-ho-acquired","clockId":"0x507c6fffff1fb484","type":"eec","temp":0}
{"timestamp":"2024-08-18T13:02:01.413781385Z","id":3,"moduleName":"ice","mode":"automatic","modeSupported":"automatic","lockStatus":"locked-ho-acquired","clockId":"0x507c6fffff1fb484","type":"pps","temp":0}
[core@cnfdg33 ~]$ sudo podman run --privileged --network=host --rm quay.io/vgrinber/tools:dpll  /usr/local/bin/dpll-cli getPinId -i 0x507c6fffff1fb484 -l GNSS-1PPS
23
[core@cnfdg33 ~]$ sudo podman run --privileged --network=host --rm quay.io/vgrinber/tools:dpll  /usr/local/bin/dpll-cli getPin -i 23 
{"timestamp":"2024-08-18T13:02:36.680281604Z","id":23,"clockId":"0x507c6fffff1fb484","boardLabel":"GNSS-1PPS","panelLabel":"","packageLabel":"","type":"gnss","frequency":1,"frequencySupported":{"frequencyMin":1,"frequencyMax":1},"capabilities":"state-can-change,priority-can-change","pinParentDevice":{"parentId":3,"direction":"input","prio":0,"state":"connected","phaseOffsetPs":-2773.38},"pinParentPin":{"parentId":0,"parentState":""},"phaseAdjustMin":-16723,"phaseAdjustMax":16723,"phaseAdjust":7000,"fractionalFrequencyOffset":0,"moduleName":"ice"}
[core@cnfdg33 ~]$ sudo podman run --privileged --network=host --rm quay.io/vgrinber/tools:dpll  /usr/local/bin/dpll-cli setPhaseAdjust -i 23 -p 3000
[core@cnfdg33 ~]$ sudo podman run --privileged --network=host --rm quay.io/vgrinber/tools:dpll  /usr/local/bin/dpll-cli getPin -i 23 | jq '.phaseAdjust'
3000
[core@cnfdg33 ~]$ 

```
or, using the linux kernel netlink tool
```bash
[core@cnfdg32 ~]$ sudo podman run --privileged --network=host --rm quay.io/vgrinber/tools:dpll   python3 cli.py --spec /linux/Documentation/netlink/specs/dpll.yaml --dump device-get
[{'clock-id': 5799633565433967664,
  'id': 0,
  'lock-status': 'holdover',
  'mode': 'automatic',
  'mode-supported': ['automatic'],
  'module-name': 'ice',
  'type': 'eec'},
 {'clock-id': 5799633565433967664,
  'id': 1,
  'lock-status': 'holdover',
  'mode': 'automatic',
  'mode-supported': ['automatic'],
  'module-name': 'ice',
  'type': 'pps'},
 {'clock-id': 5799633565433966964,
  'id': 2,
  'lock-status': 'locked-ho-acq',
  'mode': 'automatic',
  'mode-supported': ['automatic'],
  'module-name': 'ice',
  'type': 'eec'},
 {'clock-id': 5799633565433966964,
  'id': 3,
  'lock-status': 'locked-ho-acq',
  'mode': 'automatic',
  'mode-supported': ['automatic'],
  'module-name': 'ice',
  'type': 'pps'}]

```
