<div class="row">
	<div>
		<a href="/admin/casts/add"><button class="btn btn-success m-b-xs">新增</button></a>
		
		<div class="container">
			<div class="row">
				<table class="table table-striped col-md-12">
					<thead>
						<tr>
							<th>Title</th>
							<th>Author</th>
							<th>Author Url</th>
						</tr>
					</thead>
					<tbody>
						{{ if .Num }}
						{{ range .Casts }}
						<tr>
							<td>{{ .Title }}</td>
							<td>{{ .Author }}</td>
							<td>{{ .AuthorUrl }}</td>

							<td style="width:120px">
								<div>
									<span class="pull-left">
										<a href="/admin/casts/modify/{{ .Id }}"><button class="btn btn-success btn-sm">修改</button></a>
									</span>

									<span class="pull-right">
										<button class="btn btn-danger btn-sm" data-toggle="modal" data-target="#myModal{{ .Id }}">删除</button>
									</span>
								</div>
							</td>

							<div class="modal fade" id="myModal{{ .Id }}" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
								<div class="modal-dialog">
									<div class="modal-content">
										<div class="modal-header">
											<button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
											<h4 class="modal-title" id="myModalLabel">删除确认</h4>
										</div>
										<div class="modal-body">
											请确认是否删除此条数据？
										</div>
										<div class="modal-footer">
											<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
											<form class="col-md-3" role="form" method="post" action="/admin/casts/del">
												<input type="hidden" value="{{ .Id }}" name="Id" />
												<button type="submit" class="btn btn-danger" >确认删除</button>
											</form>
										</div>
									</div><!-- /.modal-content -->
								</div><!-- /.modal-dialog -->
							</div><!-- /.modal -->
						</tr>
						{{ end }}
						{{ end }}
					</tbody>
				</table>
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
		</div>
	</div>
</div>
</div>
</div>
