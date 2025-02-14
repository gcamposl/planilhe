# planilhe
A new way to control your finances

## Getting Started

### Prerequisites
- Docker
- Docker Compose

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/gcamposl/planilhe.git
    cd planilhe
    ```

2. Copy the example environment variables file and edit as needed:
    ```sh
    cp .env.example .env
    ```

3. Build and run the Docker containers:
    ```sh
    docker-compose up --build
    ```

### Usage

- Access the application at `http://localhost:3000`
- The API documentation is available at `http://localhost:3000/api-docs`

### Running Tests

To run tests, use the following command:
```sh
docker-compose exec app npm test
```

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Contact

Gabriel Campos - [@gcamposlps](https://x.com/gcamposlps)

Project Link: [https://github.com/gcamposl/planilhe](https://github.com/gcamposl/planilhe)