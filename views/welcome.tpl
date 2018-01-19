<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>欢迎页面-X-admin2.0</title>
        <meta name="renderer" content="webkit">
        <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
        <meta name="viewport" content="width=device-width,user-scalable=yes, minimum-scale=0.4, initial-scale=0.8,target-densitydpi=low-dpi" />
        <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />
        <link rel="stylesheet" href="/static/css/font.css">
        <link rel="stylesheet" href="/static/css/sysx.css">
        <script type="text/javascript" src="https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js"></script>
    </head>
    <body>
        <div class="x-body">
            <blockquote class="layui-elem-quote">欢迎使用sys-X 后台模版！</blockquote>
            <fieldset class="layui-elem-field">
              <legend>信息统计</legend>
              <div class="layui-field-box">
                <table class="layui-table" lay-even>
                    <thead>
                        <tr>
                            <th>统计</th>
                            <th>资讯库</th>
                            <th>图片库</th>
                            <th>产品库</th>
                            <th>用户</th>
                            <th>管理员</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>总数</td>
                            <td>92</td>
                            <td>9</td>
                            <td>0</td>
                            <td>8</td>
                            <td>20</td>
                        </tr>
                        <tr>
                            <td>今日</td>
                            <td>0</td>
                            <td>0</td>
                            <td>0</td>
                            <td>0</td>
                            <td>0</td>
                        </tr>
                        <tr>
                            <td>昨日</td>
                            <td>0</td>
                            <td>0</td>
                            <td>0</td>
                            <td>0</td>
                            <td>0</td>
                        </tr>
                        <tr>
                            <td>本周</td>
                            <td>2</td>
                            <td>0</td>
                            <td>0</td>
                            <td>0</td>
                            <td>0</td>
                        </tr>
                        <tr>
                            <td>本月</td>
                            <td>2</td>
                            <td>0</td>
                            <td>0</td>
                            <td>0</td>
                            <td>0</td>
                        </tr>
                    </tbody>
                </table>
                <table class="layui-table">
                <thead>
                    <tr>
                        <th colspan="2" scope="col">服务器信息</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <th width="30%">服务器计算机名</th>
                        <td><span id="ser_name"></span></td>
                    </tr>
                    <tr>
                        <td>服务器IP地址</td>
                        <td><span id="ser_ip"></span></td>
                    </tr>
                    <tr>
                        <td>服务器端口 </td>
                        <td><span id="ser_port"></span></td>
                    </tr>
                    <tr>
                        <td>服务器IIS版本 </td>
                        <td><span id="ser_version"></span></td>
                    </tr>
                    <tr>
                        <td>本文件所在文件夹 </td>
                        <td><span id="ser_path"></span></td>
                    </tr>
                    <tr>
                        <td>服务器操作系统 </td>
                        <td><span id="ser_os"></span></td>
                    </tr>
                    <tr>
                        <td>服务器当前时间 </td>
                        <td><span id="ser_time"></span></td>
                    </tr>
                    <tr>
                        <td>CPU 总数 </td>
                        <td><span id="ser_cpu"></span></td>
                    </tr>
                    <tr>
                        <td>当前线程数量 </td>
                        <td><span id="ser_thread_name"></span></td>
                    </tr>
                </tbody>
            </table>
              </div>
            </fieldset>
        </div>
    </body>
    <script type="text/javascript">
        $(function () {
            $.post("/user/sysinfo",function (data) {
                if(data){
                   json = $.parseJSON(data);
                    $("#ser_name").html(json.pcname)
                    $("#ser_ip").html(json.ip)
                    $("#ser_port").html(json.port)
                    $("#ser_version").html(json.version)
                    $("#ser_path").html(json.path)
                    $("#ser_os").html(json.os)
                    $("#ser_cpu").html(json.cpunum)
                    $("#ser_thread_name").html(json.threadnum)
                    $("#ser_time").html(json.time)
                }
            });
        })
    </script>
</html>