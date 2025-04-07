# watchdog

![WatchDog](assets/watchdog.png)

Watchdog - watches domains and certificates for expiry, and endpoints connectivity.
Work in progress.

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
Column description
* Mutated - days before last update of whois data
* Expiry - number of days remaining for the domain to expire
* Certs
  * Total - total number of certificates associated with the domain, subdomain
  * Valid - number of valid certificates (name mapping + expiry)
  * Expiry - number of days remaining for the ceritificate to expire
* IP - number of ip associated with the domain
* Reachable - endpoints can be configured in the config.yaml file, each domain can have multiple endpoints, watchdog will test reachability of each endpoint.

## What is this
This is simple tool to display expiry and connectivity information on a small set of domains, subdomains, and endpoints. I wanted a simple tool to look at the following
* Find number of certificates associated with a domain
* Find the latest expiring certificate
* Check if whois data was updated recently (in 10 days)
* Check if domain is going to expire, show a warning if it is expiring in 10 days
* Find number of IPV4/IPv6 addresses associated with the domain
* If there are sub domains, make seperate entries in the config.yaml file, as shown in this readme file
* You might be hosting dev apps, in some intenal endpoint path like example.com/dev/app1, you can add these
end point entries, and the tool will check for its reachability

## Install

### Cli Mode

Via go
```
go install github.com/binuud/watchdog/cmd/watchdog@latest
```
OR

Download the watchdog binary from the release folder (https://github.com/binuud/watchdog/releases)

### Server Mode

Via go
```
go install github.com/binuud/watchdog/cmd/watchdogServer@latest
```
OR

Download the watchdogServer binary from the release folder (https://github.com/binuud/watchdog/releases)

## Usage

Note: whois data caching coming soon, if you run this in loop, whois data providers might block your ip.

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
Hope the yaml is self explanatory.

Now run the binary

```
watchdog --file [PATH-TO-FILE]/config.yaml
```


## Using docker image


```
docker pull dronasys/watchdog
```


To start the grpc and http server 
```
docker run --name WatchDog -p 10090:9090 -p 10080:9080 -v  "./config.yaml:/configs/config.yaml" dronasys/watchdog -d
```
* container port 9090 - grpc
* container port 9080 - http
* mount config file to /configs/config.yaml

When using docker image, the default entrypoint is the server. If you want to 
run once, use the following command, this will give the table output.
```
docker run -v ./config.yaml:/configs/config.yaml --entrypoint /watchDog  dronasys/watchdog  --file /configs/config.yaml   
```

## Connectivity

* GRPC enabled, reflection enabled by default
* HTTP server enabled, [OpenApiSpec](/gen/web/v1/watchdog/openapi.json)
* Typescript interface at /gen/web/v1/watchdog/*.ts
* Docker image by default starts in server mode
* For looking at OpenApiSpec, clone the project, and run below command...
```
docker run  -p 10030:8080 -v ./gen/web/v1/watchdog/openapi.json:/tmp/swagger.json -e SWAGGER_FILE=/tmp/swagger.json docker.swagger.io/swaggerapi/swagger-editor
```
*  After running above command, launch browser http://localhost:10030/

## TODO
* Complete tests
* If user base increases, and domain count increases optimize proto file
* Simple UI
* MCP server coming soon
* In server mode, reload api/rpc has to be called by the client, once caching is enabled, watchdog will updates all domain details in fixed interval

## Credits

| For             | License     | Repo                                    | 
| :---            |    :----    |          :---                           |
| TabulationView  | MIT         | https://github.com/jedib0t/go-pretty    |
| WhoIsParser     | Apache 2.0  | https://github.com/likexian/whois       |
| Pkg Release     | NA          | https://goreleaser.com                  |

Usage videos will be uploaded here, this tool will be available as a AI Module on the BrainUI soon
* https://www.youtube.com/@dronasystems/shorts 
* https://www.youtube.com/@dronasystems/videos