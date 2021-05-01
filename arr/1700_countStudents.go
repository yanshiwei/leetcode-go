func countStudents(students []int, sandwiches []int) int {
    //队列不队列的根本无所谓，因为每个学生最终都能排到。。所以只用统计喜欢每种形状的学生个数即
    var like=make([]int,2)//0代表喜欢0，1代表喜欢1
    for i:=range students{
        like[students[i]]++
    }
    for i:=range sandwiches{
        if like[sandwiches[i]]==0{
            // 当前的形状没有学生喜欢了
            return len(students)-i
        }
        like[sandwiches[i]]--
    }
    return 0
}
