test-coverage:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out  -coverpkg=gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/... -covermode count -count 5
	grep -v 'mocks\|config' coverage.tmp.out  > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out;
	go tool cover -func=./coverage.out | grep "total";
	grep -sqFx "/coverage.out" .gitignore || echo "/coverage.out" >> .gitignore
