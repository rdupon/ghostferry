package ghostferry

import (
	"crypto/rand"
	"encoding/binary"
	"time"

	"github.com/sirupsen/logrus"
)

func WithRetries(maxRetries int, sleep time.Duration, logger *logrus.Entry, verb string, f func() error) (err error) {
	try := 1

	for {
		err = f()
		if err == nil {
			return
		}

		if maxRetries != 0 && try >= maxRetries {
			break
		}

		logger.WithError(err).Errorf("failed to %s, %d of %d max retries", verb, try, maxRetries)

		try++
		time.Sleep(sleep)
	}

	logger.WithError(err).Errorf("failed to %s after %d attempts, retry limit exceeded", verb, try)

	return
}

func randomServerId() uint32 {
	var buf [4]byte
	if _, err := rand.Read(buf[:]); err != nil {
		panic(err)
	}

	return binary.LittleEndian.Uint32(buf[:])
}