package dao

import (
	"time"

	"gorm.io/gorm"
)

type VideoInfo struct {
	// gorm段
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	// 数据段
	RVID        int64  `gorm:"column:rvid;primaryKey;autoIncrement" json:"rvid"`
	FaceUrl     string `gorm:"column:face_url;size:512" json:"face_url"`
	M3u8Url     string `gorm:"column:m3u8_url;size:512" json:"m3u8_url"`
	Mp4TempUrl  string `gorm:"column:mp4_temp_url;size:512" json:"mp4_temp_url"`
	Title       string `gorm:"column:title;size:255" json:"title"`
	Description string `gorm:"column:description;type:text" json:"description"`
	ViewNum     int64  `gorm:"column:view_num;default:0" json:"view_num"`
	TimeLength  string `gorm:"column:time_length;size:20" json:"time_length"`
	// 属性段
	UseFace bool   `gorm:"column:use_face;default:false" json:"use_face"`
	InJudge bool   `gorm:"column:in_judge;default:true" json:"in_judge"`
	Status  string `gorm:"column:status;size:20;default:'waiting'" json:"status"`
}
