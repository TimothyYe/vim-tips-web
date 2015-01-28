<div class="row">
	<div class="col-md-8 col-md-offset-2 col-xs-10 col-xs-offset-1 round" id="tip-content">
		<img src="/img/lightbulb.png" class="light-bulb hidden-xs" alt=""/>
		<a href="http://vim-tips.com/tips/{{.Id}}">

			<div class="tip-text">

				<div class="content">
					{{.Content}}
				</div>

				<div class="comment">
					{{.Comment}}
				</div>

				<div id="tip_id">
					{{.Id}}
				</div>
			</div>
		</a>
	</div>

	<div class="col-md-8 col-md-offset-2" id="comments">
		<div class="ds-thread" data-thread-key="{{.Id}}" data-title="" data-url="http://vim-tips.com/tips/{{.Id}}"></div>

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
</div>

