package main

type State struct {
	value            string
	stateMachineName string
}

func NewState(v string) State {
	return State{
		value: v,
	}
}

func (s *State) GetStateValue() string {
	return s.value
}

func (s *State) GetFullStateValue() string {
	return s.stateMachineName + ":" + s.value
}

func (s *State) SetStateValue(v string) {
	s.value = v
}

func (s *State) Equals(st State) bool {
	return s.value == st.GetStateValue()
}
