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
    long long pre=-9223372036854775808; // 官方给的数据太极端，只能long long min
// 中序遍历时，判断当前节点是否大于中序遍历的前一个节点，如果大于，说明满足 BST，继续遍历；否则直接返回 false。
    bool isValidBST(TreeNode* root) {
        if(root==nullptr){
            return true;
        }
        if(!isValidBST(root->left)){
            return false;
        }
        // 当前节点是否大于中序遍历的前一个节点
        if(root->val<=pre){
            return false;
        }
        pre=root->val;
        return isValidBST(root->right);
    }
};
