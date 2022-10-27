package service

import (
	"context"
	"strings"
	"mongorestapi3/internal/model"
	. "mongorestapi3/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentService interface {
	GetAllStudents(ctx context.Context) (*[]Student, error)
	GetStudent(ctx context.Context, id string) (*Student, error)
	InsertStudent(ctx context.Context, student *Student) (int64, error)
	UpdateStudent(ctx context.Context, student *Student) (int64, error)
	DeleteStudent(ctx context.Context, id string) (int64, error)
}

type studentService struct {
	Collection *mongo.Collection
}

func NewStudentService(db *mongo.Database) *studentService {
	collectionName := "student"
	return &studentService{Collection: db.Collection(collectionName)}
}

// GetAllStudent implements StudentService
func (s *studentService) GetAllStudents(ctx context.Context) (*[]Student, error) {
	filter := bson.M{}
	cursor, err1 := s.Collection.Find(ctx, filter)
	if err1 != nil {
		return nil, err1
	}

	var users []Student
	err2 := cursor.All(ctx, &users)
	if err2 != nil {
		return nil, err2
	}

	return &users, nil
}

// GetStudent implements StudentService
func (s *studentService) GetStudent(ctx context.Context, id string) (*Student, error) {
	filter := bson.M{"_id": id}
	res := s.Collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}
	student := Student{}
	err := res.Decode(&student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

// InsertStudent implements StudentService
func (s *studentService) InsertStudent(ctx context.Context, student *model.Student) (int64, error) {
	_, err := s.Collection.InsertOne(ctx, student)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "duplicate key error collection") {
			if strings.Contains(errMsg, "dup key: { _id: }")  {
				return 0, nil
			} else {
				return -1, nil
			}
		} else {
			return 0, err
		}
	}

	return 1, nil
}

// UpdateStudent implements StudentService
func (s *studentService) UpdateStudent(ctx context.Context, student *model.Student) (int64, error) {
	filter := bson.M{"_id": student.ID}
	update := bson.M{
		"$set": student,
	}
	res, err := s.Collection.UpdateOne(ctx, filter, update)
	if res.ModifiedCount > 0 {
		return res.ModifiedCount, err
	} else if res.UpsertedCount > 0 {
		return res.UpsertedCount, err
	} else {
		return res.MatchedCount, err
	}
}

// DeleteStudent implements StudentService
func (s *studentService) DeleteStudent(ctx context.Context, id string) (int64, error) {
	filter := bson.M{"_id": id}
	res, err := s.Collection.DeleteOne(ctx, filter)
	if res == nil || err != nil {
		return 0, err
	}

	return res.DeletedCount, err
}
