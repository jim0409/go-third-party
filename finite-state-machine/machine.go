package main

import (
	"errors"
	"strings"
)

// 定義狀態機
type StateMachine struct {
	name          string
	version       string
	transitionMap map[string]*Transition
	callBacks     map[string]CallBack
	alias         map[string]string
}

func NewStateMachine(name, ver string) *StateMachine {
	return &StateMachine{
		name:          name,
		version:       ver,
		transitionMap: make(map[string]*Transition),
		callBacks:     make(map[string]CallBack),
		alias:         make(map[string]string),
	}
}

func (m *StateMachine) Name() string {
	return m.name
}

func (m *StateMachine) Version() string {
	return m.version
}

func (m *StateMachine) PutAlias(k, v string) {
	k = strings.ToUpper(k)
	m.alias[k] = v
}

func (m *StateMachine) GetAlias(k string) (v string, ok bool) {
	k = strings.ToUpper(k)
	v, ok = m.alias[k]
	return
}

func (m *StateMachine) PutCallBacks(name string, cb CallBack) {
	m.callBacks[name] = cb
}

func (m *StateMachine) AddTransition(tName string, startState, nextState State) *StateMachine {
	tName = strings.ToUpper(tName)
	smm := m.name + ":" + m.version
	startState.stateMachineName = smm
	nextState.stateMachineName = smm
	t := NewTransition(tName, startState, nextState, m)
	m.transitionMap[tName] = t
	return m
}

func (m *StateMachine) PutTransition(t *Transition) *StateMachine {
	if t.stateMachine != m {
		t.stateMachine = m
	}
	smm := m.name + ":" + m.version
	t.startState.stateMachineName = smm
	t.nextState.stateMachineName = smm
	m.transitionMap[t.Name()] = t
	return m
}

// 透過 Name 獲取 transition
func (m *StateMachine) GetTransitionByName(tName string) (*Transition, error) {
	tName = strings.ToUpper(tName)
	if t, ok := m.transitionMap[tName]; ok {
		return t, nil
	}
	return nil, errors.New("can't find transition by name" + tName)
}

// 透過 State 獲取 transition
func (m *StateMachine) GetTransitionByState(state State) []*Transition {
	ts := make([]*Transition, 0)
	for _, v := range m.transitionMap {
		if v.startState.Equals(state) {
			ts = append(ts, v)
		}
	}
	return ts
}
