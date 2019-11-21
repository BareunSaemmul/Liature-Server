# Liature Main Server
Liature의 메인 서버 폴더입니다.

GET `/` : (HTML)메인 페이지
GET `/login` : (HTML)로그인 페이지
GET `/detail`: (HTML)상세 정보(날씨 및 자연재해에 대한)
GET `/rooms/messages` : 모든 메시지 조회
POST `/ws/room` : 채팅의 새로운 클라이언트 생성
  - area: 지역 정보(어느 지역인지 ex: "대전", "광주", ..)