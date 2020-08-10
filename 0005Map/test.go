package main

import "fmt"

func main() {
	// The map is store by hash table, so there is no return order in map
	/* 声明变量，默认 map 是 nil */
	//var map_variable map[key_data_type]value_data_type

	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	// fmt.Println(capital)
	// fmt.Println(ok)
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	//delete a element in the map
	fmt.Println("Testing delete")
	delete(countryCapitalMap, "India")
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

}
