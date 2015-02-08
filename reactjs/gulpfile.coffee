gulp = require 'gulp'

browserify = require 'browserify'
source = require 'vinyl-source-stream'

gulp.task 'js', ->
  browserify
    entries: ['./public/scripts/main.js']
    extensions: ['.coffee'] # CoffeeScriptも使えるように
  .bundle()
  .pipe source 'main.js' # 出力ファイル名を指定
  .pipe gulp.dest "./public/js/" # 出力ディレクトリを指定