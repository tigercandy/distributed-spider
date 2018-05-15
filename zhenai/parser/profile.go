package parser

import (
	"distributed-spider/engine"
	"regexp"
	"strconv"
	"distributed-spider/model"
)

var (
	ageRe        = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
	heightRe     = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
	weightRe     = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
	incomeRe     = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	marriageRe   = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	genderRe     = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	xinzuoRe     = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
	educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
	hokouRe      = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	houseRe      = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe        = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	// 姓名
	profile.Name = name
	// 年龄
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	// 身高
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = height
	}
	// 体重
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}
	// 月收入
	profile.Income = extractString(contents, incomeRe)
	// 婚况
	profile.Marriage = extractString(contents, marriageRe)
	// 性别
	profile.Gender = extractString(contents, genderRe)
	// 星座
	profile.Xinzuo = extractString(contents, xinzuoRe)
	// 学历
	profile.Education = extractString(contents, educationRe)
	// 职业
	profile.Occupation = extractString(contents, occupationRe)
	// 籍贯
	profile.Hokou = extractString(contents, hokouRe)
	// 住房条件
	profile.House = extractString(contents, houseRe)
	// 是否购车
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}
