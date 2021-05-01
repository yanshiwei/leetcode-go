func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
    hasCommon:=func(x,y int)bool{
        //判断用户x y有没有共同语言
        for i:=range languages[x-1]{
            for j:=range languages[y-1]{
                if languages[x-1][i]==languages[y-1][j]{
                    return true
                }
            }
        }
        return false
    }
    //统计哪些用户存在沟通障碍
    var needToLearn=make(map[int]struct{})//key是可比的结构体
    for i:=range friendships{
        ship:=friendships[i]
        x:=ship[0]
        y:=ship[1]
        if hasCommon(x,y)==false{
            //好友之间有沟通障碍，去重
            needToLearn[x]=struct{}{}
            needToLearn[y]=struct{}{}
        }
    }
    //没人需要学习
    if len(needToLearn)<1{
        return 0
    }
    //存在沟通障碍的人群中，每种语言会的人数
    var languageFre=make(map[int]int)
    for k:=range needToLearn{
        for j:=range languages[k-1]{
            languageFre[languages[k-1][j]]++
        }
    }
    // 找出最流行的那门语言,会的人数
    var mostLanguageCnt int
    for _,v:=range languageFre{
        if mostLanguageCnt<v {
            mostLanguageCnt=v
        }
    }
    //没掌握这门“最流行语言”的人数，就是需要教学的最少人数
    return len(needToLearn)-mostLanguageCnt
}
