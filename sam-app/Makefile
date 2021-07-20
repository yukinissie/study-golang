.PHONY: build

build:
	sam build

packege:
	sam package --template-file sam-app/template.yaml --output-template-file sam-app/output-template.yaml --s3-bucket template-store --profile koike

deploy:
	sam deploy --template-file sam-app/output-template.yaml --stack-name go-lambda --capabilities CAPABILITY_IAM --profile koike