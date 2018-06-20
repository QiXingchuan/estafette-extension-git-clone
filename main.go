package main

import (
	stdlog "log"
	"os"
	"runtime"

	"github.com/alecthomas/kingpin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	version   string
	branch    string
	revision  string
	buildDate string
	goVersion = runtime.Version()
)

var (
	// flags
	gitName      = kingpin.Flag("git-name", "The owner plus repository name.").Envar("ESTAFETTE_GIT_NAME").Required().String()
	gitURL       = kingpin.Flag("git-url", "The authenticated url to clone.").Envar("ESTAFETTE_GIT_URL").Required().String()
	gitBranch    = kingpin.Flag("git-branch", "The branch to clone.").Envar("ESTAFETTE_GIT_BRANCH").Required().String()
	gitRevision  = kingpin.Flag("git-revision", "The revision to check out.").Envar("ESTAFETTE_GIT_REVISION").Required().String()
	shallowClone = kingpin.Flag("shallow-clone", "Shallow clone git repository for improved clone time.").Default("true").OverrideDefaultFromEnvar("ESTAFETTE_EXTENSION_SHALLOW").Bool()
)

func main() {

	// parse command line parameters
	kingpin.Parse()

	// pretty print to make build logs more readable
	log.Logger = zerolog.New(os.Stdout).With().
		Logger()

	// use zerolog for any logs sent via standard log library
	stdlog.SetFlags(0)
	stdlog.SetOutput(log.Logger)

	// log startup message
	log.Info().
		Str("branch", branch).
		Str("revision", revision).
		Str("buildDate", buildDate).
		Str("goVersion", goVersion).
		Msg("Starting estafette-extension-git-clone...")

	// git clone to specific branch and revision
	err := gitCloneRevision(*gitName, *gitURL, *gitBranch, *gitRevision, *shallowClone)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error cloning git repository %v to branch %v and revision %v with shallow clone is %v...", *gitName, *gitBranch, *gitRevision, *shallowClone)
	}
}
