## FROM : FROM은 컨테이너에 사용할 기본 이미지
FROM golang:alpine

## 명령어 실행
RUN mkdir /app

## 작업 디렉토리
WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

### 컴파일 데몬
RUN go get github.com/githubnemo/CompileDaemon

## EXPOSE는 Docker 컨테이너에서 실행되는 서비스가 포트 8000에서 사용 가능함을 나타냄
EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main