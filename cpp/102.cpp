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
        queue<TreeNode*>qu;//push pop front back
        qu.push(root);
        vector<vector<int>> res;
        while(qu.empty()==false){
            vector<int>tmp;
            int curSuze=qu.size();
            for(int i=0;i<curSuze;i++){
                auto cur=qu.front();
                qu.pop();
                tmp.push_back(cur->val);
                if(cur->left){
                    qu.push(cur->left);
                }
                if(cur->right){
                    qu.push(cur->right);
                }
            }
            res.push_back(tmp);
        }
        return res;
    }
};
