/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/

class Solution {
public:
    Node* connect(Node* root) {
        if(root==NULL){
            return NULL;
        }
        queue<Node*>qu;
        qu.push(root);
        while(qu.empty()==false){
            int size=qu.size();
            Node*pre=NULL;
            for(int i=0;i<size;i++){
                Node*cur=qu.front();
                qu.pop();
                if(i>0){
                    pre->next=cur;
                }
                pre=cur;
                if(cur->left){
                    qu.push(cur->left);
                }
                if(cur->right){
                    qu.push(cur->right);
                }
            }
        }
        return root;
    }
};
