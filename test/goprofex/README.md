Profiling and optimizing Go web applications: Example
=====================================================

Example showing how to profile and optimize Go web applications.

The initial version is tagged as [v1](https://github.com/akrylysov/goprofex/tree/v1) and the optimized version is tagged as [v2](https://github.com/akrylysov/goprofex/tree/v2).
Here is [the link](https://github.com/akrylysov/goprofex/compare/v1...v2) to compare these two versions.

You can find more details at <http://artem.krylysov.com/blog/2017/03/13/profiling-and-optimizing-go-web-applications/>.


## API

```
$ http://127.0.0.1:8080/v1/leftpad/?str=hello&len=10&chr=*

$ ab -k -c 5 -n 20000 "http://127.0.0.1:8080/v1/leftpad/?str=hello&len=10&chr=*"

$ go tool pprof goprofex http://127.0.0.1:8080/debug/pprof/profile
```

## install

```
yum install graphviz xdg-utils
```

## Console
```
[monkey@zhengyscn ~]$ go tool pprof goprofex http://127.0.0.1:8080/debug/pprof/profile
Fetching profile over HTTP from http://127.0.0.1:8080/debug/pprof/profile
goprofex: open goprofex: no such file or directory
fetched 1 profiles out of 2
Saved profile in /home/monkey/pprof/pprof.goprofex.samples.cpu.001.pb.gz
File: goprofex
Type: cpu
Time: Aug 22, 2018 at 11:42pm (CST)
Duration: 30s, Total samples = 4.74s (15.80%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 2240ms, 47.26% of 4740ms total
Dropped 174 nodes (cum <= 23.70ms)
Showing top 10 nodes out of 152
      flat  flat%   sum%        cum   cum%
     990ms 20.89% 20.89%     1040ms 21.94%  syscall.Syscall /home/monkey/local/go/src/syscall/asm_linux_amd64.s
     750ms 15.82% 36.71%      750ms 15.82%  runtime._ExternalCode /home/monkey/local/go/src/runtime/proc.go
      90ms  1.90% 38.61%       90ms  1.90%  runtime.usleep /home/monkey/local/go/src/runtime/sys_linux_amd64.s
      80ms  1.69% 40.30%      280ms  5.91%  runtime.mallocgc /home/monkey/local/go/src/runtime/malloc.go
      60ms  1.27% 41.56%      110ms  2.32%  runtime.deferreturn /home/monkey/local/go/src/runtime/panic.go
      60ms  1.27% 42.83%       60ms  1.27%  runtime.getitab /home/monkey/local/go/src/runtime/iface.go
      60ms  1.27% 44.09%       90ms  1.90%  runtime.mapaccess1_faststr /home/monkey/local/go/src/runtime/hashmap_fast.go
      50ms  1.05% 45.15%       50ms  1.05%  runtime.mapiternext /home/monkey/local/go/src/runtime/hashmap.go
      50ms  1.05% 46.20%       50ms  1.05%  runtime.scanobject /home/monkey/local/go/src/runtime/mgcmark.go
      50ms  1.05% 47.26%      100ms  2.11%  time.Time.Sub /home/monkey/local/go/src/time/time.go
(pprof) web

(pprof) svg
Generating report in profile002.svg
(pprof) list leftpad
Total: 3.52s
ROUTINE ======================== main.leftpad in /home/monkey/go/src/github.com/zhengyscn/golang-doc/test/goprofex/leftpad.go
         0       10ms (flat, cum)  0.28% of Total
         .          .      4:	"strings"
         .          .      5:)
         .          .      6:
         .          .      7:func leftpad(s string, length int, char rune) string {
         .          .      8:	if len(s) < length {
         .       10ms      9:		return strings.Repeat(string(char), length-len(s)) + s
         .          .     10:	}
         .          .     11:	return s
         .          .     12:}
ROUTINE ======================== main.leftpadHandler in /home/monkey/go/src/github.com/zhengyscn/golang-doc/test/goprofex/handlers.go
      10ms      770ms (flat, cum) 21.88% of Total
         .          .     23:		statsd.Incr(countMetric)
         .          .     24:	}
         .          .     25:}
         .          .     26:
         .          .     27:func leftpadHandler(w http.ResponseWriter, r *http.Request) {
         .      200ms     28:	log.Printf("url %s\tip %s\tua %s", r.RequestURI, r.RemoteAddr, r.UserAgent())
         .      140ms     29:	str := r.FormValue("str")
         .       20ms     30:	length, err := strconv.Atoi(r.FormValue("len"))
         .          .     31:	if err != nil {
         .          .     32:		http.Error(w, err.Error(), http.StatusBadRequest)
         .          .     33:		return
         .          .     34:	}
         .      120ms     35:	statsd.Histogram("request.str.len", float64(length))
         .          .     36:	chr := ' '
      10ms       20ms     37:	if len(r.FormValue("chr")) > 0 {
         .       10ms     38:		chr = []rune(r.FormValue("chr"))[0]
         .          .     39:	}
         .       20ms     40:	w.Header().Set("Content-Type", "application/json; charset=utf-8")
         .       10ms     41:	resp := leftpadResponse{Str: leftpad(str, length, chr)}
         .       10ms     42:	enc := json.NewEncoder(w)
         .      220ms     43:	if err := enc.Encode(resp); err != nil {
         .          .     44:		http.Error(w, err.Error(), http.StatusInternalServerError)
         .          .     45:		return
         .          .     46:	}
         .          .     47:}
(pprof)
```