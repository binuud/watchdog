# Config.yaml

* name: MyDomains - Name give to this config.yaml file
* refreshinterval: 86400 - intervals, in which domain and project data will be refreshed
* domains - list of domains
  * uuid - for tracking purpose
  * name - name to be used for display purpose
  * domain - actual domain name
  * endpoints - set of endpoints, which are checked for reachability
* projects - list of github projects on the local machine, this tool needs access to all the project
  * uuid - for tracking purpose
  * name - name to be used for display purpose
  * org - the org name of the repo
  * user - current local user of the repo (token is needed)
  * path - relative path to home dir, of the local repo code
  * repo - remote git path of the repo

Sample config.yaml
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
projects:
  - uuid: ""
    name: dev-tools
    org: binuud
    user: binuud
    path: Code/local/dev-tools
    repo: https://github.com/binuud/dev-tools.git
```