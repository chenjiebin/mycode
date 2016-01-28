package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.Match("H.* ", []byte("Hello World!")))

	r := bytes.NewReader([]byte("Hello World!"))
	fmt.Println(regexp.MatchReader("H.* ", r))

	fmt.Println(regexp.MatchString("H.* ", "Hello World!"))

	fmt.Println(regexp.QuoteMeta("(?P:Hello) [a-z]"))

	reg, err := regexp.Compile(`\w+`)
	fmt.Printf("%q,%v\n", reg.FindString("Hello World!"), err)
	// "Hello",

	reg, err = regexp.CompilePOSIX(`[[:word:]]+`)
	fmt.Printf("%q,%v\n", reg.FindString("Hello World!"), err)
	// "Hello"

	reg = regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindString("Hello World!"))
	// Hello

	reg = regexp.MustCompilePOSIX(`[[:word:]].+ `)
	fmt.Printf("%q\n", reg.FindString("Hello World!"))
	// "Hello "

	reg = regexp.MustCompile(`\w+`)
	fmt.Printf("%q\n", reg.Find([]byte("Hello World!")))
	// "Hello"

	reg = regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindString("Hello World!"))
	// "Hello"

	reg = regexp.MustCompile(`\w+`)
	fmt.Printf("%q\n", reg.FindAll([]byte("Hello World!"), -1))
	// ["Hello" "World"]

	reg = regexp.MustCompile(`\w+`)
	fmt.Printf("%q\n", reg.FindAllString("Hello World!", -1))
	// ["Hello" "World"]

	reg = regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindIndex([]byte("Hello World!")))
	// [0 5]

	reg = regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindStringIndex("Hello World!"))
	// [0 5]

	r = bytes.NewReader([]byte("Hello World!"))
	reg = regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindReaderIndex(r))
	// [0 5]

	reg = regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindAllIndex([]byte("Hello World!"), -1))
	// [[0 5] [6 11]]

	reg = regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindAllStringIndex("Hello World!", -1))
	// [[0 5] [6 11]]

	reg = regexp.MustCompile(`(\w)(\w)+`)
	fmt.Printf("%q\n", reg.FindSubmatch([]byte("Hello World!")))
	// ["Hello" "H" "o"]

	reg = regexp.MustCompile(`(\w)(\w)+`)
	fmt.Printf("%q\n", reg.FindStringSubmatch("Hello World!"))
	// ["Hello" "H" "o"]

	reg = regexp.MustCompile(`(\w)(\w)+`)
	fmt.Printf("%q\n", reg.FindAllSubmatch([]byte("Hello World!"), -1))
	// [["Hello" "H" "o"] ["World" "W" "d"]]

	reg = regexp.MustCompile(`(\w)(\w)+`)
	fmt.Printf("%q\n", reg.FindAllStringSubmatch("Hello World!", -1))
	// [["Hello" "H" "o"] ["World" "W" "d"]]

	reg = regexp.MustCompile(`(\w)(\w)+`)
	fmt.Println(reg.FindSubmatchIndex([]byte("Hello World!")))
	// [0 5 0 1 4 5]

	reg = regexp.MustCompile(`(\w)(\w)+`)
	fmt.Println(reg.FindStringSubmatchIndex("Hello World!"))
	// [0 5 0 1 4 5]

	r = bytes.NewReader([]byte("Hello World!"))
	reg = regexp.MustCompile(`(\w)(\w)+`)
	fmt.Println(reg.FindReaderSubmatchIndex(r))
	// [0 5 0 1 4 5]

	reg = regexp.MustCompile(`(\w)(\w)+`)
	fmt.Println(reg.FindAllSubmatchIndex([]byte("Hello World!"), -1))
	// [[0 5 0 1 4 5] [6 11 6 7 10 11]]

	reg = regexp.MustCompile(`(\w)(\w)+`)
	fmt.Println(reg.FindAllStringSubmatchIndex("Hello World!", -1))
	// [[0 5 0 1 4 5] [6 11 6 7 10 11]]

	reg = regexp.MustCompile(`(\w+),(\w+)`)
	src := []byte("Golang,World!")           // 源文本
	dst := []byte("Say: ")                   // 目标文本
	template := []byte("Hello $1, Hello $2") // 模板
	match := reg.FindSubmatchIndex(src)      // 解析源文本
	// 填写模板，并将模板追加到目标文本中
	fmt.Printf("%q\n", reg.Expand(dst, template, src, match))
	// "Say: Hello Golang, Hello World"

	reg = regexp.MustCompile(`(\w+),(\w+)`)
	srcString := "Golang,World!"                   // 源文本
	dst = []byte("Say: ")                          // 目标文本（可写）
	templateString := "Hello $1, Hello $2"         // 模板
	match = reg.FindStringSubmatchIndex(srcString) // 解析源文本
	// 填写模板，并将模板追加到目标文本中
	fmt.Printf("%q\n", reg.ExpandString(dst, templateString, srcString, match))
	// "Say: Hello Golang, Hello World"

	reg = regexp.MustCompile(`Hello[\w\s]+`)
	fmt.Println(reg.LiteralPrefix())
	// Hello false
	reg = regexp.MustCompile(`Hello`)
	fmt.Println(reg.LiteralPrefix())
	// Hello true

	text := `Hello World, 123 Go!`
	pattern := `(?U)H[\w\s]+o` // 正则标记“非贪婪模式”(?U)
	reg = regexp.MustCompile(pattern)
	fmt.Printf("%q\n", reg.FindString(text))
	// Hello
	reg.Longest() // 切换到“贪婪模式”
	fmt.Printf("%q\n", reg.FindString(text))
	// Hello Wo

	b := []byte(`Hello World`)
	reg = regexp.MustCompile(`Hello\w+`)
	fmt.Println(reg.Match(b))
	// false
	reg = regexp.MustCompile(`Hello[\w\s]+`)
	fmt.Println(reg.Match(b))
	// true

	r = bytes.NewReader([]byte(`Hello World`))
	reg = regexp.MustCompile(`Hello\w+`)
	fmt.Println(reg.MatchReader(r))
	// false
	r.Seek(0, 0)
	reg = regexp.MustCompile(`Hello[\w\s]+`)
	fmt.Println(reg.MatchReader(r))
	// true

	s := `Hello World`
	reg = regexp.MustCompile(`Hello\w+`)
	fmt.Println(reg.MatchString(s))
	// false
	reg = regexp.MustCompile(`Hello[\w\s]+`)
	fmt.Println(reg.MatchString(s))
	// true

	reg = regexp.MustCompile(`(?U)(?:Hello)(\s+)(\w+)`)
	fmt.Println(reg.NumSubexp())
	// 2

	b = []byte("Hello World, 123 Go!")
	reg = regexp.MustCompile(`(Hell|G)o`)
	rep := []byte("${1}ooo")
	fmt.Printf("%q\n", reg.ReplaceAll(b, rep))
	// "Hellooo World, 123 Gooo!"

	s = "Hello World, 123 Go!"
	reg = regexp.MustCompile(`(Hell|G)o`)
	repString := "${1}ooo"
	fmt.Printf("%q\n", reg.ReplaceAllString(s, repString))
	// "Hellooo World, 123 Gooo!"

	b = []byte("Hello World, 123 Go!")
	reg = regexp.MustCompile(`(Hell|G)o`)
	rep = []byte("${1}ooo")
	fmt.Printf("%q\n", reg.ReplaceAllLiteral(b, rep))
	// "${1}ooo World, 123 ${1}ooo!"

	s = "Hello World, 123 Go!"
	reg = regexp.MustCompile(`(Hell|G)o`)
	repString = "${1}ooo"
	fmt.Printf("%q\n", reg.ReplaceAllLiteralString(s, repString))
	// "${1}ooo World, 123 ${1}ooo!"

	by := []byte("Hello World!")
	reg = regexp.MustCompile("(H)ello")
	rep = []byte("$0$1")
	fmt.Printf("%s\n", reg.ReplaceAll(by, rep))
	// HelloH World!

	fmt.Printf("%s\n", reg.ReplaceAllFunc(by,
		func(b []byte) []byte {
			rst := []byte{}
			rst = append(rst, b...)
			rst = append(rst, "$1"...)
			return rst
		}))
	// Hello$1 World!

	s = "Hello World!"
	reg = regexp.MustCompile("(H)ello")
	repString = "$0$1"
	fmt.Printf("%s\n", reg.ReplaceAllString(s, repString))
	// HelloH World!
	fmt.Printf("%s\n", reg.ReplaceAllStringFunc(s,
		func(b string) string {
			return b + "$1"
		}))
	// Hello$1 World!

	s = "Hello World\tHello\nGolang"
	reg = regexp.MustCompile(`\s`)
	fmt.Printf("%q\n", reg.Split(s, -1))
	// ["Hello" "World" "Hello" "Golang"]

	re := regexp.MustCompile("Hello.*$")
	fmt.Printf("%s\n", re.String())
	// Hello.*$

	//	re = regexp.MustCompile("(?PHello) (World)")
	//	fmt.Printf("%q\n", re.SubexpNames())
	// ["" "Name1" ""]

	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("bar.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("a(b", "seafood")
	fmt.Println(matched, err)

	matched, err = regexp.MatchString("^[ _0-9a-zA-Z\u4E00-\u9FFF]+$", "~你好 ")
	fmt.Println(matched, err)
}

func match() {
	//	string := "hello world!"
	//	matched, err = regexp.
}
