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
	var err error
	connectionAttemptCount := 0
	shouldOpenConnection := true

	for shouldOpenConnection {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println(fmt.Errorf("Error in DB connection: %s \n", err))
			connectionAttemptCount += 1
			time.Sleep(1 * time.Second)
			shouldOpenConnection = connectionAttemptCount < 6
		} else {
			shouldOpenConnection = false
		}

	}

	if err != nil {
		return nil, err
	}

	eventLog := EventLog{db: db}
	AutoMigrate(db)

	return &eventLog, nil
}

func (e *EventLog) FindLastSynchronisedDepth(topic string) (uint64, error) {
	event := Event{}

	result := e.db.Where("depth > 0 AND topic = ?", topic).Order("depth desc").First(&event)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, result.Error
	}

	return event.Depth, nil
}

func (e *EventLog) Insert(event Event) (Event, error) {
	result := e.db.Create(&event)
	fmt.Printf("New event with CID %v registered in database!\n", event.Cid)
	if result.Error != nil {
		return Event{}, result.Error
	}

	return event, nil
}

func (e *EventLog) InsertMany(events []Event) ([]Event, error) {
	insertedEvents := []Event{}
	for _, event := range events {
		insertedEvent, insertErr := e.Insert(event)
		if insertErr != nil {
			return []Event{}, insertErr
		}
		insertedEvents = append(insertedEvents, insertedEvent)
	}

	return insertedEvents, nil
}

func (e *EventLog) Confirm(txHash string, blockNumber uint64, timestamp uint64, depth uint64) (Event, error) {
	existingEvent := Event{}
	getResult := e.db.Where("tx_hash = ?", txHash).Take(&existingEvent)
	if getResult.Error != nil {
		return Event{}, getResult.Error
	}
	existingEvent.Timestamp = timestamp
	existingEvent.BlockNumber = blockNumber
	existingEvent.Depth = depth
	e.db.Save(&existingEvent)
	fmt.Printf("Confirmation of event with tx hash %v is made\n", txHash)
	return existingEvent, nil
}

func (e *EventLog) FindConfirmedEvents() ([]Event, error) {
	events := []Event{}

	result := e.db.Where("depth > 0").Order("depth asc").Find(&events)

	if result.Error != nil {
		return events, result.Error
	}

	return events, nil
}

func (e *EventLog) FindPendingEvents() ([]Event, error) {
	events := []Event{}

	result := e.db.Where("depth = 0").Order("id asc").Find(&events)

	if result.Error != nil {
		return events, result.Error
	}

	return events, nil
}

func (e *EventLog) ClearPendingEvents(topic string) error {
	e.db.Delete(&Event{}, "topic = ? AND depth = 0", topic)
	return nil
}

func (e *EventLog) FindPage(offset uint, pageSize uint) ([]Event, bool, error) {
	events := []Event{}

	result := e.db.Where("depth > 0").Offset(int(offset)).Limit(int(pageSize) + 1).Find(&events)

	if result.Error != nil {
		return events, false, result.Error
	}
	hasMore := len(events) > int(pageSize)

	var filteredEvents []Event
	if hasMore {
		filteredEvents = events[:pageSize]
	} else {
		filteredEvents = events
	}
	return filteredEvents, hasMore, nil
}
