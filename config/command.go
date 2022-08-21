package conf

type CommandRunner interface{}

// Command interface is used to implement sub-commands in the system.
type Command interface {

	//
	Run() CommandRunner
}
