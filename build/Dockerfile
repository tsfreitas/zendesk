# syntax=docker/dockerfile:1.3-labs
FROM public.ecr.aws/bitnami/golang:1.17 as builder
RUN yum update -y && yum install -y git

ARG BUF_USER
ARG BUF_TOKEN
ARG PERSONAL_ACCESS_TOKEN

WORKDIR /project

COPY go.mod ./go.mod
COPY go.sum ./go.sum

RUN <<EOF cat > ~/.netrc
machine buf.build
  login ${BUF_USER}
  password ${BUF_TOKEN}

machine go.buf.build
  login ${BUF_USER}
  password ${BUF_TOKEN}
EOF

RUN git config --global url.https://${PERSONAL_ACCESS_TOKEN}@github.com/.insteadOf https://github.com/

RUN go mod download

COPY ./ .

RUN go build -o /app/main cmd/main.go

FROM public.ecr.aws/lambda/go:1

COPY --from=builder /project/configs/configs.env /var/task/configs/configs.env
COPY --from=builder /app/main /var/task/main

# Command can be overwritten by providing a different command in the template directly.
CMD ["main"]
