.PHONY: build

build:
	sam build

start:
	make build
	sam local start-api

deploy-first:
	make build
	sam deploy --s3-bucket ca-techboost-05 --profile techboost --guided

deploy:
	make build
	sam deploy --s3-bucket ca-techboost-05 --profile techboost --stack-name techboost-05-evacuation
