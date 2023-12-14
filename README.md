# birdpedia

A standalone web app written in Golang to build a community encyclopedia of birds.

Explore more details in this [blog post](https://lunarwhite.notion.site/Birdpedia-a-community-encyclopedia-of-birds-56f2b8930b3840a7a2f6c7689d8d6a5a).

```
.
├── LICENSE
├── README.md
├── assets
│   ├── index.html
│   └── script.js
├── bird_handlers.go
├── bird_handlers_test.go
├── go.mod
├── go.sum
├── main.go
└── main_test.go
```

## Feature

- Display the different entries submitted by the community, with the name and details of the bird they found.
- Allow anyone to post a new entry about a bird that they saw.

## Structure

1. The web server
2. The front-end (client side) app

## Tool chain

- Golang
