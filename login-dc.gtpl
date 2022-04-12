<html>
<head>
    <title></title>
</head>
<body>
<form action="/login" method="post">
    <input type="checkbox" name="interest" value="football">Football
    <input type="checkbox" name="interest" value="basketball">Basketball
    <input type="checkbox" name="interest" value="tennis">Tennis
    Username:<input type="text" name="username">
    Password:<input type="password" name="password">
    <!-- 每次都会生成一个新的token, We use an MD5 hash (time stamp) to generate the token-->
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="Login">
</form>
</body>
</html>