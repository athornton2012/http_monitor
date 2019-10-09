# http_monitor

## How to Run
```
git clone https://github.com/athornton2012/http_monitor
cd http_monitor
go build main.go
./main --logFilePath='/Full/path/to/csv/file' --alertThreshold <threshold>
```

## TODO
1. Tests for the monitor packages are sparse. I wanted to use counterfeiter to fake out the `StatMonitor` and `RollingTrafficList` but ran out of time
2. Didn't spend a whole lot of time on displaying useful debugging information for every 10 seconds of loglines.
3. Realize I never finished the message for the high traffic alert, but given my implementation it would be trivial to add :)


## Assumptions
1. I made an assumption in the 10 second window that the clock skew would never be more than 20 seconds

## Possible Improvements
1. Could process multiple loglines at once if efficiency wasn't cutting it. Ie have several log monitors pulling values from a channel
1. Could Make the StatList more resusable
1. Perhaps combine the StatList data structure into the RollingTrafficList...though there is more to think about there
