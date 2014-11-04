<div class="row">
	<div>
		<button class="btn btn-success m-b-xs" data-toggle="modal" data-target="#addTip">新增</button>

		<div class="modal fade" id="addTip" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
			<div class="modal-dialog">
				<form role="form" method="post" action="/admin/tips/add">
					<div class="modal-content">
						<div class="modal-header">
							<button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
							<h4 class="modal-title" id="myModalLabel">新增Tip</h4>
						</div>
						<div class="modal-body">
							<div class="input-group">
								<span class="input-group-addon">Tip内容:</span>
								<input type="text" class="form-control input-lg text-center" id="tip" name="tip">
							</div>
							<div class="input-group">
								<span class="input-group-addon">Tip说明:</span>
								<input type="text" class="form-control input-lg text-center" id="tip_comment" name="comment">
							</div>
						</div>
						<div class="modal-footer">
							<button type="submit" class="btn btn-danger" >添加</button>
							<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						</div>
					</div><!-- /.modal-content -->
				</form>
			</div><!-- /.modal-dialog -->
		</div><!-- /.modal -->
		
		<div class="container">
			<div class="row">
				<table class="table table-striped col-md-12">
					<thead>
						<tr>
							<th>Tip</th>
							<th>Comments</th>
							<th>Operations</th>
						</tr>
					</thead>
					<tbody>
						{{ if .Num }}
						{{ range .Tips }}
						<tr>
							<td>{{ .Content }}</td>
							<td>{{ .Comment }}</td>

							<td style="width:120px">
								<div>
									<span class="pull-left">
										<form role="form" method="get" action="/admin/tips/modify">
											<input type="hidden" value="{{ .Id }}" name="Id" />
											<button type="submit" class="btn btn-success btn-sm">修改</button>
										</form>
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
											<form class="col-md-3" role="form" method="post" action="/admin/tips/del">
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
