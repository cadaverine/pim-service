generate:
	buf generate

run:
	go run cmd/pim-service/main.go

kube:
	./scripts/kube.sh

pim-api-build:
	eval $(minikube docker-env)
	docker build . -t cadaverine/pim-api:latest

pim-api-clear:
	docker rmi cadaverine/pim-api:latest

pim-api: pim-api-clear
	pim-api-build