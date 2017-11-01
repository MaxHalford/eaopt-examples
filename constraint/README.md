# Constrained optimization

You can do constrained optimization simply by tinkering with the objective function. For example consider the following `Evaluate` method.

```go
type Vector []float64

func (X Vector) Evaluate() (y float64) {
    y = math.Pow(X[0], 2) + math.Pow(X[1], 2)
    if math.Abs(X[0])+math.Abs(X[1]) < 4 {
        y += 10000
    }
    return
}
```

Here any value that is in the box `abs(X[0]) + abs(X[1]) < 4` will receive a penalty of 10000. This way of doing works because unwanted solutions will be considered weak. However, they will still be evaluated which wastes useless CPU cycles. Another way of doing would be to make sure that the `Mutate` and `Crossover` produce solutions that are in a desired range.

Run `go run main.go && python plot_progress.py` in a terminal to get the following kind of animation.

<div align="center">
  <img src="progress.gif" alt="progress" />
</div>
