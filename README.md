# lyfeed

An experimental Go  web application that is using the
[gb](http://getgb.io/) project.

# Goals

* Pull in a feed (RSS/Atom) and display an a page.
* Be a self contained executable including storage (like ledisdb).
* Caveat: HTML, CSS ad JavaScript can be external files.
* Allow users to comment on items in the feeds.

# goimports

To use goimports within your project, set your GOPATH:

export GOPATH=$(pwd):$(pwd)/vendor
