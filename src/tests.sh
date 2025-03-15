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
  
echo "Creating Portfolio..."
curl -X POST http://localhost:8080/portfolio/createPortfolio \
  -b cookies.txt \
  -H "Content-Type: application/json" \
  -d '{
    "money": 1000000,
    "stocksinhand": ""
  }'

echo "Creating Stock..."
curl -X POST http://localhost:8080/stock/createStock \
  -b cookies.txt \
  -H "Content-Type: application/json" \
  -d '{
    "name": "BTC",
    "currPrice": 1000,
    "allAmount": "10:1000:1,20:3000:1"
  }'
  
echo "Updating Stock..."
curl -X POST http://localhost:8080/stock/updateStock \
  -b cookies.txt \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "amountToAdd": 5,
    "price": 3000
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

