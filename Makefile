# ==================================================================================== #
# DEPLOYMENT
# ==================================================================================== #

## docker/build: build the docker image
.PHONY: docker/build
docker/build:
	docker build --tag federated-dashboard .

## docker/test: run the built docker image
.PHONY: docker/test
docker/test:
	docker run -p 8080:8080 -e TIER="best" -e DOMAIN="test.host" federated-dashboard

## docker/tag: tag the container and prepare to push
.PHONY: docker/tag
docker/tag:
	docker tag federated-dashboard:latest federatedcomputer/dashboard:latest

## docker/push: push container to feds
.PHONY: docker/push
docker/push:
	docker push federatedcomputer/dashboard:latest