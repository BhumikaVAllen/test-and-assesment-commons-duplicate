package entity

// TestRankEntity TODO :Implement with test-evaluator
type TestRankEntity struct {
	TestIDStudentID string `dynamodbav:"TestIdStudentId"` // partition key
}
