
Sample Output
```
usage: watchdog --config-file [filename-with-path] --code_root_folder [absolute path of folder containing code]
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