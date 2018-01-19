<!doctype html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>后台登录-sys-X</title>
	<meta name="renderer" content="webkit|ie-comp|ie-stand">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width,user-scalable=yes, minimum-scale=0.4, initial-scale=0.8,target-densitydpi=low-dpi" />
    <meta http-equiv="Cache-Control" content="no-siteapp" />

    <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />
    <link rel="stylesheet" href="/static/css/font.css">
	<link rel="stylesheet" href="/static/css/sysx.css">
    <script type="text/javascript" src="https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js"></script>
    <script src="/static/lib/layui/layui.js" charset="utf-8"></script>
    <script type="text/javascript" src="/static/js/sysx.js"></script>

</head>
<body class="login-bg">
    
    <div class="login">
        <div class="message">sys-X-管理系统登录</div>
        <div id="darkbannerwrap"></div>
        
        <form class="layui-form" id="login_form" method="post">
            <input name="username" placeholder="用户名"  type="text" lay-verify="required" class="layui-input" >
            <hr class="hr15">
            <input name="password" lay-verify="required" placeholder="密码"  type="password" class="layui-input">
            <hr class="hr15">
            <input value="登录" lay-submit lay-filter="login" style="width:100%;" type="submit" onclick="subLogin()">
            <hr class="hr20" >
        </form>
    </div>

    <script>
        //回车事件
        $(document).keydown(function(event){
            if(event.keyCode==13){
                subLogin();
            }
        });

        function subLogin() {
            $('#login_form').attr('action','/user/login');
            $('#login_form').submit();
        }
    </script>
    <!-- 底部结束 -->
</body>
</html>