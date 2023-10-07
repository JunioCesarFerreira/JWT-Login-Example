# Vue.js JWT Login Test UI with Golang API

This repository contains a simple example of a user login interface built using Vue.js and JWT authentication, along with a corresponding API implemented in Golang. The UI allows users to test the login functionality using fixed login credentials and demonstrates the interaction between the frontend and backend.

## Prerequisites

Before getting started, make sure you have the following tools installed:

- Node.js: To run and manage the Vue.js frontend.
- Go: To build and run the Golang backend.
- Git: To clone this repository.

## Getting Started

Follow these steps to set up and run the application:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/JunioCesarFerreira/JWT-Login-Example
   cd JWT-Login-Example
   ```

2. **Frontend Setup**:
   
   Navigate to the `ui` directory:

   ```bash
   cd ../ui
   ```

   Install the required dependencies:

   ```bash
   npm install
   ```

   Run the Vue.js development server:

   ```bash
   npm run serve
   ```

   Access the frontend UI by opening your web browser and visiting `http://localhost:8080`.

3. **Backend Setup**:

   Navigate to the `api` directory:

   ```bash
   cd ../api
   ```

   Build and run the Golang API:

   ```bash
   go run main.go
   ```

   The API will be accessible at `http://localhost:8082`.

## Usage

1. Open your web browser and go to the Vue.js frontend URL: `http://localhost:8080`.

2. You will see a login form. Use the following fixed login credentials for testing:
   - Username: `Junio`
   - Password: `pass1234`

3. Upon successful login, you will receive a JWT token.

4. You can make API requests using the obtained JWT token to the Golang backend at `http://localhost:8082/api/protected`. The backend will validate the token and respond accordingly.

## Folder Structure

- `ui`: Contains the Vue.js frontend application.
- `api`: Contains the Golang API server.

## Technologies Used

- **Frontend**:
  - Vue.js: JavaScript framework for building user interfaces.
  - Vue Router: Routing library for Vue.js applications.

- **Backend**:
  - Golang: Programming language used to build the API.
  - Gorilla Mux: Router and dispatcher for building web applications.
  - JWT: JSON Web Token library for authentication.

## License

This project is licensed under the MIT License - see file [LICENSE.md](LICENSE.md) for details.

---

This project serves as a demonstration of integrating Vue.js with Golang for user authentication using JWT. It can be used as a starting point to build more complex authentication systems or as a learning resource for understanding the interaction between frontend and backend components.