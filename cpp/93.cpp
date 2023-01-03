class Solution {
    /*
    有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
    */
public:
    //IP总共四段，每段最长3个，最短1个，且当长度非1时候起始字母不能是0;且在0-255之间
    bool isValid(string seg){
        if(seg.empty()||seg.size()>3||(seg.size()>1&&seg[0]=='0')){
            return false;
        }
        int res=atoi(seg.c_str());
        return res<=255&&res>=0;
    }
    void dfs(string s,string tmp,vector<string>&res,int cnt){
        if(cnt==0){
            //cnt代表剩余处理的次数，次数为0时候查看字符串s是否处理完钱
            if(s.empty()){
                res.push_back(tmp);
            }
            return;
        }
        // 尝试字段长度1/2/3
        for(int segLen=1;segLen<=3;segLen++){
            if(s.size()>=segLen&&isValid(s.substr(0,segLen))){
                //该字段可以处理
                if(cnt==1){
                    dfs(s.substr(segLen),tmp+s.substr(0,segLen),res,cnt-1);//substr(i)表示从下标为i开始一直到结尾
                }else{
                    dfs(s.substr(segLen),tmp+s.substr(0,segLen)+".",res,cnt-1);
                }
            }
        }
    }
    vector<string> restoreIpAddresses(string s) {
        vector<string>res;
        string tmp="";
        dfs(s,tmp,res,4);
        return res;
    }
};
