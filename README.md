# MNC Test Go

A simple API built with Go that allows customers to login, logout, and make payments.

## Prerequisites

- Go 1.18 or higher
- Git

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/fbpr/mnc-test-go.git
   cd mnc-test-go
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Create an environment file:

   ```bash
   cp .env.example .env
   ```

4. (Optional) Edit the `.env` file to configure the application, there is no important env var for now

## Running the Application

1. Start the application:

   ```bash
   go run cmd/api/main.go
   ```

2. The API will be accessible at `http://localhost:8080/api/v1`

## API Documentation

### Authentication Endpoints

#### Login

- **URL**: `/api/v1/auth/login`
- **Method**: `POST`
- **Request Body**:

  ```json
  {
    "email": "example@gmail.com",
    "password": "yourpassword"
  }
  ```

- **Response Example**:

  ```json
  {
    "status": "success",
    "message": "login successful",
    "data": {
      "id": "1",
      "email": "example@gmail.com",
      "name": "name"
    }
  }
  ```

#### Logout

- **URL**: `/api/v1/auth/logout`
- **Method**: `POST`
- **Request Body**:

  ```json
  {
    "email": "example@gmail.com"
  }
  ```

- **Response Example**:

  ```json
  {
    "status": "success",
    "message": "logout successful"
  }
  ```

### Transaction Endpoints

#### Process Payment

- **URL**: `/api/v1/transactions/:id/pay`
- **Method**: `POST`
- **URL Parameters**: `id` (Transaction ID)
- **Request Body**:

  ```json
  {
    "customer_id": "1",
    "customer_email": "example@gmail.com"
  }
  ```

- **Response Example**:

  ```json
  {
    "status": "success",
    "message": "payment successful",
    "data": {
      "transaction_id": "1",
      "customer_id": "1",
      "merchant_id": "1",
      "amount": 250.75,
      "status": "completed"
    }
  }
  ```

## Testing

You can test the API using API testing tool like Postman. Check and use mock data in the `data` folder for easy testing.
