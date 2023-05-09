<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
    用户名:<input type="text" name="username">
    年龄:<input type="text" name="age">
    密码:<input type="password" name="password">
    <input type="submit" value="登录">

    <select name="fruit">
        <option value="apple">apple</option>
        <option value="pear">pear</option>
        <option value="banane">banane</option>
    </select>

    <input type="checkbox" name="interest" value="football">足球
    <input type="checkbox" name="interest" value="basketball">篮球
    <input type="checkbox" name="interest" value="tennis">网球

    <input type="hidden" name="token" value="{{.}}">
    
</form>
</body>
</html>