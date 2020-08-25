dep1:
	kubectl apply -f deployment1.yaml

exposeDep1:
	kubectl expose deployment web --type=NodePort --port=8080

ingress:
	kubectl apply -f ingress.yaml