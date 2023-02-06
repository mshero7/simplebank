# simplebank

기본 뱅킹기능을 가진 toy 프로젝트

README.md  계속 수정중..

개발환경
OS : Window 10
requirements : go(1.19) / grpc / postgresql(sqlc) / docker

테스트
main_test.go 의 MainTest를 먼저 시행하게된다..

항상 서로 독립적인 테스트를 보장해야한다.

createRandomAccount()로 분리시킨이유
ㄴ 다른 테스트에서도 동일한 기능이 필요해서 뺀 이유와, 종속이 강해져 예기치못한 오류가 생길수있어 안전하게 다른함수로 뻄

DeadLock
ㄴ 여러 트랜잭션이 생성될때 FK 참조로 데드락이 발생할수도있다

2/6
1. Gin(HTTP framework) 을 통한 API 구축
2. 기존 const로 선언된 환경변수 Viper로 파일에서 읽어들이게끔 변경