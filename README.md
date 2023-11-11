# Simple Eth Geth Beacon
Single Eth Geth Beacon Deployment (modified by me for demo purposes outside of original spec)

### Ethereum Mainnet Node

(A) The node should is deployed using the following clients: 

- Execution Client = Geth
- Consensus Client = Prysm

(B) The node is monitored using:

- Prometheus
- Grafana
- Loki

Beacon Syncing:
![beaconsync](https://github.com/mdr0id/eth_geth_beacon/assets/36639405/a6316240-75b4-45f0-adc1-3aa86f2271f7)


Geth finding Peers on syncing:
![geth](https://github.com/mdr0id/eth_geth_beacon/assets/36639405/b1f385bf-a3e4-495f-bc8e-f5cc42e416b5)


### Go Exporter

Sample Go program that consistently outputs a list of the execution client’s connected peers.

geth_peer_exporter find valid peers:
![lokipeer](https://github.com/mdr0id/eth_geth_beacon/assets/36639405/956cdabc-bcea-477f-9bcc-c87cee18d70e)


## Setup
🔥 ⚠️ **This is intended for demo use only and not production environments. For example, it uses self signed certs for GRPC. If used in production please use a valid certificate generation with domain host etc.** ⚠️ 🔥

#### Initial Setup:

1. Ensure your system has Docker installed: https://docs.docker.com/engine/install/
2. Ensure your system has Docker Compose installed: https://docs.docker.com/compose/install/
3. Install loki plugin: `docker plugin install  grafana/loki-docker-driver:latest --alias loki --grant-all-permissions`

To start (from within eth_geth_beacon repo):

(configure volumes for Eth Geth platform; within eth_geth_beacon repo)
1. `sudo ./setup.sh`

*⚠️NOTE: If you create different volumes/users you need to modify the prometheus and grafana services accordingly in docker-compose.yaml⚠️*

## If you have all the requires software and want to use the defualt env/volumes
1. `docker compose up -d`
2. Sanity check the services/network are running gracefully per your envinronment: `docker compose ps`
![image](https://github.com/mdr0id/valcloud/assets/36639405/45e8618c-50b7-4527-bb1d-d08d245def7b)

4. Navigate to http://localhost:3000 for Graphana
5. Navigate to http://localhost:9090 for Prometheus
6. If you are not seeing proper connections, ensure the expected ports are setup correctly per your platform:
`sudo ./setup-network.sh`

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
6. Since these are not validation nodes and do not have addrs this is intended to not scale up, but is possible

Security Considerations:
1. Grafana is open; needs proper password etc but simplifies debugging
2. TLS cert is self signed for demo but for production would need proper setup
3. Using JWTS, but these have some attack vectors if not implemented correctly
4. Created own network for containers
5. HTTPS should be used on endpoints that need it
6. Would need to verify what is containers are leaving uneeded ports open, but creating docker network can help on certain deploys locally
7. Some containers would need DoS protection if the ports allowed flooding or maybe a back off algorithm

Future Improvements for Production
1. Proper TLS cert process as noted above
2. Baseline requirements for each beacon node machine to ensure money properly used in cloud
3. Again since these don't have addrs and are not validating, there is NO economic insentive to spin these up at scale to fill disk
4. If validators or addrs were implemented, proper terraform, addr generation, and machine configuration tooling would be needed.




