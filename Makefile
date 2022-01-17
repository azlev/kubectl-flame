java:
	docker build -f agent/docker/jvm/Dockerfile . -t azlev/kubectl-flame:v0.2.2-jvm
	docker push 
