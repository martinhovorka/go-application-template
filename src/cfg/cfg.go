package CFG

// -----------------------------------------------------------------------------
// module imports
// -----------------------------------------------------------------------------
import (
	LOG "app/src/log"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/apsdehal/go-logger"
)

// -----------------------------------------------------------------------------
// type definitions
// -----------------------------------------------------------------------------
type ConfigurationHolder struct {
	LogLevel        logger.LogLevel `json:"LogLevel"`
	MainLoopTimeout uint            `json:"MainLoopTimeout"`
}

// -----------------------------------------------------------------------------
// module global variables (non-exported)
// -----------------------------------------------------------------------------
var configuration *ConfigurationHolder = nil

// -----------------------------------------------------------------------------
// module global variables (exported)
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// non-exported methods
// -----------------------------------------------------------------------------

// print current configuration
func printConfiguration() {
	LOG.Inf("<<< Current configuration >>>")
	LOG.Inf("LogLevel: '%d'", configuration.LogLevel)
	LOG.Inf("MainLoopTimeout: '%d'", configuration.MainLoopTimeout)
}

// -----------------------------------------------------------------------------
// exported methods
// -----------------------------------------------------------------------------

// load JSON configuration from file
func LoadConfiguration(configurationFile string) *ConfigurationHolder {
	if configurationFile == "" {
		LOG.Crt("name of configuration file must not be empty")
		return nil
	}

	LOG.Dbg("opening configuration file '%s'", configurationFile)
	jsonFile, err := os.Open(configurationFile)

	if err != nil {
		LOG.Crt("unable to open configuration file '%s', error: '%s'", configurationFile, err)
		return nil
	}

	defer jsonFile.Close()

	LOG.Dbg("configuration file '%s' was opened; reading data...", configurationFile)
	jsonData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		LOG.Crt("unable to read configuration file '%s', error: '%s'", configurationFile, err)
		return nil
	}

	var cfg = new(ConfigurationHolder)

	LOG.Dbg("un-marshaling json data", configurationFile)
	err = json.Unmarshal(jsonData, cfg)

	if err != nil {
		LOG.Crt("unable to un-marshall configuration file '%s', error: '%s'", configurationFile, err)
		return nil
	}

	configuration = cfg

	printConfiguration()

	LOG.Dbg("configuration file was loaded")
	return configuration
}

// return pointer to configuration holder
func Get() *ConfigurationHolder {
	return configuration
}
