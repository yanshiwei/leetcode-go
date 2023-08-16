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
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        if(root==nullptr){
            return{};
        }
        vector<vector<int>> res;
        queue<TreeNode*>qu;
        qu.push(root);
        while(!qu.empty()){
           int cur_size=qu.size();
           vector<int>tmp;
           for(int i=0;i<cur_size;i++){
               TreeNode*cur=qu.front();
               qu.pop();
               tmp.push_back(cur->val);
               if(cur->left){
                   qu.push(cur->left);
               }
               if(cur->right){
                   qu.push(cur->right);
               }
           }
           res.push_back(std::move(tmp));        
        }
        return res;
    }
};
