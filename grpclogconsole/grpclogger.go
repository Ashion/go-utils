package grpclogger

import "fmt"

type GRPCLogger struct {
}

// Info returns
func (gLogger *GRPCLogger) Info(args ...interface{}) {
	fmt.Println(args...)
}

// Infoln returns
func (gLogger *GRPCLogger) Infoln(args ...interface{}) {
	fmt.Println(args...)
}

// Infof returns
func (gLogger *GRPCLogger) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Warning returns
func (gLogger *GRPCLogger) Warning(args ...interface{}) {
	fmt.Println(args...)
}

// Warningln returns
func (gLogger *GRPCLogger) Warningln(args ...interface{}) {
	fmt.Println(args...)
}

// Warningf returns
func (gLogger *GRPCLogger) Warningf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Error returns
func (gLogger *GRPCLogger) Error(args ...interface{}) {
	fmt.Println(args...)
}

// Errorln returns
func (gLogger *GRPCLogger) Errorln(args ...interface{}) {
	fmt.Println(args...)
}

// Errorf returns
func (gLogger *GRPCLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Fatal returns
func (gLogger *GRPCLogger) Fatal(args ...interface{}) {
	fmt.Println(args...)
}

// Fatalln returns
func (gLogger *GRPCLogger) Fatalln(args ...interface{}) {
	fmt.Println(args...)
}

// Fatalf logs to fatal level
func (gLogger *GRPCLogger) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// V reports whether verbosity level l is at least the requested verbose level.
func (gLogger *GRPCLogger) V(v int) bool {
	return false
}
