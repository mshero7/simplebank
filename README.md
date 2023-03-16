# simplebank

기본 뱅킹기능을 가진 toy 프로젝트

README.md  계속 수정중..

개발환경
ㄴ OS : Window 10
ㄴ requirements : go(1.19) / grpc / postgresql(sqlc) / docker

DB 생성
ㄴ https://dbdiagram.io/home

내용 정리

테스트
ㄴ main_test.go 의 MainTest를 먼저 시행하게된다.. 항상 서로 독립적인 테스트를 보장해야한다.

createRandomAccount()로 분리시킨이유
ㄴ 다른 테스트에서도 동일한 기능이 필요해서 뺀 이유와, 종속이 강해져 예기치못한 오류가 생길수있어 안전하게 다른함수로 뻄

DeadLock
ㄴ 여러 트랜잭션이 생성될때 FK 참조로 데드락이 발생할수도있다



2/6
1. Gin(HTTP framework) 을 통한 API 구축
2. 기존 const로 선언된 환경변수 Viper로 파일에서 읽어들이게끔 변경
3. Mock DB for testing HTTP API in Go
mockgen 을 통한 mock db 사용.
ㄴ 아래이슈 나올때 참고 // https://github.com/golang/mock/issues/494
prog.go:12:2: no required module provides package github.com/golang/mock/mockgen/model: go.mod file not found in current directory or any parent directory; see 'go help modules'
3-1. 맥북 mockgen 을 위한 환경변수 추가
ㄴ export PATH=$PATH:$(go env GOPATH)/bin

각 패키지에서 main_test.go 파일로 선행될 부분을 정해줄수있다.
ㄴ 여기선 gin의 testmode 세팅을 이용하기 위해 main_test.go를 썼다.

2/7
1. Custom Validator in go 
```go
    형 변환시, 형변환 값, err 가 나온다.
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
```
    Gin 의 기본 밸리데이터는 github.com/go-playground/validator/v10 이라, 새로운 밸리데이터 함수 선언한 뒤, Gin 에 등록해주면 끝
    ex) 
        var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {

2/8
1. Add users Table
    기존 account 와 동일한 내용들이 있어 그대로 사용해 생성.
    makefile 수정
        migrate down [N] > N 플래그로 이전 N 단계까지 migrate down 할 수 있다.

2/18
1. 토큰 인증(paseto) 구현 및 전체 적용 + 테스트 적용
2. 관련 엔드포인트별 인증 적용 및 테스트 적용
3. mock, sqlc 재정의.

4. bracnch check test

2/21
1. docker-compose 구성
2. DB container 구성을 위해 depends_on 활용
3. db 민감정보 app.env 로 이관

3/16
1. gRPC createUser, loginUser 구현
1-1. https://grpc.io/docs/languages/go/quickstart/ >>> Prerequisites
1-2. 메인객체가 될 User > user.proto
     User를 사용할 RPC proto 구현
     RPC를 활용한 Service(gRPC server) 구현
     pb 폴더 하위에 go 파일 생성
     gapi 폴더에 Server 객체 구현
```go
     // Server serves gRPC requests for out banking services
    type Server struct {
        pb.UnimplementedSimpleBankServer // embeding, 구현한 RPC 이용 가능
        config                           util.Config
        store                            db.Store    // interact with the database processing API requests from client
        router                           *gin.Engine // help us send each API request to the correct handler for processing
        tokenMaker                       token.Maker
    }
```
2. evans cli 활용
3. https://github.com/ktr0731/evans
