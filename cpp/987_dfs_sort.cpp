/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
    /*
    给你二叉树的根结点 root ，请你设计算法计算二叉树的 垂序遍历 序列。
对位于 (row, col) 的每个结点而言，其左右子结点分别位于 (row + 1, col - 1) 和 (row + 1, col + 1) 。树的根结点位于 (0, 0) 。
二叉树的 垂序遍历 从最左边的列开始直到最右边的列结束，按列索引每一列上的所有结点，形成一个按出现位置从上到下排序的有序列表。如果同行同列上有多个结点，则按结点的值从小到大进行排序。
返回二叉树的 垂序遍历 序列。
    */
public:
    struct Node{
        int col;
        int row;
        int val;
        Node():val(0),row(0),col(0){}
        Node(int v,int r,int c):val(v),row(r),col(c){}
    };
    struct cmp1{
	bool operator ()(const Node& a , const Node& b)
	{
        // 按照列排序，如果列相同，按照行排序
        if(a.col!=b.col){
            return a.col<b.col;
        }
        if(a.row!=b.row){
            return a.row<b.row;
        }
        // 如果同行同列上有多个结点，则按结点的值从小到大进行排序
        return a.val<b.val;
	}
};
    struct Node1{
        int col;
        int row;
        int val;
        Node1():val(0),row(0),col(0){}
        Node1(int v,int r,int c):val(v),row(r),col(c){}
    };
    static bool cmp(Node a, Node b){
        // 按照列排序，如果列相同，按照行排序
        if(a.col!=b.col){
            return a.col<b.col;
        }
        if(a.row!=b.row){
            return a.row<b.row;
        }
        // 如果同行同列上有多个结点，则按结点的值从小到大进行排序
        return a.val<b.val;
    }
    void dfs(TreeNode*root,vector<Node>&hash,int row,int col){
        if(root==nullptr){
            return;
        }
        Node node;
        node.val=root->val;
        node.row=row;
        node.col=col;
        hash.push_back(node);
        dfs(root->left,hash,row+1,col-1);
        dfs(root->right,hash,row+1,col+1);
    }
    vector<vector<int>> verticalTraversal(TreeNode* root) {
        if(root==nullptr){
            return {};
        }
        vector<Node>hash;//key is node, vaue is col,row,val
        dfs(root,hash,0,0);//递归便利BST，获得各个结点的坐标值
        //sort(hash.begin(),hash.end(),cmp);
        //sort(hash.begin(),hash.end(),cmp1());
        sort(hash.begin(),hash.end(),[](const Node& a , const Node& b){
            // 按照列排序，如果列相同，按照行排序
            if(a.col!=b.col){
                return a.col<b.col;
            }
            if(a.row!=b.row){
                return a.row<b.row;
            }
            // 如果同行同列上有多个结点，则按结点的值从小到大进行排序
            return a.val<b.val;
        });
        vector<vector<int>>res;
        // 相同列归为一组
        for(int i=0;i<hash.size();){
            vector<int>tmp;
            int j=i;
            // 相同列的在一起
            while(j<hash.size()){
                if(hash[i].col==hash[j].col){
                    tmp.push_back(hash[j].val);
                    j++;
                }else{
                    break;
                }
            }
            res.push_back(tmp);
            i=j;
        }
        return res;
    }
};
