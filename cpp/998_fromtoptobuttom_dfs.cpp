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
    // dfs自顶向下，https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/815690/yi-pian-wen-zhang-jie-jue-suo-you-er-cha-kcb0/
public:
    string smallestFromLeaf(TreeNode* root) {
        dfs(root,"");
        sort(res.begin(),res.end());
        return res[0];
    }
private:
    void dfs(TreeNode* root,string s){
        if(root==nullptr){
            return;
        }
        s+='a'+(root->val);//转为字母
        if(root->left==nullptr&&root->right==nullptr){
            // 题目要求从叶子到根，故逆序
            reverse(s.begin(),s.end());
            res.push_back(s);
            return;
        }
        dfs(root->left,s);
        dfs(root->right,s);
    }
    vector<string> res;
};
