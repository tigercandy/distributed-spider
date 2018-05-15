package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("city_list_data.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", contents)
	result := ParseCityList(contents)

	// test data len
	const resultSize = 470
	// test url
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	// test city
	expectedCities := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d"+" request, but had %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url is #%d: %s; but was %s ", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d"+" request, but had %d",
			resultSize, len(result.Items))
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city is #%d: %s; but was %s ", i, city, result.Items[i].(string))
		}
	}
}
