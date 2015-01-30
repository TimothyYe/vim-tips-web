<!DOCTYPE html>
<html>
<head>
  <meta content="Vim,Vimer,Vim技巧,Vim教程,Vim学习,Vim编辑器" name="keywords" />
  <meta content="Vim tips,Vimer的园地,来自Vim的每日技巧与点滴,让你每天都能学到来自Vim的奇技淫巧,让你的Vim更进一步..."name="description" />
  <title>VimTips - Daily tips for Vim!</title>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="description" content="">
  <meta name="author" content="Timothy">
  <!--   <link rel="shortcut icon" href="../../docs-assets/ico/favicon.png"> -->

  <title>Vim-tips.com</title>

  <!-- Bootstrap core CSS -->
<!--   <link href="http://fonts.useso.com/css?family=Lobster" rel="stylesheet" type="text/css"/> -->
  <link href="http://cdn.bootcss.com/bootstrap/3.2.0/css/bootstrap.min.css" rel="stylesheet">
  <link href="/css/admin.css" rel="stylesheet">
  <link href="/css/font-awesome-4.2.0/css/font-awesome.css" rel="stylesheet">
  <link rel="stylesheet" href="http://vimtips.qiniudn.com/css/app.css" type="text/css" />

  <script src="http://cdn.bootcss.com/jquery/2.1.1/jquery.min.js"></script>
  <script src="http://cdn.bootcss.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>

  <!-- Custom styles for this template -->
  <!--   <link href="theme.css" rel="stylesheet"> -->
  <style>
  body {
    font-family: "Microsoft YaHei", "微软雅黑", SimSun, "宋体", Heiti, "黑体", sans-serif;
  }

  h1, .h1, h2, .h2, h3, .h3, h4, .h4, .lead {
    font-family: "Microsoft YaHei", "微软雅黑", SimSun, "宋体", Heiti, "黑体", sans-serif;}
  }
  </style>
</head>

<body style="background-color: #FFFFFF;">
 <div class="navbar navbar-default" role="navigation">
  <div class="container">
    <div class="navbar-header">
      <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
        <span class="sr-only">Toggle navigation</span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a class="navbar-brand" href="/">VimTips</a>
    </div>
    <div class="collapse navbar-collapse">
      <ul class="nav navbar-nav">
        <li class="{{ if .IsIndex }}active{{ end }}"><a href="/admin/index">管理首页</a></li>
        <li class="{{ if .IsTips }}active{{ end }}"><a href="/admin/tips">Tips管理</a></li>
        <li class="{{ if .IsCasts }}active{{ end }}"><a href="/admin/casts">播客管理</a></li>
        <li class="{{ if .IsPassword }}active{{ end }}"><a href="/admin/password">密码修改</a></li>
      </ul>

      <form class="navbar-form navbar-right" role="form" method="get" action="/admin/logout">
        <button type="submit" class="btn btn-success">退出后台管理</button>
      </form>
    </div><!--/.navbar-collapse -->
  </div>
</div>

<div id="wrapper">
  <div id="content" class="container">
    {{ yield }}
  </div>

  <div id="footer" class="navbar navbar-fixed-bottom">
    <div id="copyright" class="container">
      <div class="row">
        <div class="col-md-12 col-xs-12">
				<div class="col-md-6 col-xs-6">Powered by Golang &amp; <a href="http://martini.codegangsta.io/" target="_blank" >Martini</a> | <a href="/admin/login">管理入口</a></div>
        <div class="col-md-6 col-xs-6">
					<div class="pull-right">© 2014 <a href="http://github.com/timothyye"> timothyye </a> All Rights Reserved.</div>
				</div>
				</div>
      </div>
    </div>
  </div>
</div>

</body>
</html>
