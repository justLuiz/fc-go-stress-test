# stress-test

A CLI load testing tool that sends concurrent HTTP requests to a target URL and reports results.

## Build

```bash
docker build -t stress-test .
```

## Run

```bash
docker run stress-test --url=http://google.com --requests=1000 --concurrency=10
```

### Flags

| Flag | Description | Required |
|------|-------------|----------|
| `--url` | Target URL to test | Yes |
| `--requests` | Total number of requests to send | Yes |
| `--concurrency` | Number of parallel workers | Yes |

## Output

```
Total time: 12s
Total requests: 1000
Requests with status 200: 987
Status 301: 13 requests
```
