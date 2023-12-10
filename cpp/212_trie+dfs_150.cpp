/*
题目：
给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
解法：
trie+dfs
*/
class TrieNode{
public:
    string word="";//nod used bool end, 因为我们要直接使用这个字符串
    vector<TrieNode*>nodes;//长度26，一一对应26个字母
    TrieNode():nodes(26,0){}
};
class Solution {
public:
    int rows,cols;
    vector<string>res;
    void dfs(vector<vector<char>>&board,TrieNode*root,int x,int y){
        // 边界控制
        if(x<0||x>board.size()-1||y<0||y>board[0].size()-1){
            return;
        }

        char c=board[x][y];
        // 已经访问控制
        if(c=='.'){
            return;
        }
        // 是否是匹配字符
        if(root->nodes[c-'a']==nullptr){
            return;
        }
        // 匹配当前字符，则更新到下一个
        root=root->nodes[c-'a'];
        // 判断是否结束
        if(root->word.size()>0){
            //以该节点结尾的word非空
            res.push_back(root->word);
            root->word="";//防止一个单词被重复添加
        }
        // 标志访问
        board[x][y]='.';
        // 四个方向尝试
        dfs(board,root,x-1,y);
        dfs(board,root,x+1,y);
        dfs(board,root,x,y-1);
        dfs(board,root,x,y+1);
        board[x][y]=c;//恢复现场，方便后面dfs
    }
    vector<string>findWords(vector<vector<char>>&board,vector<string>&words){
        rows=board.size();
        cols=rows?board[0].size():0;
        if(rows==0||cols==0){
            return res;
        }
        //建立字典树的模板
        TrieNode*root= new TrieNode();
        for(string word : words){
            TrieNode*node=root;
            for(char c : word){
                int idx=c-'a';
                if(node->nodes[idx]==nullptr){
                    // no data
                    node->nodes[idx]=new TrieNode();
                }
                node=node->nodes[idx];
            }
            node->word=word;
        }
        //dfs模版
        for(int i=0;i<rows;i++){
            for(int j=0;j<cols;j++){
                dfs(board,root,i,j);
            }
        }
        return res;
    }
};
