package mongodb

import (
	"context"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

var studentTestInsightEntry = entity.StudentTestInsightEntity{
	StudentID: "eK0AF0EaQoKIItakllp0w",
	TestID:    "test_2ecrPgr0eQ9o",
	SectionalInsights: []entity.SectionInsight{
		{
			SectionName:        "Botany",
			Attempted:          5,
			TotalQuestions:     50,
			CorrectQuestions:   3,
			InCorrectQuestions: 2,
			MarksScored:        10,
			MaximumMarks:       200,
			Rank:               5,
			AttemptPercentage:  10,
			Accuracy:           60,
			Percentage:         5,
			Percentile:         100,
		},
	},
	PeerInsights: nil,
	OverallInsight: entity.OverallInsight{
		Attempted:          29,
		TotalQuestions:     200,
		CorrectQuestions:   11,
		InCorrectQuestions: 18,
		MarksScored:        26,
		MaximumMarks:       800,
		Rank:               3,
		AttemptPercentage:  14.5,
		Accuracy:           37.93,
		Percentage:         3.25,
		Percentile:         33.33,
	},
}

func TestStudentTestInsightRepository_CreateStudentTestInsight(t *testing.T) {
	d := &MongoClient{
		client: mongoClient,
		dbName: "test-and-assessments",
	}
	ctx := context.Background()
	studentTestInsightRepo := NewStudentTestInsightRepository(log.NewStdLogger(os.Stdout), d)
	type assertions struct {
		wantErr    bool
		assertFunc func(t *testing.T, ctx context.Context, studentTestInsight entity.StudentTestInsightEntity)
	}
	tests := []struct {
		name        string
		assertions  assertions
		prepareFunc func(*testing.T)
		arg         entity.StudentTestInsightEntity
	}{
		{
			name: "Insert student test insight document to DB",
			assertions: assertions{
				wantErr: false,
				assertFunc: func(t *testing.T, ctx context.Context, studentTestInsightArg entity.StudentTestInsightEntity) {
					filter := bson.D{
						{Key: testIdKey, Value: studentTestInsightArg.TestID},
						{Key: studentIdKey, Value: studentTestInsightArg.StudentID},
					}
					res := studentTestInsightRepo.db.FindOne(ctx, filter)
					err := res.Err()
					assert.Nil(t, err)
					var studentTestInsight entity.StudentTestInsightEntity
					err = res.Decode(&studentTestInsight)
					assert.Nil(t, err)
					assert.Equal(t, studentTestInsightArg.TestID, studentTestInsight.TestID)
					assert.Equal(t, studentTestInsightArg.StudentID, studentTestInsight.StudentID)
					assert.Equal(t, studentTestInsightArg.OverallInsight, studentTestInsight.OverallInsight)
					assert.Equal(t, studentTestInsightArg.SectionalInsights, studentTestInsight.SectionalInsights)
					assert.Equal(t, studentTestInsightArg.PeerInsights, studentTestInsight.PeerInsights)
				},
			},
			prepareFunc: func(t *testing.T) {},
			arg: entity.StudentTestInsightEntity{
				StudentID: "Gn0emRC4Uh7ZkG8OiaU5q",
				TestID:    "test_8PpscNGgXjP4",
				SectionalInsights: []entity.SectionInsight{
					{
						SectionName:        "Physics",
						Attempted:          20,
						TotalQuestions:     50,
						CorrectQuestions:   5,
						InCorrectQuestions: 15,
						MarksScored:        5,
						MaximumMarks:       200,
						Rank:               2,
						AttemptPercentage:  40.,
						Accuracy:           25,
						Percentage:         2.5,
						Percentile:         0,
					},
				},
				PeerInsights: nil,
				OverallInsight: entity.OverallInsight{
					Attempted:          155,
					TotalQuestions:     200,
					CorrectQuestions:   41,
					InCorrectQuestions: 114,
					MarksScored:        52,
					MaximumMarks:       800,
					Rank:               1,
					AttemptPercentage:  77.5,
					Accuracy:           26.8,
					Percentage:         6.5,
					Percentile:         100,
				},
			},
		},
		{
			name: "Upsert document in DB",
			assertions: assertions{
				wantErr: false,
				assertFunc: func(t *testing.T, ctx context.Context, studentTestInsightArg entity.StudentTestInsightEntity) {
					filter := bson.D{
						{Key: testIdKey, Value: studentTestInsightArg.TestID},
						{Key: studentIdKey, Value: studentTestInsightArg.StudentID},
					}
					res := studentTestInsightRepo.db.FindOne(ctx, filter)
					var studentTestInsight entity.StudentTestInsightEntity
					err := res.Decode(&studentTestInsight)
					assert.Nil(t, err)
					assert.Equal(t, studentTestInsightArg.TestID, studentTestInsight.TestID)
					assert.Equal(t, studentTestInsightArg.TestID, studentTestInsight.TestID)
					assert.Equal(t, studentTestInsightArg.StudentID, studentTestInsight.StudentID)
					assert.Equal(t, studentTestInsightArg.OverallInsight, studentTestInsight.OverallInsight)
					assert.Equal(t, studentTestInsightArg.SectionalInsights, studentTestInsight.SectionalInsights)
					assert.Equal(t, studentTestInsightArg.PeerInsights, studentTestInsight.PeerInsights)

					assert.Equal(t, studentTestInsightArg.TestID, studentTestInsightEntry.TestID)
					assert.Equal(t, studentTestInsightArg.StudentID, studentTestInsightEntry.StudentID)
					assert.NotEqual(t, studentTestInsightArg.OverallInsight, studentTestInsightEntry.OverallInsight)
					assert.NotEqual(t, studentTestInsightArg.SectionalInsights, studentTestInsightEntry.SectionalInsights)
					assert.NotEqual(t, studentTestInsightArg.PeerInsights, studentTestInsightEntry.PeerInsights)

				},
			},
			prepareFunc: func(t *testing.T) {
				err, _ := studentTestInsightRepo.db.InsertOne(ctx, studentTestInsightEntry)
				assert.NotNil(t, err)
			},
			arg: entity.StudentTestInsightEntity{
				StudentID: "eK0AF0EaQoKIItakllp0w",
				TestID:    "test_2ecrPgr0eQ9o",
				SectionalInsights: []entity.SectionInsight{
					{
						SectionName:        "Physics",
						Attempted:          50,
						TotalQuestions:     12,
						CorrectQuestions:   12,
						InCorrectQuestions: 35,
						MarksScored:        6,
						MaximumMarks:       23,
						Rank:               3,
						AttemptPercentage:  20,
						Accuracy:           6,
						Percentage:         1,
						Percentile:         18,
					},
				},
				PeerInsights: []entity.PeerInsight{
					{
						SectionName:                       "Physics",
						PercentileSectionMarksAvg:         12,
						QuestionsForNextPercentileSection: 13,
					},
				},
				OverallInsight: entity.OverallInsight{
					Attempted:          155,
					TotalQuestions:     200,
					CorrectQuestions:   41,
					InCorrectQuestions: 114,
					MarksScored:        52,
					MaximumMarks:       800,
					Rank:               1,
					AttemptPercentage:  77.5,
					Accuracy:           26.8,
					Percentage:         6.5,
					Percentile:         100,
				},
			},
		},
	}

	t.Cleanup(func() {
		_, _ = studentTestInsightRepo.db.DeleteMany(ctx, bson.D{})
	})
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.prepareFunc(t)
			err := studentTestInsightRepo.CreateStudentTestInsight(ctx, testCase.arg)
			if testCase.assertions.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			testCase.assertions.assertFunc(t, ctx, testCase.arg)
		})
	}
}

func TestStudentTestInsightRepository_FetchStudentCountByTestID(t *testing.T) {
	d := &MongoClient{
		client: mongoClient,
		dbName: "test-and-assessments",
	}
	ctx := context.Background()
	studentTestInsightRepo := NewStudentTestInsightRepository(log.NewStdLogger(os.Stdout), d)
	type assertions struct {
		wantErr  bool
		expected int64
	}
	tests := []struct {
		name        string
		assertions  assertions
		prepareFunc func(*testing.T)
		arg         string
	}{
		{
			name: "Return 1 document count from DB",
			assertions: assertions{
				wantErr:  false,
				expected: 1,
			},
			prepareFunc: func(t *testing.T) {
				err, _ := studentTestInsightRepo.db.InsertOne(ctx, studentTestInsightEntry)
				assert.NotNil(t, err)
			},
			arg: "test_2ecrPgr0eQ9o",
		},
		{
			name: "Return 0 document count from DB",
			assertions: assertions{
				wantErr:  false,
				expected: 0,
			},
			prepareFunc: func(t *testing.T) {},
			arg:         "random_test_id",
		},
	}

	t.Cleanup(func() {
		_, _ = studentTestInsightRepo.db.DeleteMany(ctx, bson.D{})
	})
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.prepareFunc(t)
			res, err := studentTestInsightRepo.FetchStudentCountByTestID(ctx, testCase.arg)
			if testCase.assertions.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, testCase.assertions.expected, res)
		})
	}
}

func TestStudentTestInsightRepository_FindByStudentIDTestID(t *testing.T) {
	d := &MongoClient{
		client: mongoClient,
		dbName: "test-and-assessments",
	}
	ctx := context.Background()
	studentTestInsightRepo := NewStudentTestInsightRepository(log.NewStdLogger(os.Stdout), d)
	type assertions struct {
		wantErr    bool
		assertFunc func(t *testing.T, ctx context.Context, studentTestInsight *entity.StudentTestInsightEntity)
	}
	type args struct {
		testID    string
		studentID string
	}
	tests := []struct {
		name        string
		assertions  assertions
		prepareFunc func(*testing.T)
		args        args
	}{
		{
			name: "Return document from DB",
			assertions: assertions{
				wantErr: false,
				assertFunc: func(t *testing.T, ctx context.Context, studentTestInsight *entity.StudentTestInsightEntity) {
					assert.Equal(t, studentTestInsightEntry.TestID, studentTestInsight.TestID)
					assert.Equal(t, studentTestInsightEntry.StudentID, studentTestInsight.StudentID)
					assert.Equal(t, studentTestInsightEntry.OverallInsight, studentTestInsight.OverallInsight)
					assert.Equal(t, studentTestInsightEntry.SectionalInsights, studentTestInsight.SectionalInsights)
					assert.Equal(t, studentTestInsightEntry.PeerInsights, studentTestInsight.PeerInsights)
				},
			},
			prepareFunc: func(t *testing.T) {
				err, _ := studentTestInsightRepo.db.InsertOne(ctx, studentTestInsightEntry)
				assert.NotNil(t, err)
			},
			args: args{
				studentID: "eK0AF0EaQoKIItakllp0w",
				testID:    "test_2ecrPgr0eQ9o",
			},
		},
		{
			name: "Return error when no document is found",
			assertions: assertions{
				wantErr: true,
				assertFunc: func(t *testing.T, ctx context.Context, studentTestInsight *entity.StudentTestInsightEntity) {
					assert.Nil(t, studentTestInsight)
				},
			},
			prepareFunc: func(t *testing.T) {},
			args: args{
				testID:    "random_test_id",
				studentID: "random_student_id",
			},
		},
		{
			name: "Return error when decoding document fails",
			assertions: assertions{
				wantErr: true,
				assertFunc: func(t *testing.T, ctx context.Context, studentTestInsight *entity.StudentTestInsightEntity) {
					assert.Nil(t, studentTestInsight)
				},
			},
			prepareFunc: func(t *testing.T) {
				err, _ := studentTestInsightRepo.db.InsertOne(ctx, struct {
					StudentID      string `bson:"studentId"`
					TestID         string `bson:"testId"`
					OverallInsight string `bson:"overallInsight"`
				}{
					TestID:         "corrupted_test_id",
					StudentID:      "corrupted_student_id",
					OverallInsight: "string overall insights",
				})
				assert.NotNil(t, err)
			},
			args: args{
				testID:    "corrupted_test_id",
				studentID: "corrupted_student_id",
			},
		},
	}
	t.Cleanup(func() {
		_, _ = studentTestInsightRepo.db.DeleteMany(ctx, bson.D{})
	})
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.prepareFunc(t)
			res, err := studentTestInsightRepo.FindByStudentIDTestID(ctx, testCase.args.studentID, testCase.args.testID)
			if testCase.assertions.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			testCase.assertions.assertFunc(t, ctx, res)
		})
	}
}
