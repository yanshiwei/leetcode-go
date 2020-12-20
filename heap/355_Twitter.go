type Info struct {
    UserId int
    TweetId int
    Time int
}
type MSG struct {
    TweetId int
    Time int
}
type Twitter struct {
    AllTwitters []Info
    UserFollowOthers map[int][]int
    GolbalTime int
}

const LL=10

/** Initialize your data structure here. */
func Constructor() Twitter {
    tw:=Twitter{AllTwitters:make([]Info,0),UserFollowOthers:make(map[int][]int)}
    return tw
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.GolbalTime++
    info:=Info{UserId:userId,TweetId:tweetId,Time:this.GolbalTime}
    this.AllTwitters=append(this.AllTwitters,info)
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
    //var heap []int//min heap
    var raw []MSG
    var res []int
    for i:=range this.AllTwitters {
        if this.AllTwitters[i].UserId==userId {
            msg:=MSG{TweetId:this.AllTwitters[i].TweetId,Time:this.AllTwitters[i].Time}
            raw=append(raw,msg)
        }
    }
    if fList,ok:=this.UserFollowOthers[userId];ok==true {
        for _,otherUser:=range fList{
            for i:=range this.AllTwitters {
                if this.AllTwitters[i].UserId==otherUser{
                    msg:=MSG{TweetId:this.AllTwitters[i].TweetId,Time:this.AllTwitters[i].Time}
                    raw=append(raw,msg)
                }
            }
        }
    }
    fmt.Println("raw",raw)
    heapSort(raw)
    fmt.Println("sort",raw)
    for i:=range raw {
        msg:=raw[i]
        if i<LL {
            res=append(res,msg.TweetId)
        }
    }
    return res
}
/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
    if followeeId==followerId {
        //订阅自己
        return
    }
    if _,ok:=this.UserFollowOthers[followerId];ok==false {
        var tweeterList []int
        tweeterList=append(tweeterList,followeeId)
        this.UserFollowOthers[followerId]=tweeterList
    }else {
        for i:=range this.UserFollowOthers[followerId] {
            if this.UserFollowOthers[followerId][i]==followeeId {
                return //已经订阅的直接返回 防止重复
            }
        }
        this.UserFollowOthers[followerId]=append(this.UserFollowOthers[followerId],followeeId)
    }
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if fList,ok:=this.UserFollowOthers[followerId];ok==true {
        for i:=range fList {
            if followeeId==fList[i] {
                this.UserFollowOthers[followerId]=append(this.UserFollowOthers[followerId][0:i],this.UserFollowOthers[followerId][i+1:]...)//delete i
            }
        }
    }
}

func adjustHeap(heap []MSG, start int, length int) {
	var i=start
	var child int
	for {
		child=2*i+1
		if child>length-1 {
			break
		}
		if child+1<=length-1&&heap[child].Time>heap[child+1].Time {
			child+=1
		}
		if heap[i].Time>heap[child].Time {
			heap[i],heap[child]=heap[child],heap[i]
			i=child//child update
		}else {
			break
		}
	}
}
func buildHeap(heap []MSG) {
	//half before
	for i:=len(heap)/2-1;i>=0;i-- {
		adjustHeap(heap,i,len(heap))
	}
}
func heapSort(heap []MSG) {
	buildHeap(heap)
	for i:=len(heap)-1;i>=0;i-- {
		heap[0],heap[i]=heap[i],heap[0]
		adjustHeap(heap,0,i)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
