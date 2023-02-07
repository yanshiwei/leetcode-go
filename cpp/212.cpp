/*
题目：
给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
解法：
trie+dfs
*/
class TrieNode{
public:
    string word="";
    vector<TrieNode*>nodes;//长度26，一一对应26个字母
    TrieNode():nodes(26,0){}
};
class Solution {
public:
    int rows,cols;
    vector<string>res;
    void dfs(vector<vector<char>>&board,TrieNode*root,int x,int y){
        char c=board[x][y];
        //递归边界
        if(c=='.'||root->nodes[c-'a']==0){
            //已经访问过或者没有该字母则退出
            return;
        }
        root=root->nodes[c-'a'];
        if(root->word!=""){
            //以该节点结尾的word非空
            res.push_back(root->word);
            root->word="";//防止一个单词被重复添加
        }
        board[x][y]='.';//标记访问过
        if(x>0){
            dfs(board,root,x-1,y);
        }
        if(y>0){
            dfs(board,root,x,y-1);
        }
        if(x+1<rows){
            dfs(board,root,x+1,y);
        }
        if(y+1<cols){
            dfs(board,root,x,y+1);
        }
        board[x][y]=c;//恢复现场
    }
    vector<string>findWords(vector<vector<char>>&board,vector<string>&words){
        rows=board.size();
        cols=rows?board[0].size():0;
        if(rows==0||cols==0){
            return res;
        }
        //建立字典树的模板
        TrieNode*root=new TrieNode();
        for(string word:words){
            TrieNode*cur=root;
            for(int i=0;i<word.size();i++){
                int idx=word[i]-'a';
                if(cur->nodes[idx]==0){
                    // no data
                    cur->nodes[idx]=new TrieNode();
                }
                cur=cur->nodes[idx];
            }
            cur->word=word;
        }
        //DFS模板
        for(int i=0;i<rows;i++){
            for(int j=0;j<cols;j++){
                dfs(board,root,i,j);
            }
        }
        return res;
    }
};
