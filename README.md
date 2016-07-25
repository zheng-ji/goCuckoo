## goCuckoo

[![Build Status](https://travis-ci.org/zheng-ji/goCuckoo.svg)](https://travis-ci.org/zheng-ji/goCuckoo)
[![GoDoc](https://godoc.org/github.com/zheng-ji/goCuckoo?status.svg)](https://godoc.org/github.com/zheng-ji/goCuckoo)

A Cuckoo hashing, substituting for bloom filter. written by Go

一个 Cuckoo Filter 的 Go 库

![goCuckoo](https://cloud.githubusercontent.com/assets/1414745/17084380/8c3a4896-51ee-11e6-869e-b087226cc5ce.jpg)

Description
-----------

Bloom Filter 的位图模式带来两个问题：一是误报(false positives)，它能判断元素一定不存在，但只能判断可能存在，因为存在其它元素被映射到部分相同位上，导致该位置1，那么一个不存在的元素可能会被误报成存在；另一个是漏报（false nagatives），如果删除了某个元素，导致该映射位被置0，那么本来存在的元素会被漏报成不存在。 

Cuckoo Filter，它既可以确保该元素存在的必然性，又可以在不违背此前提下删除任意元素，仅仅比 Bloom Filter 牺牲了微量空间效率。

Feature
--------

* Deletion Support
* FastLoopUp O(1)
* High Space Utilization,4-way set-associative table: > 95% entries occupied
* Subsituting for Bloom Filter.

Installation
-------------

```
go get github.com/zheng-ji/goCuckoo
```

Example
-------

```go
import (
	"fmt"
	"github.com/zheng-ji/goCuckoo"
)

func main() {
    // speicify capacity 
	filter := cuckoo.NewCuckooFilter(10000)

	filter.Insert([]byte("zheng-ji"))
	filter.Insert([]byte("stupid"))
	filter.Insert([]byte("coder"))

	if filter.Find([]byte("stupid")) {
		fmt.Println("exist")
	} else {
		fmt.Println("Not exist")
	}

	filter.Del([]byte("stupid"))
	filter.Println(filter.Size())
}
```

Documentation
-------------

- [CMU Paper](http://www.cs.cmu.edu/~binfan/papers/conext14_cuckoofilter.pdf)
- [CMU PPT](http://www.cs.cmu.edu/~binfan/papers/conext14_cuckoofilter.pptx)
- [CoolShell Article](http://coolshell.cn/articles/17225.html)

License
-------

Copyright (c) 2016 by [zheng-ji](http://zheng-ji.info) released under MIT License.

