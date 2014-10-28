<div class="row">
	<div class="col-md-8 col-md-offset-2 col-xs-10 col-xs-offset-1 round" id="tip-content">
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
</div>
