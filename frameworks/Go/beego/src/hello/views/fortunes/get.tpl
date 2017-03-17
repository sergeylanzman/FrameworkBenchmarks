<!DOCTYPE html>
<html>
<head><title>Fortunes</title></head>
<body>
<table>
<tr><th>id</th><th>message</th></tr>
{{range $f:=.fortunes}}<tr><td>{{$f.Id}}</td><td>{{$f.Message}}</td></tr>{{end}}
</table>
</body>
</html>