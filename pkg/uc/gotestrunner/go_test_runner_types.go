package gotestrunner

type TestResult int

const (
	TestSuccess TestResult = iota
	TestFailure
)

type Suite struct {
	AllTests []TestGroup
}

type SuiteExecutionResult struct {
	AllTestResults []TestGroupExecutionResult
}

type TestGroupExecutionResult struct {
	Name                    string
	VersionExecutionResults []VersionExecutionResult
}

type VersionExecutionResult struct {
	ID     string
	Result TestResult
}

type TestGroup struct {
	Name     string //Ex: Postgres
	Packages []string
	Versions []Version
}

type Version struct {
	ID  string //Ex:1.9
	Env []string
	//VersionDependencyList []ContainerReference
	//TestConfig            TestConfig
}
