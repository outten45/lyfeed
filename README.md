# lyfeed

An experimental Go  web application.

# Building

Use [glide](https://github.com/Masterminds/glide) to build. Fetch dependencies
use:

    $ glide install
Use go commands to build:

    $ go build -v github.com/outten45/lyfeed/lyfeed/cmd/...
# Goals

* Pull in a feed (RSS/Atom) and display an a page.
* Be a self contained executable including storage (like sqlite).
* Caveat: HTML, CSS ad JavaScript can be external files.
* Allow users to comment on items in the feeds.

# Approach

event 
- UUID
- Date/Time
- Type
- Version
- Raw

  rss -> item -> events
              -> item (rolled up Type)
              
