package model

import (
	"time"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3" // go-sqlite3 注册sqlite3到sql
)

// Engine 数据库连接engine
var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "data.db") // 优化db path
	if err != nil {
		// TODO: logging
		panic(err)
	}
}

// BasicMeta 基础元信息
// 创建 更新 删除 乐观锁
type BasicMeta struct {
	Version int       `xorm:"version"`
	Deleted time.Time `xorm:"deleted"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

// InitDB 表结构init
func InitDB() error {
	sql := "begin;" + initScript + "commit;"
	_, err := engine.Exec(sql)
	return err
}

var initScript = `
DROP TABLE IF EXISTS 'user';
DROP TABLE IF EXISTS 'sub';
DROP TABLE IF EXISTS 'brief';
DROP TABLE IF EXISTS 'tag';

CREATE TABLE IF NOT EXISTS 'user' (
  'id' integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  'salt' text NOT NULL DEFAULT '',
  'pwd' text NOT NULL,
  'name' text NOT NULL,
  'email' text NOT NULL,
  'access' text NOT NULL,
  'version' integer NOT NULL DEFAULT 1,
  'deleted' datetime,
  'created' datetime NOT NULL,
  'updated' datetime NOT NULL
);

CREATE TABLE IF NOT EXISTS 'sub' (
  'id' integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  'user_id' integer NOT NULL,
  'alias' text NOT NULL,
  'feed_url' text NOT NULL,
  'url' text NOT NULL,
  'last_fetch' datetime,
  'version' integer NOT NULL DEFAULT 1,
  'deleted' datetime,
  'created' datetime NOT NULL,
  'updated' datetime NOT NULL
);

CREATE TABLE IF NOT EXISTS 'brief' (
  'id' integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  'sub_id' integer NOT NULL,
  'title' text NOT NULL,
  'article' text NOT NULL,
  'version' integer NOT NULL DEFAULT 1,
  'deleted' datetime,
  'created' datetime NOT NULL,
  'updated' datetime NOT NULL
);

CREATE TABLE IF NOT EXISTS 'tag' (
  'id' integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  'user_id' integer NOT NULL,
  'sub_id' integer NOT NULL,
  'brief_id' integer NOT NULL,
  'tag_name' text NOT NULL DEFAULT '',
  'version' integer NOT NULL DEFAULT 1,
  'deleted' datetime,
  'created' datetime NOT NULL,
  'updated' datetime NOT NULL
);
`
