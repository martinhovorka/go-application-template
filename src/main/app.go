package main

import (
	CFG "app/src/cfg"
	LOG "app/src/log"
	"time"
)

// -----------------------------------------------------------------------------
// module imports
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// type definitions
// -----------------------------------------------------------------------------
type AppData struct {
	runMainLoop bool
	cfgFile     string
}

const (
	rcExitSuccess = iota
	rcExitFailure
)

// -----------------------------------------------------------------------------
// module global variables (non-exported)
// -----------------------------------------------------------------------------

var appData = new(AppData)

// -----------------------------------------------------------------------------
// module global variables (exported)
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// non-exported methods
// -----------------------------------------------------------------------------

// initialize application
func initialize() int {
	LOG.Dbg("initializing application")
	appData.runMainLoop = true

	return rcExitSuccess
}

// shutdown and de-initialize application
func shutdown() int {
	LOG.Dbg("shutting down application main loop")

	return rcExitSuccess
}

// main loop
func mainLoop() int {
	LOG.Dbg("running main loop")
	for appData.runMainLoop {
		mainLoopIteration()
		time.Sleep(time.Duration(CFG.Get().MainLoopTimeout) * time.Millisecond)
	}

	return rcExitSuccess
}

// main loop iteration
func mainLoopIteration() {
	LOG.Dbg("main loop hearbeat")
}

// -----------------------------------------------------------------------------
// exported methods
// -----------------------------------------------------------------------------

// stop application execution
func StopApplication() {
	appData.runMainLoop = false
}

// run application
func RunApplication() int {
	LOG.Dbg("starting application")

	var rc = initialize()

	if rc != rcExitSuccess {
		LOG.Crt("application 'initialize()' ended with an error '%d'", rc)
		return rcExitFailure
	}

	rc += mainLoop()

	if rc != rcExitSuccess {
		LOG.Err("application 'mainLoop()' method ended with an error '%d'", rc)
	}

	rc += shutdown()

	if rc != rcExitSuccess {
		LOG.Err("application 'shutdown()' method ended with an error '%d'", rc)
	}

	return rc
}
