# multi stage build
# step1. build step
FROM golang:1.19.2-alpine3.16 AS builder
# 이미지의 실제작업 디렉토리 설정
WORKDIR /app
# 첫번째 . 는 simplebank 하위 모든파일을 루트디렉토리에 복사
# 두번째 . 는 루트에서 작업디렉토리(/app)로 이동
COPY . .
# build our app single binary executable file
RUN go build -o main 

# step2. run step
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .


EXPOSE 8080
CMD [ "/app/main"]