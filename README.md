
## Polling App

This is the backend of a polling app written purely in Go.

## Endpoints

- **`/`**
  - **Method:** GET
  - **Response:** `{ "message": "welcome to the go auth app" }`

- **`/ping`**
  - **Method:** GET
  - **Response:** `{ "message": "server is up and running" }`

- **`/v1/auth/register`**
  - **Method:** POST
  - **Request Body:** `{ "username": "your_username","email": "your_email", "password": "your_password", "confirmpassword": "your_password" }`
  - **Response:** `{ "message": "user successfully created" }`

- **`/v1/auth/login`**
  - **Method:** POST
  - **Request Body:** `{ "username": "your_username", "password": "your_password" }`
  - **Response:** `{ "message": "Login successful", "token": "your_jwt_token" }`

- **`/v1/polls`**
  - **Method:** POST
  - **Request Body:** `{ "title": "poll_title", "description": "poll_description", "options": [{"text": "option1"}, {"text": "option2"}] }`
  - **Response:** `{ "poll_id": poll_id, "message": "Poll created successfully" }`

- **`/v1/polls`**
  - **Method:** GET
  - **Response:** `[ { "poll_id": poll_id, "title": "poll_title", "description": "poll_description", "created_at": "timestamp" } ]`

- **`/v1/polls/{poll_id}`**
  - **Method:** GET
  - **Response:** `{ "poll_id": poll_id, "title": "poll_title", "description": "poll_description", "options": [{"option_id": option_id, "text": "option_text", "number_of_votes": vote_count}], "created_at": "timestamp" }`

- **`/v1/polls/{poll_id}/vote`**
  - **Method:** POST
  - **Request Body:** `{ "option_id": option_id }`
  - **Response:** `{ "message": "Vote cast successfully" }`

## Installation

1. Clone the repository:
    \`\`\`sh
    git clone https://github.com/xtasysensei/go-poll.git
    cd go-poll
    \`\`\`

2. Setup a PostgreSQL database based on the credentials in the `.env` file.

3. Install `make` and run:
    \`\`\`sh
    make migrate up
    \`\`\`

4. For complete functionality, use `make` and `air`, then run:
    \`\`\`sh
    make run
    #or
    make run-air
    \`\`\`

## Dependencies

- [Chi](https://github.com/go-chi/chi) for routing
- [JWT](https://github.com/dgrijalva/jwt-go) for token generation
- [PostgreSQL](https://www.postgresql.org/) for database
- [Migrate (for DB migrations)](https://github.com/golang-migrate/migrate/tree/v4.17.0/cmd/migrate)
- [Air](https://github.com/air-verse/air): Optional (can be installed with \`go install github.com/air-verse/air@latest\`)

## Environment Variables

- `SERVER_PORT`: Port on which the server runs (default: 8000)
- `POSTGRES_SERVER`: PostgreSQL server address (default: localhost)
- `POSTGRES_PORT`: PostgreSQL server port (default: 5432)
- `POSTGRES_DB`: PostgreSQL database name (default: go_app)
- `POSTGRES_USER`: PostgreSQL database user (default: sensei)
- `POSTGRES_PASSWORD`: PostgreSQL database password (default: 12345)
- `JWT_EXP`: JWT token expiration
- `JWT_SECRET`: JWT secret key

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
