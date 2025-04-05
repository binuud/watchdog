# watchdog
Watchdog - watches domains and certificates for expiry and connectivity.

## INSTALL

Via go
```
go install github.com/binuud/watchdog/cmd/watchdog@latest
```
OR

Download the binary from the release folder (https://github.com/binuud/watchdog/releases)


## USAGE

create a config.yaml file with the following contents.

Sample Yaml file
Sub domains have to be treated as seperate entries
```
name: MyDomains
refreshinterval: 86400
domains:
    - uuid: ""
      name: www.google.com
      domainname: google.com
      endpoints:
        - https://www.google.com
        - https://www.google.com/?client=safari
    - uuid: ""
      name: www.gmail.com
      domainname: gmail.com
      endpoints:
        - https://www.gmail.com
    - uuid: ""
      name: www.dronasys.com
      domainname: dronasys.com
      endpoints:
        - https://www.dronasys.com
```

sample entry for subdomain
```
    - uuid: ""
      name: dev.example.com
      endpoints:
        - https://dev.example.com
        - https://dev.example.com/test1
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
                                                                
    
Fetching data... (Single Thread)
┌───┬─────────────────────────────────────┬─────────────────────────────────┬───────┬───────────────┐
│   │               DOMAINS               │              CERTS              │   IP  │   REACHABLE   │
│   ├──────────────────┬─────────┬────────┼─────┬───────┬────────┬──────────┼───────┼───────┬───────┤
│   │ NAME/SUB         │ MUTATED │ EXPIRY │ TOT │ VALID │ EXPIRY │ VALIDITY │ (NUM) │ VALID │ TOTAL │
│   ├──────────────────┼─────────┴────────┼─────┴───────┼────────┼──────────┴───────┴───────┴───────┤
│   │                  │      (DAYS)      │             │ (DAYS) │                                  │
├───┼──────────────────┼─────────┬────────┼─────┬───────┼────────┼──────────┬───────┬───────┬───────┤
│ 1 │ www.google.com   │  2034 ✓ │ 1257 ✓ │  3  │   1   │ 68   ✓ │   Valid  │   2   │   2   │     2 │
│   │                  │         │        │     │       │        │          │       │       │       │
├───┼──────────────────┼─────────┼────────┼─────┼───────┼────────┼──────────┼───────┼───────┼───────┤
│ 2 │ mail.google.com  │         │        │  3  │   1   │ 68   ✓ │   Valid  │   2   │   1   │     1 │
│   │                  │         │        │     │       │        │          │       │       │       │
├───┼──────────────────┼─────────┼────────┼─────┼───────┼────────┼──────────┼───────┼───────┼───────┤
│ 3 │ www.gmail.com    │  267  ✓ │ 128  ✓ │  3  │   1   │ 68   ✓ │   Valid  │   2   │   1   │     1 │
│   │                  │         │        │     │       │        │          │       │       │       │
├───┼──────────────────┼─────────┼────────┼─────┼───────┼────────┼──────────┼───────┼───────┼───────┤
│ 4 │ www.apple.com    │  73   ✓ │ 320  ✓ │  2  │   1   │ 194  ✓ │   Valid  │   3   │   1   │     1 │
│   │                  │         │        │     │       │        │          │       │       │       │
├───┼──────────────────┼─────────┼────────┼─────┼───────┼────────┼──────────┼───────┼───────┼───────┤
│ 5 │ www.dronasys.com │  744  ✓ │ 837  ✓ │  2  │   1   │ 53   ✓ │   Valid  │   1   │   1   │     1 │
│   │                  │         │        │     │       │        │          │       │       │       │
├───┼──────────────────┼─────────┼────────┼─────┼───────┼────────┼──────────┼───────┼───────┼───────┤
│   │                  │         │        │     │       │        │          │       │       │       │
└───┴──────────────────┴─────────┴────────┴─────┴───────┴────────┴──────────┴───────┴───────┴───────┘
```
For each domain, it checks the number of endpoints which are reachable.
Then checks all the certificates associated with the domain, checks if the domain association is proper.
Marks the certificate as expiring, if the expiry date is in a 10 day window (from today).
Also displays the number of IPv4 and IPv6 addresses associated with the domain


## Using docker image

When using docker image, the default entrypoint of the docker image is the server. So if you want to 
run once, use the following command
```
docker run -v ./config.yaml:/configs/config.yaml --entrypoint /watchDog  dronasys/watchdog  --file /configs/config.yaml   
```

To start the grpc and http server 
```
docker run --name WatchDog -p 10090:9090 -p 10080:9080 -v  "./config.yaml:/configs/config.yaml" dronasys/watchdog -d
```

## TODO
* Complete tests
* If user base increases, and domain count increases optimize proto file

## Credits

| For             | License     | Repo                                    | 
| :---            |    :----    |          :---                           |
| TabulationView  | MIT         | https://github.com/jedib0t/go-pretty    |
| WhoIsParser     | Apache 2.0  | https://github.com/likexian/whois       |