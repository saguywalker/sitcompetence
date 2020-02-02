<p align="center">
  <a href="https://github.com/saguywalker/sitcompetence/actions"><img alt="GitHub Actions status" src="https://https://github.com/saguywalker/sitcompetence/workflows/Go/badge.svg"></a>
</p>

# SIT-Competence Installation Manual
Operating system: Ubuntu 18.04
## SITCOMCHAIN
Blockchain's part of a SIT-Competence using tendermint on each node. While blockchain running, each node will automatically connect to each other.

### Prerequisites
**1. Install required dependencies**
```
$ sudo apt-get update
$ sudo apt install build-essential unzip git
```
**2. Install Go and Setup GOPATH for Ubuntu and Debian distros**
Install Go
```
$ wget https://dl.google.com/go/go1.13.5.linux-amd64.tar.gz
$ mkdir -p /usr/local && tar xvf go1.13.5.linux-amd64.tar.gz && sudo mv go /usr/local
```

Setup Go environment variables
```
$ mkdir -p ~/gofolder
$ echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
$ echo 'export GOPATH=$HOME/gofolder' >> ~/.bashrc && source ~/.bashrc
$ echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> ~/.bashrc && source ~/.bashrc
```


**3. Install Tendermint v0.32.7**

> For Ubuntu and Debian distros
```
$ wget https://github.com/tendermint/tendermint/releases/download/v0.32.7/tendermint_v0.32.7_linux_amd64.zip

$ unzip tendermint_v0.32.7_linux_amd64.zip -d ~/tendermint
$ echo 'export PATH=$PATH:$HOME/tendermint' >> ~/.bashrc && source ~/.bashrc
 ```
 
 
 
 
## Usage
**1. Generate validator key, node key and genesis file**
```
$ tendermint init
```

**1.2  For more than 1 node, set node id and their corresponding IP address and port to persistent_peers variable in ~/.tendermint/config/config.toml in format => persistent_peers = "{NODEID}@{IP}:{Port}"
get node's id**
```
$ tendermint show_node_id
```
example in ~/.tendermint/config/config.toml
```
$ persistent_peers = "5a3b1b228d558235d5a8c76c28ecef13e6ad55f2@10.4.56.17:26656,31c219dd725aa371052c2d9b8c1f12de13ed4591@10.4.56.22:26656,8369dfd9f8cedf85db929186fade7054175a4cf1@10.4.56.23:26656"
```

**1.3  You may set create_empty_blocks = false in config.toml to prevent unnecessary in producing empty block.**

Run tendermint node
```
$ tendermint node
```

Open another tab to run a SITCOMCHAIN smart-contract
```
$ mkdir -p $GOPATH/src/github.com/saguywalker
$ cd $GOPATH/src/github.com/saguywalker
$ git clone https://github.com/saguywalker/sitcomchain
$ cd sitcomchain && go build && ./sitcomchain
```

## Web application
The web application part consists of API, database, user interface, and third-party service.
### Prerequisites

**1. Install Docker**
To install Docker on Ubuntu
```
$ sudo apt-get update && sudo apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
$ sudo apt-get update && sudo apt-get install docker-ce docker-ce-cli containerd.io
```

To install Docker in other operating systems, please follow the instruction on the Docker official website through this link https://docs.docker.com/install/.

**2. Install Docker-Compose (see: https://docs.docker.com/compose/install)**
```
$ sudo curl -L "https://github.com/docker/compose/releases/download/1.25.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

$ sudo chmod +x /usr/local/bin/docker-compose
```

**3. Clone SIT-Competence repository**
```
$ cd ~ && git clone https://github.com/saguywalker/sitcompetence.git
$ cd sitcompetence
```

**4. Setup environment variables**
```
$ ./setup.sh
```

**5. Set Machine IP Address to .env**

Please input your machine IP in the {machine-ip} section.
If you don’t know your machine IP, please do the following:
```
$ ifconfig
```

The result will be something look like this:

On the eth0 section, the underlined IP address is your machine IP. Copy it and paste in the {machine-ip} section.
```
$ echo -e "\nVUE_APP_BASE_IP={machine-ip}"  >> ./frontend/.env
```
Example: (IP from the result above)
```
$ echo -e "\nVUE_APP_BASE_IP=192.168.1.231"  >> ./frontend/.env
```

### Usage

**1. Whitelist the project’s port in the firewall (if applicable)**
```
$ sudo ufw allow http && sudo ufw allow 3000
$ sudo ufw reload
```

**2. To start the web application**
```
$ cd ~/sitcompetence
$ sudo docker-compose up -d
```

Check the result by visiting http://{machine-ip}

### After finish all installation above
**To summarize** <br>
1 Tab for tendermint (Blockchain)<br>
To start 
```
$ cd ~ && tendermint node
```
To stop
```
Ctrl + C
```

1 Tab for smart-contract (Blockchain)<br>
To start 
```
$ cd sitcomchain && ./sitcomchain
```
To stop
```
Ctrl + C
```

1 Tab for web application<br>
To start
```
$ cd ~/sitcompetence
$ sudo docker-compose up -d
```
To stop
```
$ cd ~/sitcompetence
$ sudo docker-compose down
```
