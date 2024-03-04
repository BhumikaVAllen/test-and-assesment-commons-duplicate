package entity

import (
	qbSetType "github.com/Allen-Career-Institute/common-protos/questionbank/v1/questionSets/type"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/constants"
	"google.golang.org/protobuf/types/known/structpb"
	"sort"
	"strings"
)
import qbQType "github.com/Allen-Career-Institute/common-protos/questionbank/v1/questions/type"
import pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_engine/v1/types"

type QuestionSet struct {
	QuestionSetID         string                 `json:"question_set_id,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	PaperCode             string                 `json:"paper_code,omitempty"`
	MaxMarks              float32                `json:"max_marks,omitempty"`
	TotalTime             *Time                  `json:"total_time,omitempty"`
	TestDate              int64                  `json:"test_date,omitempty"`
	Phase                 string                 `json:"phase,omitempty"`
	Languages             []string               `json:"languages,omitempty"`
	IsMultilingual        bool                   `json:"is_multilingual"`
	PaperUniqueIdentifier string                 `json:"paper_unique_identifier,omitempty"`
	Stream                string                 `json:"stream,omitempty"`
	Pattern               string                 `json:"pattern,omitempty"`
	TestType              string                 `json:"test_type,omitempty"`
	Session               string                 `json:"session,omitempty"`
	NumberOfQuestions     int32                  `json:"number_of_questions,omitempty"`
	Questions             []*QuestionSetQuestion `json:"questions,omitempty"`
	Sections              []*QuestionSetSection  `json:"sections,omitempty"`
	Metadata              *Metadata              `json:"test_metadata,omitempty"`
}

type Metadata struct {
	TestStartDate string `json:"test_start_date"`
}

type Question struct {
	QuestionID   string     `json:"question_id,omitempty"`
	Type         string     `json:"type,omitempty"`
	Content      []*Content `json:"content,omitempty"`
	QuestionType string     `json:"QuestionType,omitempty"`
}

type QuestionSetQuestion struct {
	Questions     *Question      `json:"question_detail,omitempty"`
	MarkingScheme *MarkingScheme `json:"marking_scheme,omitempty"`
	//QuestionSolution *QuestionSolution `json:"questionSolution,omitempty"`
	SequenceNo int32 `json:"sequence_no,omitempty"`
}

type QuestionSetSection struct {
	Name              string                 `json:"name,omitempty"`
	TopicList         []string               `json:"topic_list,omitempty"`
	HaveSubsections   bool                   `json:"have_subsections"`
	MaxMarks          float32                `json:"max_marks,omitempty"`
	Subsections       []*QuestionSetSection  `json:"subsections,omitempty"`
	NumberOfQuestions int32                  `json:"number_of_questions,omitempty"`
	Questions         []*QuestionSetQuestion `json:"questions,omitempty"`
	Instructions      []*Instruction         `json:"instructions,omitempty"`
	SectionID         int64                  `json:"section_id,omitempty"`
	Namespace         string                 `json:"namespace"`
}

type MarkingScheme struct {
	MarkingSchemeID    string                       `json:"marking_scheme_id,omitempty"`
	PartialMarkingType *PartialCorrectMarkingScheme `json:"partial_marking_type,omitempty"`
	Marks              float32                      `json:"marks,omitempty"`
	NegMarks           float32                      `json:"negMarks,omitempty"`
	SequenceNo         int32                        `json:"sequence_no,omitempty"`
}

type PartialCorrectMarkingScheme struct {
	OptionSequenceID int32   `json:"option_sequence_id,omitempty"`
	CorrectMarks     float32 `json:"correct_marks,omitempty"`
	NegativeMarks    float32 `json:"negative_marks,omitempty"`
}

type Content struct {
	Language string `json:"language,omitempty"`
	IsGroup  bool   `json:"isGroup"`
	GroupID  string `json:"group_id,omitempty"`
	//	Answer     string        `json:"answer,omitempty"`
	QnsContent          *QuestionStem `json:"qns_content,omitempty"`
	Options             []*Option     `json:"options,omitempty"`
	Nature              string        `json:"nature,omitempty"`
	QuestionLanguageStr string        `json:"QuestionLanguageStr,omitempty"`
	QuestionNatureStr   string        `json:"QuestionNatureStr,omitempty"`
	MatchOptions        *MatchOptions `json:"match_options,omitempty"`
}

type QuestionStem struct {
	Text  string   `json:"text,omitempty"`
	Media []*Media `json:"media,omitempty"`
}

type QuestionSolution struct {
	QuestionID     string           `json:"question_id,omitempty"`
	TextSolutions  []*TextSolution  `json:"text_solutions,omitempty"`
	VideoSolutions []*VideoSolution `json:"video_solutions,omitempty"`
}

type TextSolution struct {
	Language string   `json:"language,omitempty"`
	Text     string   `json:"text,omitempty"`
	Media    []*Media `json:"media,omitempty"`
}

type VideoSolution struct {
	Language  string `json:"language,omitempty"`
	VideoPath string `json:"videoPath,omitempty"`
}

type Option struct {
	Text        string   `json:"text,omitempty"`
	Explanation string   `json:"explanation,omitempty"`
	SequenceID  int32    `json:"sequence_id,omitempty"`
	Media       []*Media `json:"media,omitempty"`
}

type MatchOptions struct {
	Row   string `json:"row,omitempty"`
	Cols  string `json:"cols,omitempty"`
	Label Label  `json:"label,omitempty"`
}

type Label struct {
	Row string `json:"row,omitempty"`
	Col string `json:"col,omitempty"`
}
type Media struct {
	MediaID     string `json:"media_id,omitempty"`
	MediaPath   string `json:"media_path,omitempty"`
	MediaType   string `json:"media_type,omitempty"`
	LatexSource string `json:"latex_source,omitempty"`
}

type Instruction struct {
	Language string `json:"language,omitempty"`
	Text     string `json:"text,omitempty"`
}

type Time struct {
	Value int32  `json:"value,omitempty"`
	Unit  string `json:"unit,omitempty"`
}

func (questionSet *QuestionSet) FromPB(qQuestionSet *qbSetType.QuestionSetInformation) {
	questionSetLanguages := qQuestionSet.Languages
	questionSet.QuestionSetID = qQuestionSet.QuestionSetId
	questionSet.Name = qQuestionSet.Name
	questionSet.PaperCode = qQuestionSet.PaperCode
	questionSet.MaxMarks = qQuestionSet.MaxMarks

	time := &Time{}
	time.fromPB(qQuestionSet.TotalTime)
	questionSet.TotalTime = time

	questionSet.TestDate = qQuestionSet.TestDate
	questionSet.Phase = qQuestionSet.Phase
	questionSet.Languages = qQuestionSet.Languages
	questionSet.IsMultilingual = qQuestionSet.IsMultilingual
	questionSet.PaperUniqueIdentifier = qQuestionSet.PaperUniqueIdentifier
	questionSet.Stream = qQuestionSet.Stream
	questionSet.Pattern = qQuestionSet.Pattern
	questionSet.TestType = qQuestionSet.TestType
	questionSet.Session = qQuestionSet.Session
	questionSet.NumberOfQuestions = qQuestionSet.NumberOfQuestions

	questionSet.Questions = createQuestionSetQuestionList(qQuestionSet.Questions, questionSetLanguages)
	var questionSetSectionList []*QuestionSetSection
	for _, qS := range qQuestionSet.Sections {
		questionSetSection := &QuestionSetSection{}
		questionSetSection.fromPB(qS, questionSetLanguages)
		questionSetSectionList = append(questionSetSectionList, questionSetSection)
	}
	questionSet.Sections = questionSetSectionList
}

func (questionSetSection *QuestionSetSection) fromPB(qSection *qbSetType.QuestionSetSection, questionSetLanguages []string) {
	questionSetSection.Name = qSection.Name
	questionSetSection.TopicList = qSection.TopicList
	questionSetSection.HaveSubsections = qSection.HaveSubsections
	questionSetSection.MaxMarks = qSection.MaxMarks
	questionSetSection.NumberOfQuestions = qSection.NumberOfQuestions

	var questionSetSectionList []*QuestionSetSection
	for _, qS := range qSection.Subsections {
		section := &QuestionSetSection{}
		section.fromPB(qS, questionSetLanguages)
		questionSetSectionList = append(questionSetSectionList, section)
	}
	questionSetSection.Subsections = questionSetSectionList
	questionSetSection.Questions = createQuestionSetQuestionList(qSection.Questions, questionSetLanguages)

	var instructionList []*Instruction
	for _, qI := range qSection.Instructions {
		isContains := ContainsQuestionSetLanguage(questionSetLanguages, qI.Language)
		if !isContains {
			continue
		}
		instruction := &Instruction{}
		instruction.FromPB(qI)
		instructionList = append(instructionList, instruction)
	}
	questionSetSection.Instructions = instructionList
	questionSetSection.SectionID = qSection.SectionId
	questionSetSection.Namespace = qSection.Namespace
}

func (instruction *Instruction) FromPB(qInstruction *qbSetType.Instruction) {
	instruction.Text = qInstruction.Text
	instruction.Language = qInstruction.Language
}

func (time *Time) fromPB(qTime *qbSetType.Time) {
	time.Value = qTime.Value
	time.Unit = qTime.Unit.String()
}

func (questionSetQuestion *QuestionSetQuestion) fromPB(qQuestion *qbSetType.QuestionSetQuestion, questionSetLanguages []string) {
	question := &Question{}
	question.fromPB(qQuestion.QuestionDetail, questionSetLanguages)
	markingScheme := &MarkingScheme{}
	markingScheme.fromPB(qQuestion.MarkingScheme)

	questionSetQuestion.Questions = question
	questionSetQuestion.MarkingScheme = markingScheme
	questionSetQuestion.SequenceNo = qQuestion.SequenceNo
}

func (markingScheme *MarkingScheme) fromPB(qMarkingScheme *qbSetType.MarkingScheme) {
	partialCorrectMarkingScheme := &PartialCorrectMarkingScheme{}

	if qMarkingScheme.PartialMarkingType != nil {
		partialCorrectMarkingScheme.fromPB(qMarkingScheme.PartialMarkingType)
		markingScheme.PartialMarkingType = partialCorrectMarkingScheme
	}

	markingScheme.MarkingSchemeID = qMarkingScheme.MarkingSchemeId // Need to check
	markingScheme.Marks = qMarkingScheme.Marks
	markingScheme.NegMarks = qMarkingScheme.NegMarks
	markingScheme.SequenceNo = qMarkingScheme.SequenceNumber // Need to check
}

func (partialCorrectMarkingScheme *PartialCorrectMarkingScheme) fromPB(qPartialCorrectMarkingScheme *qbSetType.PartialCorrectMarkingScheme) {
	partialCorrectMarkingScheme.OptionSequenceID = qPartialCorrectMarkingScheme.OptionSequenceId // Need to check
	partialCorrectMarkingScheme.CorrectMarks = qPartialCorrectMarkingScheme.CorrectMarks
	partialCorrectMarkingScheme.NegativeMarks = qPartialCorrectMarkingScheme.NegativeMarks
}

func (question *Question) fromPB(qQuestion *qbQType.QuestionInformation, questionSetLanguages []string) {
	var contentList []*Content
	for _, qC := range qQuestion.Content {
		isContains := ContainsQuestionSetLanguage(questionSetLanguages, qC.Language.String())
		if !isContains {
			continue
		}
		content := &Content{}
		content.fromPB(qC)
		contentList = append(contentList, content)
	}
	question.QuestionID = qQuestion.QuestionId
	question.Type = qQuestion.Type.String()
	question.Content = contentList
}

func (matchOptions *MatchOptions) fromPB(opt *qbQType.MatchOptions) {
	matchOptions.Row = opt.Row
	matchOptions.Cols = opt.Cols
	matchOptions.Label.Row = constants.DefaultLabelRow
	matchOptions.Label.Col = constants.DefaultLabelCol
	if opt.Label != nil {
		switch v := opt.Label.GetKind().(type) {
		case *structpb.Value_StringValue:
			matchOptions.Label.Col = v.StringValue

		case *structpb.Value_StructValue:
			if val, ok := opt.Label.GetStructValue().Fields[constants.LabelRowKey]; ok && val.GetStringValue() != "" {
				matchOptions.Label.Row = val.GetStringValue()
			}
			if val, ok := opt.Label.GetStructValue().Fields[constants.LabelColKey]; ok && val.GetStringValue() != "" {
				matchOptions.Label.Col = val.GetStringValue()
			}
		}
	}
}
func (content *Content) fromPB(qContent *qbQType.Content) {
	var optionList []*Option
	for _, qO := range qContent.Options {
		option := &Option{}
		option.fromPB(qO)
		optionList = append(optionList, option)
	}

	if qContent.MatchOptions != nil {
		content.MatchOptions = &MatchOptions{}
		content.MatchOptions.fromPB(qContent.MatchOptions)
	}
	questionStem := &QuestionStem{}
	questionStem.fromPB(qContent.QnsContent)

	// TODO getAnswer from here qContent.Answer
	content.Language = qContent.Language.String()
	content.IsGroup = qContent.IsGroup
	content.GroupID = qContent.GroupId
	content.QnsContent = questionStem
	content.Options = optionList
	content.QuestionNatureStr = qContent.Nature.String()
	content.QuestionLanguageStr = qContent.Language.String()
	content.Nature = qContent.Nature.String()
}

func (option *Option) fromPB(qOption *qbQType.Option) {
	var mediaList []*Media
	for _, qM := range qOption.Media {
		media := &Media{}
		media.FromPB(qM)
		mediaList = append(mediaList, media)
	}
	option.Text = qOption.Text
	option.Explanation = qOption.Explanation
	option.SequenceID = qOption.SequenceId
	option.Media = mediaList
}

func (questionStem *QuestionStem) fromPB(qQuestionStem *qbQType.QuestionStem) {
	var mediaList []*Media
	for _, qM := range qQuestionStem.Media {
		media := &Media{}
		media.FromPB(qM)
		mediaList = append(mediaList, media)
	}
	questionStem.Text = qQuestionStem.Text
	questionStem.Media = mediaList
}

func (media *Media) FromPB(qMedia *qbQType.Media) {
	media.MediaID = qMedia.MediaId
	media.MediaPath = qMedia.MediaPath
	media.MediaType = qMedia.MediaType
	media.LatexSource = qMedia.LatexSource
}

/****************************************************ToPB****************************************************/

func (entity *QuestionSet) ToPbt(rQP *pbTypes.RenderableQuestionPaper) {
	entity.setTopLevelDetails(rQP)

	// Question Details
	if entity.Questions != nil {
		questions := make([]*pbTypes.QuestionDetails, 0, len(entity.Questions))
		for _, qpEntity := range entity.Questions {
			questionDetail := &pbTypes.QuestionDetails{}
			qpEntity.toPbt(questionDetail)
			questions = append(questions, questionDetail)
		}
		sort.Slice(questions, func(i, j int) bool {
			return questions[i].SequenceNo < questions[j].SequenceNo
		})
		rQP.Questions = questions
	}

	// section Details  subsections to be covered
	sections := make([]*pbTypes.QuestionsSection, 0, len(entity.Sections))
	for _, sectionEntity := range entity.Sections {
		section := &pbTypes.QuestionsSection{}
		section.SectionId = sectionEntity.SectionID
		section.Name = sectionEntity.Name
		section.HaveSubsections = sectionEntity.HaveSubsections
		section.MaxMarks = sectionEntity.MaxMarks
		section.NumberOfQuestions = sectionEntity.NumberOfQuestions
		section.TopicList = sectionEntity.TopicList
		section.Namespace = sectionEntity.Namespace

		instructionsList := make([]*pbTypes.Instructions, 0, len(sectionEntity.Instructions))
		for _, item := range sectionEntity.Instructions {
			instruction := &pbTypes.Instructions{
				Language: item.Language,
				Text:     item.Text,
			}
			instructionsList = append(instructionsList, instruction)
		}
		section.Instructions = instructionsList

		sectionalQuestions := make([]*pbTypes.QuestionDetails, 0, len(sectionEntity.Questions))
		for _, qpEntity := range sectionEntity.Questions {
			questionDetail := &pbTypes.QuestionDetails{}
			qpEntity.toPbt(questionDetail)
			sectionalQuestions = append(sectionalQuestions, questionDetail)
		}
		sort.Slice(sectionalQuestions, func(i, j int) bool {
			return sectionalQuestions[i].SequenceNo < sectionalQuestions[j].SequenceNo
		})
		section.Questions = sectionalQuestions
		sections = append(sections, section)
	}
	rQP.Sections = sections
}

func (entity *QuestionSet) setTopLevelDetails(rQP *pbTypes.RenderableQuestionPaper) {
	rQP.QuestionPaperId = entity.QuestionSetID
	rQP.Name = entity.Name
	rQP.PaperCode = entity.PaperCode

	if entity.TotalTime != nil {
		time := &qbSetType.Time{}
		entity.TotalTime.toPbt(time)
		rQP.TotalTime = time
	}
	rQP.TestDate = entity.TestDate
	rQP.Phase = entity.Phase
	rQP.MaxMarks = entity.MaxMarks
	rQP.Session = entity.Session
	rQP.Languages = entity.Languages
	rQP.IsMultilingual = entity.IsMultilingual
	rQP.PaperUniqueIdentifier = entity.PaperUniqueIdentifier
	rQP.Stream = entity.Stream
	rQP.Pattern = entity.Pattern
	rQP.TestType = entity.TestType
	rQP.Session = entity.Session
	rQP.NumberOfQuestions = entity.NumberOfQuestions
	if entity.Metadata != nil {
		rQP.Metadata = &pbTypes.Metadata{
			TestDate: entity.Metadata.TestStartDate,
		}
	}
}

func (entity *Time) toPbt(time *qbSetType.Time) {
	time.Value = entity.Value
	val, ok := qbSetType.Unit_value[entity.Unit]
	if ok {
		time.Unit = qbSetType.Unit(val)
	}
}

func (entity *QuestionSetQuestion) toPbt(questionSetQuestionInfo *pbTypes.QuestionDetails) {
	questionInfo := &pbTypes.Question{}
	questionInfo.QuestionId = entity.Questions.QuestionID
	val, ok := qbQType.QuestionType_value[entity.Questions.Type]
	if ok {
		questionInfo.Type = qbQType.QuestionType(val)
	}
	questionInfo.QuestionType = questionInfo.Type.String()
	questionContentInfoList := make([]*pbTypes.Content, 0, len(entity.Questions.Content))
	for _, questionContentEntity := range entity.Questions.Content {
		questionContentInfo := &pbTypes.Content{}
		questionContentEntity.toPbt(questionContentInfo)
		//TODO : Add options conversion to content.ToPB
		optionInfoList := make([]*pbTypes.Option, 0, len(questionContentEntity.Options))
		for _, optionEntity := range questionContentEntity.Options {
			optionInfo := &pbTypes.Option{}
			optionEntity.toPbt(optionInfo)
			optionInfoList = append(optionInfoList, optionInfo)
		}
		if questionContentEntity.MatchOptions != nil {
			questionContentInfo.MatchOptions = &pbTypes.MatchOptions{
				Row:  questionContentEntity.MatchOptions.Row,
				Cols: questionContentEntity.MatchOptions.Cols,
				Label: &pbTypes.Label{
					Row: questionContentEntity.MatchOptions.Label.Row,
					Col: questionContentEntity.MatchOptions.Label.Col,
				},
			}
		}
		questionContentInfo.Options = optionInfoList
		questionContentInfoList = append(questionContentInfoList, questionContentInfo)
	}
	questionInfo.Content = questionContentInfoList
	questionSetQuestionInfo.Question = questionInfo
	markingScheme := &pbTypes.QuestionMarkingScheme{}
	entity.MarkingScheme.toPbt(markingScheme)
	questionSetQuestionInfo.MarkingScheme = markingScheme
	questionSetQuestionInfo.SequenceNo = entity.SequenceNo
}

func (entity *MarkingScheme) toPbt(content *pbTypes.QuestionMarkingScheme) {
	content.Marks = entity.Marks
	content.NegMarks = entity.NegMarks
	content.SequenceNumber = entity.SequenceNo

	partialMarkingType := &qbSetType.PartialCorrectMarkingScheme{}
	entity.PartialMarkingType.toPbt(partialMarkingType)
	content.PartialMarkingType = partialMarkingType
}

func (entity *Content) toPbt(content *pbTypes.Content) {
	content.IsGroup = entity.IsGroup
	content.GroupId = entity.GroupID
	val, ok := qbQType.QuestionLanguage_value[entity.Language]
	if ok {
		content.Language = qbQType.QuestionLanguage(val)
	}
	content.QuestionNature = content.Nature.String()
	content.QuestionLanguage = content.Language.String()
	qnsStem := &pbTypes.QuestionStem{}
	entity.QnsContent.toPbt(qnsStem)
	content.QnsContent = qnsStem
}

func (entity *QuestionStem) toPbt(questionStem *pbTypes.QuestionStem) {
	questionStem.Text = entity.Text

	questionStemMediaList := make([]*pbTypes.Media, 0, len(entity.Media))
	for _, mediaEntity := range entity.Media {
		questionStemMedia := &pbTypes.Media{}
		mediaEntity.toPbt(questionStemMedia)
		questionStemMediaList = append(questionStemMediaList, questionStemMedia)
	}
	questionStem.Media = questionStemMediaList
}

func (entity *Media) toPbt(media *pbTypes.Media) {
	media.MediaPath = entity.MediaPath
	media.MediaId = entity.MediaID
	media.MediaType = entity.MediaType
	media.LatexSource = entity.LatexSource
}

func (entity *Option) toPbt(option *pbTypes.Option) {
	option.Text = entity.Text
	option.Explanation = entity.Explanation
	option.SequenceId = entity.SequenceID
	optionsMediaList := make([]*pbTypes.Media, 0, len(entity.Media))
	for _, mediaEntity := range entity.Media {
		optionMedia := &pbTypes.Media{}
		mediaEntity.toPbt(optionMedia)
		optionsMediaList = append(optionsMediaList, optionMedia)
	}
	option.Media = optionsMediaList
}

func (entity *PartialCorrectMarkingScheme) toPbt(info *qbSetType.PartialCorrectMarkingScheme) {
	if entity != nil {
		info.OptionSequenceId = entity.OptionSequenceID
		info.CorrectMarks = entity.CorrectMarks
		info.NegativeMarks = entity.NegativeMarks
	}
}

func createQuestionSetQuestionList(questions []*qbSetType.QuestionSetQuestion, questionSetLanguages []string) []*QuestionSetQuestion {
	var questionSetQuestionList []*QuestionSetQuestion
	for _, item := range questions {
		questionSetQuestion := &QuestionSetQuestion{}
		questionSetQuestion.fromPB(item, questionSetLanguages)
		questionSetQuestionList = append(questionSetQuestionList, questionSetQuestion)
	}
	sort.Slice(questionSetQuestionList, func(i, j int) bool {
		return questionSetQuestionList[i].SequenceNo < questionSetQuestionList[j].SequenceNo
	})

	return questionSetQuestionList
}
func ContainsQuestionSetLanguage(qSetLanguages []string, value string) bool {
	for _, item := range qSetLanguages {
		if strings.ToLower(item) == strings.ToLower(value) {
			return true
		}
	}
	return false
}
