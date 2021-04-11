type SubrectangleQueries struct {
    Rect [][]int
}


func Constructor(rectangle [][]int) SubrectangleQueries {
    var s =SubrectangleQueries{
        Rect: rectangle,
    }
    return s
}


func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    for i:=row1;i<=row2;i++{
        for j:=col1;j<=col2;j++{
            this.Rect[i][j]=newValue
        }
    }
}


func (this *SubrectangleQueries) GetValue(row int, col int) int {
    return this.Rect[row][col]
}


/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
