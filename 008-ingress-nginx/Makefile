dep1:
	kubectl apply -f deployment1.yaml

exposeDep1:
	kubectl expose deployment web --type=NodePort --port=8080

dep2:
	kubectl apply -f deployment2.yaml

exposeDep2:
	kubectl expose deployment web1 --type=NodePort --port=8080

ingress:
	kubectl apply -f ingress.yaml