# watchdog

![WatchDog](assets/watchdog.gif)

* [UI Documentation](https://github.com/binuud/watchdog-ui/blob/master/docs/UI-HELP.md)
* [Config.yaml Documentation](docs/CONFIG.md)
* [API Documentation](docs/API.md)
* [TODO Documentation](docs/TODO.md)

Watchdog is a dashboard, for showing domains and certificates for expiry, and endpoints connectivity. Watches a configured list of
projects in your local code base, and reads information from remote github server.

* Domains
  * Show connectivity status of configured list of domains
  * Shows certification validity 
  * Show reachability of multiple configured endpoints per domain
  * Show whois database for the domain
  * Show if whois data was modified recently in under 30 days
* Projects 
  * Show list of configured git projects from your local repo
  * Shows open issues against the project
  * Show open PR's against the project
  * Show number of modified and new files in the repo, on local machine
  * Compare against remote github, to show if master can be merged
  * Show if remote is ahead or behind local master

Work in progress.

## Install

* Local build by Cloning the project and run, docker is needed. Change the ./devcontainer/devcontainer.json to change the mount path.
This program needs to access all your git folders, that you want part of the dashboard.

```
make run
```

* Using go install
```
go install github.com/binuud/watchdog/cmd/watchdog@latest
```

* Download the watchdog binary from the release folder (https://github.com/binuud/watchdog/releases)

* Access UI, after running the binary, using address
```
http://localhost:9080
```

## Usage

Create a config.yaml file with the contents, as explained in this 
[Config.yaml Documentation](docs/CONFIG.md)

Alternatively, use the download the [sample config](configs/sampleConfig.yaml), and make changes as required

Note: whois data caching coming soon, if you run this in loop, whois data providers might block your ip.

### Go Binary
```
./watchdog -grpc_port 9091 -v -http_port 9081 -GITHUB_TOKEN $(GITHUB_TOKEN) -config_file ~/Code/watchdog/watchdog/configs/config.yaml -code_root_folder ~/Code
```
Command Line options

* -grpc_port Grpc Port, defaults to 9090
* -http_port Http Port, defaults to 9080, both the api and web ui interface is exposed in this port
* -config-file absolute path to the config file with name
* -code_root_folder relative path to the folder from HOME dir, which contains all local git rep 
* -GITHUB_TOKEN the github token, with read access to all projects, for the specified user


### Docker Image

```
docker pull dronasys/watchdog
```

To run once with the given config, use the following command, this will give the table output.
```
docker run -v ./config.yaml:/configs/config.yaml --entrypoint /watchDog  dronasys/watchdog  --file /configs/config.yaml   
```

To start the grpc and http server 
```
docker run --name WatchDog -p 10090:9090 -p 10080:9080 -d -v  "./config.yaml:/configs/config.yaml" dronasys/watchdog
```
* container port 9090 - grpc
* container port 9080 - http
* mount config file to /configs/config.yaml

When using docker image, the default entrypoint is the 'server'.

## Local Development

Checkout the UI project, is the parent folder of this project, so watchdog (this project) and watchdog-ui (ui project) are at the same level in the parent folder.

```
git clone https://github.com/binuud/watchdog-ui.git
```

When you are developing on this, or want to make changes, 
use below command to start a golang container, and mount the source code within the container
```
make start-container
```

from within the container, run the make command to select various targets
```
make
```

## Credits

| For             | License     | Repo                                    | 
| :---            |    :----    |          :---                           |
| TabulationView  | MIT         | https://github.com/jedib0t/go-pretty    |
| WhoIsParser     | Apache 2.0  | https://github.com/likexian/whois       |
| Pkg Release     | NA          | https://goreleaser.com                  |
| UI Project      | MIT         | https://github.com/binuud/watchdog-ui   |

Usage videos will be uploaded here, this tool will be available as a AI Module on the BrainUI soon
* https://www.youtube.com/@dronasystems/shorts 
* https://www.youtube.com/@dronasystems/videos


## VSCode

* Dev Containers - for launching development containers

## .local.env
```
GITHUB_TOKEN=[GITHUB API TOKEN]
```