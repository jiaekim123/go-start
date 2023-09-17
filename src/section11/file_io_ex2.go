package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 파일 I/O (2)
func errorCheck3(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}
func main() {
	// 파일 읽기, 버퍼 사용 쓰기 -> bufio 패키지 활용
	// ioutil, bufio 등은 io.Reader, io.Writer 인터페이스를 구현 함 -> 즉, writer, read 메소드를 사용 가능
	/*
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	*/
	// 즉, bufio의 NewReader, NewWriter를 통해서 객체 반환 후 메소드 사용

	//bufio(Buffered io) 패키지
	// https://golang.org/pkg/bufio

	// 파일 열기
	// 두 번째 매개변수 확인
	// https://golang.org/pkg/os/#pkg-constants

	/*
		상태 (버퍼 사이즈가 4일 경우)
		a ------> a
		b ------> ab
		c ------> abc
		d ------> abcd
		e ------> e    -----> abcd
		f ------> ef   -----> abcd
		g ------> efg  -----> abcd
		h ------> efgh -----> abcd
		i ------> i    -----> abcdefgh
	*/

	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	defer file.Close()

	// bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file) // Writer 반환 (버퍼 사용)
	wt.WriteString("Hello Golang! \n FileWrite Test1! \n")
	wt.Write([]byte("Hello Golang! \n FileWrite Test2! \n"))

	// 에러 체크
	errorCheck3(err)

	// 버퍼 정보 출력
	fmt.Printf("사용한 버퍼 사이즈 (%d bytes)\n", wt.Buffered())
	fmt.Printf("남은 버퍼 사이즈 (%d bytes)\n", wt.Available())
	fmt.Printf("전체 버퍼 사이즈 (%d bytes)\n", wt.Size())

	wt.Flush() // 버퍼 비우고 반영(버퍼의 내용을 디스크에 기록한다.)
	fmt.Println("쓰기 작업 완료")
	fmt.Println("-------------------------")

	// 버퍼 정보 출력
	fmt.Printf("사용한 버퍼 사이즈2 (%d bytes)\n", wt.Buffered())
	fmt.Printf("남은 버퍼 사이즈2 (%d bytes)\n", wt.Available())
	fmt.Printf("전체 버퍼 사이즈2 (%d bytes)\n", wt.Size())

	rt := bufio.NewReader(file) // reader가 반환
	fi, err := file.Stat()
	errorCheck3(err)

	b := make([]byte, fi.Size())
	fmt.Println("파일 정보 출력 : ", fi)
	fmt.Println("파일 이름 : ", fi.Name())
	fmt.Println("파일 크기 : ", fi.Size())
	fmt.Println("파일 수정일자 : ", fi.ModTime())

	file.Seek(0, io.SeekStart)
	data, _ := rt.Read(b) // 읽기(ReadLine, ReadByte, ReadBytes 등)

	fmt.Printf("전체 버퍼 사이즈 : (%d bytes)\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d bytes)\n", data)
	fmt.Println("====================")
	fmt.Println(string(b))
	fmt.Println("====================")
}
