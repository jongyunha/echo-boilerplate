package domain

import "gorm.io/gorm"

type HelloRepositoryDb struct {
	Client *gorm.DB
}

func (d HelloRepositoryDb) FindByName(name string) (*Hello, error) {
	var hello Hello
	if err := d.Client.Debug().Raw("select * from helloTable where name = ?;", name).Scan(&hello).Error; err != nil {
		return nil, err
	}
	return &hello, nil
}

func NewHelloRepositoryDb(client *gorm.DB) HelloRepository {
	return HelloRepositoryDb{Client: client}
}
