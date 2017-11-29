# Buffeed
[![Build Status](https://travis-ci.org/tannineo/buffeed.svg?branch=master)](https://travis-ci.org/tannineo/buffeed)
[![codecov](https://codecov.io/gh/tannineo/buffeed/branch/master/graph/badge.svg)](https://codecov.io/gh/tannineo/buffeed)  
Enjoy your self-hosted feed buffet with Buffeed - a feed reader in go.  
The project is under construction...  

1. Sharing - Feed can be added and seen by every user in Buffeed.
2. Managing - You can collect and tag your favorite items.

# Using...
- [echo](https://github.com/labstack/echo) as the framework.
- [xorm](https://github.com/go-xorm/xorm) with [sqlite3](https://github.com/mattn/go-sqlite3) for data storage.
- [configor](https://github.com/jinzhu/configor) for a [config file](https://github.com/tannineo/buffeed/blob/master/config.json).
- [goconvey](https://github.com/smartystreets/goconvey) for tests.

## Notice
1. Project dependencies are managed with [dep](https://github.com/golang/dep)
2. Notice: configure the path of sqlite3 db file, it's written in [config file](https://github.com/tannineo/buffeed/blob/master/config.json).
