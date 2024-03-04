package entity

import (
	resourceEnum "github.com/Allen-Career-Institute/common-protos/resource/v1/types/enums"
	enum "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/enums"
	"github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"github.com/Allen-Career-Institute/common-protos/test_and_assessment_service/admin/v1/reply"
	v1 "github.com/Allen-Career-Institute/common-protos/test_and_assessment_service/admin/v1/reply"
	request "github.com/Allen-Career-Institute/common-protos/test_and_assessment_service/admin/v1/request"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

// TestInfoEntity if adding any new fields need to take care of changes in get test and update test and also in bff
type TestInfoEntity struct {
	TestID             string             `bson:"testId" json:"test_id"`
	QuestionPaperID    string             `bson:"questionPaperId" json:"question_paper_id"`
	WorkflowID         string             `bson:"workflowId" json:"workflow_id"`
	DisplayName        string             `bson:"displayName" json:"display_name"`
	ScheduleID         string             `bson:"scheduleId" json:"schedule_id"`
	Assignment         Assignment         `bson:"assignment" json:"assignment"`
	Schedule           Schedule           `bson:"schedule" json:"schedule"`
	Mode               string             `bson:"mode" json:"mode"`
	Stream             string             `bson:"stream" json:"stream"`
	Class              string             `bson:"class" json:"class"`
	Category           string             `bson:"category" json:"category"`
	MetaData           MetaData           `bson:"metadata" json:"metadata"`
	Status             string             `bson:"status" json:"status"`
	QuestionSets       QuestionSets       `bson:"questionSets" json:"question_sets"`
	QuestionPaperStats QuestionPaperStats `bson:"questionPaperStats" json:"question_paper_stats"`
	CreatedBy          string             `bson:"createdBy" json:"created_by"`
	UpdatedBy          string             `bson:"updatedBy" json:"updated_by"`
	CreatedAt          time.Time          `bson:"createdAt" json:"created_at"`
	UpdatedAt          time.Time          `bson:"updatedAt" json:"updated_at"`
	Pattern            string             `bson:"pattern" json:"pattern"`
	Languages          []string           `bson:"languages" json:"languages"`
	Type               string             `bson:"type" json:"type"`
	PaperCode          string             `bson:"paperCode" json:"paper_code"`
	Syllabus           []*Syllabus        `bson:"syllabus" json:"syllabus"`
	SyllabusPdfLink    string             `bson:"syllabusPdfLink" json:"syllabus_pdf_link"`
	IsMock             bool               `bson:"isMock" json:"is_mock"`
	GenerateRanks      []string           `bson:"generateRanks" json:"generate_ranks"`
	TestSetting        *TestSetting       `bson:"TestSetting" json:"test_setting"`
}

type TestSetting struct {
	ShowCalculator       bool          `bson:"showCalculator" json:"show_calculator"`
	EnableSubmitAfterMin time.Duration `bson:"EnableSubmitAfterMin" json:"enable_submit_after_min"`
}

type QuestionPaperStats struct {
	TotalQuestions int32      `bson:"totalQuestions" json:"total_questions"`
	Sections       []Sections `bson:"sections" json:"sections"`
}

type Sections struct {
	SectionName    string `bson:"sectionName" json:"section_name"`
	TotalQuestions int32  `bson:"totalQuestions" json:"total_questions"`
	Namespace      string `bson:"namespace" json:"namespace"`
}

type QuestionSets struct {
	NoOfSets int32    `bson:"noOfSets" json:"no_of_sets"`
	Sets     []string `bson:"sets" json:"sets"`
}

type MetaData struct {
	SubCategory                     string                 `bson:"subCategory" json:"sub_category"`
	MaxAttemptableQuestionInSection []MaxQuestionInSection `bson:"maxAttemptableQuestionInSection" json:"max_attemptable_question_in_section"`
}

type MaxQuestionInSection struct {
	SectionNamespace       string `bson:"sectionNamespace" json:"section_namespace"`
	MaxAttemptableQuestion int32  `bson:"maxAttemptableQuestion" json:"max_attemptable_question"`
}

type Schedule struct {
	Type             string        `bson:"type" json:"type"`
	StartTime        time.Time     `bson:"startTime" json:"start_time"`
	EndTime          time.Time     `bson:"endTime" json:"end_time"`
	DurationInMin    time.Duration `bson:"durationInMin" json:"duration_in_min"`
	LoginWindowInMin time.Duration `bson:"loginWindowInMin" json:"login_window_in_min"`
	Extension        time.Duration `bson:"extension" json:"extension"`
}

type Duration struct {
	Value int32  `bson:"value" json:"value"`
	Unit  string `bson:"unit" json:"unit"`
}

type Assignment struct {
	Batches  []*Batch `bson:"batches" json:"batches"`
	Centers  []string `bson:"centers" json:"centers"`
	Students []string `bson:"students" json:"students"`
}

type Batch struct {
	BatchCode string `bson:"batchCode" json:"batch_code"`
	BatchID   string `bson:"batchID" json:"batch_id"`
}

type Syllabus struct {
	GroupNodeID string            `bson:"groupNodeId" json:"group_node_id"`
	NodeIds     []string          `bson:"nodeIds" json:"node_ids"`
	Metadata    *SyllabusMetadata `bson:"metadata" json:"metadata"`
	TaxonomyID  string            `bson:"taxonomyId" json:"taxonomy_id"`
}

type SyllabusMetadata struct {
	Notes string `bson:"notes" json:"notes"`
}

func (entity *TestInfoEntity) FromPB(req *request.CreateTestRequest) {
	entity.DisplayName = req.GetDisplayName()
	entity.Schedule = getSchedule(req.Schedule)
	entity.Class = req.GetClass().String()
	entity.Stream = req.GetStream().String()
	entity.Mode = req.GetMode().String()
	entity.Category = req.GetCategory().String()
	entity.MetaData = MetaData{}

	if req.GetMetaData() != nil {
		entity.MetaData.SubCategory = req.GetMetaData().GetSubCategory()
		maxQuestionInSectionList := make([]MaxQuestionInSection, 0, len(req.MetaData.MaxAttemptableQuestionInSection))
		for _, item := range req.MetaData.MaxAttemptableQuestionInSection {
			maxQuestionInSection := MaxQuestionInSection{
				SectionNamespace:       item.SectionNamespace,
				MaxAttemptableQuestion: item.MaxAttemptableQuestion,
			}
			maxQuestionInSectionList = append(maxQuestionInSectionList, maxQuestionInSection)
		}
		entity.MetaData.MaxAttemptableQuestionInSection = maxQuestionInSectionList
	}
	// assignment
	entity.Assignment = Assignment{
		Centers:  req.Assignment.Centers,
		Students: req.Assignment.Students,
	}
	batchList := make([]*Batch, 0, len(req.Assignment.Batches))
	for _, item := range req.Assignment.Batches {
		batch := &Batch{
			BatchCode: item.BatchCode,
			BatchID:   item.BatchId,
		}
		batchList = append(batchList, batch)
	}
	entity.Assignment.Batches = batchList
	entity.CreatedAt = time.Now()
	entity.UpdatedAt = time.Now()
	entity.Languages = req.GetLanguages()
	entity.Pattern = req.GetPattern()
	entity.Type = req.GetType().String()
	entity.CreatedBy = req.GetCreatedBy()
	entity.PaperCode = req.GetPaperCode()
	entity.Syllabus = getSyllabusFromPB(req.Syllabus)
	entity.SyllabusPdfLink = req.SyllabusPdfLink
	entity.IsMock = req.IsMock
}

func getSchedule(inputSchedule *types.TestSchedule) Schedule {
	schedule := Schedule{}
	startTime := time.Unix(inputSchedule.StartTime.Seconds, 0)
	schedule.StartTime = startTime.UTC()
	if inputSchedule.EndTime != nil {
		endTime := time.Unix(inputSchedule.EndTime.Seconds, 0)
		schedule.EndTime = endTime.UTC()
	}
	schedule.DurationInMin = time.Duration(inputSchedule.DurationInMin)
	schedule.LoginWindowInMin = time.Duration(inputSchedule.LoginWindowInMin)
	schedule.Extension = time.Duration(inputSchedule.ExtensionInMinutes)
	schedule.Type = inputSchedule.GetType().String()
	return schedule
}

func (entity *TestInfoEntity) ToPB(info *reply.GetTestReply) {
	info.TestId = entity.TestID
	info.QuestionPaperId = entity.QuestionPaperID
	info.DisplayName = entity.DisplayName
	info.PaperCode = entity.PaperCode
	info.ScheduleId = entity.ScheduleID
	info.Assignment = &types.TestAssignment{
		Centers:  entity.Assignment.Centers,
		Students: entity.Assignment.Students,
	}
	batchList := make([]*types.Batch, 0, len(entity.Assignment.Batches))
	for _, item := range entity.Assignment.Batches {
		batch := &types.Batch{
			BatchId:   item.BatchID,
			BatchCode: item.BatchCode,
		}
		batchList = append(batchList, batch)
	}
	info.Assignment.Batches = batchList

	info.Schedule = &types.TestSchedule{
		//Type: enum.TestScheduleType(enum.TestScheduleType_value[entity.Schedule.Type]),
		StartTime: &timestamp.Timestamp{
			Seconds: entity.Schedule.StartTime.Unix(),
		},
		EndTime: &timestamp.Timestamp{
			Seconds: entity.Schedule.EndTime.Unix(),
		},
		LoginWindowInMin:   int64(entity.Schedule.LoginWindowInMin),
		DurationInMin:      int64(entity.Schedule.DurationInMin),
		ExtensionInMinutes: int64(entity.Schedule.Extension),
		Type:               enum.TestScheduleType(enum.TestScheduleType_value[entity.Schedule.Type]),
	}
	info.Mode = enum.TestMode(enum.TestMode_value[entity.Mode])
	info.Languages = entity.Languages
	info.Pattern = entity.Pattern
	info.Type = enum.TestType(resourceEnum.TestType_value[entity.Type])
	info.Class = resourceEnum.Class(resourceEnum.Class_value[entity.Class])
	info.Stream = resourceEnum.Stream(resourceEnum.Stream_value[entity.Stream])
	info.Category = enum.TestCategory(enum.TestCategory_value[entity.Category])

	maxQuestionInSectionList := make([]*types.MaxQuestionInSection, 0, len(entity.MetaData.MaxAttemptableQuestionInSection))
	for _, item := range entity.MetaData.MaxAttemptableQuestionInSection {
		maxQuestionInSection := &types.MaxQuestionInSection{
			SectionNamespace:       item.SectionNamespace,
			MaxAttemptableQuestion: item.MaxAttemptableQuestion,
		}
		maxQuestionInSectionList = append(maxQuestionInSectionList, maxQuestionInSection)
	}

	info.MetaData = &types.TestMetaData{
		SubCategory:                     entity.MetaData.SubCategory,
		MaxAttemptableQuestionInSection: maxQuestionInSectionList,
	}
	info.QuestionSets = &v1.QuestionSets{
		NoOfSets: entity.QuestionSets.NoOfSets,
		Sets:     entity.QuestionSets.Sets,
	}

	info.Status = enum.Status(enum.Status_value[entity.Status])
	info.CreatedBy = entity.CreatedBy
	info.UpdatedBy = entity.UpdatedBy
	info.CreatedAt = &timestamp.Timestamp{
		Seconds: entity.CreatedAt.Unix(),
	}
	info.UpdatedAt = &timestamp.Timestamp{
		Seconds: entity.UpdatedAt.Unix(),
	}

	var qPSSections []*v1.Sections
	for _, section := range entity.QuestionPaperStats.Sections {
		qPSSections = append(qPSSections,
			&v1.Sections{
				SectionName:    section.SectionName,
				TotalQuestions: section.TotalQuestions,
				Namespace:      section.Namespace,
			})
	}
	info.QuestionPaperStats = &v1.QuestionPaperStats{
		TotalQuestions: entity.QuestionPaperStats.TotalQuestions,
		Sections:       qPSSections,
	}
	info.Syllabus = getSyllabusFromEntity(entity.Syllabus)
	info.SyllabusPdfLink = entity.SyllabusPdfLink
	info.WorkflowId = entity.WorkflowID
	info.PaperCode = entity.PaperCode
	info.IsMock = entity.IsMock
	for _, rank := range entity.GenerateRanks {
		info.GenerateRanks = append(info.GenerateRanks, enum.Rank(enum.Rank_value[rank]))
	}
	if entity.TestSetting != nil {
		info.TestSetting = &reply.TestSetting{
			ShowCalculator:       entity.TestSetting.ShowCalculator,
			EnableSubmitAfterMin: int64(entity.TestSetting.EnableSubmitAfterMin),
		}
	}

}

func getSyllabusFromPB(syllabusRequests []*request.Syllabus) []*Syllabus {
	syllabusDetails := make([]*Syllabus, 0)
	if syllabusRequests == nil || len(syllabusRequests) == 0 {
		return syllabusDetails
	}
	for _, syllabusRequest := range syllabusRequests {
		syllabusDetail := &Syllabus{
			GroupNodeID: syllabusRequest.GroupNodeId,
			NodeIds:     syllabusRequest.NodeIds,
			TaxonomyID:  syllabusRequest.TaxonomyId,
		}
		if syllabusRequest.Metadata != nil {
			syllabusDetail.Metadata = &SyllabusMetadata{Notes: syllabusRequest.Metadata.Notes}
		}
		syllabusDetails = append(syllabusDetails, syllabusDetail)
	}
	return syllabusDetails
}

func getSyllabusFromEntity(syllabusEntity []*Syllabus) []*v1.Syllabus {
	syllabusDetails := make([]*v1.Syllabus, 0)
	if syllabusEntity == nil || len(syllabusEntity) == 0 {
		return syllabusDetails
	}
	for _, detail := range syllabusEntity {
		syllabusDetail := &v1.Syllabus{GroupNodeId: detail.GroupNodeID,
			NodeIds:    detail.NodeIds,
			TaxonomyId: detail.TaxonomyID,
		}
		if detail.Metadata != nil {
			syllabusDetail.Metadata = &v1.SyllabusMetadata{Notes: detail.Metadata.Notes}
		}
		syllabusDetails = append(syllabusDetails, syllabusDetail)
	}
	return syllabusDetails
}

func (entity *TestInfoEntity) FromUpdateTestPBToTestInfoEntity(req *request.UpdateTestRequest) {

	entity.DisplayName = req.GetDisplayName()
	entity.Category = req.GetCategory().String()
	entity.Stream = req.GetStream().String()
	entity.Class = req.GetClass().String()
	entity.Mode = req.GetMode().String()
	entity.MetaData = MetaData{}

	if req.GetMetaData() != nil {
		entity.MetaData.SubCategory = req.GetMetaData().GetSubCategory()
		maxQuestionInSectionList := make([]MaxQuestionInSection, 0, len(req.MetaData.MaxAttemptableQuestionInSection))
		for _, item := range req.MetaData.MaxAttemptableQuestionInSection {
			maxQuestionInSection := MaxQuestionInSection{
				SectionNamespace:       item.SectionNamespace,
				MaxAttemptableQuestion: item.MaxAttemptableQuestion,
			}
			maxQuestionInSectionList = append(maxQuestionInSectionList, maxQuestionInSection)
		}
		entity.MetaData.MaxAttemptableQuestionInSection = maxQuestionInSectionList
	}
	entity.Assignment = Assignment{
		Centers:  req.Assignment.Centers,
		Students: req.Assignment.Students,
	}
	batchList := make([]*Batch, 0, len(req.Assignment.Batches))
	for _, item := range req.Assignment.Batches {
		batch := &Batch{
			BatchCode: item.BatchCode,
			BatchID:   item.BatchId,
		}
		batchList = append(batchList, batch)
	}
	entity.Assignment.Batches = batchList
	entity.Schedule = getSchedule(req.Schedule)
	entity.Languages = req.GetLanguages()
	entity.Pattern = req.GetPattern()
	entity.Type = req.GetType().String()
	entity.PaperCode = req.GetPaperCode()
	entity.Syllabus = getSyllabusFromPB(req.Syllabus)
	entity.SyllabusPdfLink = req.SyllabusPdfLink
	entity.TestID = req.GetTestId()
	entity.QuestionPaperID = req.GetQuestionPaperId()
	entity.WorkflowID = req.GetWorkFlowId()
	entity.ScheduleID = req.ScheduleId
	entity.Status = req.GetStatus().String()
	entity.QuestionSets = QuestionSets{
		Sets:     req.QuestionSets.Sets,
		NoOfSets: req.QuestionSets.NoOfSets,
	}
	var sections []Sections
	for _, section := range req.QuestionPaperStats.Sections {
		sections = append(sections, Sections{
			SectionName:    section.SectionName,
			TotalQuestions: section.TotalQuestions,
			Namespace:      section.Namespace,
		})
	}
	entity.QuestionPaperStats = QuestionPaperStats{
		TotalQuestions: req.QuestionPaperStats.TotalQuestions,
		Sections:       sections,
	}
	entity.UpdatedBy = req.GetUpdatedBy()
	entity.UpdatedAt = time.Now()
	entity.IsMock = req.IsMock
	for _, rank := range req.GenerateRanks {
		entity.GenerateRanks = append(entity.GenerateRanks, rank.String())
	}
	if req.TestSetting != nil {
		entity.TestSetting = &TestSetting{
			ShowCalculator:       req.TestSetting.ShowCalculator,
			EnableSubmitAfterMin: time.Duration(req.TestSetting.EnableSubmitAfterMin),
		}
	}
}

func (schedule *Schedule) CalculateEndTime() time.Time {
	endTime := schedule.StartTime
	endTime = endTime.Add(schedule.LoginWindowInMin * time.Minute)
	endTime = endTime.Add(schedule.DurationInMin * time.Minute)
	endTime = endTime.Add(schedule.Extension * time.Minute)
	return endTime
}

// nolint: cyclop
// TODO Pick up during update
//func (entity *TestInfoEntity) FromUpdatePB(req *request.UpdateTestRequest) {
//	if len(req.DisplayName) > 0 {
//		entity.DisplayName = req.DisplayName
//	}
//	if len(req.ScheduleId) > 0 {
//		entity.ScheduleID = req.ScheduleId
//	}
//
//	if req.Assignment != nil {
//		entity.updateAssignment(req)
//	}
//
//	//if req.Schedule != nil {
//	//	if req.Schedule.Type.Number() != 0 {
//	//		entity.Schedule.Type = req.Schedule.Type.String()
//	//	}
//	//	if req.Schedule.StartTime != nil {
//	//		entity.Schedule.StartTime = time.Unix(req.Schedule.StartTime.Seconds, 0)
//	//	}
//	//	if req.Schedule.EndTime != nil {
//	//		entity.Schedule.EndTime = time.Unix(req.Schedule.EndTime.Seconds, 0)
//	//	}
//	//	if req.Schedule.Duration != nil {
//	//		if req.Schedule.Duration.Unit.Number() != 0 {
//	//			entity.Schedule.Duration.Unit = req.Schedule.Duration.Unit.String()
//	//		}
//	//		if req.Schedule.Duration.Value > 0 {
//	//			entity.Schedule.Duration.Value = req.Schedule.Duration.Value
//	//		}
//	//	}
//	//}
//
//	if req.Stream.Number() != 0 {
//		entity.Stream = req.Stream.String()
//	}
//	if req.Mode.Number() != 0 {
//		entity.Mode = req.Mode.String()
//	}
//	if req.Category.Number() != 0 {
//		entity.Category = req.Category.String()
//	}
//	if req.MetaData != nil {
//		entity.MetaData.SubCategory = req.MetaData.SubCategory
//	}
//	entity.UpdatedBy = req.UpdatedBy
//	entity.UpdatedAt = time.Now()
//}

//func (entity *TestInfoEntity) updateAssignment(req *request.UpdateTestRequest) {
//	if req.Assignment.Batches != nil {
//		entity.Assignment.Batches = removeAll(req.Assignment.Batches.Remove, entity.Assignment.Batches)
//		entity.Assignment.Batches = append(entity.Assignment.Batches, req.Assignment.Batches.Add...)
//	}
//	if req.Assignment.Centers != nil {
//		entity.Assignment.Centers = removeAll(req.Assignment.Centers.Remove, entity.Assignment.Centers)
//		entity.Assignment.Centers = append(entity.Assignment.Centers, req.Assignment.Centers.Add...)
//	}
//	if req.Assignment.Students != nil {
//		entity.Assignment.Students = removeAll(req.Assignment.Students.Remove, entity.Assignment.Students)
//		entity.Assignment.Students = append(entity.Assignment.Students, req.Assignment.Students.Add...)
//	}
//}
//
//func removeAll(removeElements, elements []string) []string {
//	for _, element := range removeElements {
//		x := -1
//		for i, batch := range elements {
//			if element == batch {
//				x = i
//				break
//			}
//		}
//		if x != -1 {
//			elements = remove(elements, x)
//		}
//	}
//	return elements
//}
//
//func remove(s []string, i int) []string {
//	s[i] = s[len(s)-1]
//	return s[:len(s)-1]
//}
