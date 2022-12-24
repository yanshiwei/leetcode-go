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
    void dfs(TreeNode*root,map<int,vector<pair<int,int>>>&mp,int row, int col){
        if(root==nullptr){
            return;
        }
        mp[col].emplace_back(make_pair(row, root->val));
        dfs(root->left,mp,row+1,col-1);
        dfs(root->right,mp,row+1,col+1);
    }
    vector<vector<int>> verticalTraversal(TreeNode* root) {
        map<int,vector<pair<int,int>>>mp;//col:(row,value)
        dfs(root,mp,0,0);
        vector<vector<int>>res;
        for(auto&tt:mp){
            auto listForCol=tt.second;
            // vector 指定排序，先按照行 然后按照值
            sort(listForCol.begin(),listForCol.end(),[](pair<int,int>&a,pair<int,int>&b){
                if(a.first!=b.first){
                    return a.first<b.first;
                }
                return a.second<b.second;
            });
            vector<int>tmp;
            for(auto&ll:listForCol){
                tmp.emplace_back(ll.second);
            }
            res.emplace_back(tmp);
        }
        return res;
    }
};
