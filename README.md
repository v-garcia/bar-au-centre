## Config locale
- Install snap + microk8s
- Activate the following microk8s modules: storage, ingress, dashboard (optional)
- Install skaffold for local deployment and refresh
- export DOCKER_HOST="unix:///var/snap/microk8s/current/docker.sock"
- export config: `microk8s.config > $HOME/.kube/config`