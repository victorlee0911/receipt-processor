# receipt-processor

Build Docker Image:
$ docker build -t process-receipt .

Run Docker Container:
$ docker run -dp 8080:8080 process-receipt

Then freely access api:

Example:

$ curl -X POST http://localhost:8080/receipts/process \
    -H "Content-Type: application/json" \
    -d '{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}'

Replace :id with object id received from response.

$ curl http://localhost:8080/receipts/:id/points
