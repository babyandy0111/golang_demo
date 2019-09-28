package models

// Gorm 使用
type Member struct {
	MemberID int    `json:"member_id"`
	FbID     string `json:"fb_id"`
	Num      string `json:"num"`
	Account  string `json:"account"`
	Password string `json:"password"`
	RealName string `json:"real_name"`
	Phone    string `json:"phone"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Sex      string `json:"sex"`
	BirthY   string `json:"birth_y"`
	BirthM   string `json:"birth_m"`
	BirthD   string `json:"birth_d"`
	Area_1   string `json:"area_1"`
	Area_2   string `json:"area_2"`
	Ecode    string `json:"ecode"`
	Address  string `json:"address"`
	ZhiYe    string `json:"zhi_ye"`
	ErpFlag  string `json:"erp_flag"`
	OpenFlag string `json:"open_flag"`
	Score    string `json:"score"`
	Onlystr  string `json:"onlystr"`
	Udate    string `json:"udate"`
}

func GetMembers(pageNum int, pageSize int, maps interface{}) (members []Member) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&members)
	return
}

func GetMembersTotal(maps interface{}) (count int) {
	db.Model(&Member{}).Where(maps).Count(&count)
	return
}

func ExistMemberByName(account string) bool {
	var mamber Member
	db.Select("member_id").Where("account = ?", account).First(&mamber)
	if mamber.MemberID > 0 {
		return true
	}

	return false
}

func ExistMemberById(id string) bool {
	var mamber Member
	db.Select("member_id").Where("member_id = ?", id).First(&mamber)
	if mamber.MemberID > 0 {
		return true
	}

	return false
}

func AddMember(account string) bool {
	db.Create(&Member{
		Account:  account,
		BirthD:   "01",
		BirthY:   "1985",
		BirthM:   "01",
		ErpFlag:  "1",
		OpenFlag: "1",
		Score:    "123",
	})
	return true
}

func EditMember(id string, param interface{}) bool {
	db.Model(&Member{}).Where("member_id = ?", id).Updates(param)
	return true
}

func DeleteMemberByID(id string) bool {
	db.Where("member_id = ?", id).Delete(Member{})
	return true
}
