.PHONY: build install

build:
	CGO_ENABLED=0 GOOS=linux go build -a

windows:
	GOOS=windows GOARCH=386 go build -o mdmerge.exe

docker: build
	docker build -t briefbote/mdmerge:latest .
	docker push briefbote/mdmerge:latest