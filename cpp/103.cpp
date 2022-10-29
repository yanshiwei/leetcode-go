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
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        if(root==nullptr){
            return {};
        }
        vector<vector<int>> res;
        queue<TreeNode*>qu;//push pop front back
        qu.push(root);
        int i=0;
        while(qu.empty()==false){
            int curSize=qu.size();
            vector<int>tmp;
            for(int i=0;i<curSize;i++){
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
            if((i&1)==1){
                reverse(tmp.begin(),tmp.end());
            }
            res.push_back(tmp);
            i++;
        }
        return res;
    }
};
