// 天喜	天虚	天哭	天德	红鸾	龙池	凤阁	孤辰	寡宿	破碎	大耗	华盖	解神	咸池	劫杀	天马	蜚廉
// "擎羊星","陀罗星","天钺星","天魁星","禄存星","天福星","天官星","正空亡","副空亡"
package ziwei

import (
	"strings"
)

//年干系诸星
func GetYearGanStarMap(ygz string) map[string]string {
	return findYearGanStar(ygz)
}
func findYearGanStar(ygz string) map[string]string {
	starNameArr := []string{"擎羊星", "陀罗星", "天钺星", "天魁星", "禄存星", "天福星", "天官星", "正空亡", "副空亡"}
	ygmap := map[string][]string{
		"甲": {"卯", "丑", "未", "丑", "寅", "酉", "未", "申", "酉"},
		"乙": {"辰", "寅", "申", "子", "卯", "申", "辰", "未", "午"},
		"丙": {"午", "辰", "酉", "亥", "巳", "子", "巳", "辰", "巳"},
		"丁": {"未", "巳", "酉", "亥", "午", "亥", "寅", "卯", "寅"},
		"戊": {"午", "辰", "未", "丑", "巳", "卯", "卯", "子", "丑"},
		"己": {"未", "巳", "申", "子", "午", "寅", "酉", "酉", "申"},
		"庚": {"酉", "未", "未", "丑", "申", "午", "亥", "午", "未"},
		"辛": {"戌", "申", "寅", "午", "酉", "巳", "酉", "巳", "辰"},
		"壬": {"子", "戌", "巳", "卯", "亥", "午", "戌", "寅", "卯"},
		"癸": {"丑", "亥", "巳", "卯", "子", "巳", "午", "丑", "子"},
	}
	gans := GetGanS(ygz)
	var xmap = make(map[string]string) //k星名　v地支宫位
	for k, zhiArr := range ygmap {
		if strings.EqualFold(k, gans) {
			for i := 0; i < len(starNameArr); i++ {
				for j := 0; j < len(zhiArr); j++ {
					if i == j {
						xmap[starNameArr[j]] = zhiArr[j]
						break
					}
				}
			}
			break
		}
	}
	return xmap
}

//###################################################################

//年支系诸星　k星名　v地支宫位
func GetYearZhiStarMap(ygz string) map[string]string {
	return findYearZhiStar(ygz)
}
func findYearZhiStar(ygz string) map[string]string {
	starNameArr := []string{"天喜", "天虚", "天哭", "天德", "红鸾", "龙池", "凤阁", "孤辰", "寡宿",
		"破碎", "大耗", "华盖", "解神", "咸池", "劫杀", "天马", "蜚廉"} //年支诸星
	yzmap := map[string][]string{
		"子": {"酉", "午", "午", "酉", "卯", "辰", "戌", "寅", "戌", "巳", "未", "辰", "戌", "酉", "巳", "寅", "申"},
		"丑": {"申", "未", "巳", "戌", "寅", "巳", "酉", "寅", "戌", "丑", "午", "丑", "酉", "午", "寅", "亥", "酉"},
		"寅": {"未", "申", "辰", "亥", "丑", "午", "申", "巳", "丑", "酉", "酉", "戌", "申", "卯", "亥", "申", "戌"},
		"卯": {"午", "酉", "卯", "子", "子", "未", "未", "巳", "丑", "巳", "申", "未", "未", "子", "申", "巳", "巳"},
		"辰": {"巳", "戌", "寅", "丑", "亥", "申", "午", "巳", "丑", "丑", "亥", "辰", "午", "酉", "巳", "寅", "午"},
		"巳": {"辰", "亥", "丑", "寅", "戌", "酉", "巳", "申", "辰", "酉", "戌", "丑", "巳", "午", "寅", "亥", "未"},
		"午": {"卯", "子", "子", "卯", "酉", "戌", "辰", "申", "辰", "巳", "丑", "戌", "辰", "卯", "亥", "申", "寅"},
		"未": {"寅", "丑", "亥", "辰", "申", "亥", "卯", "申", "辰", "丑", "子", "未", "卯", "子", "申", "巳", "卯"},
		"申": {"丑", "寅", "戌", "巳", "未", "子", "寅", "亥", "未", "酉", "卯", "辰", "寅", "酉", "巳", "寅", "辰"},
		"酉": {"子", "卯", "酉", "午", "午", "丑", "丑", "亥", "未", "巳", "寅", "丑", "丑", "午", "寅", "亥", "亥"},
		"戌": {"亥", "辰", "申", "未", "巳", "寅", "子", "亥", "未", "丑", "巳", "戌", "子", "卯", "亥", "申", "子"},
		"亥": {"戌", "巳", "未", "申", "辰", "卯", "亥", "寅", "戌", "酉", "辰", "未", "亥", "子", "申", "巳", "丑"},
	}
	zhis := GetZhiS(ygz)
	var xmap = make(map[string]string) //k星名　v 地支
	for k, zhiArr := range yzmap {
		if strings.EqualFold(k, zhis) {
			for i := 0; i < len(starNameArr); i++ {
				for j := 0; j < len(zhiArr); j++ {
					if i == j {
						xmap[starNameArr[j]] = zhiArr[j]
						break
					}
				}
			}
			break
		}
	}
	//fmt.Printf("年支系诸星:%s年　%v\n", ygz, xmap)
	return xmap
}

//=========================================
//返回　擎羊星　擎羊星地支宫位　　陀罗星　陀罗星地支宫位
func qingYangTuoLuo(ygz string) (string, string, string, string) {
	garr := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	qingYang := []string{"卯", "辰", "午", "未", "午", "未", "酉", "戌", "子", "丑"} //擎羊星　0甲年地支
	tuoLuo := []string{"丑", "寅", "辰", "巳", "辰", "巳", "未", "申", "戌", "亥"}   //陀罗星 0甲年地支
	gs := GetGanS(ygz)
	var s1, s2 string
	for i := 0; i < len(garr); i++ {
		if strings.EqualFold(gs, garr[i]) {
			s1 = qingYang[i]
			s2 = tuoLuo[i]
			break
		}
	}
	return "擎羊星", s1, "陀罗星", s2
}
