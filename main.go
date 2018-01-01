package main

import (
    "fmt"
    "github.com/ikawaha/kagome/tokenizer"
)

func main() {
    dic := tokenizer.SysDicSimple()
    t := tokenizer.NewWithDic(dic)
    tokens := t.Tokenize("すもももももももものうち")
    words := map[string]int{}
    for _, token := range tokens {
        features := token.Features()
        // 名詞以外ならスキップ
        if len(features) == 0 || features[0] != "名詞" {
            continue
        }

        // マップに既にキーが存在する場合はインクリメント、存在しない場合は1で初期化
        _, ok := words[token.Surface]
        if ok {
            words[token.Surface]++
        } else {
            words[token.Surface] = 1
        }
    }

    for word, count := range words {
        fmt.Printf("%s\t%d\n", word, count)
    }
}
