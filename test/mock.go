package mocks

type MockTool[C any] struct {
	mockResult       map[string][]C
	mockError        map[string]error
	mockCalledTime   map[string]int
	mockCalledParams map[string][]interface{}
}

func (p *MockTool[C]) SetMockValue(funcName string, data ...C) {
	if p.mockResult == nil {
		p.mockResult = make(map[string][]C)
	}

	p.mockResult[funcName] = data
}

func (p *MockTool[C]) SetMockError(funcName string, err error) {
	if p.mockError == nil {
		p.mockError = make(map[string]error)
	}

	p.mockError[funcName] = err
}

func (p *MockTool[C]) GetMockValue(funcName string) []C {
	return p.mockResult[funcName]
}

func (p *MockTool[C]) GetMockError(funcName string) error {
	return p.mockError[funcName]
}

func (p *MockTool[C]) InitMock() {
	p.mockResult = make(map[string][]C)
	p.mockError = make(map[string]error)
	p.mockCalledTime = make(map[string]int)
	p.mockCalledParams = make(map[string][]interface{})
}

func (p *MockTool[C]) ToBeCalledTime(funcName string) int {
	return p.mockCalledTime[funcName]
}

func (p *MockTool[C]) ToBeCalledWith(funcName string) []interface{} {
	return p.mockCalledParams[funcName]
}

func (p *MockTool[C]) SetMockFunction(funcName string, data ...interface{}) {
	if _, ok := p.mockCalledTime[funcName]; ok {
		p.mockCalledTime[funcName] += 1
	} else {
		p.mockCalledTime[funcName] = 1
	}

	if _, ok := p.mockCalledParams[funcName]; ok {
		p.mockCalledParams[funcName] = append(p.mockCalledParams[funcName], data)
	} else {
		p.mockCalledParams[funcName] = []interface{}{
			data,
		}
	}
}
