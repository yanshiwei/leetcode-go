/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> neighbors;
    Node() {
        val = 0;
        neighbors = vector<Node*>();
    }
    Node(int _val) {
        val = _val;
        neighbors = vector<Node*>();
    }
    Node(int _val, vector<Node*> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};
*/

class Solution {
public:
    unordered_map<Node*,Node*>hash;//源节点到克隆节点之间的映射关系
    Node* cloneGraph(Node* node) {
        if(node==nullptr){
            return nullptr;
        }
        return dfs(node);
    }
    Node* dfs(Node* node){
        if(hash[node]){
            // node节点已经被访问过了,直接从哈希表hash中取出对应的克隆节点返回
            return hash[node];
        }
        Node* clone=new Node(node->val);//克隆节点
        hash[node]=clone;//建立源节点到克隆节点的映射
        //克隆边
        for(Node* n:node->neighbors){
            clone->neighbors.push_back(dfs(n));
        }
        return clone;
    }
};
