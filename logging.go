package p9p

import (
	"context"
	"log"
	"os"
)

type logging struct {
	session Session
	logger  log.Logger
}

var _ Session = &logging{}

func NewLogger(prefix string, session Session) Session {
	return &logging{
		session: session,
		logger:  *log.New(os.Stdout, prefix, 0),
	}
}

func (l *logging) Auth(ctx context.Context, afid Fid, uname, aname string) (Qid, error) {
	qid, err := l.session.Auth(ctx, afid, uname, aname)
	l.logger.Printf("Auth(%v, %s, %s) -> (%v, %v)", afid, uname, aname, qid, err)
	return qid, err
}

func (l *logging) Attach(ctx context.Context, fid, afid Fid, uname, aname string) (Qid, error) {
	qid, err := l.session.Attach(ctx, fid, afid, uname, aname)
	l.logger.Printf("Attach(%v, %v, %s, %s) -> (%v, %v)", fid, afid, uname, aname, qid, err)
	return qid, err
}

func (l *logging) Clunk(ctx context.Context, fid Fid) error {
	err := l.session.Clunk(ctx, fid)
	l.logger.Printf("Clunk(%v) -> %v", fid, err)
	return err
}

func (l *logging) Remove(ctx context.Context, fid Fid) (err error) {
	defer func() {
		l.logger.Printf("Remove(%v) -> %v", fid, err)
	}()
	return l.session.Remove(ctx, fid)
}

func (l *logging) Walk(ctx context.Context, fid Fid, newfid Fid, names ...string) ([]Qid, error) {
	qid, err := l.session.Walk(ctx, fid, newfid, names...)
	l.logger.Printf("Walk(%v, %v, %s) -> (%v, %v)", fid, newfid, names, qid, err)
	return qid, err
}

func (l *logging) Read(ctx context.Context, fid Fid, p []byte, offset int64) (n int, err error) {
	defer func() {
		l.logger.Printf("Read(%v, bytes[], %v) -> (%v, %v)", fid, offset, n, err)
	}()
	return l.session.Read(ctx, fid, p, offset)
}

func (l *logging) Write(ctx context.Context, fid Fid, p []byte, offset int64) (n int, err error) {
	defer func() {
		l.logger.Printf("Write(%v, bytes[], %v) -> (%v, %v)", fid, offset, n, err)
	}()
	return l.session.Write(ctx, fid, p, offset)
}

func (l *logging) Open(ctx context.Context, fid Fid, mode Flag) (q Qid, n uint32, err error) {
	defer func() {
		l.logger.Printf("Open(%v, %v) -> (%v, %v, %s)", fid, mode, q, n, err)
	}()
	return l.session.Open(ctx, fid, mode)
}

func (l *logging) Create(ctx context.Context, parent Fid, name string, perm uint32, mode Flag) (q Qid, n uint32, err error) {
	defer func() {
		l.logger.Printf("Create(%v, %s, %v, %v) -> (%v, %v, %s)", parent, name, perm, mode, q, n, err)
	}()
	return l.session.Create(ctx, parent, name, perm, mode)
}

func (l *logging) Stat(ctx context.Context, fid Fid) (Dir, error) {
	return l.session.Stat(ctx, fid)
}

func (l *logging) WStat(ctx context.Context, fid Fid, dir Dir) error {
	return l.session.WStat(ctx, fid, dir)
}

func (l *logging) Version() (int, string) {
	i, v := l.session.Version()
	l.logger.Printf("Version() -> (%v, %s)", i, v)
	return i, v
}
