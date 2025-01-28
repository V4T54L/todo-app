# Todo App

A modern Todo application built with Golang for the backend and React.js for the frontend. This app provides a clean and efficient way to manage tasks while following best practices in coding, testing, and deployment.

<!-- ![Todo App Screenshot](screenshot.png)  -->

## Features

- **User Authentication**: Secure login and registration with JWT.
- **Task Management**: Create, read, update, and delete (CRUD) tasks easily.
- **Responsive Design**: Works on various devices and screen sizes.
- **Comprehensive Testing**: Unit and integration tests for both backend and frontend.

## Tech Stack

- **Backend**: 
  - Go (Golang)
  - Go-Chi (Routing)
  - PostgreSQL (Database)

- **Frontend**: 
  - React.js
  - Axios (HTTP client)
  - Jest (Testing)

## Getting Started

To run this project locally, follow these steps:

### Prerequisites

Make sure you have the following installed:

- Go (1.20 or later)
- Node.js (18 or later)
- PostgreSQL (and a running database)

### Setup

1. **Clone the repository**:

   ```bash
   git clone https://github.com/V4T54L/todo-app.git
   cd todo-app
   ```

2. **Backend Setup**:

   Navigate to the backend folder:

   ```bash
   cd backend
   ```

   - Install dependencies:

   ```bash
   go mod tidy
   ```

   - Set up your PostgreSQL database and update your `.env` as needed.

   - Run the Go application:

   ```bash
   go run main.go
   ```

3. **Frontend Setup**:

   Navigate back to the root directory and then to the frontend folder:

   ```bash
   cd frontend
   ```

   - Install dependencies:

   ```bash
   npm install
   ```

   - Run the React application:

   ```bash
   npm start
   ```

### Testing

To run tests for the backend, navigate to the `backend` directory and use:

```bash
go test ./...
```

For the frontend tests, navigate to the `frontend` directory and use:

```bash
npm test
```

### CI/CD

This project includes automated deployment pipelines using GitHub Actions. Every push to the repository will trigger the deployment process.

## Contributing

Feel free to contribute by opening issues or submitting pull requests. Please ensure that your code follows the standard Go and React project structures and naming conventions.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.

## Acknowledgments

- Thanks to the open-source community for their invaluable resources and libraries.

---

Happy coding!
