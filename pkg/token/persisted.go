package token

import (
	"encoding/hex"
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Store interface {
	Save(any, string) error
	Find(any, string, string) error
	FindAll(any) error
	Delete(any, string) error
}

type PersistedToken struct {
	Token  string `gorm:"primaryKey"`
	Id     string
	Expiry time.Time
}

type PersistedTokenManager struct {
	Rand     rand.Rand
	Store    Store
	Duration time.Duration
}

func (t *PersistedTokenManager) GenerateToken(id string) (string, error) {
	buf := make([]byte, 128)
	t.Rand.Read(buf)
	token := PersistedToken{
		Token:  hex.EncodeToString(buf),
		Id:     id,
		Expiry: time.Now().Add(t.Duration),
	}
	if err := t.Store.Save(&token, token.Token); err != nil {
		return "", err
	}
	return token.Token, nil
}

func (t *PersistedTokenManager) Validate(s string) (string, bool) {
	var token PersistedToken
	if err := t.Store.Find(&token, "token", s); err != nil {
		return "", false
	}
	if token.Expiry.Before(time.Now()) {
		return "", false
	}
	return token.Id, true
}

func (t *PersistedTokenManager) Invalidate(s string) bool {
	return t.Store.Delete(&PersistedToken{Token: s}, s) == nil
}

func (t *PersistedTokenManager) GetExpiry(s string) time.Time {
	var token PersistedToken
	if err := t.Store.Find(&token, "token", s); err != nil {
		return time.Time{}
	}
	return token.Expiry
}

func (t *PersistedTokenManager) CleanupExpiredTokens() error {
	var tokens []PersistedToken
	if err := t.Store.FindAll(&tokens); err != nil {
		return err
	}
	total, failed := 0, 0
	wg := &sync.WaitGroup{}
	c := make(chan struct{}, 2)
	failedTokens := make(chan string)
	for _, token := range tokens {
		if token.Expiry.Before(time.Now()) {
			c <- struct{}{}
			wg.Add(1)
			total++
			go func(token PersistedToken) {
				defer func() {
					wg.Done()
					<-c
				}()
				if !t.Invalidate(token.Token) {
					failedTokens <- token.Token
				}
			}(token)
		}
	}
	go func() {
		for token := range failedTokens {
			log.Println("could not cleanup expired token", token)
			failed++
		}
	}()
	wg.Wait()
	close(failedTokens)
	log.Printf("Expired tokens: %d, Cleaned up: %d, Failed: %d\n", total, total-failed, failed)
	if failed > 0 {
		return errors.New("could not cleanup all expired tokens")
	}
	return nil
}

func (t *PersistedTokenManager) RunCleaner(stop <-chan struct{}) {
	ticker := time.NewTicker(time.Second * 20)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := t.CleanupExpiredTokens(); err != nil {
					log.Println(err.Error())
				}
			case <-stop:
				return
			}
		}
	}()
}
