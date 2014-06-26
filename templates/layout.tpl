<!DOCTYPE html>
<html lang="zh-cn">
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="description" content="">
  <meta name="author" content="">
  <!--   <link rel="shortcut icon" href="../../docs-assets/ico/favicon.png"> -->

  <title>Vim-tips.com</title>

  <!-- Bootstrap core CSS -->
<!--   <link href="../static/css/style.css" rel="stylesheet">
-->  <link href="/css/bootstrap.min.css" rel="stylesheet">

<script type="text/javascript" src="http://cdn.bootcss.com/jquery/1.10.2/jquery.min.js"></script>

<!-- Custom styles for this template -->
<!--   <link href="theme.css" rel="stylesheet"> -->
<style>
body{font-family: Helvetica, Tahoma, Arial, STXihei, "华文细黑", "Microsoft YaHei", "微软雅黑", SimSun, "宋体", Heiti, "黑体", sans-serif;}
h1, .h1, h2, .h2, h3, .h3, h4, .h4, .lead {font-family: Helvetica, Tahoma, Arial, STXihei, "华文细黑", "Microsoft YaHei", "微软雅黑", SimSun, "宋体", Heiti, "黑体", sans-serif;}
pre code { background: transparent; }
@media (min-width: 768px) {
  .bs-docs-home .bs-social, 
  .bs-docs-home .bs-masthead-links {
    margin-left: 0;
  }
}

.bs-docs-section p {
  line-height: 2;
}

.bs-docs-section p.lead {
  line-height: 1.4;
}

</style>
</head>

<body>
 <div class="navbar navbar-default" role="navigation">
  <div class="container">
    <div class="navbar-header">
      <a class="navbar-brand" href="/"><img src="/img/vim.png" /></a>
      <a class="navbar-brand" href="/"><h3>Daily tips for Vim</h3></a>
    </div>
     <div class="navbar-right">
      <ul class="nav navbar-nav">
        <li><a href="/">Home</a></li>
        <li class="active"><a href="/about">About</a></li>
        <li><a href="#"></a></li>
      </ul>
    </div><!--/.navbar-collapse -->
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
