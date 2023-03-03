package main

import (
	"github.com/spf13/viper"
	assert2 "github.com/stretchr/testify/assert"
	"os"
	"sort"
	"testing"
)

// Note: Set(Override) > flags > env variables > config file > defaults

func TestMain(m *testing.M) {
	// Ensure no environment variables are in the mix.
	os.Clearenv()

	m.Run()

}

func runAssertions(t *testing.T, ks sort.StringSlice, kvmap map[string]interface{}) {
	allkeys := sort.StringSlice(viper.AllKeys())
	allkeys.Sort()
	ks.Sort()

	assert2.Equal(t, ks, allkeys)
	assert2.Equal(t, kvmap, viper.AllSettings())
}

func TestKr8Main(t *testing.T) {

	// Nothing here yet

	viper.Reset()

	// Expected outputs
	ks := sort.StringSlice{}
	all := map[string]interface{}{}

	runAssertions(t, ks, all)
}
