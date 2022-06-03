# point-of-sale

How To Run This Project

 - GraphQL Server
 
      go run .
 
 
 - GRPC Server 
      
      
      go run main.go


Test The Project
- Open GraphQL Playground in the Browser (type http://localhost:8080/), GraphQL server uses port 8080
- just excecute this query (full fields version)

mutation {
  signIn(input:{username:"ok", password:"1234567"}) {
    respon {
      status
      code
      message
    }
    user {
          id
          name
          username
          phonenumber
          email
          password
    }
  }
}


- Output's Gonna Be

{
  "data": {
    "signIn": {
      "respon": {
        "status": true,
        "code": 200,
        "message": "Success"
      },
      "user": {
        "id": 123,
        "name": "test",
        "username": "okccccc",
        "phonenumber": "081234567890",
        "email": "test@test.com",
        "password": "1234567"
      }
    }
  }
}



