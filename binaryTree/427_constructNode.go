/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
    var m=len(grid)
    var n=len(grid[0])
    return dfs(grid,0,0,m-1,n-1)
}
//upper x/y && down x/y, 上，左，下，右
func dfs(grid [][]int,ux,uy,dx,dy int)*Node{
    var isLeaf=true
    for i:=ux;i<=dx;i++{
        for j:=uy;j<=dy;j++{
            if grid[i][j]!=grid[ux][uy]{
                //大小不一致
                isLeaf=false
                break
            }
        }
    }
    if isLeaf==true{
        node:= &Node{IsLeaf:isLeaf}
        if grid[ux][uy]==1{
            node.Val=true
        }else{
            node.Val=false
        }
        return node
    }
    //切分
    var midX=(ux+dx)>>1
    var midY=(uy+dy)>>1

    upLeftNode:=dfs(grid,ux,uy,midX,midY)
    upRightNode:=dfs(grid,ux,midY+1,midX,dy)
    downLeftNode:=dfs(grid,midX+1,uy,dx,midY)
    downRightNode:=dfs(grid,midX+1,midY+1,dx,dy)
    return &Node{
        IsLeaf:isLeaf,
        Val:false,
        TopLeft:upLeftNode,
        TopRight:upRightNode,
        BottomLeft:downLeftNode,
        BottomRight:downRightNode,
    }
}
