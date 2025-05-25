# Trend Summary Engine

Trend Summary Engine is a full-stack SaaS application designed to extract and summarize insights from online articles. Users can register, submit article URLs, and receive AI-generated summaries via a simple interface. The goal is to build a portfolio-quality backend-focused project with a production-grade architecture using Go, GraphQL, and AWS.

## 🚀 Features

- User registration and login (JWT auth)
- Submit article URLs to be summarized
- Background scraping and summarization pipeline (planned)
- GraphQL API for interaction
- Frontend built with React + Vite + TailwindCSS (WIP)

## 🧱 Tech Stack

**Backend:**

- Go
- GraphQL (gqlgen)
- PostgreSQL
- JWT Authentication
- Docker

**Frontend:**

- React
- TypeScript
- Vite
- TailwindCSS

**Infrastructure:**

- AWS (EC2, S3, RDS) — planned
- Docker Compose (for local dev)
- Terraform (for IaC) — planned
- Message queue (RabbitMQ or NATS) — planned

---

## ✅ MVP Status

### 🔒 Authentication

- [x] User registration
- [x] User login
- [x] JWT-based route protection

### 🌐 Article Handling

- [x] Save article URLs (GraphQL mutation)
- [ ] Scrape full content from URLs (pending)
- [ ] Summarize content via OpenAI API (pending)

### 🔎 GraphQL API

- [x] Auth-protected `saveArticle` mutation
- [ ] `getSavedArticles` query
- [ ] `getSummary` query

### 🖥️ Frontend

- [x] Vite + React + TailwindCSS set up
- [ ] Login/Registration form
- [ ] Article URL submission form
- [ ] Display summaries

### ☁️ Infrastructure

- [x] Docker setup for backend
- [ ] Docker Compose for multi-service local dev
- [ ] AWS deployment (EC2, RDS, S3)
- [ ] Terraform provisioning scripts

---

## 🛠️ Next Steps to Complete MVP

1. **Backend Worker**

   - Implement `cmd/worker/main.go`
   - Extract article content using readability libraries
   - Call OpenAI (or placeholder) to generate summaries
   - Store summaries in the database

2. **Frontend UI**

   - Build forms for registration, login, and article submission
   - Display saved articles and summaries

3. **GraphQL Enhancements**

   - Add `getSavedArticles` and `getSummary` queries

4. **Infrastructure & DevOps**
   - Add Docker Compose config
   - Add Terraform config for AWS deployment
   - Deploy full stack on AWS

---

## 📁 Repository Structure

```
trend-summary-engine/
├── backend/
│   ├── cmd/              # Entrypoints for API and worker
│   ├── graph/            # GraphQL schema and resolvers
│   ├── internal/         # Auth, DB, config logic
│   └── Dockerfile
├── frontend/
│   ├── src/              # React + Tailwind app
│   └── vite.config.ts
└── docker-compose.yml    # (Planned)
```

---

## 📌 Project Goal

This project is being developed to demonstrate backend and infrastructure engineering expertise, with a focus on GraphQL APIs, background processing, and modern cloud-native architecture. It serves as a portfolio piece for applying to full stack and backend-focused roles at companies like Roblox, Plaid, and Grammarly.

---

## 👤 Author

Developed by [markiskorova](https://github.com/markiskorova)
