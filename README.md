Run the example code at [ethereum-bn256/main.go](https://github.com/SadPencil/an-issue-of-go-bn256/blob/master/ethereum-bn256/main.go).

## The issue

`p` and `q` are both type `*bn256.G1`. 

Now let's calculate p=p+q.

```go
p.Add(p,q)
```

## What happend

Point `p` should be calculate correctly, and thus on the curve.

## What actually happend

The value of `p` is incorrect, and it's not on the curve. The following code fails:

```go
var t bn256.G1
_, err := t.Unmarshal(p.Marshal())
```

## Note that

In `bn256_test.go` file inside the library,

```go
func TestOrderG1(t *testing.T) {
    // omited codes...
    g.Add(g, one)
    // omited codes...    
}
```

`g.Add(g, one)` shows that invoking Add() method like this is allowed.

