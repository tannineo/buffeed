package model

import "time"

// Feed 订阅关系
type Feed struct {
	ID      int64     `xorm:"pk autoincr 'id'"`
	Version int       `xorm:"version"`
	Deleted time.Time `xorm:"deleted"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`

	Hash string `xorm:"varchar(64) not null unique 'hash'"` // 上传时产生的hash值 作为暴露在外的key

	UserID   int64  `xorm:"'user_id' not null"`               // 创建用户id
	UserName string `xorm:"varchar(64) not null 'user_name'"` // 创建用户昵称

	Alias    string `xorm:"varchar(64) 'alias'"`               // Alias feed别称
	Link     string `xorm:"varchar(256) not null 'link'"`      // FeedURL feed地址
	JumpLink string `xorm:"varchar(256) not null 'jump_link'"` // URL 跳转地址

	LastFetch time.Time `xorm:"'last_fetch'"` // LastFetch 最后一次更新
}

// AllFeeds 获取所有sub条目
// 同样鉴于小规模 不分页了
func AllFeeds() (*[]Feed, error) {
	allFeeds := &[]Feed{}
	err := engine.Find(allFeeds)
	return allFeeds, err
}

// GetFeed 根据条件获取单条Sub
func (f *Feed) GetFeed() (has bool, err error) {
	has, err = engine.Get(f)
	return
}

// InsertIn 插入Feed
func (f *Feed) InsertIn() (int, error) {
	affected, err := engine.Insert(f)
	return int(affected), err
}

// DeleteByID 根据删除sub
// 看情况能不能重构一个统一的deleteById 方法
func (f *Feed) DeleteByID() (int, error) {
	affected, err := engine.Id(f.ID).Delete(f)
	return int(affected), err
}

// ModifyAliasAndLinkByID 根据id修改alias和feed_url
func (f *Feed) ModifyAliasAndLinkByID() (int, error) {
	affected, err := engine.Id(f.ID).Cols("alias", "link").Update(f)
	return int(affected), err
}
