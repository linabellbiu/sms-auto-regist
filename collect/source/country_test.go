package source

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseCountryCode(t *testing.T) {
	err := json.Unmarshal([]byte(countryCodeData), &CountryCodeData)
	if err != nil {
		fmt.Errorf("解析手机区号失败:%v", err)
		return
	}
	t.Log(CountryCodeData)
	for _, v := range CountryCodeData {
		t.Log(v.Code)
	}

}
