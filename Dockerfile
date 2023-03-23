# multi stage build
# step1. build step
FROM golang:1.19.2-alpine3.16 AS builder
# 이미지의 실제작업 디렉토리 설정
WORKDIR /app
# 첫번째 dot = 현재 디렉토리
# 두번째 dot = /app 디렉토리 
COPY . .
# build our app single binary executable file
RUN go build -o main main.go
      
# step2. run step
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
# go build 명령은 .go관련 파일들만 빌드 되기에 그외의 확장자 파일들은 필요에 따라 run 단계에세 복사해주어야함
COPY app.env .
COPY start.sh .
COPY db/migration ./db/migration

EXPOSE 8080
# Keep in mind that when the CMD instruction is used together with ENTRYPOINT, 
CMD [ "/app/main"]
ENTRYPOINT [ "/app/start.sh" ]