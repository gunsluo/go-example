



minikube config set driver docker
minikube config set cpus 4
minikube config set memory 8192

minikube start
#minikube start --cpus 4 --memory 8192 --driver=docker


curl -sL https://istio.io/downloadIstioctl | sh -
export PATH=$PATH:$HOME/.istioctl/bin
istioctl install --set profile=demo

minikube addons enable istio


kubectl label namespace postgres istio-injection=enabled

kubectl get svc istio-ingressgateway -n istio-system

istioctl -n istio-system proxy-config listener $(kubectl get po -l istio=ingressgateway -n istio-system -o jsonpath='{.items[0]..metadata.name}')

kubectl get namespace -L istio-injection

export INGRESS_HOST=$(minikube ip)
export INGRESS_HOST=$(kubectl get po -l istio=ingressgateway -n istio-system -o jsonpath='{.items[0].status.hostIP}')

export TCP_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="tcp")].nodePort}')
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')

