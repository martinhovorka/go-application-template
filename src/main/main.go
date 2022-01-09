package main

// -----------------------------------------------------------------------------
// module imports
// -----------------------------------------------------------------------------
import (
	CFG "app/src/cfg"
	LOG "app/src/log"
	"os"
	"os/signal"
)

// -----------------------------------------------------------------------------
// type definitions
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// module global variables (non-exported)
// -----------------------------------------------------------------------------

const (
	maxNumberOfArgument = 2
	cfgFileMinSize      = 2       // 2 bytes/chars for valid JSON
	cfgFileMaxSize      = 1048576 // 1 MB
	cfgUserReadable     = 0x100
)

// -----------------------------------------------------------------------------
// module global variables (exported)
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// non-exported methods
// -----------------------------------------------------------------------------

// setup signal handling
func setupSignalHandling() {
	LOG.Dbg("setting up signal handling")

	signalChannel := make(chan os.Signal, 1)

	// TODO: not necessary to exit on all signals
	signal.Notify(signalChannel)

	go func() {
		s := <-signalChannel
		LOG.Ntc("received '%s' signal; signum = '%d'", s.String(), s)
		StopApplication()
	}()
}

// check and process configuration file argument
func processCfgFileArgument() int {
	if len(os.Args) != maxNumberOfArgument {
		LOG.Crt("invalid number or arguments; expected '%d'", maxNumberOfArgument)
		return rcExitFailure
	}

	appData.cfgFile = os.Args[1]
	fileInfo, error := os.Stat(appData.cfgFile)

	if error != nil {
		LOG.Crt("unable to find/open JSON configuration file '%s'", appData.cfgFile)
		return rcExitFailure
	}

	if fileInfo.IsDir() || !fileInfo.Mode().IsRegular() {
		LOG.Crt("configuration file '%s' is not a regular file", appData.cfgFile)
		return rcExitFailure
	}

	if (fileInfo.Size() < cfgFileMinSize) || (fileInfo.Size() > cfgFileMaxSize) {
		LOG.Crt("configuration file '%s' size invalid; valid size is between '%d' and '%d'", appData.cfgFile, cfgFileMinSize, cfgFileMaxSize)
		return rcExitFailure
	}

	if fileInfo.Mode()&cfgUserReadable == 0 {
		LOG.Crt("configuration file '%s' is not readable", appData.cfgFile)
		return rcExitFailure
	}
	return rcExitSuccess
}

// initialize application runtime
func initApplication() int {
	if !LOG.Initialize(LOG.GetDefaultLevel()) {
		return rcExitFailure
	}

	setupSignalHandling()

	if processCfgFileArgument() != rcExitSuccess {
		return rcExitFailure
	}

	if CFG.LoadConfiguration(appData.cfgFile) == nil {
		return rcExitFailure
	}

	LOG.SetLevel(CFG.Get().LogLevel)

	return rcExitSuccess
}

// perform application runtime shutdown
func shutdownApplication() {
	LOG.Dbg("shutting down main application")
}

// main function and entry point
func main() {
	var rc = initApplication()

	if rc != rcExitSuccess {
		LOG.Crt("unable to initialize application")
		os.Exit(rc)
	}

	rc = RunApplication()

	if rc == rcExitSuccess {
		LOG.Inf("application runtime ended without errors; code '%d'", rc)
	} else {
		LOG.Err("application runtime ended with an error; code '%d'", rc)
	}

	shutdownApplication()

	os.Exit(rc)
}

// -----------------------------------------------------------------------------
// exported methods
// -----------------------------------------------------------------------------
