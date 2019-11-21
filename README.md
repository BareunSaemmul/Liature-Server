# Liature Main Server
Liature의 메인 서버 폴더입니다.

GET `/` : (HTML)메인 페이지

GET `/login` : (HTML)로그인 페이지

GET `/detail`: (HTML)상세 정보(날씨 및 자연재해에 대한)

GET `/rooms`: 모든 채팅방의 정보 조회

POST `/rooms/messages` : 특정 채팅방의 메시지 조회

GET `/ws/room/{area}` : 채팅의 새로운 클라이언트 생성({}로 감싸진 부분에 지역Data -URL 예시 => /ws/room/Daejeon)