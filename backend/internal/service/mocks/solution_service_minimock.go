// Code generated by http://github.com/gojuno/minimock (v3.3.10). DO NOT EDIT.

package mocks

//go:generate minimock -i gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service.SolutionService -o solution_service_minimock.go -n SolutionServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
)

// SolutionServiceMock implements service.SolutionService
type SolutionServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcDeploy          func(ctx context.Context) (err error)
	inspectFuncDeploy   func(ctx context.Context)
	afterDeployCounter  uint64
	beforeDeployCounter uint64
	DeployMock          mSolutionServiceMockDeploy

	funcList          func(ctx context.Context) (spa1 []*model.Solution, err error)
	inspectFuncList   func(ctx context.Context)
	afterListCounter  uint64
	beforeListCounter uint64
	ListMock          mSolutionServiceMockList

	funcStatus          func(ctx context.Context) (err error)
	inspectFuncStatus   func(ctx context.Context)
	afterStatusCounter  uint64
	beforeStatusCounter uint64
	StatusMock          mSolutionServiceMockStatus
}

// NewSolutionServiceMock returns a mock for service.SolutionService
func NewSolutionServiceMock(t minimock.Tester) *SolutionServiceMock {
	m := &SolutionServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeployMock = mSolutionServiceMockDeploy{mock: m}
	m.DeployMock.callArgs = []*SolutionServiceMockDeployParams{}

	m.ListMock = mSolutionServiceMockList{mock: m}
	m.ListMock.callArgs = []*SolutionServiceMockListParams{}

	m.StatusMock = mSolutionServiceMockStatus{mock: m}
	m.StatusMock.callArgs = []*SolutionServiceMockStatusParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mSolutionServiceMockDeploy struct {
	optional           bool
	mock               *SolutionServiceMock
	defaultExpectation *SolutionServiceMockDeployExpectation
	expectations       []*SolutionServiceMockDeployExpectation

	callArgs []*SolutionServiceMockDeployParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// SolutionServiceMockDeployExpectation specifies expectation struct of the SolutionService.Deploy
type SolutionServiceMockDeployExpectation struct {
	mock      *SolutionServiceMock
	params    *SolutionServiceMockDeployParams
	paramPtrs *SolutionServiceMockDeployParamPtrs
	results   *SolutionServiceMockDeployResults
	Counter   uint64
}

// SolutionServiceMockDeployParams contains parameters of the SolutionService.Deploy
type SolutionServiceMockDeployParams struct {
	ctx context.Context
}

// SolutionServiceMockDeployParamPtrs contains pointers to parameters of the SolutionService.Deploy
type SolutionServiceMockDeployParamPtrs struct {
	ctx *context.Context
}

// SolutionServiceMockDeployResults contains results of the SolutionService.Deploy
type SolutionServiceMockDeployResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option by default unless you really need it, as it helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmDeploy *mSolutionServiceMockDeploy) Optional() *mSolutionServiceMockDeploy {
	mmDeploy.optional = true
	return mmDeploy
}

// Expect sets up expected params for SolutionService.Deploy
func (mmDeploy *mSolutionServiceMockDeploy) Expect(ctx context.Context) *mSolutionServiceMockDeploy {
	if mmDeploy.mock.funcDeploy != nil {
		mmDeploy.mock.t.Fatalf("SolutionServiceMock.Deploy mock is already set by Set")
	}

	if mmDeploy.defaultExpectation == nil {
		mmDeploy.defaultExpectation = &SolutionServiceMockDeployExpectation{}
	}

	if mmDeploy.defaultExpectation.paramPtrs != nil {
		mmDeploy.mock.t.Fatalf("SolutionServiceMock.Deploy mock is already set by ExpectParams functions")
	}

	mmDeploy.defaultExpectation.params = &SolutionServiceMockDeployParams{ctx}
	for _, e := range mmDeploy.expectations {
		if minimock.Equal(e.params, mmDeploy.defaultExpectation.params) {
			mmDeploy.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeploy.defaultExpectation.params)
		}
	}

	return mmDeploy
}

// ExpectCtxParam1 sets up expected param ctx for SolutionService.Deploy
func (mmDeploy *mSolutionServiceMockDeploy) ExpectCtxParam1(ctx context.Context) *mSolutionServiceMockDeploy {
	if mmDeploy.mock.funcDeploy != nil {
		mmDeploy.mock.t.Fatalf("SolutionServiceMock.Deploy mock is already set by Set")
	}

	if mmDeploy.defaultExpectation == nil {
		mmDeploy.defaultExpectation = &SolutionServiceMockDeployExpectation{}
	}

	if mmDeploy.defaultExpectation.params != nil {
		mmDeploy.mock.t.Fatalf("SolutionServiceMock.Deploy mock is already set by Expect")
	}

	if mmDeploy.defaultExpectation.paramPtrs == nil {
		mmDeploy.defaultExpectation.paramPtrs = &SolutionServiceMockDeployParamPtrs{}
	}
	mmDeploy.defaultExpectation.paramPtrs.ctx = &ctx

	return mmDeploy
}

// Inspect accepts an inspector function that has same arguments as the SolutionService.Deploy
func (mmDeploy *mSolutionServiceMockDeploy) Inspect(f func(ctx context.Context)) *mSolutionServiceMockDeploy {
	if mmDeploy.mock.inspectFuncDeploy != nil {
		mmDeploy.mock.t.Fatalf("Inspect function is already set for SolutionServiceMock.Deploy")
	}

	mmDeploy.mock.inspectFuncDeploy = f

	return mmDeploy
}

// Return sets up results that will be returned by SolutionService.Deploy
func (mmDeploy *mSolutionServiceMockDeploy) Return(err error) *SolutionServiceMock {
	if mmDeploy.mock.funcDeploy != nil {
		mmDeploy.mock.t.Fatalf("SolutionServiceMock.Deploy mock is already set by Set")
	}

	if mmDeploy.defaultExpectation == nil {
		mmDeploy.defaultExpectation = &SolutionServiceMockDeployExpectation{mock: mmDeploy.mock}
	}
	mmDeploy.defaultExpectation.results = &SolutionServiceMockDeployResults{err}
	return mmDeploy.mock
}

// Set uses given function f to mock the SolutionService.Deploy method
func (mmDeploy *mSolutionServiceMockDeploy) Set(f func(ctx context.Context) (err error)) *SolutionServiceMock {
	if mmDeploy.defaultExpectation != nil {
		mmDeploy.mock.t.Fatalf("Default expectation is already set for the SolutionService.Deploy method")
	}

	if len(mmDeploy.expectations) > 0 {
		mmDeploy.mock.t.Fatalf("Some expectations are already set for the SolutionService.Deploy method")
	}

	mmDeploy.mock.funcDeploy = f
	return mmDeploy.mock
}

// When sets expectation for the SolutionService.Deploy which will trigger the result defined by the following
// Then helper
func (mmDeploy *mSolutionServiceMockDeploy) When(ctx context.Context) *SolutionServiceMockDeployExpectation {
	if mmDeploy.mock.funcDeploy != nil {
		mmDeploy.mock.t.Fatalf("SolutionServiceMock.Deploy mock is already set by Set")
	}

	expectation := &SolutionServiceMockDeployExpectation{
		mock:   mmDeploy.mock,
		params: &SolutionServiceMockDeployParams{ctx},
	}
	mmDeploy.expectations = append(mmDeploy.expectations, expectation)
	return expectation
}

// Then sets up SolutionService.Deploy return parameters for the expectation previously defined by the When method
func (e *SolutionServiceMockDeployExpectation) Then(err error) *SolutionServiceMock {
	e.results = &SolutionServiceMockDeployResults{err}
	return e.mock
}

// Times sets number of times SolutionService.Deploy should be invoked
func (mmDeploy *mSolutionServiceMockDeploy) Times(n uint64) *mSolutionServiceMockDeploy {
	if n == 0 {
		mmDeploy.mock.t.Fatalf("Times of SolutionServiceMock.Deploy mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmDeploy.expectedInvocations, n)
	return mmDeploy
}

func (mmDeploy *mSolutionServiceMockDeploy) invocationsDone() bool {
	if len(mmDeploy.expectations) == 0 && mmDeploy.defaultExpectation == nil && mmDeploy.mock.funcDeploy == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmDeploy.mock.afterDeployCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmDeploy.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Deploy implements service.SolutionService
func (mmDeploy *SolutionServiceMock) Deploy(ctx context.Context) (err error) {
	mm_atomic.AddUint64(&mmDeploy.beforeDeployCounter, 1)
	defer mm_atomic.AddUint64(&mmDeploy.afterDeployCounter, 1)

	if mmDeploy.inspectFuncDeploy != nil {
		mmDeploy.inspectFuncDeploy(ctx)
	}

	mm_params := SolutionServiceMockDeployParams{ctx}

	// Record call args
	mmDeploy.DeployMock.mutex.Lock()
	mmDeploy.DeployMock.callArgs = append(mmDeploy.DeployMock.callArgs, &mm_params)
	mmDeploy.DeployMock.mutex.Unlock()

	for _, e := range mmDeploy.DeployMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDeploy.DeployMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeploy.DeployMock.defaultExpectation.Counter, 1)
		mm_want := mmDeploy.DeployMock.defaultExpectation.params
		mm_want_ptrs := mmDeploy.DeployMock.defaultExpectation.paramPtrs

		mm_got := SolutionServiceMockDeployParams{ctx}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmDeploy.t.Errorf("SolutionServiceMock.Deploy got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeploy.t.Errorf("SolutionServiceMock.Deploy got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeploy.DeployMock.defaultExpectation.results
		if mm_results == nil {
			mmDeploy.t.Fatal("No results are set for the SolutionServiceMock.Deploy")
		}
		return (*mm_results).err
	}
	if mmDeploy.funcDeploy != nil {
		return mmDeploy.funcDeploy(ctx)
	}
	mmDeploy.t.Fatalf("Unexpected call to SolutionServiceMock.Deploy. %v", ctx)
	return
}

// DeployAfterCounter returns a count of finished SolutionServiceMock.Deploy invocations
func (mmDeploy *SolutionServiceMock) DeployAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeploy.afterDeployCounter)
}

// DeployBeforeCounter returns a count of SolutionServiceMock.Deploy invocations
func (mmDeploy *SolutionServiceMock) DeployBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeploy.beforeDeployCounter)
}

// Calls returns a list of arguments used in each call to SolutionServiceMock.Deploy.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeploy *mSolutionServiceMockDeploy) Calls() []*SolutionServiceMockDeployParams {
	mmDeploy.mutex.RLock()

	argCopy := make([]*SolutionServiceMockDeployParams, len(mmDeploy.callArgs))
	copy(argCopy, mmDeploy.callArgs)

	mmDeploy.mutex.RUnlock()

	return argCopy
}

// MinimockDeployDone returns true if the count of the Deploy invocations corresponds
// the number of defined expectations
func (m *SolutionServiceMock) MinimockDeployDone() bool {
	if m.DeployMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.DeployMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.DeployMock.invocationsDone()
}

// MinimockDeployInspect logs each unmet expectation
func (m *SolutionServiceMock) MinimockDeployInspect() {
	for _, e := range m.DeployMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SolutionServiceMock.Deploy with params: %#v", *e.params)
		}
	}

	afterDeployCounter := mm_atomic.LoadUint64(&m.afterDeployCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.DeployMock.defaultExpectation != nil && afterDeployCounter < 1 {
		if m.DeployMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SolutionServiceMock.Deploy")
		} else {
			m.t.Errorf("Expected call to SolutionServiceMock.Deploy with params: %#v", *m.DeployMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeploy != nil && afterDeployCounter < 1 {
		m.t.Error("Expected call to SolutionServiceMock.Deploy")
	}

	if !m.DeployMock.invocationsDone() && afterDeployCounter > 0 {
		m.t.Errorf("Expected %d calls to SolutionServiceMock.Deploy but found %d calls",
			mm_atomic.LoadUint64(&m.DeployMock.expectedInvocations), afterDeployCounter)
	}
}

type mSolutionServiceMockList struct {
	optional           bool
	mock               *SolutionServiceMock
	defaultExpectation *SolutionServiceMockListExpectation
	expectations       []*SolutionServiceMockListExpectation

	callArgs []*SolutionServiceMockListParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// SolutionServiceMockListExpectation specifies expectation struct of the SolutionService.List
type SolutionServiceMockListExpectation struct {
	mock      *SolutionServiceMock
	params    *SolutionServiceMockListParams
	paramPtrs *SolutionServiceMockListParamPtrs
	results   *SolutionServiceMockListResults
	Counter   uint64
}

// SolutionServiceMockListParams contains parameters of the SolutionService.List
type SolutionServiceMockListParams struct {
	ctx context.Context
}

// SolutionServiceMockListParamPtrs contains pointers to parameters of the SolutionService.List
type SolutionServiceMockListParamPtrs struct {
	ctx *context.Context
}

// SolutionServiceMockListResults contains results of the SolutionService.List
type SolutionServiceMockListResults struct {
	spa1 []*model.Solution
	err  error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option by default unless you really need it, as it helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmList *mSolutionServiceMockList) Optional() *mSolutionServiceMockList {
	mmList.optional = true
	return mmList
}

// Expect sets up expected params for SolutionService.List
func (mmList *mSolutionServiceMockList) Expect(ctx context.Context) *mSolutionServiceMockList {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("SolutionServiceMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &SolutionServiceMockListExpectation{}
	}

	if mmList.defaultExpectation.paramPtrs != nil {
		mmList.mock.t.Fatalf("SolutionServiceMock.List mock is already set by ExpectParams functions")
	}

	mmList.defaultExpectation.params = &SolutionServiceMockListParams{ctx}
	for _, e := range mmList.expectations {
		if minimock.Equal(e.params, mmList.defaultExpectation.params) {
			mmList.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmList.defaultExpectation.params)
		}
	}

	return mmList
}

// ExpectCtxParam1 sets up expected param ctx for SolutionService.List
func (mmList *mSolutionServiceMockList) ExpectCtxParam1(ctx context.Context) *mSolutionServiceMockList {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("SolutionServiceMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &SolutionServiceMockListExpectation{}
	}

	if mmList.defaultExpectation.params != nil {
		mmList.mock.t.Fatalf("SolutionServiceMock.List mock is already set by Expect")
	}

	if mmList.defaultExpectation.paramPtrs == nil {
		mmList.defaultExpectation.paramPtrs = &SolutionServiceMockListParamPtrs{}
	}
	mmList.defaultExpectation.paramPtrs.ctx = &ctx

	return mmList
}

// Inspect accepts an inspector function that has same arguments as the SolutionService.List
func (mmList *mSolutionServiceMockList) Inspect(f func(ctx context.Context)) *mSolutionServiceMockList {
	if mmList.mock.inspectFuncList != nil {
		mmList.mock.t.Fatalf("Inspect function is already set for SolutionServiceMock.List")
	}

	mmList.mock.inspectFuncList = f

	return mmList
}

// Return sets up results that will be returned by SolutionService.List
func (mmList *mSolutionServiceMockList) Return(spa1 []*model.Solution, err error) *SolutionServiceMock {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("SolutionServiceMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &SolutionServiceMockListExpectation{mock: mmList.mock}
	}
	mmList.defaultExpectation.results = &SolutionServiceMockListResults{spa1, err}
	return mmList.mock
}

// Set uses given function f to mock the SolutionService.List method
func (mmList *mSolutionServiceMockList) Set(f func(ctx context.Context) (spa1 []*model.Solution, err error)) *SolutionServiceMock {
	if mmList.defaultExpectation != nil {
		mmList.mock.t.Fatalf("Default expectation is already set for the SolutionService.List method")
	}

	if len(mmList.expectations) > 0 {
		mmList.mock.t.Fatalf("Some expectations are already set for the SolutionService.List method")
	}

	mmList.mock.funcList = f
	return mmList.mock
}

// When sets expectation for the SolutionService.List which will trigger the result defined by the following
// Then helper
func (mmList *mSolutionServiceMockList) When(ctx context.Context) *SolutionServiceMockListExpectation {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("SolutionServiceMock.List mock is already set by Set")
	}

	expectation := &SolutionServiceMockListExpectation{
		mock:   mmList.mock,
		params: &SolutionServiceMockListParams{ctx},
	}
	mmList.expectations = append(mmList.expectations, expectation)
	return expectation
}

// Then sets up SolutionService.List return parameters for the expectation previously defined by the When method
func (e *SolutionServiceMockListExpectation) Then(spa1 []*model.Solution, err error) *SolutionServiceMock {
	e.results = &SolutionServiceMockListResults{spa1, err}
	return e.mock
}

// Times sets number of times SolutionService.List should be invoked
func (mmList *mSolutionServiceMockList) Times(n uint64) *mSolutionServiceMockList {
	if n == 0 {
		mmList.mock.t.Fatalf("Times of SolutionServiceMock.List mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmList.expectedInvocations, n)
	return mmList
}

func (mmList *mSolutionServiceMockList) invocationsDone() bool {
	if len(mmList.expectations) == 0 && mmList.defaultExpectation == nil && mmList.mock.funcList == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmList.mock.afterListCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmList.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// List implements service.SolutionService
func (mmList *SolutionServiceMock) List(ctx context.Context) (spa1 []*model.Solution, err error) {
	mm_atomic.AddUint64(&mmList.beforeListCounter, 1)
	defer mm_atomic.AddUint64(&mmList.afterListCounter, 1)

	if mmList.inspectFuncList != nil {
		mmList.inspectFuncList(ctx)
	}

	mm_params := SolutionServiceMockListParams{ctx}

	// Record call args
	mmList.ListMock.mutex.Lock()
	mmList.ListMock.callArgs = append(mmList.ListMock.callArgs, &mm_params)
	mmList.ListMock.mutex.Unlock()

	for _, e := range mmList.ListMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.spa1, e.results.err
		}
	}

	if mmList.ListMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmList.ListMock.defaultExpectation.Counter, 1)
		mm_want := mmList.ListMock.defaultExpectation.params
		mm_want_ptrs := mmList.ListMock.defaultExpectation.paramPtrs

		mm_got := SolutionServiceMockListParams{ctx}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmList.t.Errorf("SolutionServiceMock.List got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmList.t.Errorf("SolutionServiceMock.List got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmList.ListMock.defaultExpectation.results
		if mm_results == nil {
			mmList.t.Fatal("No results are set for the SolutionServiceMock.List")
		}
		return (*mm_results).spa1, (*mm_results).err
	}
	if mmList.funcList != nil {
		return mmList.funcList(ctx)
	}
	mmList.t.Fatalf("Unexpected call to SolutionServiceMock.List. %v", ctx)
	return
}

// ListAfterCounter returns a count of finished SolutionServiceMock.List invocations
func (mmList *SolutionServiceMock) ListAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.afterListCounter)
}

// ListBeforeCounter returns a count of SolutionServiceMock.List invocations
func (mmList *SolutionServiceMock) ListBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.beforeListCounter)
}

// Calls returns a list of arguments used in each call to SolutionServiceMock.List.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmList *mSolutionServiceMockList) Calls() []*SolutionServiceMockListParams {
	mmList.mutex.RLock()

	argCopy := make([]*SolutionServiceMockListParams, len(mmList.callArgs))
	copy(argCopy, mmList.callArgs)

	mmList.mutex.RUnlock()

	return argCopy
}

// MinimockListDone returns true if the count of the List invocations corresponds
// the number of defined expectations
func (m *SolutionServiceMock) MinimockListDone() bool {
	if m.ListMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.ListMock.invocationsDone()
}

// MinimockListInspect logs each unmet expectation
func (m *SolutionServiceMock) MinimockListInspect() {
	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SolutionServiceMock.List with params: %#v", *e.params)
		}
	}

	afterListCounter := mm_atomic.LoadUint64(&m.afterListCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.ListMock.defaultExpectation != nil && afterListCounter < 1 {
		if m.ListMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SolutionServiceMock.List")
		} else {
			m.t.Errorf("Expected call to SolutionServiceMock.List with params: %#v", *m.ListMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcList != nil && afterListCounter < 1 {
		m.t.Error("Expected call to SolutionServiceMock.List")
	}

	if !m.ListMock.invocationsDone() && afterListCounter > 0 {
		m.t.Errorf("Expected %d calls to SolutionServiceMock.List but found %d calls",
			mm_atomic.LoadUint64(&m.ListMock.expectedInvocations), afterListCounter)
	}
}

type mSolutionServiceMockStatus struct {
	optional           bool
	mock               *SolutionServiceMock
	defaultExpectation *SolutionServiceMockStatusExpectation
	expectations       []*SolutionServiceMockStatusExpectation

	callArgs []*SolutionServiceMockStatusParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// SolutionServiceMockStatusExpectation specifies expectation struct of the SolutionService.Status
type SolutionServiceMockStatusExpectation struct {
	mock      *SolutionServiceMock
	params    *SolutionServiceMockStatusParams
	paramPtrs *SolutionServiceMockStatusParamPtrs
	results   *SolutionServiceMockStatusResults
	Counter   uint64
}

// SolutionServiceMockStatusParams contains parameters of the SolutionService.Status
type SolutionServiceMockStatusParams struct {
	ctx context.Context
}

// SolutionServiceMockStatusParamPtrs contains pointers to parameters of the SolutionService.Status
type SolutionServiceMockStatusParamPtrs struct {
	ctx *context.Context
}

// SolutionServiceMockStatusResults contains results of the SolutionService.Status
type SolutionServiceMockStatusResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option by default unless you really need it, as it helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmStatus *mSolutionServiceMockStatus) Optional() *mSolutionServiceMockStatus {
	mmStatus.optional = true
	return mmStatus
}

// Expect sets up expected params for SolutionService.Status
func (mmStatus *mSolutionServiceMockStatus) Expect(ctx context.Context) *mSolutionServiceMockStatus {
	if mmStatus.mock.funcStatus != nil {
		mmStatus.mock.t.Fatalf("SolutionServiceMock.Status mock is already set by Set")
	}

	if mmStatus.defaultExpectation == nil {
		mmStatus.defaultExpectation = &SolutionServiceMockStatusExpectation{}
	}

	if mmStatus.defaultExpectation.paramPtrs != nil {
		mmStatus.mock.t.Fatalf("SolutionServiceMock.Status mock is already set by ExpectParams functions")
	}

	mmStatus.defaultExpectation.params = &SolutionServiceMockStatusParams{ctx}
	for _, e := range mmStatus.expectations {
		if minimock.Equal(e.params, mmStatus.defaultExpectation.params) {
			mmStatus.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmStatus.defaultExpectation.params)
		}
	}

	return mmStatus
}

// ExpectCtxParam1 sets up expected param ctx for SolutionService.Status
func (mmStatus *mSolutionServiceMockStatus) ExpectCtxParam1(ctx context.Context) *mSolutionServiceMockStatus {
	if mmStatus.mock.funcStatus != nil {
		mmStatus.mock.t.Fatalf("SolutionServiceMock.Status mock is already set by Set")
	}

	if mmStatus.defaultExpectation == nil {
		mmStatus.defaultExpectation = &SolutionServiceMockStatusExpectation{}
	}

	if mmStatus.defaultExpectation.params != nil {
		mmStatus.mock.t.Fatalf("SolutionServiceMock.Status mock is already set by Expect")
	}

	if mmStatus.defaultExpectation.paramPtrs == nil {
		mmStatus.defaultExpectation.paramPtrs = &SolutionServiceMockStatusParamPtrs{}
	}
	mmStatus.defaultExpectation.paramPtrs.ctx = &ctx

	return mmStatus
}

// Inspect accepts an inspector function that has same arguments as the SolutionService.Status
func (mmStatus *mSolutionServiceMockStatus) Inspect(f func(ctx context.Context)) *mSolutionServiceMockStatus {
	if mmStatus.mock.inspectFuncStatus != nil {
		mmStatus.mock.t.Fatalf("Inspect function is already set for SolutionServiceMock.Status")
	}

	mmStatus.mock.inspectFuncStatus = f

	return mmStatus
}

// Return sets up results that will be returned by SolutionService.Status
func (mmStatus *mSolutionServiceMockStatus) Return(err error) *SolutionServiceMock {
	if mmStatus.mock.funcStatus != nil {
		mmStatus.mock.t.Fatalf("SolutionServiceMock.Status mock is already set by Set")
	}

	if mmStatus.defaultExpectation == nil {
		mmStatus.defaultExpectation = &SolutionServiceMockStatusExpectation{mock: mmStatus.mock}
	}
	mmStatus.defaultExpectation.results = &SolutionServiceMockStatusResults{err}
	return mmStatus.mock
}

// Set uses given function f to mock the SolutionService.Status method
func (mmStatus *mSolutionServiceMockStatus) Set(f func(ctx context.Context) (err error)) *SolutionServiceMock {
	if mmStatus.defaultExpectation != nil {
		mmStatus.mock.t.Fatalf("Default expectation is already set for the SolutionService.Status method")
	}

	if len(mmStatus.expectations) > 0 {
		mmStatus.mock.t.Fatalf("Some expectations are already set for the SolutionService.Status method")
	}

	mmStatus.mock.funcStatus = f
	return mmStatus.mock
}

// When sets expectation for the SolutionService.Status which will trigger the result defined by the following
// Then helper
func (mmStatus *mSolutionServiceMockStatus) When(ctx context.Context) *SolutionServiceMockStatusExpectation {
	if mmStatus.mock.funcStatus != nil {
		mmStatus.mock.t.Fatalf("SolutionServiceMock.Status mock is already set by Set")
	}

	expectation := &SolutionServiceMockStatusExpectation{
		mock:   mmStatus.mock,
		params: &SolutionServiceMockStatusParams{ctx},
	}
	mmStatus.expectations = append(mmStatus.expectations, expectation)
	return expectation
}

// Then sets up SolutionService.Status return parameters for the expectation previously defined by the When method
func (e *SolutionServiceMockStatusExpectation) Then(err error) *SolutionServiceMock {
	e.results = &SolutionServiceMockStatusResults{err}
	return e.mock
}

// Times sets number of times SolutionService.Status should be invoked
func (mmStatus *mSolutionServiceMockStatus) Times(n uint64) *mSolutionServiceMockStatus {
	if n == 0 {
		mmStatus.mock.t.Fatalf("Times of SolutionServiceMock.Status mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmStatus.expectedInvocations, n)
	return mmStatus
}

func (mmStatus *mSolutionServiceMockStatus) invocationsDone() bool {
	if len(mmStatus.expectations) == 0 && mmStatus.defaultExpectation == nil && mmStatus.mock.funcStatus == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmStatus.mock.afterStatusCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmStatus.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Status implements service.SolutionService
func (mmStatus *SolutionServiceMock) Status(ctx context.Context) (err error) {
	mm_atomic.AddUint64(&mmStatus.beforeStatusCounter, 1)
	defer mm_atomic.AddUint64(&mmStatus.afterStatusCounter, 1)

	if mmStatus.inspectFuncStatus != nil {
		mmStatus.inspectFuncStatus(ctx)
	}

	mm_params := SolutionServiceMockStatusParams{ctx}

	// Record call args
	mmStatus.StatusMock.mutex.Lock()
	mmStatus.StatusMock.callArgs = append(mmStatus.StatusMock.callArgs, &mm_params)
	mmStatus.StatusMock.mutex.Unlock()

	for _, e := range mmStatus.StatusMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmStatus.StatusMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmStatus.StatusMock.defaultExpectation.Counter, 1)
		mm_want := mmStatus.StatusMock.defaultExpectation.params
		mm_want_ptrs := mmStatus.StatusMock.defaultExpectation.paramPtrs

		mm_got := SolutionServiceMockStatusParams{ctx}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmStatus.t.Errorf("SolutionServiceMock.Status got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmStatus.t.Errorf("SolutionServiceMock.Status got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmStatus.StatusMock.defaultExpectation.results
		if mm_results == nil {
			mmStatus.t.Fatal("No results are set for the SolutionServiceMock.Status")
		}
		return (*mm_results).err
	}
	if mmStatus.funcStatus != nil {
		return mmStatus.funcStatus(ctx)
	}
	mmStatus.t.Fatalf("Unexpected call to SolutionServiceMock.Status. %v", ctx)
	return
}

// StatusAfterCounter returns a count of finished SolutionServiceMock.Status invocations
func (mmStatus *SolutionServiceMock) StatusAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStatus.afterStatusCounter)
}

// StatusBeforeCounter returns a count of SolutionServiceMock.Status invocations
func (mmStatus *SolutionServiceMock) StatusBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStatus.beforeStatusCounter)
}

// Calls returns a list of arguments used in each call to SolutionServiceMock.Status.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmStatus *mSolutionServiceMockStatus) Calls() []*SolutionServiceMockStatusParams {
	mmStatus.mutex.RLock()

	argCopy := make([]*SolutionServiceMockStatusParams, len(mmStatus.callArgs))
	copy(argCopy, mmStatus.callArgs)

	mmStatus.mutex.RUnlock()

	return argCopy
}

// MinimockStatusDone returns true if the count of the Status invocations corresponds
// the number of defined expectations
func (m *SolutionServiceMock) MinimockStatusDone() bool {
	if m.StatusMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.StatusMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.StatusMock.invocationsDone()
}

// MinimockStatusInspect logs each unmet expectation
func (m *SolutionServiceMock) MinimockStatusInspect() {
	for _, e := range m.StatusMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SolutionServiceMock.Status with params: %#v", *e.params)
		}
	}

	afterStatusCounter := mm_atomic.LoadUint64(&m.afterStatusCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.StatusMock.defaultExpectation != nil && afterStatusCounter < 1 {
		if m.StatusMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SolutionServiceMock.Status")
		} else {
			m.t.Errorf("Expected call to SolutionServiceMock.Status with params: %#v", *m.StatusMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStatus != nil && afterStatusCounter < 1 {
		m.t.Error("Expected call to SolutionServiceMock.Status")
	}

	if !m.StatusMock.invocationsDone() && afterStatusCounter > 0 {
		m.t.Errorf("Expected %d calls to SolutionServiceMock.Status but found %d calls",
			mm_atomic.LoadUint64(&m.StatusMock.expectedInvocations), afterStatusCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SolutionServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockDeployInspect()

			m.MinimockListInspect()

			m.MinimockStatusInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SolutionServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *SolutionServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeployDone() &&
		m.MinimockListDone() &&
		m.MinimockStatusDone()
}