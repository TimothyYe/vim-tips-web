	<link href="http://vjs.zencdn.net/4.7/video-js.css" rel="stylesheet">
	<script src="http://vjs.zencdn.net/4.7/video.js"></script>

	<div class="row" id="cast-container">
		{{ with .ViewCast }}
		<div class="col-md-8 col-md-offset-2" id="api-json">

			<div class="m-t m-b"><h4>{{ .Title }}     [主播: <a href="{{ .AuthorUrl }}" target="_blank">{{.Author}}</a>]</h4></div>

			<video class="video-js vjs-default-skin" controls preload="auto" width="750" height="483" data-setup="{}">
				<source src="{{ .Url }}" type='video/mp4'>
				</video>

				<ul id="mytab" class="nav nav-tabs m-t">
					<li class="active"><a href="#intro" data-toggle="tab">介绍</a></li>
					<li><a href="#resource" data-toggle="tab">相关资源</a></li>
				</ul>

				<div class="tab-content" id="show-notes">
					<div class="tab-pane active" id="intro">
						{{ .Intro }}		
					</div>
					<div class="tab-pane" id="resource">
						{{ .ShowNotes }}
					</div>
				</div>

				<div id="comments">
					<div class="ds-thread" data-thread-key="{{.Id}}" data-title="" data-url="http://vim-tips.com/casts/{{.Id}}"></div>

					<script type="text/javascript">
					var duoshuoQuery = {short_name:"vim-tips"};
					(function() {
						var ds = document.createElement('script');
						ds.type = 'text/javascript';ds.async = true;
						ds.src = (document.location.protocol == 'https:' ? 'https:' : 'http:') + '//static.duoshuo.com/embed.js';
						ds.charset = 'UTF-8';
						(document.getElementsByTagName('head')[0] 
							|| document.getElementsByTagName('body')[0]).appendChild(ds);
					})();
					</script>
				</div>

				{{ end }}
			</div>
