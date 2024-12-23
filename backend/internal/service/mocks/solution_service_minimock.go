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

	funcDeploy          func(ctx context.Context, deploy *model.Deploy) (i1 int64, err error)
	inspectFuncDeploy   func(ctx context.Context, deploy *model.Deploy)
	afterDeployCounter  uint64
	beforeDeployCounter uint64
	DeployMock          mSolutionServiceMockDeploy

	funcDeployStatus          func(ctx context.Context, id int64) (s1 string, s2 string, err error)
	inspectFuncDeployStatus   func(ctx context.Context, id int64)
	afterDeployStatusCounter  uint64
	beforeDeployStatusCounter uint64
	DeployStatusMock          mSolutionServiceMockDeployStatus

	funcList          func(ctx context.Context) (spa1 []*model.Solution, err error)
	inspectFuncList   func(ctx context.Context)
	afterListCounter  uint64
	beforeListCounter uint64
	ListMock          mSolutionServiceMockList
}

// NewSolutionServiceMock returns a mock for service.SolutionService
func NewSolutionServiceMock(t minimock.Tester) *SolutionServiceMock {
	m := &SolutionServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeployMock = mSolutionServiceMockDeploy{mock: m}
	m.DeployMock.callArgs = []*SolutionServiceMockDeployParams{}

	m.DeployStatusMock = mSolutionServiceMockDeployStatus{mock: m}
	m.DeployStatusMock.callArgs = []*SolutionServiceMockDeployStatusParams{}

	m.ListMock = mSolutionServiceMockList{mock: m}
	m.ListMock.callArgs = []*SolutionServiceMockListParams{}

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
	ctx    context.Context
	deploy *model.Deploy
}

// SolutionServiceMockDeployParamPtrs contains pointers to parameters of the SolutionService.Deploy
type SolutionServiceMockDeployParamPtrs struct {
	ctx    *context.Context
	deploy **model.Deploy
}

// SolutionServiceMockDeployResults contains results of the SolutionService.Deploy
type SolutionServiceMockDeployResults struct {
	i1  int64
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
func (mmDeploy *mSolutionServiceMockDeploy) Expect(ctx context.Context, deploy *model.Deploy) *mSolutionServiceMockDeploy {
	if mmDeploy.mock.funcDeploy != nil {
		mmDeploy.mock.t.Fatalf("SolutionServiceMock.Deploy mock is already set by Set")
	}

	if mmDeploy.defaultExpectation == nil {
		mmDeploy.defaultExpectation = &SolutionServiceMockDeployExpectation{}
	}

	if mmDeploy.defaultExpectation.paramPtrs != nil {
		mmDeploy.mock.t.Fatalf("SolutionServiceMock.Deploy mock is already set by ExpectParams functions")
	}

	mmDeploy.defaultExpectation.params = &SolutionServiceMockDeployParams{ctx, deploy}
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

// ExpectDeployParam2 sets up expected param deploy for SolutionService.Deploy
func (mmDeploy *mSolutionServiceMockDeploy) ExpectDeployParam2(deploy *model.Deploy) *mSolutionServiceMockDeploy {
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
	mmDeploy.defaultExpectation.paramPtrs.deploy = &deploy

	return mmDeploy
}

// Inspect accepts an inspector function that has same arguments as the SolutionService.Deploy
func (mmDeploy *mSolutionServiceMockDeploy) Inspect(f func(ctx context.Context, deploy *model.Deploy)) *mSolutionServiceMockDeploy {
	if mmDeploy.mock.inspectFuncDeploy != nil {
		mmDeploy.mock.t.Fatalf("Inspect function is already set for SolutionServiceMock.Deploy")
	}

	mmDeploy.mock.inspectFuncDeploy = f

	return mmDeploy
}

// Return sets up results that will be returned by SolutionService.Deploy
func (mmDeploy *mSolutionServiceMockDeploy) Return(i1 int64, err error) *SolutionServiceMock {
	if mmDeploy.mock.funcDeploy != nil {
		mmDeploy.mock.t.Fatalf("SolutionServiceMock.Deploy mock is already set by Set")
	}

	if mmDeploy.defaultExpectation == nil {
		mmDeploy.defaultExpectation = &SolutionServiceMockDeployExpectation{mock: mmDeploy.mock}
	}
	mmDeploy.defaultExpectation.results = &SolutionServiceMockDeployResults{i1, err}
	return mmDeploy.mock
}

// Set uses given function f to mock the SolutionService.Deploy method
func (mmDeploy *mSolutionServiceMockDeploy) Set(f func(ctx context.Context, deploy *model.Deploy) (i1 int64, err error)) *SolutionServiceMock {
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
func (mmDeploy *mSolutionServiceMockDeploy) When(ctx context.Context, deploy *model.Deploy) *SolutionServiceMockDeployExpectation {
	if mmDeploy.mock.funcDeploy != nil {
		mmDeploy.mock.t.Fatalf("SolutionServiceMock.Deploy mock is already set by Set")
	}

	expectation := &SolutionServiceMockDeployExpectation{
		mock:   mmDeploy.mock,
		params: &SolutionServiceMockDeployParams{ctx, deploy},
	}
	mmDeploy.expectations = append(mmDeploy.expectations, expectation)
	return expectation
}

// Then sets up SolutionService.Deploy return parameters for the expectation previously defined by the When method
func (e *SolutionServiceMockDeployExpectation) Then(i1 int64, err error) *SolutionServiceMock {
	e.results = &SolutionServiceMockDeployResults{i1, err}
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
func (mmDeploy *SolutionServiceMock) Deploy(ctx context.Context, deploy *model.Deploy) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmDeploy.beforeDeployCounter, 1)
	defer mm_atomic.AddUint64(&mmDeploy.afterDeployCounter, 1)

	if mmDeploy.inspectFuncDeploy != nil {
		mmDeploy.inspectFuncDeploy(ctx, deploy)
	}

	mm_params := SolutionServiceMockDeployParams{ctx, deploy}

	// Record call args
	mmDeploy.DeployMock.mutex.Lock()
	mmDeploy.DeployMock.callArgs = append(mmDeploy.DeployMock.callArgs, &mm_params)
	mmDeploy.DeployMock.mutex.Unlock()

	for _, e := range mmDeploy.DeployMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmDeploy.DeployMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeploy.DeployMock.defaultExpectation.Counter, 1)
		mm_want := mmDeploy.DeployMock.defaultExpectation.params
		mm_want_ptrs := mmDeploy.DeployMock.defaultExpectation.paramPtrs

		mm_got := SolutionServiceMockDeployParams{ctx, deploy}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmDeploy.t.Errorf("SolutionServiceMock.Deploy got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.deploy != nil && !minimock.Equal(*mm_want_ptrs.deploy, mm_got.deploy) {
				mmDeploy.t.Errorf("SolutionServiceMock.Deploy got unexpected parameter deploy, want: %#v, got: %#v%s\n", *mm_want_ptrs.deploy, mm_got.deploy, minimock.Diff(*mm_want_ptrs.deploy, mm_got.deploy))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeploy.t.Errorf("SolutionServiceMock.Deploy got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeploy.DeployMock.defaultExpectation.results
		if mm_results == nil {
			mmDeploy.t.Fatal("No results are set for the SolutionServiceMock.Deploy")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmDeploy.funcDeploy != nil {
		return mmDeploy.funcDeploy(ctx, deploy)
	}
	mmDeploy.t.Fatalf("Unexpected call to SolutionServiceMock.Deploy. %v %v", ctx, deploy)
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

type mSolutionServiceMockDeployStatus struct {
	optional           bool
	mock               *SolutionServiceMock
	defaultExpectation *SolutionServiceMockDeployStatusExpectation
	expectations       []*SolutionServiceMockDeployStatusExpectation

	callArgs []*SolutionServiceMockDeployStatusParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// SolutionServiceMockDeployStatusExpectation specifies expectation struct of the SolutionService.DeployStatus
type SolutionServiceMockDeployStatusExpectation struct {
	mock      *SolutionServiceMock
	params    *SolutionServiceMockDeployStatusParams
	paramPtrs *SolutionServiceMockDeployStatusParamPtrs
	results   *SolutionServiceMockDeployStatusResults
	Counter   uint64
}

// SolutionServiceMockDeployStatusParams contains parameters of the SolutionService.DeployStatus
type SolutionServiceMockDeployStatusParams struct {
	ctx context.Context
	id  int64
}

// SolutionServiceMockDeployStatusParamPtrs contains pointers to parameters of the SolutionService.DeployStatus
type SolutionServiceMockDeployStatusParamPtrs struct {
	ctx *context.Context
	id  *int64
}

// SolutionServiceMockDeployStatusResults contains results of the SolutionService.DeployStatus
type SolutionServiceMockDeployStatusResults struct {
	s1  string
	s2  string
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option by default unless you really need it, as it helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmDeployStatus *mSolutionServiceMockDeployStatus) Optional() *mSolutionServiceMockDeployStatus {
	mmDeployStatus.optional = true
	return mmDeployStatus
}

// Expect sets up expected params for SolutionService.DeployStatus
func (mmDeployStatus *mSolutionServiceMockDeployStatus) Expect(ctx context.Context, id int64) *mSolutionServiceMockDeployStatus {
	if mmDeployStatus.mock.funcDeployStatus != nil {
		mmDeployStatus.mock.t.Fatalf("SolutionServiceMock.DeployStatus mock is already set by Set")
	}

	if mmDeployStatus.defaultExpectation == nil {
		mmDeployStatus.defaultExpectation = &SolutionServiceMockDeployStatusExpectation{}
	}

	if mmDeployStatus.defaultExpectation.paramPtrs != nil {
		mmDeployStatus.mock.t.Fatalf("SolutionServiceMock.DeployStatus mock is already set by ExpectParams functions")
	}

	mmDeployStatus.defaultExpectation.params = &SolutionServiceMockDeployStatusParams{ctx, id}
	for _, e := range mmDeployStatus.expectations {
		if minimock.Equal(e.params, mmDeployStatus.defaultExpectation.params) {
			mmDeployStatus.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeployStatus.defaultExpectation.params)
		}
	}

	return mmDeployStatus
}

// ExpectCtxParam1 sets up expected param ctx for SolutionService.DeployStatus
func (mmDeployStatus *mSolutionServiceMockDeployStatus) ExpectCtxParam1(ctx context.Context) *mSolutionServiceMockDeployStatus {
	if mmDeployStatus.mock.funcDeployStatus != nil {
		mmDeployStatus.mock.t.Fatalf("SolutionServiceMock.DeployStatus mock is already set by Set")
	}

	if mmDeployStatus.defaultExpectation == nil {
		mmDeployStatus.defaultExpectation = &SolutionServiceMockDeployStatusExpectation{}
	}

	if mmDeployStatus.defaultExpectation.params != nil {
		mmDeployStatus.mock.t.Fatalf("SolutionServiceMock.DeployStatus mock is already set by Expect")
	}

	if mmDeployStatus.defaultExpectation.paramPtrs == nil {
		mmDeployStatus.defaultExpectation.paramPtrs = &SolutionServiceMockDeployStatusParamPtrs{}
	}
	mmDeployStatus.defaultExpectation.paramPtrs.ctx = &ctx

	return mmDeployStatus
}

// ExpectIdParam2 sets up expected param id for SolutionService.DeployStatus
func (mmDeployStatus *mSolutionServiceMockDeployStatus) ExpectIdParam2(id int64) *mSolutionServiceMockDeployStatus {
	if mmDeployStatus.mock.funcDeployStatus != nil {
		mmDeployStatus.mock.t.Fatalf("SolutionServiceMock.DeployStatus mock is already set by Set")
	}

	if mmDeployStatus.defaultExpectation == nil {
		mmDeployStatus.defaultExpectation = &SolutionServiceMockDeployStatusExpectation{}
	}

	if mmDeployStatus.defaultExpectation.params != nil {
		mmDeployStatus.mock.t.Fatalf("SolutionServiceMock.DeployStatus mock is already set by Expect")
	}

	if mmDeployStatus.defaultExpectation.paramPtrs == nil {
		mmDeployStatus.defaultExpectation.paramPtrs = &SolutionServiceMockDeployStatusParamPtrs{}
	}
	mmDeployStatus.defaultExpectation.paramPtrs.id = &id

	return mmDeployStatus
}

// Inspect accepts an inspector function that has same arguments as the SolutionService.DeployStatus
func (mmDeployStatus *mSolutionServiceMockDeployStatus) Inspect(f func(ctx context.Context, id int64)) *mSolutionServiceMockDeployStatus {
	if mmDeployStatus.mock.inspectFuncDeployStatus != nil {
		mmDeployStatus.mock.t.Fatalf("Inspect function is already set for SolutionServiceMock.DeployStatus")
	}

	mmDeployStatus.mock.inspectFuncDeployStatus = f

	return mmDeployStatus
}

// Return sets up results that will be returned by SolutionService.DeployStatus
func (mmDeployStatus *mSolutionServiceMockDeployStatus) Return(s1 string, s2 string, err error) *SolutionServiceMock {
	if mmDeployStatus.mock.funcDeployStatus != nil {
		mmDeployStatus.mock.t.Fatalf("SolutionServiceMock.DeployStatus mock is already set by Set")
	}

	if mmDeployStatus.defaultExpectation == nil {
		mmDeployStatus.defaultExpectation = &SolutionServiceMockDeployStatusExpectation{mock: mmDeployStatus.mock}
	}
	mmDeployStatus.defaultExpectation.results = &SolutionServiceMockDeployStatusResults{s1, s2, err}
	return mmDeployStatus.mock
}

// Set uses given function f to mock the SolutionService.DeployStatus method
func (mmDeployStatus *mSolutionServiceMockDeployStatus) Set(f func(ctx context.Context, id int64) (s1 string, s2 string, err error)) *SolutionServiceMock {
	if mmDeployStatus.defaultExpectation != nil {
		mmDeployStatus.mock.t.Fatalf("Default expectation is already set for the SolutionService.DeployStatus method")
	}

	if len(mmDeployStatus.expectations) > 0 {
		mmDeployStatus.mock.t.Fatalf("Some expectations are already set for the SolutionService.DeployStatus method")
	}

	mmDeployStatus.mock.funcDeployStatus = f
	return mmDeployStatus.mock
}

// When sets expectation for the SolutionService.DeployStatus which will trigger the result defined by the following
// Then helper
func (mmDeployStatus *mSolutionServiceMockDeployStatus) When(ctx context.Context, id int64) *SolutionServiceMockDeployStatusExpectation {
	if mmDeployStatus.mock.funcDeployStatus != nil {
		mmDeployStatus.mock.t.Fatalf("SolutionServiceMock.DeployStatus mock is already set by Set")
	}

	expectation := &SolutionServiceMockDeployStatusExpectation{
		mock:   mmDeployStatus.mock,
		params: &SolutionServiceMockDeployStatusParams{ctx, id},
	}
	mmDeployStatus.expectations = append(mmDeployStatus.expectations, expectation)
	return expectation
}

// Then sets up SolutionService.DeployStatus return parameters for the expectation previously defined by the When method
func (e *SolutionServiceMockDeployStatusExpectation) Then(s1 string, s2 string, err error) *SolutionServiceMock {
	e.results = &SolutionServiceMockDeployStatusResults{s1, s2, err}
	return e.mock
}

// Times sets number of times SolutionService.DeployStatus should be invoked
func (mmDeployStatus *mSolutionServiceMockDeployStatus) Times(n uint64) *mSolutionServiceMockDeployStatus {
	if n == 0 {
		mmDeployStatus.mock.t.Fatalf("Times of SolutionServiceMock.DeployStatus mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmDeployStatus.expectedInvocations, n)
	return mmDeployStatus
}

func (mmDeployStatus *mSolutionServiceMockDeployStatus) invocationsDone() bool {
	if len(mmDeployStatus.expectations) == 0 && mmDeployStatus.defaultExpectation == nil && mmDeployStatus.mock.funcDeployStatus == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmDeployStatus.mock.afterDeployStatusCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmDeployStatus.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// DeployStatus implements service.SolutionService
func (mmDeployStatus *SolutionServiceMock) DeployStatus(ctx context.Context, id int64) (s1 string, s2 string, err error) {
	mm_atomic.AddUint64(&mmDeployStatus.beforeDeployStatusCounter, 1)
	defer mm_atomic.AddUint64(&mmDeployStatus.afterDeployStatusCounter, 1)

	if mmDeployStatus.inspectFuncDeployStatus != nil {
		mmDeployStatus.inspectFuncDeployStatus(ctx, id)
	}

	mm_params := SolutionServiceMockDeployStatusParams{ctx, id}

	// Record call args
	mmDeployStatus.DeployStatusMock.mutex.Lock()
	mmDeployStatus.DeployStatusMock.callArgs = append(mmDeployStatus.DeployStatusMock.callArgs, &mm_params)
	mmDeployStatus.DeployStatusMock.mutex.Unlock()

	for _, e := range mmDeployStatus.DeployStatusMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.s2, e.results.err
		}
	}

	if mmDeployStatus.DeployStatusMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeployStatus.DeployStatusMock.defaultExpectation.Counter, 1)
		mm_want := mmDeployStatus.DeployStatusMock.defaultExpectation.params
		mm_want_ptrs := mmDeployStatus.DeployStatusMock.defaultExpectation.paramPtrs

		mm_got := SolutionServiceMockDeployStatusParams{ctx, id}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmDeployStatus.t.Errorf("SolutionServiceMock.DeployStatus got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.id != nil && !minimock.Equal(*mm_want_ptrs.id, mm_got.id) {
				mmDeployStatus.t.Errorf("SolutionServiceMock.DeployStatus got unexpected parameter id, want: %#v, got: %#v%s\n", *mm_want_ptrs.id, mm_got.id, minimock.Diff(*mm_want_ptrs.id, mm_got.id))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeployStatus.t.Errorf("SolutionServiceMock.DeployStatus got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeployStatus.DeployStatusMock.defaultExpectation.results
		if mm_results == nil {
			mmDeployStatus.t.Fatal("No results are set for the SolutionServiceMock.DeployStatus")
		}
		return (*mm_results).s1, (*mm_results).s2, (*mm_results).err
	}
	if mmDeployStatus.funcDeployStatus != nil {
		return mmDeployStatus.funcDeployStatus(ctx, id)
	}
	mmDeployStatus.t.Fatalf("Unexpected call to SolutionServiceMock.DeployStatus. %v %v", ctx, id)
	return
}

// DeployStatusAfterCounter returns a count of finished SolutionServiceMock.DeployStatus invocations
func (mmDeployStatus *SolutionServiceMock) DeployStatusAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeployStatus.afterDeployStatusCounter)
}

// DeployStatusBeforeCounter returns a count of SolutionServiceMock.DeployStatus invocations
func (mmDeployStatus *SolutionServiceMock) DeployStatusBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeployStatus.beforeDeployStatusCounter)
}

// Calls returns a list of arguments used in each call to SolutionServiceMock.DeployStatus.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeployStatus *mSolutionServiceMockDeployStatus) Calls() []*SolutionServiceMockDeployStatusParams {
	mmDeployStatus.mutex.RLock()

	argCopy := make([]*SolutionServiceMockDeployStatusParams, len(mmDeployStatus.callArgs))
	copy(argCopy, mmDeployStatus.callArgs)

	mmDeployStatus.mutex.RUnlock()

	return argCopy
}

// MinimockDeployStatusDone returns true if the count of the DeployStatus invocations corresponds
// the number of defined expectations
func (m *SolutionServiceMock) MinimockDeployStatusDone() bool {
	if m.DeployStatusMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.DeployStatusMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.DeployStatusMock.invocationsDone()
}

// MinimockDeployStatusInspect logs each unmet expectation
func (m *SolutionServiceMock) MinimockDeployStatusInspect() {
	for _, e := range m.DeployStatusMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SolutionServiceMock.DeployStatus with params: %#v", *e.params)
		}
	}

	afterDeployStatusCounter := mm_atomic.LoadUint64(&m.afterDeployStatusCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.DeployStatusMock.defaultExpectation != nil && afterDeployStatusCounter < 1 {
		if m.DeployStatusMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SolutionServiceMock.DeployStatus")
		} else {
			m.t.Errorf("Expected call to SolutionServiceMock.DeployStatus with params: %#v", *m.DeployStatusMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeployStatus != nil && afterDeployStatusCounter < 1 {
		m.t.Error("Expected call to SolutionServiceMock.DeployStatus")
	}

	if !m.DeployStatusMock.invocationsDone() && afterDeployStatusCounter > 0 {
		m.t.Errorf("Expected %d calls to SolutionServiceMock.DeployStatus but found %d calls",
			mm_atomic.LoadUint64(&m.DeployStatusMock.expectedInvocations), afterDeployStatusCounter)
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

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SolutionServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockDeployInspect()

			m.MinimockDeployStatusInspect()

			m.MinimockListInspect()
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
		m.MinimockDeployStatusDone() &&
		m.MinimockListDone()
}
