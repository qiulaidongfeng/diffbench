package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	old := flag.String("old", "", "old benchmark output file")
	new := flag.String("new", "", "new benchmark output file")
	flag.Parse()
	if *old == "" || *new == "" {
		flag.Usage()
		return
	}

	oldf, err := os.ReadFile(*old)
	if err != nil {
		panic(err)
	}
	newf, err := os.ReadFile(*new)
	if err != nil {
		panic(err)
	}

	olds := string(oldf)
	news := string(newf)

	oldd := keepInline(olds)
	newd := keepInline(news)

	for k := range newd.m {
		if _, ok := oldd.m[k]; ok {
			delete(oldd.m, k)
			delete(newd.m, k)
		}
	}
	if len(oldd.m) != 0 {
		fmt.Println("// old inline , but new not inline")
		out := make([]string, 0, len(oldd.m))
		for _, v := range oldd.m {
			out = append(out, v.raw)
		}
		slices.Sort(out)
		for i := range out {
			fmt.Println(out[i])
		}
	}

	fmt.Println("----------------")

	if len(newd.m) != 0 {
		fmt.Println("// new inline , but old not inline")
		out := make([]string, 0, len(newd.m))
		for _, v := range newd.m {
			out = append(out, v.raw)
		}
		slices.Sort(out)
		for i := range out {
			fmt.Println(out[i])
		}
	}
}

func keepInline(s string) data {
	v := strings.Split(s, "\n")
	ret := data{m: make(map[string]struct {
		raw    string
		caller string
	})}
	v = slices.DeleteFunc(v, func(s string) bool {
		if strings.Contains(s, "escape") {
			ret.m[s] = struct {
				raw    string
				caller string
			}{raw: s}
		}
		return !strings.Contains(s, "check allows inlining") && !strings.Contains(s, "inlining call to")
	})
	for i := range v {
		l := strings.Split(v[i], " ")
		for j := range l {
			if l[j] == "inlining" && l[j+1] == "call" {
				at := l[j-1]
				funcname := l[j+3]
				ret.m[at+funcname] = struct {
					raw    string
					caller string
				}{
					raw: v[i],
				}
			}
			if l[j] == "call" {
				caller := ""
				at := ""
				for k := j + 2; k < len(l); k++ {
					if l[k] == "function" {
						caller = l[k+1]
					}
					if l[k] == "at" {
						at = l[k+1]
					}
				}
				ret.m[l[j+1]+at+caller] = struct {
					raw    string
					caller string
				}{
					raw:    v[i],
					caller: caller,
				}
			}
		}
	}
	return ret
}

type data struct {
	m map[string]struct {
		raw    string
		caller string
	}
}
