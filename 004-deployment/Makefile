create: 
	kubectl create -f deployment.yaml

show-rs:
	kubectl get replicasets --selector=run=kuard

scale: 
	kubectl scale deployments kuard --replicas=2

export-deployment:
	kubectl get deployments kuard --export -o yaml > kuard-deployment.yaml

apply:
	kubectl apply -f deployment.yaml

history:
	kubectl rollout history deployments kuard

undo:
	kubectl rollout undo deployments kuard

delete:
	kubectl delete deployments kuard