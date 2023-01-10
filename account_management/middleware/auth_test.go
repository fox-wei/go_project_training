package middleware

import (
	"testing"
)

func Test_parseToken(t *testing.T) {
	t.Run("parse", func(t *testing.T) {
		claim, err := parseToken2("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NzMxMDA1ODcsImlkIjoiMyIsIm5iZiI6MTY3MzEwMDU4NywidXNlcm5hbWUiOiJ2YW4td2VpIn0.EJo4kJ6trnYtHUXP42Z8uPuGUdAYfhferfNvwbX3Cpk")
		if err != nil {
			t.Errorf("error:%v\n", err)
			return
		}

		// if claim.Username == "" {
		// 	t.Errorf("claim:%v\n", claim)
		// 	return
		// }

		want := "van-wei"
		got := claim.Username
		if want != got {
			t.Errorf("got:%s want:%s\n", got, want)
		}
	})

	t.Run("parse2", func(t *testing.T) {
		claim, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NzMxMDA1ODcsImlkIjoiMyIsIm5iZiI6MTY3MzEwMDU4NywidXNlcm5hbWUiOiJ2YW4td2VpIn0.EJo4kJ6trnYtHUXP42Z8uPuGUdAYfhferfNvwbX3Cpk")
		if err != nil {
			t.Errorf("error:%v\n", err)
			return
		}

		want := "3"
		got := claim["id"]

		if want != got {
			t.Errorf("got:%s, want:%s\n", got, want)
		}
	})
}
