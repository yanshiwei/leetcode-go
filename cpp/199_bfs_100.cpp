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
    // 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
public:
    vector<int> rightSideView(TreeNode* root) {
        if(root==nullptr){
            return {};
        }
        vector<int>res;
        queue<TreeNode*>qu;
        qu.push(root);
        while(!qu.empty()){
            int n=qu.size();
            TreeNode*last=nullptr;
            for(int i=0;i<n;i++){
                TreeNode*cur=qu.front();
                if(i==n-1){
                    last=cur;
                }
                if(cur->left){
                    qu.push(cur->left);
                }
                if(cur->right){
                    qu.push(cur->right);
                }
                qu.pop();
            }
            res.push_back(last->val);
        }
        return res;
    }
};
