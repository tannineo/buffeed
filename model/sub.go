package model

import "time"

// Sub 订阅关系
type Sub struct {
	Hash string `xorm:"varchar(64) not null unique 'hash'"` // 上传时产生的hash值 作为暴露在外的key

	UserID   int64  `xorm:"'user_id' not null"`               // 创建用户id
	UserName string `xorm:"varchar(64) not null 'user_name'"` // 创建用户昵称

	Alias   string `xorm:"varchar(64) 'alias'"`              // Alias feed别称
	FeedURL string `xorm:"varchar(128) not null 'feed_url'"` // FeedURL feed地址
	URL     string `xorm:"varchar(128) not null 'url'"`      // URL 跳转地址

	LastFetch time.Time `xorm:"'last_fetch'"` // LastFetch 最后一次更新

	BasicMeta `xorm:"extends"`
}

// AllSubs 获取所有sub条目
// 同样鉴于小规模 不分页了
func AllSubs() (*[]Sub, error) {
	allSubs := &[]Sub{}
	err := engine.Find(allSubs)
	return allSubs, err
}

// GetSub 根据条件获取单条Sub
func (s *Sub) GetSub() (has bool, err error) {
	has, err = engine.Get(s)
	return
}

// InsertIn 插入sub
func (s *Sub) InsertIn() (int, error) {
	affected, err := engine.Insert(s)
	return int(affected), err
}

// DeleteByID 根据删除sub
// 看情况能不能重构一个统一的deleteById 方法
func (s *Sub) DeleteByID() (int, error) {
	affected, err := engine.Id(s.ID).Delete(s)
	return int(affected), err
}

// ModifyAliasAndFeedURLByID 根据id修改alias和feed_url
func (s *Sub) ModifyAliasAndFeedURLByID() (int, error) {
	affected, err := engine.Id(s.ID).Cols("alias", "feed_url").Update(s)
	return int(affected), err
}
