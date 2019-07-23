package main

import (
	"fmt"
	"github.com/huichen/sego"
	"github.com/pathvar"
)

const (
	defaultDicPath = "${GOPATH}/src/github.com/huichen/sego/data/dictionary.txt"
)

func main() {
	// 载入词典
	var segmenter sego.Segmenter
	dicpath := pathvar.Subst(defaultDicPath)
	segmenter.LoadDictionary(dicpath)

	// 分词
	text := []byte("哈哈，真是傻")
	segments := segmenter.Segment(text)

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	a := sego.SegmentsToSlice(segments, false)
	fmt.Println(a)
	for _, v := range a {
		fmt.Println(v)
	}

}
