# watchdog
Watchdog - watches domains and certificates for expiry and connectivity.


## NOT READY FOR CONSUMPTION


Sample Yaml file
```
name: TestConfig
refreshinterval: 86400
domains:
    - uuid: ""
      name: www.google.com
      endpoints: ["https://www.google.com"]
    - uuid: ""
      name: www.gmail.com
      endpoints: ["https://www.gmail.com]
```