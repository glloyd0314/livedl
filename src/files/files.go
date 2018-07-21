package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"regexp"
)

func ChangeExtention(fileName, ext string) string {
	e := filepath.Ext(fileName)
	base := strings.TrimSuffix(fileName, e)
	return base + "." + ext
}

func MkdirByFileName(fileName string) {
	dir := filepath.Dir(fileName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func GetFileNameNext(name string) (fileName string, err error) {
	fileName = name
	_, test := os.Stat(fileName)
	if test == nil {
		// file Exists
		ext := filepath.Ext(fileName)
		base := strings.TrimSuffix(fileName, ext)

		var i int
		for i = 2; i < 10000000 ; i++ {
			fileName = fmt.Sprintf("%s-%d%s", base, i, ext)
			_, test := os.Stat(fileName)
			if test != nil {
				return
			}
		}
		err = fmt.Errorf("too many files: %s", name)
	}
	return
}

func ReplaceForbidden(name string) (fileName string) {
	fileName = name
	fileName = regexp.MustCompile(`\\`).ReplaceAllString(fileName, "￥")
	fileName = regexp.MustCompile(`/`).ReplaceAllString(fileName, "∕")
	fileName = regexp.MustCompile(`:`).ReplaceAllString(fileName, "：")
	fileName = regexp.MustCompile(`\*`).ReplaceAllString(fileName, "＊")
	fileName = regexp.MustCompile(`\?`).ReplaceAllString(fileName, "？")
	fileName = regexp.MustCompile(`"`).ReplaceAllString(fileName, `゛`)
	fileName = regexp.MustCompile(`<`).ReplaceAllString(fileName, "＜")
	fileName = regexp.MustCompile(`>`).ReplaceAllString(fileName, "＞")
	fileName = regexp.MustCompile(`\|`).ReplaceAllString(fileName, "｜")

	fileName = regexp.MustCompile(`）`).ReplaceAllString(fileName, ")")
	fileName = regexp.MustCompile(`（`).ReplaceAllString(fileName, "(")

	fileName = regexp.MustCompile(`\p{Zs}+`).ReplaceAllString(fileName, " ")
	fileName = regexp.MustCompile(`\A\p{Zs}+|\p{Zs}+\z`).ReplaceAllString(fileName, "")
	return
}