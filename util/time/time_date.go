package time

import (
	"time"
)

func TimeNow() time.Time {
	now := time.Now()
	local, _ := time.LoadLocation("Asia/Taipei") //修改成台北時間
	return now.In(local)
}
