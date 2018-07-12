package server

import (
	"flag"
	"fmt"
	. "todofinder/todofinder"
	"log"
)

//todofinder server mode command tool properties
const (
	CommandName  = "server"
	Description  = "run in server mode"
	usageMessage = "%s [-config path_to_config] \n"
)

//Options structure for server mode
type option struct {
	configPath string
	flagSet    *flag.FlagSet
}

//Initialise command options for server mode
func newOption(eh flag.ErrorHandling) (opt *option) {
	opt = &option{
		flagSet: flag.NewFlagSet(CommandName, eh),
	}
	// option settings
	opt.flagSet.StringVar(&opt.configPath, "config", "", "configuration file path")

	return opt
}

//Validate the flags for server mode
func (opt *option) parse(args []string) (err error) {
	if err = opt.flagSet.Parse(args); err != nil {
		return
	}
	//Check extra flags
	if nonFlag := opt.flagSet.Args(); len(nonFlag) != 0 {
		return fmt.Errorf("invalid argument: %v", nonFlag)
	}

	if opt.configPath == "" {
		return fmt.Errorf("missing flag: %v", "config")
	}
	return
}

// Provide usage for server mode
func Usage() {
	fmt.Printf(usageMessage, CommandName)
}

// Print the default flag for server mode
func PrintDefaults(eh flag.ErrorHandling) {
	opt := newOption(eh)
	opt.flagSet.PrintDefaults()
}

// Execute todofinder in server mode
// It receives arg from command line and validate them
// and run the command
func Run(args []string) error {
	opt := newOption(flag.ExitOnError)
	if e := opt.parse(args); e != nil {
		Usage()
		PrintDefaults(flag.ExitOnError)
		return fmt.Errorf("%v, %v", CommandName, e)
	}
	return command(opt)
}

// Execute todofinder in server mode
func command(opt *option) error {
	config, err := LoadConfiguration(&opt.configPath)
	if err != nil {
		return err
	}

	server := &Server{}
	server.Init(config)
	if rerr := server.Run(); rerr != nil {
		log.Fatal("[ERROR] Couldn't run: ", rerr)
	}
	return nil
}
