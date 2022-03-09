include configs/configs.env

# this hack, export all the variables declared by the Makefile as Environment Variables
export

gen:
	@go build -o out/main cmd/main.go

lambda/build:
	@rm -rf build/*
	@GOOS=linux GOARCH=amd64 go build -o out/main cmd/main.go
	@zip -jrm build/main.zip build/main
	echo "lambda zip file created"

lambda/local:
	@ rm -rf build/*
	@GOOS=linux GOARCH=amd64 go build -o out/main cmd/main.go
	sam local invoke -e events/event-sqs.json

docker/build:
	@ docker image build \
		--build-arg PERSONAL_ACCESS_TOKEN=${PERSONAL_ACCESS_TOKEN} \
		--build-arg BUF_USER=${BUF_USER} \
		--build-arg BUF_TOKEN=${BUF_TOKEN} \
		-t $(PROJECT_NAME) -f ./build/Dockerfile  .

ecr/push:
	@ docker tag zendesk-integration:latest 713307779200.dkr.ecr.us-east-2.amazonaws.com/zendesk-integration:latest
	@ docker push 713307779200.dkr.ecr.us-east-2.amazonaws.com/zendesk-integration:latest