# blockchain-
졸업과제 - 클라우드/블록체인/MSA 중고차 거래 플랫폼, 블록체인 플랫폼 레포지토리

### 전체 아키텍처

<img width="627" alt="스크린샷 2023-10-17 오후 11 23 00" src="https://github.com/Cloud-Chain/blockchain-repo/assets/78259314/ec7ad1fb-ded7-40fa-a40e-75e64efc25b6">

### 블록체인 아키텍처

<img width="481" alt="스크린샷 2023-10-17 오후 11 23 32" src="https://github.com/Cloud-Chain/blockchain-repo/assets/78259314/1dcc583c-cf85-4775-bc92-ef4752150cbd">

### 주요 디렉토리 구조

중요 디렉토리만 반영되었습니다.

📦 blockchain-repo    
 ┣ 📂 cloud_chain : 블록체인 네트워크 구성 파일 <br>
 ┃ ┣ 📂 chaincodes : 스마트 컨트랙트 파일 <br>
 ┃ ┣ 📂 compose : 블록체인 네트워크 docker compose 파일 <br>
 ┃ ┣ 📂 prometheus-grafana : 블록체인 네트워크 + 내부 컨테이너 모니터링을 위한 구성 파일 <br>
 ┃ ┣ 📂 scripts : 블록체인 관리를 위한 쉘스크립트 파일 <br>
 ┣ 📂 config : 블록체인 설정 파일 <br>
 ┣ 📂 interface : 블록체인 <-> 시스템 통신을 위한 인터페이스 서버 파일 <br>
 ┃ ┣ 📂 config : 블록체인 통신을 위한 Fabric-Gatewat SDK 설정 파일 <br>
 ┃ ┣ 📂 handlers : 요청 처리를 담당하는 파일 <br>
 ┃ ┣ 📂 models :  블록체인 데이터를 다루는 파일 <br>
 ┗ ┗ 📂 router : 라우팅 및 URL 처리를 담당하는 파일 <br>
