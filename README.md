# watchdog
Watchdog - watches domains and certificates for expiry and connectivity.


Sample Yaml file
Sub domains have to be treated as seperate entries
```
name: TestConfig
refreshinterval: 86400
domains:
    - uuid: ""
      name: www.google.com
      endpoints: ["https://www.google.com"]
    - uuid: ""
      name: www.gmail.com
      endpoints: ["https://www.gmail.com"]
```
## INSTALL

Via go
```
go install github.com/binuud/watchdog/cmd/watchdog@latest
```
OR

Download the binary from the release folder (https://github.com/binuud/watchdog/releases)


## USAGE

create a config.yaml file with the following contents, and place the yaml file as 
he same directory where you are invoking the binary.
```
name: TestConfig
refreshinterval: 86400
domains:
    - uuid: ""
      name: www.google.com
      endpoints: ["https://www.google.com"]
    - uuid: ""
      name: www.gmail.com
      endpoints: ["https://www.gmail.com"]
```

Now run the binary

```
watchdog --file [PATH-TO-FILE]/config.yaml
```

Sample Output
```
usage: watchdog --file [filename-with-path]
Using config file  config.yaml



 _  _  _ _______ _______ _______ _     _ ______   _____   ______
 |  |  | |_____|    |    |       |_____| |     \ |     | |  ____
 |__|__| |     |    |    |_____  |     | |_____/ |_____| |_____|
                                                                
    
┌───┬──────────────────┬─────────────────────────────────────┬───────┬───────────────┐
│   │ DOMAINS          │                CERTS                │   IP  │   REACHABLE   │
│   │                  ├─────┬───────┬─────┬──────┬──────────┼───────┼───────┬───────┤
│   │                  │ TOT │ VALID │ EXP │ DAYS │ VALIDITY │ (NUM) │ VALID │ TOTAL │
├───┼──────────────────┼─────┼───────┼─────┼──────┼──────────┼───────┼───────┼───────┤
│ 1 │ www.google.com   │  3  │   1   │  0  │  64  │ Valid    │   2   │   1   │   1   │
├───┼──────────────────┼─────┼───────┼─────┼──────┼──────────┼───────┼───────┼───────┤
│ 2 │ www.gmail.com    │  3  │   1   │  0  │  64  │ Valid    │   2   │   1   │   1   │
├───┼──────────────────┼─────┼───────┼─────┼──────┼──────────┼───────┼───────┼───────┤
│ 3 │ www.apple.com    │  2  │   1   │  0  │  200 │ Valid    │   3   │   1   │   1   │
├───┼──────────────────┼─────┼───────┼─────┼──────┼──────────┼───────┼───────┼───────┤
│ 4 │ www.dronasys.com │  2  │   1   │  0  │  59  │ Valid    │   1   │   1   │   1   │
├───┼──────────────────┼─────┼───────┼─────┼──────┼──────────┼───────┼───────┼───────┤
│   │                  │     │       │     │      │          │       │       │       │
└───┴──────────────────┴─────┴───────┴─────┴──────┴──────────┴───────┴───────┴───────┘
```
For each domain, it checks the number of endpoints which are reachable.
Then checks all the certificates associated with the domain, checks if the domain association is proper.
Marks the certificate as expiring, if the expiry date is in a 10 day window (from today).
Also displays the number of IPv4 and IPv6 addresses associated with the domain