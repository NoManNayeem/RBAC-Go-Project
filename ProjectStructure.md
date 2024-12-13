# Project Structure:

```
RBAC-Go-Project/
├── .env
├── .gitignore
├── docker-compose.yml
├── env.example
├── ProjectStructure.md
├── Go-Backend/
│   ├── .env
│   ├── env.example
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── main
│   ├── cmd/
│   │   └── main.go
│   ├── config/
│   │   ├── db.go
│   │   └── config.go
│   ├── internal/
│   │   ├── handlers/
│   │   ├── middlewares/
│   │   │   └── rbac.go
│   │   └── models/
│   │       └── user.go
│   └── scripts/
│       └── seed_data.sql
├── react-frontend/
│   ├── .env
│   ├── env.example
│   ├── Dockerfile
│   ├── package.json
│   ├── package-lock.json
│   ├── README.md
│   ├── node_modules/ (ignored)
│   ├── public/
│   │   ├── index.html
│   │   └── favicon.ico
│   └── src/
│       ├── App.js
│       ├── index.js
│       ├── components/
│       │   ├── AdminPage.js
│       │   ├── LoginPage.js
│       │   └── PublicPage.js
│       └── utils/
│           ├── api.js
│           └── roles.js

```