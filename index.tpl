<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>go-word-counter</title>
</head>
<body>
    <ul>
        {{range $word, $count := .wordCount}}
            <li>{{$word}}：{{$count}}回</li>
        {{end}}
    </ul>
</body>
</html>
