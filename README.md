# GOS

<img width="1267" height="762" alt="gos" src="https://github.com/user-attachments/assets/18cc8a38-405a-4f5d-a5e8-f37c65b79613" />

*gos* is a simple CLI tool that help to make a project structure fast and quick!

# Design choices and goals

Gos is not intended to compete with any other project structuring tool it's a fun hobby project made because of my interest in CLI tools.

# Quick Start

### Golang

```
gos Make-project -l go -t std   --Creates the standard go project structure 

gos Make-project -l go -t api     --Creates a go boilerplate structure for API projects 

gos Make-project -l go -t cli     --Creates a go boilerplate structure for cli projects <lxsh-s approved :D>
```

### C++

```
gos Make-project -l cpp -t std    --Creates the standard cpp project strucutre 

gos Make-project -l cpp -t app    --Creates a cpp boilerplate structure for app/game

gos Make-project -l cpp -t lib    --Creates a cpp boilerplate structure for a reusable library
```

### Typescript

```
gos Make-project -l ts -t std   --Creates the standard Typescript project strucutre

gos Make-project -l ts -t lib   --Creates a Typescript boilerplate for a library

gos Make-project -l ts -t api   --Creates a Typescript boilerplate for a API projects

gos Make-project -l ts -t nxtjs --Creates a standard Typescript biolerplate for Next.js projects
```

# File structure

## Go

### `std` — Standard Go project

```
<project>/
├── cmd/
├── internal/
└── pkg/
```

### `api` — Go API server

```
<project>/
├── cmd/
│   └── api/
├── internal/
│   ├── handlers/
│   └── models/
└── pkg/
```

### `cli` — Go CLI application

```
<project>/
├── cmd/
├── internal/
│   └── config/
└── pkg/
    └── utils/
```

---

## TypeScript

### `std` — Standard TypeScript project

```
<project>/
├── src/
└── dist/
```

### `api` — TypeScript API server

```
<project>/
├── src/
│   ├── config/
│   ├── controllers/
│   ├── middlewares/
│   ├── models/
│   ├── routes/
│   ├── services/
│   └── types/
└── tests/
```

### `nxtjs` — Next.js project

```
<project>/
├── public/
└── src/
    ├── app/
    ├── components/
    ├── features/
    ├── lib/
    ├── hooks/
    └── types/
```

### `lib` — TypeScript library

```
<project>/
└── src/
    ├── utils/
    ├── types/
    └── __tests__/
```

---

## C++

### `std` — Standard C++ project

```
<project>/
├── src/
├── tests/
└── build/
```

### `app` — C++ application/games

```
<project>/
├── src/
├── include/
├── tests/
├── assets/
└── cmake/
```

### `lib` — C++ library

```
<project>/
├── src/
├── include/
│   └── <project>/
├── tests/
└── examples/
```

# Installation

### AUR (Arch User Repository)

```
yay -S gos-bin 
```

### Using Go

```bash
go install github.com/lxsh-S/gos@latest 
```

# Contribute

```bash
git clone https://github.com/lxsh-S/gos.git
cd gos
go run .
```
