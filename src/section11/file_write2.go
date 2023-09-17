package main

import (
	_ "bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 파일 쓰기
	// CSV 파일 쓰기 예제
	// 패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 기능
	// 패키지 Github 주소 : https://github.com/tealeg/xlsx
	// bufio : 파일이 용량이 클 경우 버퍼 사용 권장

	// 파일 생성
	file, err := os.Create("test_write.csv")
	errorCheck2(err)

	// 리소스 해제
	defer file.Close()

	// csv write 생성
	wr := csv.NewWriter(file)
	//wr := csv.NewWriter(bufio.NewWriter(file))

	// csv 내용 쓰기
	wr.Write([]string{"Kim", "4.8"})
	wr.Write([]string{"Lee", "4.8"})
	wr.Write([]string{"Park", "4.8"})
	wr.Write([]string{"Cho", "4.8"})
	wr.Write([]string{"Hong", "4.8"})

	wr.Flush() // 버퍼 -> 파일로 쓰기

	fi, err := file.Stat()
	errorCheck2(err)

	fmt.Printf("CSV 쓰기 작업 후 작업 크기 (%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일 명 : ", fi.Name())
	fmt.Println("운영체제 파일 권한 : ", fi.Mode())
}

// 에러 체크 방식 1
//func errorCheck1(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

// 에러 체크 방식 2
func errorCheck2(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
