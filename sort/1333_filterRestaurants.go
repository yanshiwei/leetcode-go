func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
    var res []int
    for i:=range restaurants{
        oneInfo:=restaurants[i]
        //id:=oneInfo[0]
        //ranting:=oneInfo[1]
        isVegan:=oneInfo[2]
        price:=oneInfo[3]
        distance:=oneInfo[4]
        if veganFriendly==1{
            //要求素食
            if isVegan==1{
                if price<=maxPrice&&distance<=maxDistance{
                    res=append(res,i)
                }
            }
        }else{
            if price<=maxPrice&&distance<=maxDistance{
                res=append(res,i)
            }
        }
    }
    //fmt.Println(res)
    sort.Slice(res,func(i,j int)bool{
        //按照 rating 从高到低排序。如果 rating 相同，那么按 id 从高到低排序
        if restaurants[res[i]][1]!=restaurants[res[j]][1]{
            return restaurants[res[i]][1]>restaurants[res[j]][1]
        }
        return restaurants[res[i]][0]>restaurants[res[j]][0]
    })
    //fmt.Println("sort",res)
    for i:=range res{
        idx:=res[i]
        res[i]=restaurants[idx][0]
    }
    return res
}
