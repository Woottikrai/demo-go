package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSchema struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"` // ใช้สำหรับเก็บ ObjectID ของ MongoDB
	FirstName string             `bson:"first_name"`    // Field ชื่อ "first_name" ใน MongoDB
	LastName  string             `bson:"last_name"`     // Field ชื่อ "last_name" ใน MongoDB
	Email     string             `bson:"email"`         // Field ชื่อ "email" ใน MongoDB
	Age       int                `bson:"age"`           // Field ชื่อ "age" ใน MongoDB
}
