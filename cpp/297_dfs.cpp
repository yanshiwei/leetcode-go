/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Codec {
public:

    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        string res="";
        serializeCore(root,res);
        return res;
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        list<string>dataArr;////push_back pop_back;push_front pop_front;erase/insert
        string str="";
        for(auto&ch:data){
            if(ch==','){
                dataArr.push_back(str);
                str.clear();
            }else{
                str.push_back(ch);
            }
        }
        return deserializeCore(dataArr);
    }
private:
    void serializeCore(TreeNode*root,string&res){
        if(root==nullptr){
            res+="None,";
            return;
        }
        res+=to_string(root->val)+",";
        serializeCore(root->left,res);
        serializeCore(root->right,res);
        return;
    }
    TreeNode*deserializeCore(list<string>&dataArr){
        if(dataArr.front()=="None"){
            dataArr.pop_front();
            // null
            return nullptr;
        }
        TreeNode*root=new TreeNode(stoi(dataArr.front()));
        dataArr.pop_front();
        root->left=deserializeCore(dataArr);
        root->right=deserializeCore(dataArr);
        return root;
    }
};

// Your Codec object will be instantiated and called as such:
// Codec ser, deser;
// TreeNode* ans = deser.deserialize(ser.serialize(root));
