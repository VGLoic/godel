package eventlog

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type EventLog struct {
	db *gorm.DB
}

type EventLogConfiguration struct {
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
	PostgresPort     string
}

func NewEventLog(eventLogConfiguration EventLogConfiguration) (*EventLog, error) {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Europe/Paris",
		eventLogConfiguration.PostgresHost,
		eventLogConfiguration.PostgresUser,
		eventLogConfiguration.PostgresPassword,
		eventLogConfiguration.PostgresDb,
		eventLogConfiguration.PostgresPort,
	)

	var db *gorm.DB
	var dbErr error
	connectionAttemptCount := 0
	shouldOpenConnection := true

	for shouldOpenConnection {
		db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if dbErr != nil {
			fmt.Println(fmt.Errorf("Error in DB connection: %s \n", dbErr))
			connectionAttemptCount += 1
			time.Sleep(1 * time.Second)
			shouldOpenConnection = connectionAttemptCount < 6
		} else {
			shouldOpenConnection = false
		}

	}

	if dbErr != nil {
		return nil, dbErr
	}

	eventLog := EventLog{db: db}
	AutoMigrate(db)

	return &eventLog, nil
}

func (e *EventLog) FindLastSynchronisedBlockNumber(topic string) (uint64, error) {
	event := Event{}

	result := e.db.Where("block_number > 0 AND topic = ?", topic).Order("block_number desc").First(&event)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, result.Error
	}

	return event.BlockNumber, nil
}

func (e *EventLog) ResetFromBlockAndInsert(topic string, fromBlock uint64, events []Event) error {
	e.db.Delete(&Event{}, "topic = ? AND block_number >= ?", topic, fromBlock)

	for _, event := range events {
		_, insertErr := e.Insert(event)
		if insertErr != nil {
			return insertErr
		}
	}

	return nil
}

func (e *EventLog) Insert(event Event) (Event, error) {
	result := e.db.Create(&event)
	fmt.Printf("New event with CID %v registered in database!\n", event.Cid)
	if result.Error != nil {
		return Event{}, result.Error
	}

	return event, nil
}

func (e *EventLog) Confirm(txHash string, blockNumber uint64, timestamp uint64) (Event, error) {
	existingEvent := Event{}
	getResult := e.db.Where("tx_hash = ?", txHash).Take(&existingEvent)
	if getResult.Error != nil {
		return Event{}, getResult.Error
	}
	existingEvent.Timestamp = timestamp
	existingEvent.BlockNumber = blockNumber
	e.db.Save(&existingEvent)
	fmt.Printf("Confirmation of event with tx hash %v is made\n", txHash)
	return existingEvent, nil
}

func (e *EventLog) FindConfirmedEvents() ([]Event, error) {
	events := []Event{}

	result := e.db.Where("block_number > 0").Order("block_number asc").Find(&events)

	if result.Error != nil {
		return events, result.Error
	}

	return events, nil
}

func (e *EventLog) FindPendingEvents() ([]Event, error) {
	events := []Event{}

	result := e.db.Where("block_number = 0").Order("block_number asc").Find(&events)

	if result.Error != nil {
		return events, result.Error
	}

	return events, nil
}

func (e *EventLog) FindPage(offset uint, pageSize uint) ([]Event, error) {
	events := []Event{}

	result := e.db.Offset(int(offset)).Limit(int(pageSize)).Find(&events)

	if result.Error != nil {
		return events, result.Error
	}

	return events, nil
}

func (e *EventLog) ClearPendingEvents(topic string) error {
	e.db.Delete(&Event{}, "topic = ? AND block_number = 0", topic)
	return nil
}
