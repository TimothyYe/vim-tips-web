<link href="http://vjs.zencdn.net/4.7/video-js.css" rel="stylesheet">
<script src="http://vjs.zencdn.net/4.7/video.js"></script>

<div class="row" id="api-container">
	<div class="col-md-8 col-md-offset-2" id="api-json">
		<video class="video-js vjs-default-skin" autoplay controls preload="auto" width="860" height="483" data-setup="{}">
			<source src="http://media.happycasts.net/assets/episodes/videos/097-togetherjs.mov" type='video/mp4'>
		</video>
	</div>

		<div class="col-md-8 col-md-offset-2" id="disqus">
			<div id="disqus_thread"></div> 
			<script type="text/javascript">
			var disqus_shortname = 'vim-tips'; 
			var disqus_identifier = '{{.Id}}';
			var disqus_url = 'http://vim-tips.com/casts/{{.Id}}';

			(function() {
				var dsq = document.createElement('script'); dsq.type = 'text/javascript'; dsq.async = true;
				dsq.src = '//' + disqus_shortname + '.disqus.com/embed.js';
				(document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(dsq);
			})();
			</script>

			<noscript>Please enable JavaScript to view the <a href="http://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
			<a href="http://disqus.com" class="dsq-brlink">comments powered by <span class="logo-disqus">Disqus</span></a>

		</div>
		<br/>
		<br/>
	</div>