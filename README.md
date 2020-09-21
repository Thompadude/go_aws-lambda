# What's This?
An AWS Lambda project written in GoLang for learning purposes.
## Commands
### Create
```bash
> cd cmd
> zip function.zip main
> aws lambda create-function --function-name my-function \
--runtime go1.x \
--zip-file fileb://function.zip \
--handler main --role arn:aws:iam::<id>:role/<role>
```
### Update
```bash
> aws lambda update-function-code --function-name  my-function \
--zip-file fileb://function.zip
```
### Test
#### With CLI
```bash
aws lambda invoke \
--cli-binary-format raw-in-base64-out \
--function-name my-function \
--payload '{ "name": "Bob" }' response.json
```
#### With HTTPS (needs API Gateway)
```bash
curl --request POST 'https://<id>.execute-api.eu-north-1.amazonaws.com/prod/' -d '{"name":"bob"}'
```