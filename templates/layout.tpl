<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="description" content="">
  <meta name="author" content="">
  <!--   <link rel="shortcut icon" href="../../docs-assets/ico/favicon.png"> -->

  <title>Vim-tips.com</title>

  <!-- Bootstrap core CSS -->
  <link href="/css/style.css" rel="stylesheet">
  <link href="/css/bootstrap.min.css" rel="stylesheet">

  <script type="text/javascript" src="http://cdn.bootcss.com/jquery/1.10.2/jquery.min.js"></script>

<!-- Custom styles for this template -->
<!--   <link href="theme.css" rel="stylesheet"> -->
<style>
body{font-family: Helvetica, Tahoma, Arial, STXihei, "华文细黑", "Microsoft YaHei", "微软雅黑", SimSun, "宋体", Heiti, "黑体", sans-serif;}
h1, .h1, h2, .h2, h3, .h3, h4, .h4, .lead {font-family: Helvetica, Tahoma, Arial, STXihei, "华文细黑", "Microsoft YaHei", "微软雅黑", SimSun, "宋体", Heiti, "黑体", sans-serif;}
pre code { background: transparent; }
}

</style>
</head>

<body>
 <div id="top_banner" class="navbar navbar-default" role="navigation">
  <div class="container">
    <div class="navbar-header">
      <a href="/"><img src="/img/vim.png" /></a>
    </div>
    <div>
      <h1>Vim-Tips.com</h1>
      <small>Vim的技巧点滴与分享</small>
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

<div class="container">
  {{ yield }}
</div>

<footer>
  <div class="container">
  </div>
</footer>

</div> <!-- /container -->
</body>
</html>
