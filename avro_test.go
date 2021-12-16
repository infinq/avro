package avro_test

import "github.com/infinq/avro"

func ConfigTeardown() {
	// Reset the caches
	avro.DefaultConfig = avro.Config{}.Freeze()
}
