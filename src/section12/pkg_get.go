// 사용자 패키지 설치 및 활용 예제
package main

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	// 외부 저장소 패키지 설치
	// 2가지 방법
	// 1. import 선언 후 폴더 이동 후 go get 설치 -> 사용
	// 2. go get 패키지 주소 설치 -> 선언

	// 선언 후 Go get 설치 예제(엑셀 파일 읽기)
	xfile := "src/section12/sample.xlsx"
	xlFile, err := xlsx.OpenFile(xfile)

	if err != nil {
		panic(errors.New("Excel Load Error!"))
	}

	for _, sheet := range xlFile.Sheets {
		fmt.Println("Max Row : ", sheet.MaxRow)

		err := sheet.ForEachRow(rowVisitor)
		if err != nil {
			fmt.Println("Error Occurred! ", err.Error())
		}
	}
}

func rowVisitor(r *xlsx.Row) error {
	err := r.ForEachCell(cellVisitor)
	if err != nil {
		return err
	}
	return nil
}

func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(value)
	}
	return nil
}
