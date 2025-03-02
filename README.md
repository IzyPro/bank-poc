# Bank POC

Bank Account/Wallet proof-of-concept with the following features:

• Ability to deposit/withdraw money

• View balance

• View transaction history


## Run

Navigate to the `src/server` folder and run the following command:

```bash
go run .
```


### Deposit

##### POST /transaction/deposit

Request:
```json
{
    "reference": "tinybank1",
    "amount": 65.78,
    "surcharge": 6.578,
    "narration": "Test"
}
```

Response:
```json
{
    "Successful": true,
    "Message": "Deposit successful",
    "Data": {
        "Id": "ad5ba12a-a452-4d87-9a17-cf5a85591d92",
        "Reference": "tinybank1",
        "CreatedAt": "2025-02-27T22:51:30.701537+01:00",
        "Amount": 65.78,
        "Surcharge": 6.578,
        "Narration": "Test",
        "TransactionType": 1
    },
    "Code": "00"
}
```


### Withdrawal

##### POST /transaction/withdraw

Request:
```json
{
    "reference": "tinybank1",
    "amount": 15.78,
    "surcharge": 1.578,
    "narration": "Test"
}
```

Response:
```json
{
    "Successful": true,
    "Message": "Withdrawal successful",
    "Data": {
        "Id": "e06993c1-a848-4d4c-b408-37aeb8bc6d32",
        "Reference": "tinybank1",
        "CreatedAt": "2025-02-27T22:51:37.035372+01:00",
        "Amount": 15.78,
        "Surcharge": 1.578,
        "Narration": "Test",
        "TransactionType": 2
    },
    "Code": "00"
}
```


### Retrieve Balance

##### GET /transaction/balance

Response:
```json
{
    "Successful": true,
    "Message": "Account balance retrieved successfully",
    "Data": 344.98199999999997,
    "Code": "00"
}
```


### Retrieve Transactions

##### GET /transaction

Response:
```json
{
    "Successful": true,
    "Message": "Transaction history retrieved successfully",
    "Data": [
        {
            "Id": "d3eb1e04-0409-417e-a25b-5b9ae6761d8d",
            "Reference": "tinybank1",
            "CreatedAt": "2025-02-27T22:49:56.065467+01:00",
            "Amount": 65.78,
            "Surcharge": 6.578,
            "Narration": "Test",
            "TransactionType": 1
        },
        {
            "Id": "36bf182b-2da4-41a0-a781-3e2455176b2e",
            "Reference": "tinybank1",
            "CreatedAt": "2025-02-27T22:50:50.968732+01:00",
            "Amount": 65.78,
            "Surcharge": 6.578,
            "Narration": "Test",
            "TransactionType": 1
        },
        {
            "Id": "74fc561c-70d8-4845-91f5-52b5905c2cd7",
            "Reference": "tinybank1",
            "CreatedAt": "2025-02-27T22:51:06.141218+01:00",
            "Amount": 15.78,
            "Surcharge": 1.578,
            "Narration": "Test",
            "TransactionType": 2
        }
    ],
    "Code": "00"
}
```
