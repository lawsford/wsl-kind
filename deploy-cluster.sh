kind create cluster --name local --config=$(dirname "$0")/kind/cluster-config.yaml
kubectl cluster-info --context kind-local

println("Setting up Nginx Ingress Controller...")
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

println("Setting up Tekton...")
kubectl apply -f https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
kubectl apply -f https://storage.googleapis.com/tekton-releases/triggers/latest/interceptors.yaml
kubectl apply -f https://storage.googleapis.com/tekton-releases/dashboard/latest/tekton-dashboard-release.yaml
kubectl apply -f resources/tekton-dashboard-ing.yaml
