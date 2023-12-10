/*
// Definition for a QuadTree node.
class Node {
public:
    bool val;
    bool isLeaf;
    Node* topLeft;
    Node* topRight;
    Node* bottomLeft;
    Node* bottomRight;
    
    Node() {
        val = false;
        isLeaf = false;
        topLeft = NULL;
        topRight = NULL;
        bottomLeft = NULL;
        bottomRight = NULL;
    }
    
    Node(bool _val, bool _isLeaf) {
        val = _val;
        isLeaf = _isLeaf;
        topLeft = NULL;
        topRight = NULL;
        bottomLeft = NULL;
        bottomRight = NULL;
    }
    
    Node(bool _val, bool _isLeaf, Node* _topLeft, Node* _topRight, Node* _bottomLeft, Node* _bottomRight) {
        val = _val;
        isLeaf = _isLeaf;
        topLeft = _topLeft;
        topRight = _topRight;
        bottomLeft = _bottomLeft;
        bottomRight = _bottomRight;
    }
};
*/

class Solution {
    /*
    题目直接看示例2就明白，
    首先将矩阵划分为topLeft，topRight，bottomLeft，bottomRight四个子网格，
    如果网格内值相同，则已经到达叶节点，isLeaf=1，val就是该值
    如果网格内值不同，则还没有到叶节点，isLead=0，var任意值，递归划分子网格
    */
public:
    Node* helper(vector<vector<int>>& grid,int x,int y,int dx, int dy){
        bool isLeaf=true;
        for(int i=x;i<=dx;i++){
            for(int j=y;j<=dy;j++){
                if(grid[i][j]!=grid[dx][dy]){
                    // 子网格值不同
                    isLeaf=false;
                    break;
                }
            }
            if(isLeaf==false){
                break;
            }
        }
        // 当前网格的值相同（即，全为 0 或者全为 1）
        if(isLeaf){
            Node*node=new Node();
            node->isLeaf=isLeaf;// 将 isLeaf 设为 True
            // val 设为网格相应的值, 并将四个子节点都设为 Null 然后停止
            if(grid[dx][dy]==1){
                node->val=true;
            }else{
                node->val=false;
            }
            return node;
        }
        // 构造当前节点
        Node*node=new Node();
        // 当前网格的值不同，继续切分
        int midx=(x+dx)/2;
        int midy=(y+dy)/2;
        Node* up_left_node=helper(grid,x,y,midx,midy);
        Node* up_right_node=helper(grid,x,midy+1,midx,dy);
        Node* down_left_node=helper(grid,midx+1,y,dx,midy);
        Node* down_right_node=helper(grid,midx+1,midy+1,dx,dy);
        // 补充子节点
        node->isLeaf=isLeaf;
        node->val=false;
        node->topLeft=up_left_node;
        node->topRight=up_right_node;
        node->bottomLeft=down_left_node;
        node->bottomRight=down_right_node;
        return node;
    }
    Node* construct(vector<vector<int>>& grid) {
        int m=grid.size();
        if(m<1){
            return nullptr;
        }
        int n=grid[0].size();
        if(n<1){
            return nullptr;
        }
        return helper(grid,0,0,m-1,n-1);
    }
};
