package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"log"
	"time"
)

var client *firestore.Client

type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	}
}

type FirestoreValue struct {
	CreateTime time.Time `json:"createTime"`
	Name       string    `json: "name"`
	UpdateTime time.Time `json:"updateTime"`
	Fields     User      `json: "fields"`
}

type User struct {
	Username   StringValue  `json:"userId"`
	Email      StringValue  `json:"email"`
	DateEdited IntegerValue `json:"date_edited"`
}

type IntegerValue struct {
	IntegerValue string `json:"integerValue"`
}

type StringValue struct {
	StringValue string `json:"stringValue"`
}

// Simple init to have a firestore client available
func init() {
	ctx := context.Background()
	var err error
	client, err = firestore.NewClient(ctx, "plzjeoveme")
	if err != nil {
		log.Fatalf("Firestore: %v", err)
	}
}

// Handles the rollback to a previous document
func handleRollback(ctx context.Context, e FirestoreEvent) error {
	return errors.New("Should have rolled back to a previous version")
}

// The function that runs with the cloud function itself
func HandleUserChange(ctx context.Context, e FirestoreEvent) error {
	// This is the data that's in the database itself
	newFields := e.Value.Fields
	oldFields := e.OldValue.Fields

	// As our goal is simply to check if the username has changed
	if newFields.Username.StringValue == oldFields.Username.StringValue {
		log.Printf("Bad username: %s - %s", newFields.Username.StringValue, oldFields.Username.StringValue)
		return handleRollback(ctx, e)
	}

	// Check if the email is the same as previously
	if newFields.Email.StringValue != oldFields.Email.StringValue {
		log.Printf("Bad email: %s - %s", newFields.Email.StringValue, oldFields.Email.StringValue)
		return handleRollback(ctx, e)
	}

	return nil
}
