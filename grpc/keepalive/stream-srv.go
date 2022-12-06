package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/gunsluo/go-example/grpc/keepalive/pb"
	"google.golang.org/grpc"
)

func main() {
	gSrv := NewG()

	gSrv.Run()
}

type gServer struct {
	pb.UnimplementedGreeterServer

	port int
	gRPC *grpc.Server
}

func NewG() *gServer {
	gRPC := grpc.NewServer(
	// grpc.KeepaliveParams(keepalive.ServerParameters{
	// 	// MaxConnectionIdle:     24 * time.Hour,   // The current default value is infinity.
	// 	MaxConnectionAge:      2 * time.Minute,  // The current default value is infinity.
	// 	MaxConnectionAgeGrace: 30 * time.Second, // The current default value is infinity.
	// 	// Time:                  2 * time.Hour,    // The current default value is 2 hours.
	// 	// Timeout:               20 * time.Second, // The current default value is 20 seconds.
	// }),
	// grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
	// 	MinTime:             time.Minute, // The current default value is 5 minutes.
	// 	PermitWithoutStream: false,       // false by default.
	// }),
	)

	s := &gServer{
		port: 30000,
		gRPC: gRPC,
	}
	pb.RegisterGreeterServer(gRPC, s)

	return s
}

func (s *gServer) Run() {
	fmt.Printf("Starting gRPC server on listening %d.\n", s.port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		panic(err)
	}

	s.gRPC.Serve(listener)
}

func (s *gServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	st := time.Now()
	time.Sleep(time.Minute)
	fmt.Println("say hello:", time.Now().Sub(st))
	return &pb.HelloReply{Message: "hello" + req.Name}, nil
}

func (s *gServer) SayHelloProgress(req *pb.HelloRequest, srv pb.Greeter_SayHelloProgressServer) error {
	st := time.Now()

	p := NewProgressBar()
	err := p.Start(30*time.Second,
		s.handleSayHello,
		func(v int) error {
			if err := srv.Send(&pb.SayHelloProgressReply{Progress: int64(v)}); err != nil {
				fmt.Println("err:", err)
				return err
			}
			fmt.Println("current progress:", v)
			return nil
		},
	)
	if err != nil {
		fmt.Println("err:", time.Now().Sub(st), err)
		return err
	}

	fmt.Println("say hello:", time.Now().Sub(st))
	return nil
}

func (s *gServer) handleSayHello(ctx context.Context) error {
	// use ctx to cancel gorouting
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	select {
	case <-time.After(time.Minute):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("cancel")
	}

	// time.Sleep(time.Minute)
	return nil
}

type ProgressBar struct {
	progress     chan int
	stopProgress chan struct{}

	task chan error
	exit chan error

	finished bool
	lock     sync.RWMutex
}

func NewProgressBar() *ProgressBar {
	return &ProgressBar{
		progress:     make(chan int),
		stopProgress: make(chan struct{}),
		task:         make(chan error),
		exit:         make(chan error),
	}
}

func (p *ProgressBar) Start(period time.Duration, do func(context.Context) error, notify func(int) error) error {
	ctx := context.Background()

	// a timer and progress
	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		defer func() {
			fmt.Println("progress stop")
		}()
		var count int
		timer := time.NewTimer(period)
		for {
			select {
			case <-timer.C:
				if p.hasDone() {
					return
				}

				count++
				v := count * 20
				if v >= 100 {
					v = 99
				}

				p.progress <- v
				timer.Reset(period)
			case <-p.stopProgress:
				break
			}
		}
	}(ctx)

	// running task
	go func(ctx context.Context) {
		defer func() {
			fmt.Println("task stop")
		}()
		err := do(ctx)
		if p.hasDone() {
			return
		}
		p.task <- err
	}(ctx)

	// progress executor
	go func(ctx context.Context) {
		defer func() {
			fmt.Println("progress executor stop")
		}()

		for {
			select {
			case v, ok := <-p.progress:
				if !ok {
					return
				}

				if err := notify(v); err != nil {
					p.done()
					p.stopProgress <- struct{}{}
					p.exit <- err
					cancel()
					return
				}
			}
		}
	}(ctx)

	// task executor
	go func(ctx context.Context) {
		defer func() {
			fmt.Println("task executor stop")
		}()

		select {
		case err, ok := <-p.task:
			if !ok {
				return
			}

			if err != nil {
				p.done()
				p.stopProgress <- struct{}{}
				p.exit <- err
				return
			}

			p.done()
			p.stopProgress <- struct{}{}
			p.exit <- notify(100)
			return
		}
	}(ctx)

	defer func() {
		close(p.progress)
		close(p.task)
		close(p.stopProgress)
		close(p.exit)
	}()

	return <-p.exit
}

func (p *ProgressBar) hasDone() bool {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.finished
}

func (p *ProgressBar) done() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.finished = true
}
