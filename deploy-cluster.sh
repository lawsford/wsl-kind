mkdir -p cluster-volume

kind create cluster --name local --config=$(dirname "$0")/kind/cluster-config.yaml
kubectl cluster-info --context kind-local

printf "Setting up Nginx Ingress Controller..."
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

printf "Setting up Tekton..."
kubectl apply -f https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
kubectl apply -f https://storage.googleapis.com/tekton-releases/triggers/latest/interceptors.yaml
kubectl apply -f https://storage.googleapis.com/tekton-releases/triggers/latest/webhooks.yaml

printf "Setting up Tekton Dashboard..."
kubectl apply -f https://storage.googleapis.com/tekton-releases/dashboard/latest/tekton-dashboard-release.yaml

until [ $(kubectl get pods -l app.kubernetes.io/component=controller -n ingress-nginx -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') == "True" ]; do
    echo "Waiting for Nginx Ingress Controller to be ready..."
    sleep 15
done

if kubectl get po -n ingress-nginx -o wide | grep controller | grep -q 'Running'; then
   kubectl apply -f resources/tekton-dashboard-ing.yaml
fi

