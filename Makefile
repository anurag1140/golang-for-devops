version: "3"

tasks:
  fmt:
    cmds:
      - go fmt ./...

  lint:
    cmds:
      - golangci-lint run

  test:
    cmds:
      - go test ./...

  swagger:
    cmds:
      - swag init

  docker:
    cmds:
      - docker compose build --no-cache

  run:
    cmds:
      - docker compose up