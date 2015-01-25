<link href="http://vjs.zencdn.net/4.7/video-js.css" rel="stylesheet">
<script src="http://vjs.zencdn.net/4.7/video.js"></script>

<div class="row" id="cast-container">
	{{ with .ViewCast }}
	<div class="col-md-8 col-md-offset-2" id="api-json">
		<video class="video-js vjs-default-skin" controls preload="auto" width="750" height="483" data-setup="{}">
			<source src="{{ .Url }}" type='video/mp4'>
		</video>
	</div>

		<div class="col-md-8 col-md-offset-2" id="comments">
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
		<br/>
		<br/>
		{{ end }}
</div>