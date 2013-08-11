# Fan In Stragegy

## Basic Example

The output should look something like this...

```
$ ./fanin
2013/08/11 13:44:20 Fanning in input...0
2013/08/11 13:44:20 Fanning in input...1
2013/08/11 13:44:20 Fanning in input...2
2013/08/11 13:44:20 Quote for CRM: USD 10
2013/08/11 13:44:20 Quote for RHT: USD 151
2013/08/11 13:44:20 Quote for MSFT: USD 21
2013/08/11 13:44:20 Quote for CRM: USD 158
2013/08/11 13:44:20 Quote for CRM: USD 16
2013/08/11 13:44:20 Quote for MSFT: USD 84
2013/08/11 13:44:21 Quote for MSFT: USD 174
2013/08/11 13:44:21 Quote for CRM: USD 115
2013/08/11 13:44:21 Quote for RHT: USD 168
2013/08/11 13:44:21 Quote for RHT: USD 190
2013/08/11 13:44:21 Quote for RHT: USD 73
2013/08/11 13:44:21 Quote for MSFT: USD 111
2013/08/11 13:44:21 Quote for CRM: USD 178
2013/08/11 13:44:22 Quote for CRM: USD 72
2013/08/11 13:44:22 Quote for MSFT: USD 188
2013/08/11 13:44:22 Quote for CRM: USD 144
2013/08/11 13:44:22 Quote for RHT: USD 23
2013/08/11 13:44:23 Quote for CRM: USD 3
2013/08/11 13:44:23 Quote for RHT: USD 83
2013/08/11 13:44:23 Quote for MSFT: USD 132
2013/08/11 13:44:23 Quote for CRM: USD 115
2013/08/11 13:44:24 Quote for RHT: USD 20
2013/08/11 13:44:24 Quote for RHT: USD 52
2013/08/11 13:44:24 Quote for MSFT: USD 119
2013/08/11 13:44:24 Quote for MSFT: USD 110
2013/08/11 13:44:24 Quote for CRM: USD 105
2013/08/11 13:44:24 Quote for MSFT: USD 117
2013/08/11 13:44:25 Quote for RHT: USD 49
2013/08/11 13:44:25 Quote for CRM: USD 133
2013/08/11 13:44:25 Quote for MSFT: USD 146
2013/08/11 13:44:25 Quote for RHT: USD 118
2013/08/11 13:44:26 Quote for CRM: USD 104
2013/08/11 13:44:26 Quote for MSFT: USD 24
2013/08/11 13:44:26 Quote for RHT: USD 69
2013/08/11 13:44:26 Quote for CRM: USD 69
2013/08/11 13:44:26 Quote for MSFT: USD 57
2013/08/11 13:44:27 Quote for RHT: USD 112
2013/08/11 13:44:27 fin
panic: fin

goroutine 1 [running]:
log.Panic(0x2210228f78, 0x1, 0x1)
        GOHOME/src/pkg/log/log.go:307 +0xaa
main.main()
        OSSPATH/funwithgo/fanin/main.go:56 +0x1c4

goroutine 3 [sleep]:
time.Sleep(0x2cb41780)
        GOHOME/src/pkg/runtime/ztime_darwin_amd64.c:19 +0x2f
main.func·002(0xb27e0, 0x3)
        OSSPATH/funwithgo/fanin/main.go:76 +0xd1
created by main.quoteSource
        OSSPATH/funwithgo/fanin/main.go:78 +0x93

goroutine 4 [sleep]:
time.Sleep(0x29d7ab80)
        GOHOME/src/pkg/runtime/ztime_darwin_amd64.c:19 +0x2f
main.func·002(0xb3880, 0x3)
        OSSPATH/funwithgo/fanin/main.go:76 +0xd1
created by main.quoteSource
        OSSPATH/funwithgo/fanin/main.go:78 +0x93

goroutine 5 [sleep]:
time.Sleep(0x3703dcc0)
        GOHOME/src/pkg/runtime/ztime_darwin_amd64.c:19 +0x2f
main.func·002(0xb3200, 0x4)
        OSSPATH/funwithgo/fanin/main.go:76 +0xd1
created by main.quoteSource
        OSSPATH/funwithgo/fanin/main.go:78 +0x93

goroutine 6 [chan receive]:
main.func·001(0x21023d000)
        OSSPATH/funwithgo/fanin/main.go:20 +0x37
created by main.FanIn
        OSSPATH/funwithgo/fanin/main.go:20 +0x1ab

goroutine 7 [chan receive]:
main.func·001(0x21023d060)
        OSSPATH/funwithgo/fanin/main.go:20 +0x37
created by main.FanIn
        OSSPATH/funwithgo/fanin/main.go:20 +0x1ab

goroutine 8 [chan receive]:
main.func·001(0x21023d0c0)
        OSSPATH/funwithgo/fanin/main.go:20 +0x37
created by main.FanIn
        OSSPATH/funwithgo/fanin/main.go:20 +0x1ab

goroutine 9 [chan receive]:
main.printQuotes(0x21023d120)
        OSSPATH/funwithgo/fanin/main.go:62 +0x4e
created by main.main
        OSSPATH/funwithgo/fanin/main.go:53 +0x123
```


