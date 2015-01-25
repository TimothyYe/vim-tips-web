<div class="row" id="tools-container">
<div class="col-md-12">
<ul class="catCardList">
{{ if .Num }}
{{ range .Casts }}
<li class="catCardList">
<div class="catCard"><a href="{{ .Url }}" target="_blank"><img src="{{ .LogoUrl }}" alt=""></a>
<div class="lowerCatCard">
<h3>{{ .Title }}</h3>
<p>{{ .Intro }}</p>
<div id="catCardButton" class="button"><a href="{{ .Url }}" target="_blank">点击查看</a></div>
</div>
</div>
</li>
{{ end }}
{{ end }}
</ul>
</div>
</div>

	<div class="row">
				<div class="m-b-xl">
					{{if gt .Paginator.PageNums 1}}
					<ul class="pagination pagination-sm">
						{{if .Paginator.HasPrev}}
						<li><a href="{{.Paginator.PageLinkFirst}}">第一页</a></li>
						<li><a href="{{.Paginator.PageLinkPrev}}">&lt;</a></li>
						{{else}}
						<li class="disabled"><a>第一页</a></li>
						<li class="disabled"><a>&lt;</a></li>
						{{end}}
						{{range $index, $page := .Paginator.Pages}}
						<li{{if $.Paginator.IsActive .}} class="active"{{end}}>
						<a href="{{$.Paginator.PageLink $page}}">{{$page}}</a>
					</li>
					{{end}}
					{{if .Paginator.HasNext}}
					<li><a href="{{.Paginator.PageLinkNext}}">&gt;</a></li>
					<li><a href="{{.Paginator.PageLinkLast}}">最后一页</a></li>
					{{else}}
					<li class="disabled"><a>&gt;</a></li>
					<li class="disabled"><a>最后一页</a></li>
					{{end}}
				</ul>
				{{end}}
			</div>