#### inspect

inspect command line is a utility that gives a
brief overview on current state of system resource
usage.

Supported platforms: linux, MacOSX 10.9

inspect aims to evolve to be an intelligent tool that
can spot problems.

examples: 

  * process X is throttled on CPU because of cgroup restrictions
  * process Y is causing IO contention
  * process X rejecting connections because TCP is out of memory
  * CPU spike at 14:00 UTC. Top users: A, B


###### Installation

1. Get go
2. go get -v -u github.com/measure/inspect # fetches packages and builds

###### Documentation (WIP)

http://godoc.org/github.com/measure/inspect

http://godoc.org/github.com/measure/os

http://godoc.org/github.com/measure/metrics


###### Usage

###### Command line

./bin/inspect

```
--------------------------
total: cpu: 100.0%, mem: 9.9% (379.86MB/3.74GB)
Top processes by CPU usage:
cpu: 65.7%  command: (perl) user: s pid: 23140
cpu: 14.9%  command: (fio) user: s pid: 23214
cpu: 10.0%  command: (perl) user: root pid: 23162
cpu: 6.0%  command: (fio) user: s pid: 23212
cpu: 3.0%  command: (inspect) user: s pid: 23116
---
Top processes by Mem usage:
mem: 16.68MB command: (fio) user: s pid: 23212
mem: 11.36MB command: (tmux) user: s pid: 29769
mem: 10.97MB command: (bash) user: s pid: 15146
mem: 7.95MB command: (zsh) user: s pid: 13572
mem: 7.34MB command: (bash) user: s pid: 6478
---
Top processes by IO usage:
io: 14.47MB/s command: (carbon-cache.py) user: apache pid: 19747
io: 8.00KB/s command: (svlogd) user: nobody pid: 1757
io: 0.00B/s command: (carbon-relay.py) user: apache pid: 1792
io: 0.00B/s command: (carbon-relay.py) user: apache pid: 1796
io: 0.00B/s command: (watchdog/1) user: root pid: 10
---
diskio: sr0 usage: 0.0%
diskio: sda usage: 0.0%
diskio: sdb usage: 92.7%
fs: / usage: 84.5%
fs: /data usage: 63.1%
---
iface: lo TX: 0.00b/s, RX: 0.00b/s
iface: eth0 TX: 6.77Kb/s, RX: 1.03Kb/s
---
cgroup:small cpu: 10.0% cpu_throttling: 79.6% (0.1/1) mem: 0.1% (308.00KB/409.60MB)
---
Problem:  Disk IO usage on (sdb): 92.7%
Problem:  CPU throttling on cgroup(small): 79.6%
Problem:  CPU usage > 80%
```

###### Server 

*inspect* can be run in server mode to run continously and expose metrics via HTTP JSON api

./bin/inspect  -server -address :12345

```
s@c62% curl localhost:12345/metrics.json 2>/dev/null
[
{"type": "gauge", "name": "memstat.Mapped", "value": 16314368.000000},
{"type": "gauge", "name": "memstat.HugePages_Rsvd", "value": 0.000000},
{"type": "gauge", "name": "diskstat.sr0.IOInProgress", "value": 0.000000},
{"type": "gauge", "name": "memstat.cgroup.small.Inactive_anon", "value": 0.000000},
....... truncated
{"type": "counter", "name": "diskstat.sdb.ReadSectors", "value": 7288530, "rate": 0.000000},
{"type": "counter", "name": "interfacestat.eth0.TXpackets", "value": 6445308, "rate": 4.333320},
{"type": "counter", "name": "interfacestat.eth0.TXframe", "value": 0, "rate": 0.000000},
{"type": "counter", "name": "pidstat.pid1.Utime", "value": 31, "rate": 0.000000},
{"type": "counter", "name": "pidstat.pid29769.Utime", "value": 74296, "rate": 0.000000}]
```

###### Todo
  * Refactor to make inspect command line work with any application type
  * TESTS
  * Rules for inspection need to seperated out into user supplied code/config. Currently inspect command line has hard-coded guesswork
  * PerProcessStat on darwin doesn't include optimizations done for Linux. 
  * Add intelligence to find problems. Start with easy ones like CPU usage
  * Command line utility needs much nicer formatting and options to dig into per process/cgroup details
  * Add io metrics per process (need root priviliges)
  * Add caching support to reduce load when multiple invocations of inspect happen.
  * API to collect and expose historical/current statistics
