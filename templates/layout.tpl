<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="description" content="">
  <meta name="author" content="Timothy">
  <!--   <link rel="shortcut icon" href="../../docs-assets/ico/favicon.png"> -->

  <title>Vim-tips.com</title>

  <!-- Bootstrap core CSS -->
  <link href="http://cdn.bootcss.com/bootstrap/3.2.0/css/bootstrap.min.css" rel="stylesheet">
  <link href="/css/style.css" rel="stylesheet">

  <script src="http://cdn.bootcss.com/jquery/2.1.1/jquery.min.js"></script>
  <script src="/js/tips.js"></script>

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

<body style="background-color: #ECE5CE;">
 <div id="top_banner" class="navbar navbar-default" role="navigation">
  <div class="container">
    <div class="navbar-header">
      <a href="/"><img src="/img/vim.png" /></a>
    </div>
    <div>
      <div id="title">
        <h1>Vim-Tips.com</h1>
        <small>Vim的技巧点滴与分享</small>
      </div>
    </div>
    <div class="collapse navbar-collapse">
      <ul class="nav navbar-nav navbar-right">
        <li class="active"><a href="/">Home</a></li>
        <li><a href="/tools">Tools</a></li>
        <li><a href="/api">API</a></li>
        <li><a href="/about">About</a></li>
      </ul>
    </div>
  </div>
</div>

<div id="wrapper">
  <div id="content" class="container">
    {{ yield }}
  </div>

  <div id="footer" class="navbar navbar-fixed-bottom">
    <div id="copyright" class="container">
      <div class="row">
        <div class="col-md-12">
				<div class="col-md-6"><a href="http://vim-tips.com">Vim-Tips.com</a>, powered by Golang &amp; Martini.</div>
        <div class="col-md-6">
					<div class="pull-right">© 2014 <a href="http://github.com/timothyye"> timothyye </a> All Rights Reserved.</div>
				</div>
				</div>
      </div>
    </div>
  </div>
</div>

</body>
</html>
