# valcloud
Validation Cloud Take Home Challenge

### Part 1 - Template the deployment of an Ethereum Mainnet Node

(A) The node should be deployed using the following clients: 

- Execution Client = Geth
- Consensus Client = Prysm

Note: the node does NOT need to be fully sync’d, only evidence of it properly syncing

(B) The node should be monitored using:

- Prometheus
- Grafana

Evidence of Beacon properly syncing:
![image](https://github.com/mdr0id/valcloud/assets/36639405/bfd467dd-5777-48f2-8360-c304e9777d16)

Geth finding Peers on syncing:
![image](https://github.com/mdr0id/valcloud/assets/36639405/7d0ccf7f-6698-442f-b820-e6b22181ffb4)

### Part 2 - Monitor the Connected Peers

Write a Go program that consistently outputs a list of the execution client’s connected peers.

Evidence of geth_peer_exporter:
![image](https://github.com/mdr0id/valcloud/assets/36639405/b52db156-c7ab-4e4a-8885-0026e82bf0bf)

## Setup
🔥 ⚠️ **This is intended for demo use only and not production environments. For example, it uses self signed certs for GRPC.** ⚠️ 🔥

#### Part 1:

1. Ensure your system has Docker installed: https://docs.docker.com/engine/install/
2. Ensure your system has Docker Compose installed: https://docs.docker.com/compose/install/
3. Install loki plugin: `docker plugin install  grafana/loki-docker-driver:latest --alias loki --grant-all-permissions`

To start (from within valcloud repo):

(configure volumes for valcloud platform; within valcloud repo)
1. `sudo ./setup.sh`

*⚠️NOTE: If you create different volumes/users you need to modify the prometheus and grafana services accordingly in docker-compose.yaml⚠️*

Once volumes are configured for your platform:

1. `docker compose up -d`
2. Sanity check the services/network are running gracefully per your envinronment: `docker compose ps`
![image](https://github.com/mdr0id/valcloud/assets/36639405/45e8618c-50b7-4527-bb1d-d08d245def7b)

4. Navigate to http://localhost:3000 for Graphana
5. Navigate to http://localhost:9090 for Prometheus
6. If you are not seeing proper connections, ensure the expected ports are setup correctly per your platform:
`sudo ./setup-network.sh`

#### Part 2:

The geth_peer_exporter container is withing the docker-compose.yaml, so it will run when you issued the above `docker compose up -d`

##### Environment Version (e.g. works on my machine 🤓)

Docker version:
```
Client: Docker Engine - Community
 Version:           24.0.5
 API version:       1.43
 Go version:        go1.20.6
 Git commit:        ced0996
 Built:             Fri Jul 21 20:35:23 2023
 OS/Arch:           linux/amd64
 Context:           default

Server: Docker Engine - Community
 Engine:
  Version:          24.0.5
  API version:      1.43 (minimum version 1.12)
  Go version:       go1.20.6
  Git commit:       a61e2b4
  Built:            Fri Jul 21 20:35:23 2023
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.6.22
  GitCommit:        8165feabfdfe38c65b599c4993d227328c231fca
 runc:
  Version:          1.1.8
  GitCommit:        v1.1.8-0-g82f18fe
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
```
Docker Compose version:
```
Docker Compose version v2.3.3
```
Go version:
```
go version go1.20.7 linux/amd64
```

NOTES:
Design Process: 
1. Look at the Prysm stack to see if there are existing docker nodes and tooling
2. Determine what docker images are needed and what is exposed for each image
3. Determine what configs and volumes each docker image needs
4. Draft the templated and spin up each container
5. Review logs to fix/resolve bugs in logging

Technology choices:
1. Docker is widely used
2. Docker compose is often used for these deployments and scales well in cloud/prod environments
3. Loki is used to help correlate logs and other monitoring
4. json rpcs go package to simplify down some of the over the wire translations from http calls into Geth

Security Considerations:
1. Grafana is open; needs proper password etc but simplifies debugging
2. TLS cert is self signed for demo but for production would need proper setup
3. Using JWTS, but these have some attack vectors if not implemented correctly
4. Created own network for containers
5. HTTPS should be used on endpoints that need it
6. Would need to verify what is containers are leaving uneeded ports open, but creating docker network can help on certain deploys locally
7. Some containers would need DoS protection if the ports allowed flooding or maybe a back off algorithm

Future Improvements for Production
1. Ideally each service or binary deployed should have a config to aid deployment production
2. Some of the ideas mentioned in Security Considerations would need to be upheld
3. Proper hardware and resources would need to be allocated to ensure the services were reliable (e.g. so nodes don't get slashed etc)
4. If the services deployed required keys for validator wallets, then additional steps would be needed to ensure these are done proper so funds are not lost on mainnet
5. Proper alerts would be needed to ensure alert manager informed node operators




