package test

// IntTestRecord for store int test data
type IntTestRecord struct {
	TargetValue int
	Expected    bool
}

// IntTestRecordWithIntReturn for store int test data
type IntTestRecordWithIntReturn struct {
	TargetValue int
	Expected    int
}

// StrTestRecord for store string test data
type StrTestRecord struct {
	TargetValue string
	Expected    bool
}
