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
        int level=0;
        while(!qu.empty()){
            vector<int>tmp;
            int cur_size=qu.size();
            for(int i=0;i<cur_size;i++){
                TreeNode*node=qu.front();
                qu.pop();
                tmp.push_back(node->val);
                if(node->left){
                    qu.push(node->left);
                }
                if(node->right){
                    qu.push(node->right);
                }
            }
            if(level%2!=0){
                reverse(tmp.begin(),tmp.end());
            }
            level++;
            res.push_back(std::move(tmp));
        }
        return res;
    }
};
