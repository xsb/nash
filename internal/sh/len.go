package sh

import (
	"io"
	"strconv"

	"github.com/NeowayLabs/nash/ast"
	"github.com/NeowayLabs/nash/errors"
)

type (
	LenFn struct {
		stdin          io.Reader
		stdout, stderr io.Writer

		done    chan struct{}
		err     error
		results int

		arg []*Obj
	}
)

func NewLenFn(env *Shell) *LenFn {
	return &LenFn{
		stdin:  env.stdin,
		stdout: env.stdout,
		stderr: env.stderr,
	}
}

func (lenfn *LenFn) Name() string {
	return "len"
}

func (lenfn *LenFn) ArgNames() []string {
	return append(make([]string, 0, 1), "list")
}

func (lenfn *LenFn) run() error {
	lenfn.results = len(lenfn.arg)
	return nil
}

func (lenfn *LenFn) Start() error {
	lenfn.done = make(chan struct{})

	go func() {
		lenfn.err = lenfn.run()
		lenfn.done <- struct{}{}
	}()

	return nil
}

func (lenfn *LenFn) Wait() error {
	<-lenfn.done
	return lenfn.err
}

func (lenfn *LenFn) Results() *Obj {
	return NewStrObj(strconv.Itoa(lenfn.results))
}

func (lenfn *LenFn) SetArgs(args []ast.Expr, envShell *Shell) error {
	if len(args) != 1 {
		return errors.NewError("lenfn expects one argument")
	}

	obj, err := envShell.evalExpr(args[0])

	if err != nil {
		return err
	}

	if obj.Type() != ListType {
		return errors.NewError("lenfn expects a list, but a %s was provided", obj.Type())
	}

	lenfn.arg = obj.List()
	return nil
}

func (lenfn *LenFn) SetEnviron(env []string) {
	// do nothing
}

func (lenfn *LenFn) SetStdin(r io.Reader)  { lenfn.stdin = r }
func (lenfn *LenFn) SetStderr(w io.Writer) { lenfn.stderr = w }
func (lenfn *LenFn) SetStdout(w io.Writer) { lenfn.stdout = w }
func (lenfn *LenFn) StdoutPipe() (io.ReadCloser, error) {
	return nil, errors.NewError("lenfn doesn't works with pipes")
}
func (lenfn *LenFn) Stdin() io.Reader  { return lenfn.stdin }
func (lenfn *LenFn) Stdout() io.Writer { return lenfn.stdout }
func (lenfn *LenFn) Stderr() io.Writer { return lenfn.stderr }

func (lenfn *LenFn) String() string { return "<builtin fn len>" }