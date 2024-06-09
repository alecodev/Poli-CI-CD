# Integración Continua

### Steps to execute

> Create `.env` files in frontend and backend directories based on `.env-sample`

- To run it locally, execute `docker compose up -d`
- To run it in production, execute `docker compose -f docker-compose.yml -f docker-compose.production.yml up -d`

### Task list

#### Frontend

- [x] Initialize project with Vite
- [x] Add linter for TypeScript + React
- [x] Create TODO App - based on [TODO APP midudev](https://github.com/midudev/aprendiendo-react/tree/master/projects/08-todo-app-typescript)

#### Backend

- [x] Create API with Golang
- [x] Create unique session for each user
- [x] Store data in NoSQL database - Redis

#### Deploy

- [x] Dockerize application
- [x] Build Docker images and deploy with GitHub Actions
