<p align="center">
  <a href="https://github.com/saguywalker/sitcompetence/actions"><img alt="GitHub Actions status" src="https://https://github.com/saguywalker/sitcompetence/workflows/Go/badge.svg"></a>
</p>

# SIT-COMPETENCE

## Prerequisites
0. Install required dependencies
    ```bash
    sudo apt get update
    sudo apt get install build-essential git
    ```
1. Install Docker (see: https://docs.docker.com/install)
    ```bash
    # Example in Ubuntu
    sudo apt-get update
    sudo apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
    sudo add-apt-repository \
      "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
      $(lsb_release -cs) \
      stable"
    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io
    ```
 2. Install Docker-Compose (see: https://docs.docker.com/compose/install)
    ```bash
    sudo curl -L "https://github.com/docker/compose/releases/download/1.25.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
    ```
    
## Usage
1. Clone repository
    ```bash
    git clone https://github.com/saguywalker/sitcompetence
    cd sitcompetence
    ```
2. Downlaod config.yaml and init.sql and paste to root directory, and .env to frontend directory
3. Run docker-compose with [local-docker-compose.yml](https://github.com/saguywalker/sitcompetence/blob/master/local-docker-compose.yml) 
    ```bash
    sudo docker-compose -f local-docker-compose.yml up
    ```
4. Web application can be accessed with the mapped IP-Address with Port 8080.
   You can run the below command to see the IP-Address.
    ```bash
    sudo docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ${CONTAINTER_ID}
    ```
   
