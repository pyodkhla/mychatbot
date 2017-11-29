package main

import (
	"github.com/nlopes/slack"
	"fmt"
)

func main() {
	api := slack.New("CE+CUQ05UkV2Tf4560LKEPfeIC04eV1ntv94mmXzMfgXvhpFSfiw10vfSYgnExyVR1+7WyRyybSooDDigEp5ujLwKwpqzFkuxOuoSFXIbVfLpVBV5mZnR+uzRHpyrDL9V7SuzC4JRMJ2/AnQ5EBBLgdB04t89/1O/w1cDnyilFU=")
	user, err := api.GetUserInfo("U2eb7de42e18db7d4ce3ced698eb291ae")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}
