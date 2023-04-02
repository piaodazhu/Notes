package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type segTree struct {
    l, r   int
    maxVal int
    addVal int
    lc     *segTree
    rc     *segTree
}

func build(l, r int, a []int) *segTree {
    if l == r {
        return &segTree{l, r, a[l], 0, nil, nil}
    }
    mid := (l + r) >> 1
    lc := build(l, mid, a)
    rc := build(mid+1, r, a)
    return &segTree{l, r, max(lc.maxVal, rc.maxVal), 0, lc, rc}
}

func (t *segTree) pushup() {
    t.maxVal = max(t.lc.maxVal, t.rc.maxVal)
}

func (t *segTree) pushdown() {
    if t.addVal != 0 {
        t.lc.maxVal += t.addVal
        t.rc.maxVal += t.addVal
        t.lc.addVal += t.addVal
        t.rc.addVal += t.addVal
        t.addVal = 0
    }
}

func (t *segTree) modify(l, r, x int) {
    if l <= t.l && t.r <= r {
        t.maxVal += x
        t.addVal += x
        return
    }
    t.pushdown()
    mid := (t.l + t.r) >> 1
    if l <= mid {
        t.lc.modify(l, r, x)
    }
    if r > mid {
        t.rc.modify(l, r, x)
    }
    t.pushup()
}

func (t *segTree) query(l, r int) int {
    if l <= t.l && t.r <= r {
        return t.maxVal
    }
    t.pushdown()
    mid := (t.l + t.r) >> 1
    res := math.MinInt
    if l <= mid {
        res = max(res, t.lc.query(l, r))
    }
    if r > mid {
        res = max(res, t.rc.query(l, r))
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    var n, m int
    fmt.Fscan(reader, &n, &m)

    a := make([]int, n+1)
    for i := 1; i <= n; i++ {
        fmt.Fscan(reader, &a[i])
    }

    root := build(1, n, a)

    for i := 0; i < m; i++ {
        var op, l, r, x int
        fmt.Fscan(reader, &op, &l, &r)
        if op == 1 {
            fmt.Fscan(reader, &x)
            root.modify(l, r, x)
        } else {
            res := root.query(l, r)
            fmt.Fprintln(writer, res)
        }
    }
}

// 10 4
// 3 2 1 4 5 6 7 8 9 10
// 1 3 6 3
// 1 7 8 1
// 2 3 8