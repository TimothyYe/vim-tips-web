<div class="row">
	<div class="col-md-8 col-md-offset-2 round" id="tip-content">
		<img src="/img/lightbulb.png" class="light-bulb" alt=""/>
		<a href="http://localhost:3000/tips/{{.Id}}">

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

	<div class="col-md-8 col-md-offset-2" id="disqus">
		<div id="disqus_thread"></div> 
		<script type="text/javascript">
        var disqus_shortname = 'vim-tips'; 
	    var disqus_identifier = '{{.Id}}';
	    var disqus_url = 'http://vim-tips.com/tips/{{.Id}}';

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

