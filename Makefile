STACK_NAME=developing-sam-applications
S3_BUCKET=developing-sam-applications

build:
	GOOS=linux go build -o bin/replicate handlers/replicate/main.go
	zip -r main.zip bin/*
.PHONY: build

package:
	sam package --template-file template.yml --output-template-file packaged.yml --s3-bucket $(S3_BUCKET)
.PHONY: package

deploy:
	sam deploy --template-file ./packaged.yml --stack-name $(STACK_NAME) --capabilities CAPABILITY_IAM
.PHONY: deploy

logs:
	sam logs -t -n Replicate --stack-name $(STACK_NAME)
.PHONY: logs

destroy:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
.PHONY: destroy
