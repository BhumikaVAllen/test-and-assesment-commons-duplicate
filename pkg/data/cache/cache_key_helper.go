package cache

import "fmt"

const (
	QuestionPaperInstructionsKeyFmt = "QUESTION_PAPER_INSTRUCTIONS_%s"
	QuestionPaperKeyFmt             = "QUESTION_PAPER_%s"
	TestInfoKeyFmt                  = "%s_TEST_INFO"
	StudentTestOverviewKeyFmt       = "%s#%s_STUDENT_TEST_OVERVIEW"
	StudentTestPermitKeyFmt         = "%s#%s_TEST_PERMIT"
)

func getQuestionPaperInstructionsCacheKey(testID string) string {
	return fmt.Sprintf(QuestionPaperInstructionsKeyFmt, testID)
}

func getQuestionPaperCacheKey(testID string) string {
	return fmt.Sprintf(QuestionPaperKeyFmt, testID)
}

func getTestInfoCacheKey(testID string) string {
	return fmt.Sprintf(TestInfoKeyFmt, testID)
}

func getStudentTestOverviewCacheKey(testID string, studentID string) string {
	return fmt.Sprintf(StudentTestOverviewKeyFmt, testID, studentID)
}

func getStudentTestPermitKey(testID string, studentID string) string {
	return fmt.Sprintf(StudentTestPermitKeyFmt, testID, studentID)
}
