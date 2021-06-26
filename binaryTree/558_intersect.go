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

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
    //quadTree1是叶子
    if quadTree1.IsLeaf{
        if quadTree1.Val{
            //全是1
            return quadTree1
        }else{
            //全是0
            return quadTree2
        }
    }
    //quadTree2是叶子
    if quadTree2.IsLeaf{
        if quadTree2.Val{
            //全是1
            return quadTree2
        }else{
            //全是0
            return quadTree1
        }
    }
    //quadTree1,quadTree2都不是叶子,递归比较子树
    quadTree1.TopLeft=intersect(quadTree1.TopLeft,quadTree2.TopLeft)
    quadTree1.TopRight=intersect(quadTree1.TopRight,quadTree2.TopRight)
    quadTree1.BottomLeft=intersect(quadTree1.BottomLeft,quadTree2.BottomLeft)
    quadTree1.BottomRight=intersect(quadTree1.BottomRight,quadTree2.BottomRight)
    //quadTree1子节点比较是否相同,是则结合
    if quadTree1.TopLeft.IsLeaf&&quadTree1.TopRight.IsLeaf&&quadTree1.BottomLeft.IsLeaf&&quadTree1.BottomRight.IsLeaf&&quadTree1.TopLeft.Val==quadTree1.TopRight.Val&&quadTree1.TopRight.Val==quadTree1.BottomLeft.Val&&quadTree1.BottomLeft.Val==quadTree1.BottomRight.Val{
        quadTree1.IsLeaf=true
        quadTree1.Val=quadTree1.TopLeft.Val
        quadTree1.TopLeft=nil
        quadTree1.TopRight=nil
        quadTree1.BottomLeft=nil
        quadTree1.BottomRight=nil
    }
    return quadTree1
}
