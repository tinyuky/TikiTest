# TikiTest
This is my app.Sorry about the UI, I haven't finish UI. I have just finish the api
- Run:
  - Install Go Lang in computer
  - Clone project
  - Go to project source
  - In terminal: go run main.go
- API List:
  - localhost:3000/login :is for login action
  - localhost:3000/add   : is for add new user to system
  - localhost:3000/changepassword: is for update user's password
  * The port 3000 you can configure in file main.go
  * You should use JSON in body request when send POST request to my APIs
    *Example: {
                "username" : "example"
                "password" : "example"
              }
  * Header for request should have Content-Type: application/json
- File Struct:
  - main.go : application running file
  - routers: folder have config about router for api. I use Gin in this project
  - dto: data tranfer object. I store data type I used in this project
  - controllers: have functions to handle the request. They use function in model or service
  - models: have functions to work with file password.txt like: Add, Get,..
  - services: have functions work with password: validate, verify, set password
