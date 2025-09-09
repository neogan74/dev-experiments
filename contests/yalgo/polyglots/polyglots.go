package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	readLine := func() string {
		s, _ := in.ReadString('\n')
		return strings.TrimRight(s, "\r\n")
	}
	// читаем N
	Nline := readLine()
	for Nline == "" {
		Nline = readLine()
	}
	N, _ := strconv.Atoi(strings.TrimSpace(Nline))

	var common map[string]struct{}
	all := make(map[string]struct{})

	for i := 0; i < N; i++ {
		// читаем Mi
		line := readLine()
		for line == "" {
			line = readLine()
		}
		M, _ := strconv.Atoi(strings.TrimSpace(line))

		// читаем Mi строк-языков
		curr := make(map[string]struct{}, M)
		for j := 0; j < M; j++ {
			lang := readLine()
			// допускаем пустые строки как валидное имя? обычно нет, но на всякий случай трим
			lang = strings.TrimSpace(lang)
			curr[lang] = struct{}{}
			all[lang] = struct{}{}
		}

		if i == 0 {
			// для первого школьника common = его множество
			common = curr
		} else {
			// пересечение: оставляем только те, что есть и в curr
			for k := range common {
				if _, ok := curr[k]; !ok {
					delete(common, k)
				}
			}
		}
	}

	// подготовим отсортированные списки к выводу
	commonList := make([]string, 0, len(common))
	for k := range common {
		commonList = append(commonList, k)
	}
	sort.Strings(commonList)

	allList := make([]string, 0, len(all))
	for k := range all {
		allList = append(allList, k)
	}
	sort.Strings(allList)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Println("RESULT")
	// 1) языки, которые знают все
	fmt.Fprintln(out, len(commonList))
	for _, s := range commonList {
		fmt.Fprintln(out, s)
	}
	// 2) языки, которые знает хотя бы один
	fmt.Fprintln(out, len(allList))
	for _, s := range allList {
		fmt.Fprintln(out, s)
	}
}

func Solution2() {

	reader := bufio.NewReaderSize(os.Stdin, 1024)

	var N int
	fmt.Fscan(reader, &N)

	m := make(map[string]int, 500)

	for range N {
		var n int
		fmt.Fscan(reader, &n)
		for range n {
			var lang string
			fmt.Fscan(reader, &lang)
			m[lang]++
		}
	}

	allLang := make([]string, 0)
	commonLang := make([]string, 0)
	for k, v := range m {
		if v == N {
			allLang = append(allLang, k)
		}
		commonLang = append(commonLang, k)
	}

	fmt.Println(len(allLang))
	for _, v := range allLang {
		fmt.Println(v)
	}

	fmt.Println(len(commonLang))
	for _, v := range commonLang {
		fmt.Println(v)
	}
}
