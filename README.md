# dev-with-compose

## 1. Unit tests 

- Add a service definition `unit-tests` in `compose.yml`
- Run the unit tests with Compose
- Generate test reports in the `./reports` directory
- Use this image: `golang:1.25.2-alpine`
Commands to run the tests and generate reports:
```bash
go mod download
if go test -v > ./reports/test-output.txt; then
    echo "✅ Tests completed"
    echo "📄 Test report generated at ./reports/test-output.txt"
    go test -cover > ./reports/test-coverage.txt
    exit 0
else
    echo "❌ Tests failed"
    exit 1
fi
```

## 2. Build and Start the application with Compose

- Read the Dockerfile
- Add a service definition `web-service` in `compose.yml`
- Build and Start the application using Docker Compose
- Ensure `unit-tests` is successful before running `web-service`

## 3. Tests the services with a curl

Test the home page of the web application with a curl command

```bash
curl --fail http://domain:port
````

- Add a service definition `test-home-endpoint` in `compose.yml` to run the `curl` command
- Use this image: `curlimages/curl:latest`
- Ensure `web-service` is healthy before running `test-home-endpoint`

✋ you have to add the healthcheck to the `web-service` definition in `compose.yml`

```yaml
healthcheck:
    test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:8080/health"]
    interval: 30s
    timeout: 30s
    retries: 5
    start_period: 40s
```

## 4. Modularize your Compose file

Use the `extends` feature of Docker Compose to split your `compose.yml` file into multiple files: https://docs.docker.com/compose/how-tos/multiple-compose-files/extends/
