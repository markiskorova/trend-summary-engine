# 🧠 Trend Summary Engine

**Trend Summary Engine** is a backend-focused SaaS app that allows users to save article URLs and prepares them for AI-powered summarization. Built with Go, PostgreSQL, GraphQL, and JWT-based authentication, it's designed for modular growth and fast iteration.

---

## 🔍 Problem

In a world of information overload, professionals struggle to stay informed. Trend Summary Engine enables users to submit articles and (in future phases) receive concise summaries using AI.

---

## 🚀 Features (Phase 1 MVP)

- ✅ JWT-secured user authentication (register + login)
- ✅ Submit and store article URLs via GraphQL
- ✅ Modular Go backend with graphql-go and GORM
- ✅ Dockerized setup with PostgreSQL and Docker Compose
- ✅ Clean `.env`/`.env.example` configuration and Git hygiene

---

## 🧱 Tech Stack

- **Language**: Go
- **Frameworks**: graphql-go · GORM
- **Auth**: JWT · bcrypt
- **Database**: PostgreSQL
- **Infrastructure**: Docker · Docker Compose
- **Future**: OpenAI API · RabbitMQ · React + Vite · Terraform

---

## 📁 Project Structure

```
trend-summary-engine/
├── backend/
│   ├── cmd/api/              # GraphQL API server
│   ├── internal/
│   │   ├── auth/             # JWT and password handling
│   │   ├── config/           # Environment loading
│   │   ├── db/               # GORM models and DB setup
│   │   └── graph/            # GraphQL schema and resolvers
│   ├── .env                  # Local-only environment config
│   ├── .env.example          # Shared environment template
│   └── Dockerfile            # Multi-stage backend build
├── docker-compose.yml
└── README.md
```

---

## 🧪 Running Locally

### Prerequisites
- Docker & Docker Compose installed

### Run the stack

```bash
docker-compose up --build
```

GraphQL endpoint: [http://localhost:8080/graphql](http://localhost:8080/graphql)  
Use GraphiQL UI to register, login, and submit articles.

---

## 📅 Roadmap

- ✅ Phase 1: Backend auth + GraphQL API
- 🚧 Phase 2: Background scraping & summarization
- 🚧 Phase 3: Frontend UI (React + Vite)
- 🚧 Phase 4: Enhanced GraphQL features & tests
- 🚧 Phase 5: AWS deployment via Terraform

---

## 🤝 Contributing

Contributions are welcome! Future phases will include frontend scaffolding and infrastructure extensions.

---

## 📄 License

This project is licensed under the MIT License.