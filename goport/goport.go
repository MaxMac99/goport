package goport

import (
	"io"
	"log"
	"sync"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/config"
	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/cli/cli/context/store"
	"github.com/docker/cli/cli/flags"
	"github.com/docker/cli/cli/streams"
	"github.com/moby/term"
)

var once sync.Once

type GoPort interface {
	ConfigFile() *configfile.ConfigFile
	ContextStore() store.Store
	CurrentContext() string
	Out() *streams.Out
	Err() io.Writer
	In() *streams.In
	SetIn(in *streams.In)
}

type GoPortServer struct {
	configFile   *configfile.ConfigFile
	contextStore store.Store
	in           *streams.In
	out          *streams.Out
	err          io.Writer
}

var goPort *GoPortServer

func (g *GoPortServer) ContextStore() store.Store {
	return g.contextStore
}

func (g *GoPortServer) CurrentContext() string {
	return "default"
}

func (g *GoPortServer) ConfigFile() *configfile.ConfigFile {
	return g.configFile
}

// Out returns the writer used for stdout
func (g *GoPortServer) Out() *streams.Out {
	return g.out
}

// Err returns the writer used for stderr
func (g *GoPortServer) Err() io.Writer {
	return g.err
}

// SetIn sets the reader used for stdin
func (g *GoPortServer) SetIn(in *streams.In) {
	g.in = in
}

// In returns the reader used for stdin
func (g *GoPortServer) In() *streams.In {
	return g.in
}

func GetGoPort() *GoPortServer {
	once.Do(func() {
		goPort = &GoPortServer{}
		stdin, stdout, stderr := term.StdStreams()
		goPort.in = streams.NewIn(stdin)
		goPort.out = streams.NewOut(stdout)
		goPort.err = stderr
		goPort.configFile = config.LoadDefaultConfigFile(goPort.err)
		baseContextStore := store.New(config.ContextStoreDir(), command.DefaultContextStoreConfig())
		commonOpts := flags.CommonOptions{
			Debug: true,
		}
		goPort.contextStore = &command.ContextStoreWithDefault{
			Store: baseContextStore,
			Resolver: func() (*command.DefaultContext, error) {
				return command.ResolveDefaultContext(&commonOpts, goPort.configFile, command.DefaultContextStoreConfig(), stderr)
			},
		}
	})

	log.Printf("Config File: " + goPort.ConfigFile().Filename)
	return goPort
}
