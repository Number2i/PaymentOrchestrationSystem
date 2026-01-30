package worker

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func StartWorkers(repo workerRepository) {
	if err := godotenv.Load(); err != nil {
		log.Print(err)
	}
	workerCount := 5 //Default
	workerCount, err := strconv.Atoi(os.Getenv("WORKER_COUNT"))
	if err != nil {
		log.Print(err)
	}

	for i := 0; i < workerCount; i++ {
		go worker(repo)
	}
}

func worker(repo workerRepository) {
	for {
		//claim the payment
		payment, err := repo.ClaimPayment()
		if err == sql.ErrNoRows {
			//sleep fop 2 seconds
			log.Println("no work available")
			time.Sleep(time.Second * 2)
			continue
		}
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(payment)
	}
}
