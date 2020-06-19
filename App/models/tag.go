package models

import "time"
type Tag struct {
    Model

    Name string `json:"name"`
    CreatedBy int64 `json:"created_by"`
    ModifiedBy int64 `json:"modified_by"`
    State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
    db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
    
    return
}

func GetTagTotal(maps interface {}) (count int){
    db.Model(&Tag{}).Where(maps).Count(&count)

    return
}

func ExistTagByName(name string) bool {
    var tag Tag
    db.Select("id").Where("name = ?", name).First(&tag)
    if tag.ID > 0 {
        return true
    }

    return false
}

func AddTag(name string, state int) bool{
    db.Create(&Tag {
        Name : name,
        State : state,
        CreatedBy : time.Now().Unix(),
    })

    return true
}