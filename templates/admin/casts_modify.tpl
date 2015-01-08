  <div class="row m-t-xl">
  
  <div class="col-md-8 col-md-offset-2" >
    <div class="m-t-l m-b-xl">
      <h3>更新播客</h3>
    </div>

    {{ with .Cast }}
    <form class="form-horizontal" role="form" method="post" action="/admin/casts">
      <div class="form-group">
        <div class="col-md-12">
          <label for="author">作者</label>
          <input type="text" class="form-control" id="author" name="author" placeholder="填入作者" value="{{ .Author }}">
        </div>
      </div>

      <div class="form-group">
        <div class="col-md-12">
          <label for="authorurl">作者链接</label>
          <input type="text" class="form-control" id="authorurl" name="authorurl" placeholder="填入作者链接" value="{{ .AuthorUrl }}">
        </div>
      </div>

      <div class="form-group">
        <div class="col-md-12">
          <label for="title">标题</label>
          <input type="text" class="form-control" id="title" name="title" placeholder="填入标题" value="{{ .Title }}">
        </div>
      </div>

      <div class="form-group">
        <div class="col-md-12">
          <label for="intro">简介</label>
          <input type="text" class="form-control" id="intro" name="intro" placeholder="填入简介" value="{{ .Intro }}">
        </div>
      </div>

      <div class="form-group">
        <div class="col-md-12">
          <label for="logourl">Logo链接</label>
          <input type="text" class="form-control" id="logourl" name="logourl" placeholder="填入Logo URL" value="{{ .LogoUrl }}">
        </div>
      </div>

      <div class="form-group">
        <div class="col-md-12">
          <label for="url">播客链接</label>
          <input type="text" class="form-control" id="url" name="url" placeholder="填入播客视频URL" value="{{ .Url }}">
        </div>
      </div>

      <div class="form-group">
        <div class="col-md-12">
          <label for="shownotes">相关资料</label>
          <textarea class="form-control" id="shownotes" name="shownotes" rows="6">{{ .ShowNotes }}</textarea>
        </div>
      </div>

      {{ end }}

      <button type="submit" class="btn btn-success">更新</button>
    </form>
  </div>
</div>

