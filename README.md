# File Processor API

This project implements an API for processing files, providing endpoints for user and staff functionalities.

## Getting Started

To get started with this project, follow these steps:

1. Clone the repository.
2. Install the necessary dependencies.
3. Run the project locally.
4. Explore the API endpoints.

## Installation

To install the necessary dependencies, use the following command:

```bash
npm install
```

## Running the Project

To run the project locally, execute the following command:

```bash
npm start
```

The API will be accessible at `http://localhost:8080`.

## API Endpoints

### User Endpoints

#### User Signup

- **Endpoint:** `/user/signup`
- **Method:** POST
- **Description:** Allows users to sign up.
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Responses:**
  - 200: User created successfully

#### User Login

- **Endpoint:** `/user/login`
- **Method:** POST
- **Description:** Allows users to log in.
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Responses:**
  - 200: Successful operation
  - Response body contains a JWT token for authentication.

#### Process a File

- **Endpoint:** `/fileProcess`
- **Method:** POST
- **Description:** Processes a file.
- **Request Body:** Form data with a file and routines.
- **Responses:**
  - 200: Successful operation
  - Response body contains processed file information.

#### Get User Processes

- **Endpoint:** `/user/user_processes`
- **Method:** GET
- **Description:** Retrieves processes associated with the user.
- **Responses:**
  - 200: Successful operation
  - Response body contains a list of user processes.

#### Get User Process by ID

- **Endpoint:** `/user/get_process/{processId}`
- **Method:** GET
- **Description:** Retrieves a user process by ID.
- **Parameters:**
  - `processId`: ID of the process to retrieve
- **Responses:**
  - 200: Successful operation
  - Response body contains the user process.

### Staff Endpoints

#### Staff Login

- **Endpoint:** `/staff/staffLogin`
- **Method:** POST
- **Description:** Allows staff members to log in.
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Responses:**
  - 200: Successful operation
  - Response body contains a JWT token for authentication.

#### Get Statistics for a File

- **Endpoint:** `/staff/statistics`
- **Method:** POST
- **Description:** Retrieves statistics for a file.
- **Request Body:**
  ```json
  {
    "filename": "string"
  }
  ```
- **Responses:**
  - 200: Successful operation
  - Response body contains file statistics.

#### Get All Staff Processes

- **Endpoint:** `/staff/get_all_processes`
- **Method:** GET
- **Description:** Retrieves all processes handled by staff members.
- **Responses:**
  - 200: Successful operation
  - Response body contains a list of staff processes.

## Security

This API uses JSON Web Tokens (JWT) for authentication. All endpoints requiring authentication expect the token to be provided in the `Authorization` header as a Bearer token.

## Technologies Used

- OpenAPI 3.0.3
- Node.js
- Express
- JWT for authentication

## Contributors

- Abdul Rafay Zia


## License

This project is licensed under the [License Name] License - see the [LICENSE.md](LICENSE.md) file for details.
