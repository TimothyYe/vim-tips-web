<div class="row">
	<div>

		<form role="form" method="get" action="/admin/tips/add">
			<button type="submit" class="btn btn-success">新增</button>
		</form>

		<br/>
		<div class="container">
			<div class="row">
				<table class="table table-striped col-md-12">
					<thead>
						<tr>
							<th>Tip</th>
							<th>Comments</th>
						</tr>
					</thead>
					<tbody>
						{{ if .Num }}
						{{ range .Tips }}
						<tr>
							<td>{{ .Content }}</td>
							<td>{{ .Comment }}</td>

							<td>
								<form role="form" method="get" action="/admin/tips/modify">
									<input type="hidden" value="{{ .Id }}" name="Id" />
									<button type="submit" class="btn btn-success btn-sm">修改</button>
								</form>

								<button class="btn btn-danger btn-sm" data-toggle="modal" data-target="#myModal{{ .Id }}">删除</button>
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
											<form class="col-md-3" role="form" method="post" action="/admin/order/del">
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
	<hr>
</div>
</div>