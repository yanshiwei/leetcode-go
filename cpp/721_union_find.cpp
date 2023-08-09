class Solution {
class UnionFind{
    public:
    // 初始化，n个结点，每个单独就是一个集合
    UnionFind(int n){
        father.resize(n);
        for(int i=0;i<n;i++){
            // https://blog.csdn.net/qq_57469718/article/details/125416286
            // 按照定义，根结点保存自身引用
            father[i]=i;
        }
    }
    // 合并
    void merge(int a,int b){
        // 合并a,b所在的集合
        // 1 先找到集合所在的代表元素，也就是根结点
        int fa=find(a);
        int fb=find(b);
        if(fa==fb){
            // self
            return;
        }
        father[fa]=fb;//a所在集合合入b所在集合
    }
    // 查询
    int find(int a){
        int origin=a;
        while(father[a]!=a){
            a=father[a];
        }
        father[origin]=a;//路径压缩,origin的代表结点也是a
        return a;
    }
    private:
        vector<int>father;
};
public:
// https://leetcode.cn/problems/accounts-merge/solutions/564451/tu-jie-yi-ran-shi-bing-cha-ji-by-yexiso-5ncf/
    vector<vector<string>> accountsMerge(vector<vector<string>>& accounts) {
        vector<vector<string>>res;
        unordered_map<string,int>mail2userid;
        int n=accounts.size();
        UnionFind uf(n);//每个账户信息先单独成集合
        for(int i=0;i<n;i++){
            int mailCnt=accounts[i].size();
            for(int j=1;j<mailCnt;j++){
                string mail=accounts[i][j];
                if(mail2userid.find(mail)==mail2userid.end()){
                    // first
                    mail2userid[mail]=i;
                }else{
                    // 如果账户A、B下出现了相同的邮箱email，那么将账户A和账户B两个连通分量进行合并
                    uf.merge(i,mail2userid[mail]);//合并
                }
            }
        }
        // 遍历并查集中每个连通分量，将所有连通分量内部账户的邮箱全部合并(相同的去重，不同的合并)
        unordered_map<int,vector<string>>id2mails;
        for(auto &kv:mail2userid){
            int uid=kv.second;
            string mail=kv.first;
            int father_uid=uf.find(uid);
            id2mails[father_uid].emplace_back(mail);
        }
        for(auto &kv:id2mails){
            int uid=kv.first;
            vector<string>mail_list=kv.second;
            // 按顺序输出
            sort(mail_list.begin(),mail_list.end());
            vector<string> tmp(1,accounts[uid][0]);//get user_name
            tmp.insert(tmp.end(),mail_list.begin(),mail_list.end());
            res.emplace_back(tmp);
        }
        return res;
    }
};
