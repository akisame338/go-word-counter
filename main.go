package main

import (
    "fmt"
    "strings"
    "github.com/ikawaha/kagome/tokenizer"
)

func main() {
    dic := tokenizer.SysDicSimple()
    t := tokenizer.NewWithDic(dic)
    tokens := t.Tokenize("すもももももももものうち")
    for _, token := range tokens {
        features := strings.Join(token.Features(), ",")
        fmt.Printf("%s\t%v\n", token.Surface, features)
    }
}
