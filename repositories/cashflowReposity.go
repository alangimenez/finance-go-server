package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/alangimenez/finance-go-server/conexion"
	"github.com/alangimenez/finance-go-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var cashflowCollection *mongo.Collection = conexion.Client.Database("investment-project").Collection("cashflows")

func GetAllCashflowsWithTickets() ([]model.Cashflow, []string) {
	thirtyDaysLater := time.Now().AddDate(0, 0, 30)

	filter := bson.M{
		"finish": bson.M{"$gte": thirtyDaysLater},
	}

	// Consultar todos los documentos en la colección
	cursor, err := cashflowCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var cashflows []model.Cashflow
	var tickets []string
	for cursor.Next(context.Background()) {
		var persona model.Cashflow
		err := cursor.Decode(&persona)
		if err != nil {
			log.Fatal(err)
		}
		cashflows = append(cashflows, persona)
		tickets = append(tickets, persona.Ticket)
	}

	return cashflows, tickets
}

func GetCashflowByTicket(ticket string) (model.Cashflow, error) {
	filter := bson.M{"ticket": ticket}
	var cashflow model.Cashflow
	var err = cashflowCollection.FindOne(context.Background(), filter).Decode(&cashflow)
	if err == mongo.ErrNoDocuments {
		message := fmt.Sprintf("The cashflow for the ticket %s does not exist", ticket)
		fmt.Print(message)
		return cashflow, errors.New(message)
	} else if err != nil {
		log.Fatal(err)
		return cashflow, err
	} else {
		return cashflow, nil
	}
}
