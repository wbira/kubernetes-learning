create-pod:
	kubectl apply -f helloworld-pod.yaml

delete-pod:
	kubectl delete -f helloworld-pod.yaml

port-forward:
	kubectl port-forward helloworld 8080:8080

run: create-pod port-forward