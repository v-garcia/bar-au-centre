
#!/bin/bash

# Specific to my setup using microk8s

sudo iptables -P FORWARD ACCEPT
sudo chmod 777 /var/snap/microk8s/current/docker.sock
DOCKER_HOST="unix:///var/snap/microk8s/current/docker.sock" skaffold dev -p=local --port-forward=false