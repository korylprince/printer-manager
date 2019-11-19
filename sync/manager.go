package sync

import (
	"database/sql"
	"fmt"
	"log"
	gosync "sync"
	"time"
)

//Manager represents a synchronization manager
type Manager struct {
	config   *Config
	db       *sql.DB
	once     gosync.Once
	Interval time.Duration
	timer    *time.Timer
	trigger  chan *gosync.WaitGroup
	lastRun  time.Time
	nextRun  time.Time
	mu       *gosync.RWMutex
}

//NewManager returns a new Manager with the given parameters
func NewManager(config *Config, db *sql.DB, interval time.Duration) (*Manager, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("Unable to validate configuration: %v", err)
	}

	m := &Manager{
		config:   config,
		db:       db,
		Interval: interval,
		trigger:  make(chan *gosync.WaitGroup),
		nextRun:  time.Now().Add(interval),
		mu:       new(gosync.RWMutex),
	}
	return m, nil
}

//Start starts the sync manager. Multiple calls have no effect
func (m *Manager) Start() {
	m.once.Do(func() {
		go m.run()
	})
}

//Trigger triggers a synchronization and waits for it to complete. Must be called after Start
func (m *Manager) Trigger() {
	wg := new(gosync.WaitGroup)
	wg.Add(1)
	m.trigger <- wg
	wg.Wait()
}

//Stats returns stats about the Manager
func (m *Manager) Stats() (lastRun, nextRun time.Time) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.lastRun, m.nextRun
}

func (m *Manager) run() {
	m.timer = time.NewTimer(m.Interval)
	var wg *gosync.WaitGroup
	for {
		log.Println("SYNC: Starting")
		tx, err := m.db.Begin()
		if err != nil {
			log.Println("SYNC: Unable to start database transaction:", err)
			goto wait
		}
		if err = sync(m.config, tx); err != nil {
			log.Println("SYNC: Failure:", err)
			if rErr := tx.Rollback(); rErr != nil {
				log.Println("SYNC: Unable to roll back database transaction:", rErr)
			}
			goto wait
		}
		if err = tx.Commit(); err != nil {
			log.Println("SYNC: Unable to commit database transaction:", err)
			goto wait
		}
		m.mu.Lock()
		m.lastRun = time.Now()
		m.mu.Unlock()
		log.Println("SYNC: Completed successfully")

	wait:
		if wg != nil {
			wg.Done()
		}
		wg = nil
		select {
		case <-m.timer.C:
		case wg = <-m.trigger:
		}
		m.mu.Lock()
		m.nextRun = time.Now().Add(m.Interval)
		m.mu.Unlock()
		m.timer.Reset(m.Interval)
	}
}
