package main

import (
	"log"
)

type callbackextension struct{}

func (*callbackextension) BeforeRunCallBack(s State, data interface{}, metaData map[string]interface{}) error {
	m, ok := data.(map[string]interface{})
	if !ok {
		return nil
	}
	for k, v := range m {
		log.Printf("before : %v - %v\n", k, v)
	}
	return nil
}

func (*callbackextension) RunEffectCallBack(s State, data interface{}, metaData map[string]interface{}) error {
	return nil
}

func (*callbackextension) AfterRunCallBack(s State, data interface{}, metaData map[string]interface{}) error {
	return nil
}

// not_commit -> commit_pass -> review_pass
//					|-> commit_fail
func createCodingStateMachine() *StateMachine {
	st := NewStateMachine("start_coding", "1")
	nt_cmt := NewState("not_commit")
	nt_rev := NewState("not_review")
	nt_pas := NewState("review_pass")
	rv_fai := NewState("review_fail")
	st.AddTransition("coding", nt_cmt, nt_rev)
	st.AddTransition("commit_pass", nt_rev, nt_pas)
	st.AddTransition("review_pass", nt_rev, rv_fai)
	st.PutCallBacks("callback_ext", &callbackextension{})

	// log.Println(st.GetTransitionByName("coding"))

	return st
}

// 從 start_coding 開始 ..
func main() {
	// 1. 創建一個狀態機
	st := createCodingStateMachine()

	// 2. 提取該狀態機要使用的 transition
	cmt, err := st.GetTransitionByName("coding")

	// 3. 設定該狀態機的 callback 名稱
	cmt.SetCallBackName("callback_ext")

	if err != nil {
		panic(err)
	}
	result, _ := cmt.Execute(map[string]interface{}{
		"testKey1": "testValue1",
		"testKey2": "testValue2",
	})
	if result.GetStateValue() != "not_review" {
		panic("[FAIL] result : " + result.GetStateValue())
	}
	log.Println("[PASS] result : " + result.GetStateValue())

}
