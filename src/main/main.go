package main

// -----------------------------------------------------------------------------
// module imports
// -----------------------------------------------------------------------------
import (
	CFG "app/src/cfg"
	LOG "app/src/log"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

// -----------------------------------------------------------------------------
// type definitions
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// module global variables (non-exported)
// -----------------------------------------------------------------------------

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

	signal.Notify(signalChannel)

	go func() {
		for {
			signalNumber := <-signalChannel
			LOG.Dbg("received '%s' signal; signum = '%d'", signalNumber.String(), signalNumber)

			switch signalNumber {
			case syscall.SIGABRT:
			case syscall.SIGALRM:
			case syscall.SIGBUS:
			case syscall.SIGCHLD:
			// case syscall.SIGCLD: // duplicate value (SIGCHLD)
			case syscall.SIGCONT:
			case syscall.SIGFPE:
			case syscall.SIGHUP:
				StopApplication()
			case syscall.SIGILL:
			case syscall.SIGINT:
				StopApplication()
			// case syscall.SIGIO: // duplicate value (SIGPOLL)
			// case syscall.SIGIOT: // duplicate value (SIGABRT)
			case syscall.SIGKILL:
			case syscall.SIGPIPE:
			case syscall.SIGPOLL:
			case syscall.SIGPROF:
			case syscall.SIGPWR:
			case syscall.SIGQUIT:
				StopApplication()
			case syscall.SIGSEGV:
			case syscall.SIGSTKFLT:
			case syscall.SIGSTOP:
			// case syscall.SIGSYS: // duplicate value (SIGUNUSED)
			case syscall.SIGTERM:
				StopApplication()
			case syscall.SIGTRAP:
			case syscall.SIGTSTP:
			case syscall.SIGTTIN:
			case syscall.SIGTTOU:
			case syscall.SIGUNUSED:
			case syscall.SIGURG:
			case syscall.SIGUSR1:
			case syscall.SIGUSR2:
			case syscall.SIGVTALRM:
			case syscall.SIGWINCH:
			case syscall.SIGXCPU:
			case syscall.SIGXFSZ:
			default:
				LOG.Wrn("received signal '%s' is unknown; signum = '%d'", signalNumber.String(), signalNumber)
			}
		}
	}()
}

// check and process configuration file argument
func processArgumets() int {
	var cfgFile string
	flag.StringVar(&cfgFile, "c", "", "[-c cfg_file.json | --c cfg_file.json | -c=cfg_file.json | --c=cfg_file.json]")
	flag.StringVar(&cfgFile, "cfg", "", "[-cfg cfg_file.json | --cfg cfg_file.json | -cfg=cfg_file.json | --cfg=cfg_file.json]")

	flag.Parse()

	if cfgFile == "" {
		flag.Usage()
		return rcExitFailure
	} else {
		appData.cfgFile = cfgFile
	}

	return rcExitSuccess
}

// initialize application runtime
func initApplication() int {
	// setup logging as a first step
	if !LOG.Initialize(LOG.GetDefaultLevel()) {
		panic("unable to initialize logging!")
	}

	LOG.Inf("starting application...")

	setupSignalHandling()

	if processArgumets() != rcExitSuccess {
		return rcExitFailure
	}

	if CFG.LoadConfiguration(appData.cfgFile) == nil {
		return rcExitFailure
	}

	LOG.SetLevel(CFG.Get().LogLevel)

	// TODO: put another init steps here

	// TODO: ---------------------------

	return rcExitSuccess
}

// perform application runtime shutdown
func shutdownApplication() {

	// TODO: put another shutdown steps here

	// TODO: -------------------------------

	LOG.Inf("application runtime stopped...")
}

// main function and entry point
func main() {

	// initialize application; if initialization failed end immediately
	var rc = initApplication()

	if rc != rcExitSuccess {
		LOG.Crt("unable to initialize application")
		os.Exit(rc)
	}

	// run main application
	rc = RunApplication()

	if rc == rcExitSuccess {
		LOG.Inf("application runtime ended without errors; code '%d'", rc)
	} else {
		LOG.Err("application runtime ended with an error; code '%d'", rc)
	}

	// perform application shutdown
	shutdownApplication()

	os.Exit(rc)
}

// -----------------------------------------------------------------------------
// exported methods
// -----------------------------------------------------------------------------
