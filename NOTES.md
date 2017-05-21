HOW TO USE CURL:


curl -i -X  POST http://localhost:4020/api/v1/user/create \
  -H "Accept: application/json" -H "Content-type: application/json" \
  -d '{ "user": "coral1123", "password": "PWTEST", "email": "swappingio@gmail.com" }'

 curl -i -X  POST http://localhost:4020/api/v1/user/create   -H "Accept: application/json" -H "Content-type: application/json"   -d '{ "username": "coral", "password": "asdfasdfd", "email": "swappingio@gmail.com" }'

 curl -i -X  POST http://localhost:4020/api/v1/user/create   -H "Accept: application/json" -H "Content-type: application/json"   -d '{ "username": "coral", "activationcode": "ozkuaNyhzYggKnWQntxIHgZtSpkPXIEzcAikHXTwHxoqupnGVJUVdQyuvAlHLqTErnoeFokLkEChbVqyQzFGEwZwZQDNyahbiDTi"}'
