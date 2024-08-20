# diffbench
Analysis of different versions of go compilers, pgo debug information display inline optimization differences

# Use
flag information:
<pre>
  -new string
        new benchmark output file
  -old string
        old benchmark output file
</pre>

use example:

a command like this gets pgo debug information:

> go build -gcflags=all=-d=pgodebug=1 2> old.txt

diffbench -new /new.txt -old /old.txt

output structare is:

<pre>
// old inline , but new not inline
......(if have difference)
----------------
// new inline , but old not inline
......(if have difference)
</pre>
