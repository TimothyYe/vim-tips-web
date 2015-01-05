  <div class="row m-t-xl">

<div class="col-md-4 col-md-offset-4" >
  {{ if .IsPost}}
  {{ if .IsSuccess}}
  <div class="alert alert-success" role="alert">密码更新成功</div>
  {{ else }}
  <div class="alert alert-warning" role="alert">密码更新失败</div>
  {{ end }}
  {{ end}}
</div>
  
  <div class="col-md-4 col-md-offset-4" > 
    <form class="form-horizontal" role="form" method="post" action="/admin/password">
      <div class="form-group form-group-lg">
        
        <div class="col-md-12">
          <input type="password" class="form-control input-lg text-center" id="inputEmail3" name="password" placeholder="输入新密码">
        </div>
      </div>

      <div class="form-group">
        <div class="col-md-12">
          <input type="password" class="form-control input-lg text-center" id="inputPassword3" name="verify" placeholder="确认新密码">
        </div>
      </div>

      <div class="form-group">
        <div class="col-md-12 col-xs-12">
          <button type="submit" class="btn btn-primary btn-lg col-md-12 col-xs-12">更新密码</button>
        </div>
      </div>
    </form>
  </div>
</div>

