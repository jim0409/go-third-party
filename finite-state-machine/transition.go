package main

import (
	"errors"
	"fmt"
	"strings"
)

type CallBack interface {
	BeforeRunCallBack(runableState State, data interface{}, metaData map[string]interface{}) error
	RunEffectCallBack(runableState State, data interface{}, metaData map[string]interface{}) error
	AfterRunCallBack(runableState State, data interface{}, metaData map[string]interface{}) error
}

type callBackBlank struct{}

func (*callBackBlank) BeforeRunCallBack(s State, data interface{}, metaData map[string]interface{}) error {
	return nil
}

func (*callBackBlank) RunEffectCallBack(s State, data interface{}, metaData map[string]interface{}) error {
	return errors.New("callBackBlank is default, please reset it!")
}

func (*callBackBlank) AfterRunCallBack(s State, data interface{}, metaData map[string]interface{}) error {
	return nil
}

// 定義過濾器
type Transition struct {
	startState   State
	nextState    State
	runableState State
	name         string
	callBack     string
	stateMachine *StateMachine
	inputs       []interface{}
	metaData     map[string]interface{}
}

func NewTransition(transitionName string, startState, nextState State, sm *StateMachine) *Transition {
	transitionName = strings.ToUpper(transitionName)
	return &Transition{
		name:         transitionName,
		startState:   startState,
		nextState:    nextState,
		stateMachine: sm,
		inputs:       make([]interface{}, 0),
		metaData:     make(map[string]interface{}),
	}
}

func (t *Transition) Name() string {
	return t.name
}

func (t *Transition) AppendInput(input interface{}) {
	t.inputs = append(t.inputs, input)
}

func (t *Transition) GetInputs() []interface{} {
	return t.inputs
}

func (t *Transition) PutMetaData(key string, data interface{}) {
	t.metaData[key] = data
}

func (t *Transition) GetMetaDataByKey(key string) (interface{}, bool) {
	data, ok := t.metaData[key]
	return data, ok
}

func (t *Transition) SetCallBackName(name string) {
	t.callBack = name
}

func (t *Transition) Execute(data interface{}) (State, error) {
	var callBack CallBack
	if c, ok := t.stateMachine.callBacks[t.callBack]; ok {
		callBack = c
	} else {
		callBack = &callBackBlank{}
	}
	t.runableState = t.startState
	fmt.Println("do before run call back")
	if err := callBack.BeforeRunCallBack(t.runableState, data, t.metaData); err != nil {
		return t.runableState, err
	}
	fmt.Println("do effect call back")
	if err := callBack.RunEffectCallBack(t.runableState, data, t.metaData); err != nil {
		fmt.Println("err: ", err)
		return t.runableState, err
	}
	fmt.Println("do after call back")
	t.runableState = t.nextState
	if err := callBack.AfterRunCallBack(t.runableState, data, t.metaData); err != nil {
		return t.runableState, err
	}
	return t.runableState, nil
}
