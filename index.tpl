<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>go-word-counter</title>
</head>
<body>
    <form action="/" method="post">
        <div>
            <textarea name="text" placeholder="テキストを入力してください"></textarea>
        </div>
        <div>
            <button type="submit">送信</button>
        </div>
    </form>
    <hr>
    <ul>
        {{range $word, $count := .wordCount}}
            <li>{{$word}}：{{$count}}回</li>
        {{end}}
    </ul>
</body>
</html>
