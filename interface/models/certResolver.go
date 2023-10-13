package models

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

const SCRIPT_PATH = "/home/jeho/blockchain-repo/cloud_chain/scripts/"

func EnrollUser(request CertRequest) string {
	// 셸 스크립트 파일 경로
	scriptPath := SCRIPT_PATH

	// 명령 준비
	cmd := exec.Command("bash", scriptPath+"makeUser.sh", request.Org, request.UserID, request.Password)
	// 환경 변수 설정 (선택사항)
	//cmd.Env = append(os.Environ(), "VAR_NAME=VAR_VALUE")
	// 명령 실행 및 결과 수집
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(output)
		fmt.Println(scriptPath+"makeUser.sh")
		fmt.Printf("에러 발생: %s\n", err)
		return err.Error()
	}
	fmt.Printf("명령 실행 결과:\n%s\n", string(output))
	filePath := fmt.Sprintf("/home/jeho/blockchain-repo/cloud_chain/organizations/peerOrganizations/%s.pnu.cse/users/%s@%s.pnu.cse/msp/signcerts/cert.pem",
		request.Org, request.UserID, request.Org)
	// 파일 읽기
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("파일을 읽을 수 없음")
		return err.Error()
	}
	return string(fileContent)
}
