package entity

import (
	qbSetType "github.com/Allen-Career-Institute/common-protos/questionbank/v1/questionSets/type"
	qbQType "github.com/Allen-Career-Institute/common-protos/questionbank/v1/questions/type"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuestionSet_FromPB(t *testing.T) {
	type args struct {
		req *qbSetType.QuestionSetInformation
	}

	tests := []struct {
		name          string
		args          args
		want          *QuestionSet
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test QuestionSet Entity From PB",
			args: args{
				req: &qbSetType.QuestionSetInformation{
					Instructions: []*qbSetType.Instruction{getqInstructions()},
					Questions:    []*qbSetType.QuestionSetQuestion{getqQuestionSetQuestion()},
					Sections:     []*qbSetType.QuestionSetSection{getqQuestionSetSection()},
					TotalTime: &qbSetType.Time{
						Value: 1,
						Unit:  qbSetType.Unit_HOUR,
					},
					Languages: []string{"ENGLISH"},
				},
			},
			want: &QuestionSet{
				Questions: []*QuestionSetQuestion{getQuestionSetQuestion()},
				Sections:  []*QuestionSetSection{getQuestionSetSection()},
				TotalTime: &Time{
					Value: 1,
					Unit:  "HOUR",
				},
				Languages: []string{"ENGLISH"},
			},
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &QuestionSet{}
			entity.FromPB(tt.args.req)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func TestQuestionSetSection_FromPB(t *testing.T) {
	type args struct {
		req                  *qbSetType.QuestionSetSection
		questionSetLanguages []string
	}

	tests := []struct {
		name          string
		args          args
		want          *QuestionSetSection
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test QuestionSetSection Entity From PB",
			args: args{
				req: &qbSetType.QuestionSetSection{
					Instructions: []*qbSetType.Instruction{getqInstructions()},
					Questions:    []*qbSetType.QuestionSetQuestion{getqQuestionSetQuestion()},
					Subsections:  []*qbSetType.QuestionSetSection{getqQuestionSetSection()},
				},
				questionSetLanguages: []string{
					"ENGLISH",
				},
			},
			want: &QuestionSetSection{
				Instructions: []*Instruction{getInstructions()},
				Questions:    []*QuestionSetQuestion{getQuestionSetQuestion()},
				Subsections:  []*QuestionSetSection{getQuestionSetSection()},
			},
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &QuestionSetSection{}
			entity.fromPB(tt.args.req, tt.args.questionSetLanguages)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func getqQuestionSetSection() *qbSetType.QuestionSetSection {
	return &qbSetType.QuestionSetSection{
		Instructions: []*qbSetType.Instruction{getqInstructions()},
		Questions:    []*qbSetType.QuestionSetQuestion{getqQuestionSetQuestion()},
	}
}

func getQuestionSetSection() *QuestionSetSection {
	return &QuestionSetSection{
		Instructions: []*Instruction{getInstructions()},
		Questions:    []*QuestionSetQuestion{getQuestionSetQuestion()},
	}
}

func TestMarkingScheme_FromPB(t *testing.T) {
	type args struct {
		req *qbSetType.MarkingScheme
	}

	tests := []struct {
		name          string
		args          args
		want          *MarkingScheme
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test MarkingScheme Entity From PB",
			args: args{
				req: getqMarkingScheme(),
			},
			want:          getMarkingScheme(),
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &MarkingScheme{}
			entity.fromPB(tt.args.req)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func TestQuestionSetQuestion_FromPB(t *testing.T) {
	type args struct {
		req                  *qbSetType.QuestionSetQuestion
		questionSetLanguages []string
	}

	tests := []struct {
		name          string
		args          args
		want          *QuestionSetQuestion
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test QuestionSetQuestion Entity From PB",
			args: args{
				req: getqQuestionSetQuestion(),
				questionSetLanguages: []string{
					"ENGLISH",
				},
			},
			want:          getQuestionSetQuestion(),
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &QuestionSetQuestion{}
			entity.fromPB(tt.args.req, tt.args.questionSetLanguages)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func getqQuestionSetQuestion() *qbSetType.QuestionSetQuestion {
	return &qbSetType.QuestionSetQuestion{
		MarkingScheme:  getqMarkingScheme(),
		QuestionDetail: getqQuestion(),
	}
}

func getQuestionSetQuestion() *QuestionSetQuestion {
	return &QuestionSetQuestion{
		MarkingScheme: getMarkingScheme(),
		Questions:     getQuestion(),
	}
}

func TestPartialCorrectMarkingScheme_FromPB(t *testing.T) {
	type args struct {
		req *qbSetType.PartialCorrectMarkingScheme
	}

	tests := []struct {
		name          string
		args          args
		want          *PartialCorrectMarkingScheme
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test PartialCorrectMarkingScheme Entity From PB",
			args: args{
				req: &qbSetType.PartialCorrectMarkingScheme{
					CorrectMarks: 1,
				},
			},
			want: &PartialCorrectMarkingScheme{
				CorrectMarks: 1,
			},
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &PartialCorrectMarkingScheme{}
			entity.fromPB(tt.args.req)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func TestTime_FromPB(t *testing.T) {
	type args struct {
		req *qbSetType.Time
	}

	tests := []struct {
		name          string
		args          args
		want          *Time
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test Time Entity From PB",
			args: args{
				req: &qbSetType.Time{
					Value: 1,
					Unit:  qbSetType.Unit_HOUR,
				},
			},
			want: &Time{
				Value: 1,
				Unit:  "HOUR",
			},
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &Time{}
			entity.fromPB(tt.args.req)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func TestMedia_FromPB(t *testing.T) {
	type args struct {
		req *qbQType.Media
	}

	tests := []struct {
		name          string
		args          args
		want          *Media
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test Media Entity From PB",
			args: args{
				req: &qbQType.Media{
					MediaId: "abc123",
				},
			},
			want: &Media{
				MediaID: "abc123",
			},
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &Media{}
			entity.FromPB(tt.args.req)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func TestInstructions_FromPB(t *testing.T) {
	type args struct {
		req *qbSetType.Instruction
	}

	tests := []struct {
		name          string
		args          args
		want          *Instruction
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test Instruction Entity From PB",
			args: args{
				req: getqInstructions(),
			},
			want:          getInstructions(),
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &Instruction{}
			entity.FromPB(tt.args.req)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func getqInstructions() *qbSetType.Instruction {
	return &qbSetType.Instruction{
		Language: "ENGLISH",
		Text:     "{\\\"1\\\":\\\"<div style=\\\\\\\"text-align:center\\\\\\\"><span style=\\\\\\\"color:#a52a2a\\\\\\\"><strong>Attempt All 35 questions<\\\\/strong><\\\\/span><\\\\/div>\\\\n\\\",\\\"2\\\":\\\"<div style=\\\\\\\"text-align:center\\\\\\\"><span style=\\\\\\\"color:#a52a2a\\\\\\\"><strong>\\\\u0938\\\\u092d\\\\u0940 35 \\\\u092a\\\\u094d\\\\u0930\\\\u0936\\\\u094d\\\\u0928 \\\\u0905\\\\u0928\\\\u093f\\\\u0935\\\\u093e\\\\u0930\\\\u094d\\\\u092f \\\\u0939\\\\u0948\\\\u0902<\\\\/strong><\\\\/span><\\\\/div>\\\\n\\\",\\\"3\\\":\\\"\\\",\\\"4\\\":\\\"\\\",\\\"5\\\":\\\"\\\",\\\"6\\\":\\\"\\\",\\\"7\\\":\\\"\\\",\\\"8\\\":\\\"\\\",\\\"9\\\":\\\"\\\",\\\"10\\\":\\\"\\\",\\\"11\\\":\\\"\\\"}",
	}
}

func getInstructions() *Instruction {
	return &Instruction{
		Language: "ENGLISH",
		Text:     "{\\\"1\\\":\\\"<div style=\\\\\\\"text-align:center\\\\\\\"><span style=\\\\\\\"color:#a52a2a\\\\\\\"><strong>Attempt All 35 questions<\\\\/strong><\\\\/span><\\\\/div>\\\\n\\\",\\\"2\\\":\\\"<div style=\\\\\\\"text-align:center\\\\\\\"><span style=\\\\\\\"color:#a52a2a\\\\\\\"><strong>\\\\u0938\\\\u092d\\\\u0940 35 \\\\u092a\\\\u094d\\\\u0930\\\\u0936\\\\u094d\\\\u0928 \\\\u0905\\\\u0928\\\\u093f\\\\u0935\\\\u093e\\\\u0930\\\\u094d\\\\u092f \\\\u0939\\\\u0948\\\\u0902<\\\\/strong><\\\\/span><\\\\/div>\\\\n\\\",\\\"3\\\":\\\"\\\",\\\"4\\\":\\\"\\\",\\\"5\\\":\\\"\\\",\\\"6\\\":\\\"\\\",\\\"7\\\":\\\"\\\",\\\"8\\\":\\\"\\\",\\\"9\\\":\\\"\\\",\\\"10\\\":\\\"\\\",\\\"11\\\":\\\"\\\"}",
	}
}

func TestQuestionStem_FromPB(t *testing.T) {
	type args struct {
		req *qbQType.QuestionStem
	}
	tests := []struct {
		name          string
		args          args
		want          *QuestionStem
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test QuestionStem Entity From PB",
			args: args{
				req: &qbQType.QuestionStem{
					Text: "abc123",
				},
			},
			want: &QuestionStem{
				Text: "abc123",
			},
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &QuestionStem{}
			entity.fromPB(tt.args.req)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func TestQuestion_FromPB(t *testing.T) {
	type args struct {
		req                  *qbQType.QuestionInformation
		questionSetLanguages []string
	}
	tests := []struct {
		name          string
		args          args
		want          *Question
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test Question Entity From PB",
			args: args{
				req: getqQuestion(),
				questionSetLanguages: []string{
					"ENGLISH",
				},
			},
			want:          getQuestion(),
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &Question{}
			entity.fromPB(tt.args.req, tt.args.questionSetLanguages)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func TestOption_FromPB(t *testing.T) {
	type args struct {
		req *qbQType.Option
	}
	tests := []struct {
		name          string
		args          args
		want          *Option
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test Option Entity From PB",
			args: args{
				req: &qbQType.Option{
					Text: "abc123",
					Media: []*qbQType.Media{
						{
							MediaId: "1233",
						},
					},
				},
			},
			want: &Option{
				Text: "abc123",
				Media: []*Media{
					{
						MediaID: "1233",
					},
				},
			},
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &Option{}
			entity.fromPB(tt.args.req)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func TestContent_FromPB(t *testing.T) {
	type args struct {
		req *qbQType.Content
	}
	tests := []struct {
		name          string
		args          args
		want          *Content
		wantErr       bool
		errorResponse *errors.Error
	}{
		{
			name: "Test Content Entity From PB",
			args: args{
				req: getqContent(),
			},
			want:          getContent(),
			wantErr:       false,
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &Content{}
			entity.fromPB(tt.args.req)
			assert.Equal(t, tt.want, entity)
		})
	}
}

func getqQuestion() *qbQType.QuestionInformation {
	return &qbQType.QuestionInformation{
		Type:    qbQType.QuestionType_SUBJECTIVE,
		Content: []*qbQType.Content{getqContent()},
	}
}

func getQuestion() *Question {
	return &Question{
		Type:    "SUBJECTIVE",
		Content: []*Content{getContent()},
	}
}

func getqMarkingScheme() *qbSetType.MarkingScheme {
	return &qbSetType.MarkingScheme{
		PartialMarkingType: &qbSetType.PartialCorrectMarkingScheme{
			CorrectMarks: 1,
		},
	}
}

func getMarkingScheme() *MarkingScheme {
	return &MarkingScheme{
		PartialMarkingType: &PartialCorrectMarkingScheme{
			CorrectMarks: 1,
		},
	}
}

func getqContent() *qbQType.Content {
	return &qbQType.Content{
		GroupId:  "abc123",
		Language: qbQType.QuestionLanguage_ENGLISH,
		QnsContent: &qbQType.QuestionStem{
			Text: "abc",
			Media: []*qbQType.Media{
				{
					MediaId: "1233",
				},
			},
		},
		Nature: qbQType.QuestionNature_QN_OBJECTIVE,
		Options: []*qbQType.Option{
			{
				Text: "abc123",
				Media: []*qbQType.Media{
					{
						MediaId: "1233",
					},
				},
			},
		},
	}
}

func getContent() *Content {
	return &Content{
		GroupID:  "abc123",
		Language: "ENGLISH",
		QnsContent: &QuestionStem{
			Text: "abc",
			Media: []*Media{
				{
					MediaID: "1233",
				},
			},
		},
		Nature: "QN_OBJECTIVE",
		Options: []*Option{
			{
				Text: "abc123",
				Media: []*Media{
					{
						MediaID: "1233",
					},
				},
			},
		},
		QuestionLanguageStr: "ENGLISH",
		QuestionNatureStr:   "QN_OBJECTIVE",
	}
}
