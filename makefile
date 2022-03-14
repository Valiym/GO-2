BINARY_NAME=Lesson-2

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin Lesson-2.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux Lesson-2.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows Lesson-2.go

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows