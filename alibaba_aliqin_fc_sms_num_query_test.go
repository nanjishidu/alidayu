package alidayu

import (
	"fmt"
	"testing"
	"time"
)

func TestAlibabaAliqinFcSmsNumQuery(t *testing.T) {
	resp, err := m.AlibabaAliqinFcSmsNumQuery(rec_num, time.Now(), 1, 50)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.AlibabaAliqinFcSmsNumQueryResponePageInfo)
}
