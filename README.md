# go-challenge

## How to run
- clone this repository
- cd into the cmd directory `cd ./cmd`
- run `go mod init github.com/shuoros/go-challenge`
- run `go mod tidy`
- run `env GOOS=linux GOARCH=amd64 go build -o ../build/main` to build the project
- cd back into the build directory `cd ../build`
- zip the built binary `zip -j main.zip main`
- login to your aws console and head over to `Lambda`
- click on Create function button.
- fill your desired name, set the runtime to `go1.x`, and set the architecture to `x86_64`
- click on change default execution role
- select `Create a new role from AWS policy templates`
- give your role a name
- from policy templates select `Simple microservice permissions`
- click on create
- on Runtime settings change the default handler to `main`
- from the Code source click on Upload from the select .zip file and at the end select the zip file you just created
- click on save
- again from the console go to the `DynamoDB` tab and click on Create Table
- fill in the name `VeryImportantTable`
- in partition key fill in `id` and type `String`
- click on create table
- go back to the console and click on the `API Gateway` tab and click on create api
- choose the `REST API`
- select `REST` as protocol, `New API` as create new api, give it a name and select `Regional` as Endpoint Type
- from the Actions to create Resource, give it a name and fill `devises` as the resource path and click on create Resource
- click on devices path and from actions to create Method, select `POST`, then select `Lambda Function` as Integration type, make sure `Use Lambda Proxy Integration` and `Use Default Timeout` are checked, and choose your lambda function from the dropdown
- click on save
- again from the Actions to create Resource, give it a name and fill `{id}` as the resource path and click on create Resource
- click on {id} path and from actions to create Method, select `GET`, then select `Lambda Function` as Integration type, make sure `Use Lambda Proxy Integration` and `Use Default Timeout` are checked, and choose your lambda function from the dropdown
- click on save
- finally, for deploying the api, from Actions click on `Deploy API`, select New Stage, fill in the name `dev` and select `api` as the stage name, then click on deploy
- go to stages tab and get your api url

### addDevice
Request:
```curl
curl -X POST -H "Content-Type: application/json" -d '{
  "id": "id1",
  "deviceModel": "model1",
  "name": "Sensor",
  "note": "Testing a sensor.",
  "serial": "A020000102"
  }' https://<api-gateway-url>/api/devices
```

Expected Response:
```json
{
  "timestamp": "Sun, 15 May 2022 20:31:55 +0000",
  "ok": true,
  "status": 201,
  "message": "Created",
  "data": "{\"id\":\"id3\",\"name\":\"Sensor\",\"deviceModel\":\"/devicemodels/id1\",\"serial\":\"A020000102\",\"note\":\"Testing a sensor.\"}"
}
```

### getDevice
Request:
```curl
curl -X GET https://<api-gateway-url>/api/devices/id1
```

Expected Response:
```json
{
  "timestamp": "Sun, 15 May 2022 20:35:03 +0000",
  "ok": true,
  "status": 200,
  "message": "OK",
  "data": "{\"id\":\"id1\",\"name\":\"Sensor\",\"deviceModel\":\"/devicemodels/id1\",\"serial\":\"A020000102\",\"note\":\"Testing a sensor.\"}"
}
```

## How to test
- cd into the test directory `cd ./test`
- run `go test -v`

if you want to generate reports run:
```bash
go test -v -coverprofile=coverage.out
go tool cover -html=cover.out -o cover.html
wkhtmltopdf cover.html cover.pdf
```
