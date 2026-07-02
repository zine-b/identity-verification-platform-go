FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# CGO_ENABLED=0 : cgo permet à Go d’appeler du code C. En le désactivant, tu obtiens un binaire plus simple, souvent plus portable 
# GOOS=linux : force la compilation pour Linux.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api ./cmd/api


FROM alpine:3.22

WORKDIR /app

# créer un groupe et un utilisateur
# Crée un groupe système appelé : app
# Crée un utilisateur système appelé : app Et l’ajoute au groupe : app
RUN addgroup -S app && adduser -S app -G app

COPY --from=builder /app/api /app/api

USER app

EXPOSE 8080

ENTRYPOINT ["/app/api"]