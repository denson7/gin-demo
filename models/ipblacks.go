package models

import (
	"fmt"
	"gin-demo/utils"
	"time"
)

type IpBlack struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	IP       string    `json:"ip"`
	UserId   string    `json:"user_id"`
	UserName string    `json:"user_name"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`
}

func CreateIpBlack(ip string, userId string) uint {
	black := &IpBlack{
		IP:       ip,
		UserId:   userId,
		CreateAt: time.Now(),
	}
	utils.Db.Table("ip_blacks").Create(black)
	return black.ID
}

func UpdateIpBlack(ip string, userId string, id int) uint  {
	info := &IpBlack{
		IP: ip,
		UserId: userId,
	}
	utils.Db.Table("ip_blacks").Model(info).Where("id = ?", id).Update(info)
	fmt.Println(info)
	return info.ID
}

func FindIp(ip string) IpBlack {
	var ipBlack IpBlack
	utils.Db.Table("ip_blacks").Select("ip_blacks.*, b.username, b.email").Joins("left join users b on ip_blacks.user_id = b.id").Where("ip_blacks.ip = ?", ip).First(&ipBlack)
	//utils.Db.Table("ip_blacks").Where("ip = ?", ip).First(&ipBlack)
	return ipBlack
}

func FindIps(query interface{}, args []interface{}, page uint, pageSize uint) []IpBlack {
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}
	var ipBlacks []IpBlack
	if query != nil {
		utils.Db.Table("ip_blacks").Where(query, args...).Offset(offset).Limit(pageSize).Find(&ipBlacks)
	} else {
		utils.Db.Table("ip_blacks").Offset(offset).Limit(pageSize).Find(&ipBlacks)
	}
	return ipBlacks
}

func DeleteIpBlackByIp(ip string) {
	utils.Db.Table("ip_blacks").Where("ip = ?", ip).Delete(IpBlack{})
}
