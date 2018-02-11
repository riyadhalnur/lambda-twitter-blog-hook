build:
	go get github.com/aws/aws-lambda-go/lambda
	go get github.com/aws/aws-lambda-go/events
	go get github.com/ChimeraCoder/anaconda
	env GOOS=linux go build -ldflags="-s -w" -o bin/blogposter blogposter/main.go
