#!/usr/bin/env sh

# @link http://stackoverflow.com/questions/19439914/how-to-convert-directory-sass-scss-to-css-via-command-line
# gem install sass compass sass-globbing

sass --update $1/scss:$1/css
