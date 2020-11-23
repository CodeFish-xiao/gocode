package main

import (
	"encoding/json"
	"fmt"
)

var MemberCartVo struct {
	Id                int      `json:"id"`              //id
	BelongCoachId     int64    `json:"belong_coach_id"` //属于哪个专家
	CartType          int      `json:"cart_type"`       //卡包类型
	City              string   `json:"city"`            //卡包城市
	EffectiveTime     int      `json:"effective_time"`
	GiftCourseId      int      `json:"gift_course_id"`
	GiftRotationPic   string   `json:"gift_rotation_pic"`
	Introduce         string   `json:"introduce"`
	IsDelete          int      `json:"is_delete"`
	Name              string   `json:"name"`
	Price             float32  `json:"price"`
	RealPrice         float32  `json:"real_price"`
	Weight            int      `json:"weight"`
	Cover             []string `json:"cover"`
	GiftIntroducePic  []string `json:"gift_introduce_pic"`
	RotationIntroduce []string `json:"rotation_introduce"`
}

func main() {
	reqv, _ := json.Marshal(MemberCartVo)
	fmt.Println(string(reqv))
}
