
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

### Dependencies

- [Chi](https://github.com/go-chi/chi) for routing
- [JWT](https://github.com/dgrijalva/jwt-go) for token generation
- [PostgreSQL](https://www.postgresql.org/) for database
- [Migrate (for DB migrations)](https://github.com/golang-migrate/migrate/tree/v4.17.0/cmd/migrate)
- [Air](https://github.com/air-verse/air): Optional (For live reload during development. It can be installed with `go install github.com/air-verse/air@latest`)

### Environment Variables

- `SERVER_PORT`: Port on which the server runs (default: 8000)
- `POSTGRES_SERVER`: PostgreSQL server address (default: localhost)
- `POSTGRES_PORT`: PostgreSQL server port (default: 5432)
- `POSTGRES_DB`: PostgreSQL database name (default: go_app)
- `POSTGRES_USER`: PostgreSQL database user (default: sensei)
- `POSTGRES_PASSWORD`: PostgreSQL database password (default: 12345)
- `JWT_EXP`: JWT token expiration
- `JWT_SECRET`: JWT secret key

### Manual installation
1. Clone the repository:
    ```sh
    git clone https://github.com/xtasysensei/go-poll.git
    cd go-poll
    ```

2. Setup a PostgreSQL database based on the credentials in the `.env` file.

3. Install `make` and run:
    ```sh
    make migrate up
    ```

4. For complete functionality, use `make` and `air`, for live reload during development, then run:
    ```sh
    make run
    #or
    make run-air
   ```
### Docker

A docker image is available at dockerhub via [go-poll]().

#### Getting Started

Prerequisites: Docker and Docker Compose.

- If you are using Docker Desktop, both of these should be already installed.
- If you prefer Docker Engine on Linux, make sure to follow their [installation guide](https://docs.docker.com/engine/install/#server).

**We provide support for the latest Docker release as shown above.**
If you are using Linux and the Docker package that came with your package manager, it will probably work too, but support will only be best-effort.

Upgrading Docker from the package manager version to upstream requires that you uninstall the old versions as seen in their manuals for [Ubuntu](https://docs.docker.com/engine/install/ubuntu/#uninstall-old-versions), [Fedora](https://docs.docker.com/engine/install/fedora/#uninstall-old-versions) and others.

Then, to get started:

1. Run `docker version` and `docker compose version` to see if you have Docker and Docker Compose properly installed. You should be able to see their versions in the output.

    For example:

    ```text
    >>> docker version
    Client:
     [...]
     Version:           23.0.5
     [...]

    Server: Docker Desktop 4.19.0 (106363)
     Engine:
      [...]
      Version:          23.0.5
      [...]

    >>> docker compose version
    Docker Compose version v2.17.3
    ```

    If you don't see anything or get a command not found error, follow the prerequisites to setup Docker and Docker Compose.

2. Clone or download this repository and extract,
3. Open a terminal of your choice and change its working directory into this folder (`go-polll`).
4. Run `docker compose up -d`.

go-polll is now available at `http://localhost:9000`.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
