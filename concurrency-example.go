package main

import (
  "context"
)

func main() {
  var (
    ctx    context.Context
    cancel context.CancelFunc
  )
  ctx, cancel = context.WithCancel(context.Background())
  defer cancel()

  in := gen(ctx)

  ch1 := process(ctx, in)
  ch2 := process(ctx, in)

  for v := range merge(ch1, ch2) {

  }
}

func process(ctx context.Context, in <-chan int) <-chan int {
  out := make(<-chan int)
  go func() {
    defer close(out)

    for v := range in {
      select {
      case <-ctx.Done():
        return
      default:
        out <- v
      }
    }
  }

  return out
}

func merge(ctx context.Context, pipe ...<-chan int) <-chan int {
  trace := "main#merge"
  internal.Logs("Enter", trace, nil)
  defer internal.Logs("Exit", trace, nil)

  out := make(chan *github.Event)
  var wg sync.WaitGroup

  output := func(c <-chan *github.Event) {
    defer wg.Done()
    for n := range c {
      select {
        case <-ctx.Done():
          return
        case out <- n:
      }
    }
  }
  wg.Add(len(cs))

  for _, c := range cs {
    go output(c)
  }

  go func() {
    wg.Wait()
    close(out)
  }()
  return out
}