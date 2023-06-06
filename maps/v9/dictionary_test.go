package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
)

func TestSyncMapLog(t *testing.T) {
	syncMap := sync.Map{}
	syncMap.Store("apple", 1)
	syncMap.Store("banana", 2)
	syncMap.Store("cherry", 3)

	// 序列化为JSON格式
	jsonStr, err := json.Marshal(syncMap)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return
	}

	// 打印输出JSON字符串
	fmt.Println(string(jsonStr))
}

func TestSyncMapLog2(t *testing.T) {
	allPromotionsMap := make(map[string]*sync.Map, 0)
	syncMap := sync.Map{}
	syncMap.Store("apple", 1)
	syncMap.Store("banana", 2)
	syncMap.Store("cherry", 3)
	allPromotionsMap["test"] = &syncMap

	// 序列化为JSON格式
	jsonStr, err := json.Marshal(allPromotionsMap)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return
	}
	// 打印输出JSON字符串
	fmt.Println("allPromotionsMap:" + string(jsonStr))

	// 序列化为JSON格式
	jsonStr2, err := json.Marshal(syncMap)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return
	}
	// 打印输出JSON字符串
	fmt.Println("syncMap:" + string(jsonStr2))

	// 正常打印syncMap
	m := map[string]interface{}{}
	syncMap.Range(func(key, value interface{}) bool {
		m[fmt.Sprint(key)] = value
		return true
	})
	b, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	for k, v := range allPromotionsMap {

		m := map[string]interface{}{}
		v.Range(func(key, value interface{}) bool {
			m[fmt.Sprint(key)] = value
			return true
		})
		b, err := json.MarshalIndent(m, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println("k:" + k)

		fmt.Println(" v:" + string(b))

	}

}
