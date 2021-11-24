func originalDigits(s string) string {
    var fre=make(map[byte]int)
    for i:=range s{
        fre[s[i]]++
    }
    var cnt=make([]int,10)
    cnt[0]=fre['z']// only zero has z
    cnt[2]=fre['w']// only two has w
    cnt[4]=fre['u']// only four has u
    cnt[6]=fre['x']// only six has x
    cnt[8]=fre['g']// only eight has g
    // for h, only 3 and 8 include
    cnt[3]=fre['h']-cnt[8]
    // for f, only 4 and 5 include
    cnt[5]=fre['f']-cnt[4]
    // for s, only 6 and 7 include 
    cnt[7]=fre['s']-cnt[6]
    // for o. only 0,1,2,4 include
    cnt[1]=fre['o']-cnt[0]-cnt[2]-cnt[4]
    // for i, only 5,6,8,9 include
    //最后的 9 就可以通过 n, i, e 中的任一字符计算得到了。这里推荐使用i 进行计算，因为n 在 9中出现了 2 次，e 在 3 中出现了 2 次，容易在计算中遗漏。
    cnt[9]=fre['i']-cnt[5]-cnt[6]-cnt[8]
    var res string
    for i:=0;i<10;i++{
        // from 0 to 9 one by one
        for j:=0;j<cnt[i];j++{
            // 几个cnt[i]
            res+=string(i+'0')
        }
    }
    return res
}
