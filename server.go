package cmds

import "errors"

// Code for an action
type Code = string

// Action do something
type Action = func(s string) error

// Set is the set list for code and action
type Set = map[Code]Action

var (
	// ErrRegistered is the code has been registered
	ErrRegistered = errors.New("this code has beed registered")
)

// CMDServer interface
type CMDServer interface {
	Send(code Code, param string) error
	Register(code Code, action Action, force bool) error
}

type srvImpl struct {
	actSet Set
}

// InitCMDS initial a commander service
func InitCMDS(set Set) CMDServer {
	return &srvImpl{
		actSet: set,
	}
}

func (s *srvImpl) Send(code Code, param string) error {
	err := s.actSet[code](param)
	return err
}

func (s *srvImpl) Register(code Code, action Action, force bool) error {
	if force {
		s.actSet[code] = action
		return nil
	}

	_, ok := s.actSet[code]
	if ok {
		return ErrRegistered
	}

	s.actSet[code] = action
	return nil
}
