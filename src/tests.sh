#!/bin/bash

echo "Registering User..."
curl -X POST "http://localhost:8080/users/registerUser" \
     -H "Content-Type: application/json" \
     -d '{
           "email": "user@example.com",
           "password": "securepassword",
           "portfoliokey": null
         }'

echo "Logging in..."
curl -X POST http://localhost:8080/users/loginUser \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword"
  }' \
  -c cookies.txt

echo "Creating Stock..."
curl -X POST http://localhost:8080/stock/createStock \
  -b cookies.txt \
  -H "Content-Type: application/json" \
  -d '{
    "name": "BTC",
    "currPrice": 1000,
    "allAmount": "10:1000, 20:3000"
  }'

echo "Creating Portfolio..."
curl -X POST http://localhost:8080/portfolio/createPortfolio \
  -b cookies.txt \
  -H "Content-Type: application/json" \
  -d '{
    "money": 1000000,
    "stocksinhand": ""
  }'

echo "Buying Stock..."
curl -X POST http://localhost:8080/stock/buyStock \
  -b cookies.txt \
  -H "Content-Type: application/json" \
  -d '{
    "name": "BTC",
    "amount": 30,
    "maxPrice": 3000
  }'
# echo "Logging out..."
# curl -X POST http://localhost:8080/users/logoutUser \
#   -b cookies.txt \
#   -H "Content-Type: application/json"

# echo "Deleting Portfolio..."
# curl -X POST http://localhost:8080/portfolio/deletePortfolio \
#   -b cookies.txt \
#   -H "Content-Type: application/json" \
#   -d '{}'

# echo "Removing User..."
# curl -X POST http://localhost:8080/users/removeUser \
#   -b cookies.txt \
#   -H "Content-Type: application/json" \
#   -d '{}'

echo "Done!"

