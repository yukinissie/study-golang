deploy_prd:
	zip -r arc.zip lambda_function
	aws lambda update-function-code --profile helloGoLangFunction --region ap-northeast-1 --function-name helloGoLangFunction --zip-file fileb://./arc.zip --publish
