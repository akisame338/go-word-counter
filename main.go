package main

import (
    "net/http"
    "github.com/ikawaha/kagome/tokenizer"
    "github.com/gin-gonic/gin"
)

func getWordCount(text string) (map[string]int) {
    dic := tokenizer.SysDicSimple()
    t := tokenizer.NewWithDic(dic)
    tokens := t.Tokenize(text)
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

    return words
}

func main() {
    router := gin.Default()
    router.LoadHTMLFiles("index.tpl")

    router.GET("/", func(context *gin.Context) {
        context.HTML(http.StatusOK, "index.tpl", gin.H{})
    })

    router.POST("/", func(context *gin.Context) {
        text := context.PostForm("text")
        context.HTML(http.StatusOK, "index.tpl", gin.H{
            "wordCount": getWordCount(text),
        })
    })

    router.Run()
}
