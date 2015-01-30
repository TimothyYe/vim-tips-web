vim-tips-web
============

[![Build Status](https://drone.io/github.com/TimothyYe/vim-tips-web/status.png)](https://drone.io/github.com/TimothyYe/vim-tips-web/latest)

[Vim-Tips.com](https://vim-tips.com) is a site to share tips for Vim. 

The first version of [Vim-Tips.com](https://vim-tips.com) is written by Rails.

Now I re-designed this whole site and implement it by Go.

## Prerequisite

* Go environment is needed. 
* MongoDB is installed.

## Build it

* Get source code from Github:

```bash
git clone https://github.com/TimothyYe/vim-tips-web.git
```
* Go into the source code directory, get related library and then build it:

```bash
cd vim-tips-web
go get
go build
```

## Run it

* Run vim-tips-web site:
```bash
nohup ./vim-tips-web &
```
