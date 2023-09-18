// 최종 실습 예제
// 대상 사이트: 루리웹(ruliweb.com)
package main

import (
	"bufio"
	"fmt"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"os"
	"strings"
	_ "strings"
	"sync"
)

// go get github.com/yhat/scrape : 설치
// https://github.com/yhat/scrape : Go simple scrap package -> 사용하기 어렵지만, 코드 학습 위해서 사용
// http://go-colly.org/docs : Go scrap and crawling library -> goquery 기반 굉장히 강력하고 쉬운 패키지 (가장 많이 사용)
// https://github.com/PuerkitoBio/goquery : 쉬운 HTML Parsing 지원

// 스크랩핑 대상 URL
const (
	urlRootRuri            = "https://ruliweb.com"
	urlRootBoardGameGeek   = "https://boardgamegeek.com/browse/boardgame"
	urlBoardGameGeekDomain = "https://boardgamegeek.com"
)

// 동기화를 위한 작업 그룹 선언
var wg sync.WaitGroup

// 에러 체크 공통 함수
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// 루리웹 사이트 (구)
// 첫 번째 방문(메인 페이지) 대상으로 원하는 URL을 파싱 후 존재하는지 반환하는 함수
func parseMainNodesForRuliweb(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n.Parent, "class") == "row"
	}
	return false
}

// 보드게임 긱 사이트
// 첫 번째 방문(메인 페이지) 대상으로 원하는 URL을 파싱 후 존재하는지 반환하는 함수
func parseMainNodesForBoardGameGeek(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil &&
		n.Parent.Parent.DataAtom == atom.Td {
		return scrape.Attr(n.Parent.Parent, "class") == "collection_objectname \n\t\t\t\t\t\t browse"
	}
	return false
}

// 사이트 타이틀을 가져온다.
func parseSiteTitle(n *html.Node) (string, error) {
	title, ok := scrape.Find(n, scrape.ByTag(atom.Title))
	if ok {
		siteTitle := scrape.Text(title)
		fmt.Println("사이트 타이틀 : ", siteTitle)
		return siteTitle, nil
	}
	return "", fmt.Errorf("can not find title")
}

// 보드게임 긱
// url 대상이 되는 페이지(서브 페이지) 대상으로 원하는 내용을 파싱 후 저장
func scrapeContentsForGeek(url string, fileName string) {
	// 작업 종료 알림
	defer wg.Done()

	// GET 방식 요청
	resp, err := http.Get(url)
	errCheck(err)

	// 요청 Body 닫기
	defer resp.Body.Close()

	// 응답 데이터(html)
	root, err := html.Parse(resp.Body)
	errCheck(err)

	// Response 데이터(html)의 원하는 부분 파싱
	matchNode := func(n *html.Node) bool {
		return n.DataAtom == atom.A && n.Parent.DataAtom == atom.H1 && scrape.Attr(n.Parent.Parent, "class") == "game-header-title-info"
	}

	// 파일 스트림 생성(열기) -> 파일명, 옵션, 권한
	file, err := os.OpenFile("src/section13/scrape/"+fileName+".txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errCheck(err)

	// 메소드 종료 시 파일 닫기
	defer file.Close()

	// 쓰기 버퍼 선언
	w := bufio.NewWriter(file)

	// matchNode 메소드를 사용해서 원하는 노드 순회 (iterator)하면서 출력
	for _, g := range scrape.FindAll(root, matchNode) {
		// URL 및 해당 데이터 출력
		fmt.Println("result : ", scrape.Text(g))
		// 파싱 데이터 -> 버퍼에 기록
		w.WriteString(scrape.Text(g) + "\r\n")
	}
	w.Flush()
}

// 루리웹
// url 대상이 되는 페이지(서브 페이지) 대상으로 원하는 내용을 파싱 후 저장
func scrapeContentsForRuliweb(url string, fileName string) {
	// 작업 종료 알림
	defer wg.Done()

	// GET 방식 요청
	resp, err := http.Get(url)
	errCheck(err)

	// 요청 Body 닫기
	defer resp.Body.Close()

	// 응답 데이터(html)
	root, err := html.Parse(resp.Body)
	errCheck(err)

	// Response 데이터(html)의 원하는 부분 파싱
	matchNode := func(n *html.Node) bool {
		return n.DataAtom == atom.A && scrape.Attr(n, "class") == "deco"
	}

	// 파일 스트림 생성(열기) -> 파일명, 옵션, 권한
	file, err := os.OpenFile("scrape/"+fileName+".txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errCheck(err)

	// 메소드 종료 시 파일 닫기
	defer file.Close()

	// 쓰기 버퍼 선언
	w := bufio.NewWriter(file)

	// matchNode 메소드를 사용해서 원하는 노드 순회 (iterator)하면서 출력
	for _, g := range scrape.FindAll(root, matchNode) {
		// URL 및 해당 데이터 출력
		fmt.Println("result : ", scrape.Text(g))
		// 파싱 데이터 -> 버퍼에 기록
		w.WriteString(scrape.Text(g) + "\r\n")
	}
	w.Flush()

}

func main() {
	// 목표: 루리웹의 각 게임 게시판의 타이틀들을 모두 읽어와서 txt 파일로 저장한다.
	// 각 게시판당 txt 파일 하나로 저장할 수 있도록!
	// 2023.09 기준 루리웹 형식이 바뀌어서 다른 예제로 변경해서 진행

	// 메인 페이지 GET 방식 요청
	resp, err := http.Get(urlRootBoardGameGeek)
	errCheck(err)

	// 요청 Body 닫기
	defer resp.Body.Close()

	// 응답 데이터(Html)
	root, err := html.Parse(resp.Body)
	errCheck(err)
	_, err = parseSiteTitle(root)
	errCheck(err)

	// ParseMainNodes 메소드를 크롤링(스크랩핑) 대상 URL 추출
	urlList := scrape.FindAll(root, parseMainNodesForBoardGameGeek)

	for _, urlNode := range urlList {
		// 대상 URL 1차 출력
		//fmt.Println("Check Main Link : ", urlNode, idx)
		//fmt.Println("TargetUrl : ", scrape.Attr(urlNode, "href"))
		url := scrape.Attr(urlNode, "href")
		targetUrl := urlBoardGameGeekDomain + url
		fileName := strings.Split(url, "/")[3]
		fmt.Println("Filename : ", fileName)
		fmt.Println("TargetUrl : ", targetUrl)
		//fmt.Println(strings.Replace("oink oink oink oink", "k", "ky", 2)) // Replace 사용도 자주 쓰임.

		// 작업 대기열에 추가
		wg.Add(1) // Done 개수와 일치
		// 고루틴 시작 -> 작업 대기열 개수와 같아야 함

		go scrapeContentsForGeek(targetUrl, fileName)

	}
	// 모든 작업이 종료될 때까지 대기
	wg.Wait()
}
