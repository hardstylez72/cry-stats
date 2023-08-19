
dev := '170.64.160.93'
deploy-dev:
	rm -rf build
	mkdir build
	cp -r migrations build/migrations/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./build/pay ./cmd/main.go
	upx ./build/pay --best --lzma
	cp .env.dev ./build/.env
	cp Dockerfile ./build/Dockerfile
	cp docker-compose.yml ./build/docker-compose.yaml
	rsync --progress -re ssh ./build/ root@$(dev):/root/app/pay/
	ssh -t root@$(dev) "docker-compose -f /root/app/pay/docker-compose.yaml build --force"
	ssh -t root@$(dev) "docker-compose -f /root/app/pay/docker-compose.yaml up -d"

prod := '128.199.142.47'
deploy-prod:
	rm -rf build
	mkdir build
	cp -r migrations build/migrations/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./build/pay ./cmd/main.go
	upx ./build/pay --best --lzma
	cp .env.prod ./build/.env
	cp Dockerfile ./build/Dockerfile
	cp docker-compose.yml ./build/docker-compose.yaml
	rsync --progress -re ssh ./build/ root@$(prod):/root/app/pay/
	ssh -t root@$(prod) "docker-compose -f /root/app/pay/docker-compose.yaml build --force"
	ssh -t root@$(prod) "docker-compose -f /root/app/pay/docker-compose.yaml up -d"

build-standalone:
	rm -rf build
	mkdir build
	cp -r migrations build/migrations/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./build/pay ./cmd/main.go
	upx ./build/pay --best --lzma
	cp Dockerfile ./build/Dockerfile
	cp standalone.sh ./build/standalone.sh