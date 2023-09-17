package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// 파일 읽기
	// CSV 파일 읽기 예제
	// 패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 가능
	// 패키지 Github 주소 : https://github.com/tealeg/xlsx
	// bufio : 파일이 용량이 클 경우 버퍼 사용 건장

	// 파일 생성
	file, err := os.Open("src/section11/sample.csv")
	checkErr(err)
	defer file.Close()

	// CSV Reader 생성
	//rr := csv.NewReader(file)
	rr := csv.NewReader(bufio.NewReader(file)) // 권장

	// CSV 내용 읽기(1)
	row, err := rr.Read() // 1개의 row 단위로 읽기 (위치가 이동됨.)
	checkErr(err)

	fmt.Println("csv read example")
	fmt.Println(row)
	fmt.Println(row[0], row[1], row[2:5])
	fmt.Println("----------------------------------")

	// CSV 내용 읽기(2)
	rows, err := rr.ReadAll() // 전체 Row 읽기 - 2차원 배열 형태
	checkErr(err)
	fmt.Println("CSV ReadAll Example")
	fmt.Println(rows)
	fmt.Println(rows[6])
	fmt.Println(rows[6][1])
	fmt.Println("----------------------------------")

	fileInfo, err := file.Stat() // 파일 사이즈 확인을 위해서 정보 획득
	checkErr(err)
	fmt.Println("파일 정보 출력(1) : ", fileInfo)
	fmt.Println("파일 이름(2) : ", fileInfo.Name())
	fmt.Println("파일 크기(3) : ", fileInfo.Size())
	fmt.Println("파일 수정 시간(4) : ", fileInfo.ModTime())

	// Row 단위로 CSV 파일 읽기

	for i, row := range rows {
		// fmt.Println(i, row)
		for j := range row {
			fmt.Printf("%s	", rows[i][j])
		}
		fmt.Println()
	}
}
