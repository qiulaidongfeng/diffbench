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

diffbench -new /new.txt -old /old.txt

output structare is:

<pre>
// old inline , but new not inline
......(if have difference)
----------------
// new inline , but old not inline
......(if have difference)
</pre>
