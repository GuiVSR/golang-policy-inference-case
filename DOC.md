1. Rodar este comando para buildar a imagem para lambda aws
`GOOS=linux CGO_ENABLED=0 GOARCH=arm64 go build -tags lambda.norpc -o bootstrap cmd/lambda/main.go`

2. Zipar desta forma para subir na aws
`zip go_lambda.zip bootstrap`

3. Subir na aws o go_lambda.zip