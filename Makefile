build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./cmd/server/server ./cmd/server

run:build
	docker build -t test_task_strafov_net-scratch -f Dockerfile.scratch .
	docker run -it -p 55555:55555 test_task_strafov_net-scratch
