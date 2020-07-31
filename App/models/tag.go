package models

type Tag struct {
    Model

    Name string `json:"name"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
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

func ExistTagByMaps(maps interface{}) bool {
    var tag Tag
    db.Select("id").Where(maps).First(&tag)
    if tag.ID > 0 {
        return true
    }

    return false
}

func AddTag(tags map[string]interface{}) bool{
    tag := Tag {
        Name : tags["Name"].(string),
        State : tags["State"].(int),
        CreatedBy : "白翀华",
    }
    db.Create(&tag)
    return !db.NewRecord(tag)
}

func ExistTagByID(id int) bool {
    var tag Tag
    db.Select("id").Where("id = ?", id).First(&tag)
    if tag.ID > 0 {
        return true
    }

    return false
}

func DeleteTag(maps interface{}) bool {
    db.Where(maps).Delete(&Tag{})

    return true
}

func EditTag(id int, data interface {}) bool {
    db.Model(&Tag{}).Where("id = ?", id).Updates(data)

    return true
}