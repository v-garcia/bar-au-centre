
#!/bin/bash

# Specific to my setup using microk8s

DOCKER_HOST="unix:///var/snap/microk8s/current/docker.sock" skaffold dev -p=local --tail=false     