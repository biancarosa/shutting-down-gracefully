# shutting-down-gracefully

make build
make push
kubectl apply -f k8s.yml
watch -n 1 kubectl get pods 