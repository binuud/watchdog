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

## USAGE

Download the binary from the release folder (https://github.com/binuud/watchdog/releases)

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
./watchdog
```
Note: The binary has to be invoked from the same path, as where the config.yaml file is placed

Sample Output
```
 _  _  _ _______ _______ _______ _     _ ______   _____   ______
 |  |  | |_____|    |    |       |_____| |     \ |     | |  ____
 |__|__| |     |    |    |_____  |     | |_____/ |_____| |_____|


INFO[0000] Received response 200 https://www.google.com
INFO[0000] Resolving IP Address ([2404:6800:4007:80d::2004 142.250.195.100])
INFO[0000] Ip address of domain www.google.com 2404:6800:4007:80d::2004
INFO[0000] Ip address of domain www.google.com 142.250.195.100
INFO[0001] Received response 200 https://www.gmail.com
INFO[0001] Resolving IP Address ([2404:6800:4007:82a::2005 142.250.196.37])
INFO[0001] Ip address of domain www.gmail.com 2404:6800:4007:82a::2005
INFO[0001] Ip address of domain www.gmail.com 142.250.196.37


PrintSummary
--------------------------------------------------------------------------------
Domain                        Certs              IP   Reachability
                      Tot  Val  Exp   Status    Num   ( Vld/Tot )
--------------------------------------------------------------------------------
www.google.com          3    1    0    Valid      2   (   1/1   )
www.gmail.com           3    1    0    Valid      2   (   1/1   )

--------------------------------------------------------------------------------
```
For each domain, it checks the number of endpoints which are reachable.
Then checks all the certificates associated with the domain, checks if the domain association is proper.
Marks the certificate as expiring, if the expiry date is in a 10 day window (from today).
Also displays the number of IPv4 and IPv6 addresses associated with the domain