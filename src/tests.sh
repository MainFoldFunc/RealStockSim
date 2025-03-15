curl -X POST "http://localhost:8080/users/registerUser" \
     -H "Content-Type: application/json" \
     -d '{
           "email": "user@example.com",
           "password": "securepassword",
           "portfoliokey": null
         }'
curl -X POST http://localhost:8080/users/loginUser \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword"
  }' \
  -c cookies.txt
curl -X POST http://localhost:8080/users/logoutUser \
  -b cookies.txt
curl -X POST http://localhost:8080/portfolio/createPortfolio \
  -b cookies.txt \
  -H "Content-Type: application/json" \
  -d '{
    "User": "some_user",
    "money": 1000,
    "stocksinhand": "AAPL:10,GOOGL:5"
  }'
  curl -X POST http://localhost:8080/portfolio/deletePortfolio \
    -b cookies.txt
