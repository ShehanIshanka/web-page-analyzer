# Web Page Analyzer

This project is a web application that does an analysis of a web-page/URL.

- [Running project with Docker](#running-project-with-docker)
- [Architecture](#architecture)
    - [Backend](#backend)
    - [Frontend](#frontend)
- [Possible Improvements](#possible-improvements)
    - [Backend](#backend-1)
    - [Frontend](#frontend-1)
    - [Deployment](#deployment)

<img src="./video.webm" width="80%" height="80%"/>

## Running project with Docker

Ensure that you have Docker Engine installed locally. Refer to platform specific instructions.

1. `docker compose build`
2. `docker compose up -d`

To view the web page: <http://localhost:8888>.

## Architecture

The system includes two components; Backend and Frontend.

### Backend

Backend is built with `Go language`. Backend exposes `POST /analyze` endpoint to analyze web page.

- **Sample Request**

```json
{
  "url": "www.example.com"
}
```

- **Sample Response**

```json
{
  "data": {
    "html_version": "",
    "title": "Medium icon",
    "headings": {
      "h1": 1,
      "h2": 4,
      "h3": 3,
      "h4": 12
    },
    "internal_links": 22,
    "external_links": 117,
    "inaccessible_links": 22,
    "login_form": false
  },
  "success": true
}
```

| Status Code | Description            |
|-------------|------------------------|
| 200         | Successful Response    |
| 400         | URL cannot be empty    |
| 400         | failed to reach url    |
| 404         | page is not accessible |
| 422         | failed to parse page   |

- **Application logic**

1. Validate URL.
2. Fetch Page using `net/http` parser.
3. Parse HTML page and scrape following info.
    - HTML Version from DOCTYPE (Currently not working)
    - Title
    - Headings : "h1", "h2", "h3", "h4", "h5", "h6" tags
    - External Links : `a` tags with `http` or `https` prefix
    - Internal Links : `a` tags without `http` or `https` prefix
    - Inaccessible Links : Validate that all external and internal links are accessible and successfully retrieve their
      respective pages.
    - Has Login Form : Search for Button with `log in` or `sign in` or `sign up` text

### Frontend

Frontend is built with `Vue.js` and `PrimeVue`

## Possible Improvements

### Backend

1. Implement API Docs (Swagger).
2. Implement rate limiter to prevent excessive load.
3. Implement cache.
4. Add authentication with API key and Secret.
5. Add Correlation ID for API requests for further debugging.
6. Improve scraping logic
    - Use `goquery` library
    - More conditional checks to find login forms
7. Miscellaneous
    - Unit tests
    - Linting

### Frontend

1. UI/UX improvements.
2. Use openapi client to connect to BE.
3. More input validation.
4. More error handling.
5. Miscellaneous
    - E2E tests
    - Linting

### Deployment

1. Add API Gateway
2. Add Load balancer
3. Deployment with K8s
4. CI/CD Pipelines
