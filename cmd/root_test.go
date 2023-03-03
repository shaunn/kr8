package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	assert2 "github.com/stretchr/testify/assert"
	"os"
	"path"
	"path/filepath"
	"sort"
	"testing"
)

var (
	codeBaseDir string = getBaseDir()
	testBaseDir string = codeBaseDir + "/testdata/"
)

// Note: Set(Override) > flags > env variables > config file > defaults

func TestMain(m *testing.M) {

	m.Run()
}

func getBaseDir() string {
	currdir, err := os.Getwd()
	if err != nil {
		fatalog(err).Str("source", "getBaseDir>os.Getwd()").Msg(currdir)
	}
	return path.Dir(currdir)
}

func runSettingsAssertions(t *testing.T, ks sort.StringSlice, kvmap map[string]interface{}) {
	allkeys := sort.StringSlice(viper.AllKeys())
	allkeys.Sort()
	ks.Sort()

	assert2.Equal(t, ks, allkeys)
	assert2.Equal(t, kvmap, viper.AllSettings())

}

func TestDefaultSettings(t *testing.T) {

	// Clear out all environment variables
	os.Clearenv()

	// Expected outputs
	ks := sort.StringSlice{
		"base",
		"color",
		"config",
		"jpath",
		"loglevel",
		"long",
		"noexit",
		"warn",
	}
	all := map[string]interface{}{
		"base":     ".",
		"color":    true,
		"config":   ".kr8.env",
		"jpath":    []string{},
		"loglevel": "info",
		"long":     false,
		"noexit":   false,
		"warn":     false,
	}

	// The combo to success
	os.Clearenv()
	viper.Reset()
	initConfig()

	runSettingsAssertions(t, ks, all)

}

func TestDotEnvSettings(t *testing.T) {

	// Clear out all environment variables
	os.Clearenv()

	// Define every variable, flipping default bools
	var envars = map[string]string{
		"KR8_LOGLEVEL":     "debug",
		"KR8_LONG":         "true",
		"KR8_NOEXIT":       "true",
		"KR8_DEBUG":        "true",
		"KR8_TRACE":        "true",
		"KR8_WARN":         "true",
		"KR8_EXTSTRFILE":   "/dev/null",
		"KR8_BASE":         "/dev/null",
		"KR8_CLEXCLUDES":   "rightshark.+",
		"KR8_CLINCLUDES":   "leftshark.+",
		"KR8_CLUSTER":      ".+",
		"KR8_CLUSTERDIR":   "/dev/null",
		"KR8_COLOR":        "false",
		"KR8_COMPONENTDIR": "/dev/null",
		"KR8_CONFIG":       "/dev/null",
		"KR8_JPATH":        "/dev/null",
	}

	// Now loop through and set via os.SetEnv
	for k, v := range envars {
		tracelog(nil).Msg(k + " " + v)
		os.Setenv(k, v)
	}

	// Expected outputs

	ks := sort.StringSlice{
		"base",
		"color",
		"config",
		"jpath",
		"loglevel",
		"long",
		"noexit",
		"warn",
	}

	all := map[string]interface{}{
		"base":     "/dev/null",
		"color":    false,
		"config":   "/dev/null",
		"jpath":    []string{"/dev/null"},
		"loglevel": "debug",
		"long":     true,
		"noexit":   true,
		"warn":     true,
	}

	// The combo to success
	viper.Reset()
	initConfig()

	runSettingsAssertions(t, ks, all)

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func TestConfigFileViaDefaultFile(t *testing.T) {
	// change working directory to a kr8-configs base with .kr8.env (alt)
	// Note: This is NOT changing the base, just assuming a different dir
	//  and the code is supposed to use the local .kr8.env (alt)

	// Change working directory
	origDir, _ := os.Getwd()
	workDir := testBaseDir + "alt/kr8-configs/"
	os.Chdir(workDir)

	os.Clearenv()
	viper.Reset()
	initConfig()

	ks := sort.StringSlice{
		"base",
		"cluster",
		"color",
		"config",
		"jpath",
		"loglevel",
		"long",
		"noexit",
		"warn",
	}
	all := map[string]interface{}{
		"base":     ".",
		"cluster":  ".+",
		"color":    true,
		"config":   ".kr8.env",
		"jpath":    []string{"/tmp/TestConfigFileViaDefaultFile/lib"},
		"loglevel": "info",
		"long":     false,
		"noexit":   true,
		"warn":     false,
	}

	// Tests/assertions
	configFileUsed, _ := filepath.Abs(viper.ConfigFileUsed())
	assert2.Equal(t, workDir+".kr8.env", configFileUsed)

	runSettingsAssertions(t, ks, all)

	// Post test cleanup
	os.Chdir(origDir) // This may not be necessary
}

func TestChangeBaseViaEnv(t *testing.T) {
	// change working directory to a kr8-configs base with NO .kr8.env (default)

	// Clear out all environment variables
	os.Clearenv()

	// Set up config file and check existence

	base := testBaseDir + "default/kr8-configs/"

	// Set env with the config file path
	os.Setenv("KR8_BASE", base)

	// Expected outputs
	ks := sort.StringSlice{
		"base",
		"color",
		"config",
		"jpath",
		"loglevel",
		"long",
		"noexit",
		"warn",
	}

	all := map[string]interface{}{
		"base":     base,
		"color":    true,
		"config":   ".kr8.env",
		"jpath":    []string{},
		"loglevel": "info",
		"long":     false,
		"noexit":   false,
		"warn":     false,
	}

	viper.Reset()
	initConfig()

	// Tests/assertions

	configFileExpected := codeBaseDir + "/cmd/.kr8.env"
	configFileUsed, _ := filepath.Abs(viper.ConfigFileUsed())
	assert2.Equal(t, configFileExpected, configFileUsed)

	runSettingsAssertions(t, ks, all)
}

func TestConfigFileViaFlag(t *testing.T) {
	// How is this done wothout adding flags?
	// Would adding flags here limit the integrity of the test?

	// Clear out all environment variables
	os.Clearenv()
	//
	// os.Args = append(os.Args, "--addr=http://b.com:566/something.avsc")
	// os.Args = append(os.Args, "Get")
	// os.Args = append(os.Args, `./some/resource/fred`)

	configFile := testBaseDir + "/default/kr8-configs/.kr8.test.default.env"

	// os.Args = append(os.Args, "generate")
	// os.Args = append(os.Args, "--clusters gke")
	// os.Args = append(os.Args, "--config", configFile)
	os.Args = append(os.Args, "--config "+configFile)
	// os.Args = append(os.Args, "--config="+configFile)
	// os.Args = append(os.Args, "Get")
	// os.Args = append(os.Args, `./some/resource/fred`)

	fmt.Println(os.Args)
	fmt.Println(viper.Get("cluster"))
	fmt.Println(viper.AllSettings())
	infolog(nil).Msg(cluster)
	infolog(nil).Msg(clusters)

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	ks := sort.StringSlice{}
	all := map[string]interface{}{}

	// Run init to take in the subject parms
	viper.Reset()
	initConfig()

	fmt.Println(os.Args)
	fmt.Println(viper.Get("cluster"))
	fmt.Println(viper.AllSettings())

	infolog(nil).Msg(cluster)
	infolog(nil).Msg(clusters)

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
	// Tests/assertions
	configFileUsed, _ := filepath.Abs(viper.ConfigFileUsed())
	errorlog(nil).Msg(configFile + ":::" + configFileUsed)

	assert2.Equal(t, configFile, configFileUsed)

	runSettingsAssertions(t, ks, all)

	// Post test cleanup
}
