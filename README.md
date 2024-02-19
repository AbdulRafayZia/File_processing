# File Processor API

This is an API designed for processing files, with endpoints for user and staff functionalities. It allows users to process files, login, and sign up, while staff members can access statistics and retrieve all processes.

## Technologies Used
- **Language:** Go (Golang)
- **OpenAPI Specification:** 3.0.3
- **Dependencies:** Standard Go libraries

## Endpoints

### User Endpoints

#### 1. Process a File
- **Method:** POST
- **Path:** `/fileProcess`
- **Description:** Process a file with specified routines.
- **Request Body:**
  - `file`: The file to be processed (multipart/form-data)
  - `routines`: Type of routines to apply to the file (string)
- **Response:**
  - `Total_lines`: Number of lines in the file
  - `Total_words`: Number of words in the file
  - `Total_punctuations`: Number of punctuations in the file
  - `Total_vowels`: Number of vowels in the file
  - `Execution_Time`: Time taken for execution
  - `No_of_Routines`: Number of routines applied
  - `file_name`: Name of the processed file
  - `username`: Username of the user who processed the file

#### 2. User Login
- **Method:** POST
- **Path:** `/user/login`
- **Description:** Authenticate user login.
- **Request Body:**
  - `username`: Username of the user (string)
  - `password`: Password of the user (string)
- **Response:**
  - `token`: JWT token for authentication

#### 3. User Signup
- **Method:** POST
- **Path:** `/user/signup`
- **Description:** Register a new user.
- **Request Body:**
  - `username`: Username of the user (string)
  - `password`: Password of the user (string)
- **Response:**
  - `"User created successfully"` (string)

#### 4. Get User Processes
- **Method:** GET
- **Path:** `/user/user_processes`
- **Description:** Get processes associated with the authenticated user.
- **Response:**
  - List of processed files with details including ID, username, file name, metrics, and execution time.

#### 5. Get User Process by ID
- **Method:** GET
- **Path:** `/user/get_process/{processId}`
- **Description:** Get a specific process by its ID.
- **Request Parameter:**
  - `processId`: ID of the process to retrieve (integer)
- **Response:**
  - Details of the specified process including ID, username, file name, metrics, and execution time.

### Staff Endpoints

#### 1. Staff Login
- **Method:** POST
- **Path:** `/staff/staffLogin`
- **Description:** Authenticate staff login.
- **Request Body:**
  - `username`: Username of the staff (string)
  - `password`: Password of the staff (string)
- **Response:**
  - `token`: JWT token for authentication

#### 2. Get Statistics for a File
- **Method:** POST
- **Path:** `/staff/statistics`
- **Description:** Get statistics for a specific file.
- **Request Body:**
  - `filename`: Name of the file to get statistics for (string)
- **Response:**
  - `average_execution_time`: Average execution time for the file processing

#### 3. Get All Staff Processes
- **Method:** GET
- **Path:** `/staff/get_all_processes`
- **Description:** Get all processes handled by staff members.
- **Response:**
  - List of processed files with details including ID, username, file name, metrics, and execution time.

## Security
- **Bearer Token Authentication:** JWT tokens are used for authentication. Users and staff members must include a valid token in the request headers to access protected endpoints.

## Setup
1. Clone the repository.
2. Install Go if not already installed.
3. Install dependencies if required.
4. Build and run the application.

## Usage
- Ensure to have valid JWT tokens for accessing user and staff endpoints.
- Use appropriate endpoints for file processing, user management, and staff functionalities.

## Contributing
Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License
This project is licensed under the [MIT License](LICENSE).

---

Feel free to customize this README according to your project's specific needs and features. If you have any questions or need further assistance, please let me know!
