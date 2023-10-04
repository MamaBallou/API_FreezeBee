# Étape de construction
FROM golang:latest as builder

WORKDIR /app

# Copie seulement les fichiers nécessaires
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o main .

# Étape de production
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copie de l'exécutable depuis l'étape de construction
COPY --from=builder /app/main .

CMD ["./main"]
