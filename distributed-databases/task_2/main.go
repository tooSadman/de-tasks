package main

import (
	"context"
	"flag"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	num   int
	sleep bool
)

func init() {
	flag.IntVar(&num, "num", 1, "number of transactions to run")
	flag.IntVar(&num, "n", 1, "number of transactions to run")
	flag.BoolVar(&sleep, "sleep", false, "enable sleep on transaction")
	flag.BoolVar(&sleep, "s", false, "enable sleep on transaction")
	flag.Parse()
}

func main() {
	conn, err := pgxpool.Connect(context.Background(), "postgresql://postgres:mysecretpassword@localhost/lab2")
	if err != nil {
		log.Errorf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close()

	var wg sync.WaitGroup
	var i int

	for i = 1; i <= num; i++ {
		wg.Add(1)
		i := i
		transactLogger := *log.WithFields(log.Fields{
			"transaction": i,
		})

		go func() {
			defer wg.Done()

			err = transaction(conn, &transactLogger)
			if err != nil {
				log.Error(err)
				transactLogger.Warn("Transaction finished unsuccessfully.")
			}
		}()
	}

	wg.Wait()
}

func transaction(conn *pgxpool.Pool, transactLogger *log.Entry) error {
	var amount string

	err := conn.QueryRow(context.Background(), "select amount::varchar from account.customer_account where account_id=$1", 2).Scan(&amount)
	if err != nil {
		return err
	}

	insert_hotel := `
insert
	into
	hotel.hotel_booking (client_acc,
	hotel_name,
	arrival,
	departure,
	price)
values ($1,
'Hilton',
'2017-08-05',
'2017-08-10',
$2);
	`

	insert_flight := `
insert
	into
	flight.flight_booking (client_acc,
	flight_number,
	flight_from,
	flight_to,
	flight_date,
	price)
values ($1,
'randomNum',
'BAL',
'WOW',
'2017-08-05',
$2);
	`

	count_amount := `
update
	account.customer_account
set
	amount = amount - $1::money
where
	account_id = $2;
	`

	start := time.Now()

	transactLogger.WithFields(log.Fields{
		"amount": amount,
	}).Info("Starting transaction...")
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(context.Background())

	transactLogger.Info("Inserting flight and hotel data...")
	_, err = tx.Exec(context.Background(), insert_flight, 2, 2500)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), insert_hotel, 2, 2500)
	if err != nil {
		return err
	}
	transactLogger.Info("Inserted flight and hotel data.")

	nestedTx, err := tx.Begin(context.Background())
	if err != nil {
		return err
	}
	defer nestedTx.Rollback(context.Background())

	transactLogger.Info("Updating customer money amount...")
	_, err = nestedTx.Exec(context.Background(), count_amount, 250, 2)
	if err != nil {
		return err
	}
	transactLogger.Info("Updated customer data.")

	transactLogger.Info("Commiting transaction...")
	err = nestedTx.Commit(context.Background())
	if err != nil {
		return err
	}

	if sleep {
		transactLogger.Info("Started sleep...")
		time.Sleep(10 * time.Second)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	err = conn.QueryRow(context.Background(), "select amount::varchar from account.customer_account where account_id=$1", 2).Scan(&amount)
	if err != nil {
		return err
	}

	transactLogger.WithFields(log.Fields{
		"elapsed": time.Since(start),
		"amount":  amount,
	}).Info("Transaction successfully commited.")
	return nil
}
