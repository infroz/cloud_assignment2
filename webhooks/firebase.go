package webhooks

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type FirestoreDatabase struct {
	Ctx    context.Context
	Client *firestore.Client
}

var DB = FirestoreDatabase{}

func Init() error {

	// Firebase initialisation
	DB.Ctx = context.Background()
	// We use a service account, load credentials file that you downloaded from your project's settings menu.
	// Make sure this file is gitignored, it is the access token to the database.
	sa := option.WithCredentialsFile("../repocheck-a9524-firebase-adminsdk-37aot-2505e40bf2.json")
	app, err := firebase.NewApp(DB.Ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	DB.Client, err = app.Firestore(DB.Ctx)

	// Alternative setup, directly through Firestore
	// client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

// Close closes the DB connection
func Close() {
	DB.Client.Close()
}

func Read() ([]Webhook, error) {
	var Temp []Webhook
	iter := DB.Client.Collection("webhooks").Documents(DB.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(doc.Data())
		err = doc.DataTo(&Wh)
		if err != nil {
			log.Fatalln(err)
		}

		Temp = append(Temp, Wh)
	}

	return Temp, nil
}

func Delete(ID string) error {
	docref := DB.Client.Collection("webhooks").Doc(ID)
	_, err := docref.Delete(DB.Ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		return errors.Wrap(err, "Error in FirebaseDatabase.Delete()")
	}
	return nil
}

func Add() error {
	ref := DB.Client.Collection("webhooks").NewDoc()
	Wh.ID = ref.ID
	_, err := ref.Set(DB.Ctx, Wh)
	if err != nil {
		return errors.Wrap(err, "Error in FirebaseDatabase.Save()")
	}
	return nil
}
