<script src="/js/ws.js"></script>

<div class="row" id="api-container">
	<div class="col-md-8 col-md-offset-2" id="api-json">
		<img src="/img/rocket.png" height="70" width="70" alt=""/>
		<span class="api-url"><a href="/random_tips/json">http://vim-tips.com/random_tips/json</a></span>
		<div class="pull-right api-counter">已被调用<span id="json-counter">{{.JsonCounter}}</span>次</div>
	</div>
	<div class="col-md-8 col-md-offset-2" id="api-txt">

		<img src="/img/rocket.png" height="70" width="70" alt=""/>
		<span class="api-url"><a href="/random_tips/txt">http://vim-tips.com/random_tips/txt</a></span>

		<div class="pull-right api-counter">已被调用<span id="txt-counter">{{.TxtCounter}}</span>次</div>
	</div>
</div>