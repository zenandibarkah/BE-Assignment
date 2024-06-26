# BE-Assignment
Task to create 2 microservice with Docker for containerization. 

## Before Begin:
Make sure you have **Docker** in your local.

### Clone repository to your local:
```bash
$ git clone https://github.com/zenandibarkah/BE-Assignment.git
```
### Run script `script.sh` with command `linux` or `wsl`:
```bash
$ bash script.sh
```
After running the script, 3 Docker Container has been created: 
- `account-manager` container
- `payment-manager` container
- `postgres_dev` container (Table and Function trigger has been created)

You can import postman collection to testing endpoint url. 


## API Reference

### Account-Manager Service:
- **Register User** 

```http
  POST /public/register
```
&emsp;Body:
```json
  {
    "username": "zenandib",
    "email": "zeinandibarkah19@gmail.com",
    "password": "secretpassword"
  }
```
&emsp;Response:
```json
  {
    "statusCode": 200,
    "status": "Success",
    "message": "Register Successfull",
    "displayMessage": "warning.sucess",
    "response": "zenandib"
  }
```
- **Login User**

```http
  POST /public/login
```
&emsp;Body:
```json
  {
    "email":"zeinandibarkah19@gmail.com",
    "password":"secretpassword"
  }
```
&emsp;Response:
```json
  {
    "statusCode": 200,
    "status": "Success",
    "message": "Login Successfull",
    "displayMessage": "warning.sucess",
    "response": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzOTM0MjcsImlhdCI6MTcxNTM4OTgyNywiaWRfdXNlciI6IjEiLCJuYmYiOjE3MTUzODk4Mjd9.gLKLzMN_K6pIZRl59oSnbrJMyRjMrrEbv8B55ICXkWc"
    }
  }
```

- **Update User**

```http
  POST /private/user/update
```
&emsp;Body:
```json
  {
    "username":"zenandib",
    "phone":"08854523467"
  }
```
&emsp;Response:
```json
  {
    "statusCode": 200,
    "status": "Success",
    "message": "Update User Successfull",
    "displayMessage": "warning.sucess",
    "response": null
  }
```
- **Current User**

```http
  GET /private/myuser
```
&emsp;Response:
```json
  {
    "statusCode": 200,
    "status": "Success",
    "message": "Get Current User Successfull",
    "displayMessage": "warning.sucess",
    "response": {
        "id": "1",
        "username": "zenandib",
        "email": "zeinandibarkah19@gmail.com",
        "phone": "0885"
    }
  }
```
- **Add Account Bank**

```http
  POST /private/user/bank/add
```
&emsp;Body:
```json
  {
    "account_name":"Zenandi Barkah",
    "bank_name":"MANDIRI",
    "account_number":"12308346127312"
    // "account_name":"Zenandi Barkah",
    // "bank_name":"BCA",
    // "account_number":"889761237"
  }
```
&emsp;Response:
```json
  {
    "statusCode": 200,
    "status": "Success",
    "message": "Add Account Bank Successfull",
    "displayMessage": "warning.sucess",
    "response": null
  }
```
- **List Account Bank**
&emsp;Get list of banks based on user ID.

```http
  GET /private/user/banks
```
&emsp;Response:
```json
  {
    "statusCode": 200,
    "status": "Success",
    "message": "Get List Account Bank Successfull",
    "displayMessage": "warning.sucess",
    "response": [
        {
            "id": "1",
            "bank_name": "BCA",
            "account_number": "889761237"
        },
        {
            "id": "2",
            "bank_name": "MANDIRI",
            "account_number": "12308346127312"
        }
    ]
  }
```
- **Detail Account Bank**
&emsp;Get Detail of banks based on parameter `account_number`.

```http
  GET /private/user/bank/:accNumber
```
&emsp;Response:
```json
  {
    "statusCode": 200,
    "status": "Success",
    "message": "Get Detail Account Bank Successfull",
    "displayMessage": "warning.sucess",
    "response": {
        "id": "2",
        "account_name": "Zenandi Barkah",
        "bank_name": "MANDIRI",
        "account_number": "12308346127312",
        "saldo": 1000000
    }
  }
```
### Payment-Manager Service:
- **Withdraw Transaction** 

```http
  POST /private/trans/withdraw
```
&emsp;Body:
```json
  {
    "destination_bank":"MANDIRI",
    "destination_account_number":"12308346127312",
    "amount": 3000000
  }
```
&emsp;Response:
```json
  {
    "statusCode": 200,
    "status": "Success",
    "message": "Withdraw Transaction Successfull",
    "displayMessage": "warning.sucess",
    "response": null
  }
```
&emsp;**Note:** Balance in account bank with `account_number == destination_account_number` updated.

- **Send Transaction**

```http
  POST /private/trans/send
```
&emsp;Body:
```json
  {
    "source_bank":"MANDIRI",
    "source_account_number": "12308346127312",
    "destination_bank":"BCA",
    "destination_account_number":"889761237",
    "amount": 2000000
  }
```
&emsp;Response:
```json
  {
    "statusCode": 200,
    "status": "Success",
    "message": "Send Transaction Successfull",
    "displayMessage": "warning.sucess",
    "response": null
  }
```
&emsp;**Note:** Balance in account bank with `account_number == destination_account_number` and `account_number == source_account_number` updated.

- **List History Trans**

```http
  GET /private/history/trans
```
&emsp;Response:
```json
  {
    "statusCode": 200,
    "status": "Success",
    "message": "Get List History Trans Successfull",
    "displayMessage": "warning.sucess",
    "response": [
        {
            "id": "1",
            "source_bank": "",
            "destination_bank": "MANDIRI",
            "trans_type": "DEBIT"
        }
    ]
 }
```

- **Get Detail History Trans**
&emsp;Get Detail of History Trans based on parameter `id transaction`.
```http
  GET /private/history/trans/:idTrans
```
&emsp;Response:
```json
  {
    "statusCode": 200,
    "status": "Success",
    "message": "Get Detail History Trans Successfull",
    "displayMessage": "warning.sucess",
    "response": {
        "id": "1",
        "source_bank": "",
        "source_account_number": "",
        "destination_bank": "MANDIRI",
        "destination_account_number": "12308346127312",
        "amount": 3000000,
        "trans_type": "DEBIT",
        "trans_date": "2024-05-11T01:11:19.816781Z"
    }
  }
```

## Tech Stack:
- `Golang` for API server (Gin framework)
- `PosgreSQL` for Database.
- `Docker` for containerization with `docker-compose`

