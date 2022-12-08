class Solution {
    /*
    给你一个字符串 path ，表示指向某一文件或目录的 Unix 风格 绝对路径 （以 '/' 开头），请你将其转化为更加简洁的规范路径。
在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。任意多个连续的斜杠（即，'//'）都被视为单个斜杠 '/' 。 对于此问题，任何其他格式的点（例如，'...'）均被视为文件/目录名称。
请注意，返回的 规范路径 必须遵循下述格式：
始终以斜杠 '/' 开头。
两个目录名之间必须只有一个斜杠 '/' 。
最后一个目录名（如果存在）不能 以 '/' 结尾。
此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含 '.' 或 '..'）。
返回简化后得到的 规范路径 。
    */
public:
    void splitStr(string path,char split,vector<string>&res){
        if(path.size()<1){
            return;
        }
        string strs=path+split;//方便统一处理
        size_t pos=strs.find(split);
        while(pos!=string::npos){
            string tmp=strs.substr(0,pos);
            res.push_back(tmp);
            strs=strs.substr(pos+1);
            pos=strs.find(split);
        }
        return;
    }
    string joinStr(vector<string>&strs,char split){
        if(strs.size()<1){
            return "";
        }
        string res="";
        int i=0;
        for(;i<strs.size()-1;i++){
            res+=strs[i];
            res+=split;
        }
        res+=strs[i];
        return res;
    }
    string simplifyPath(string path) {
        if(path.size()<1){
            return "";
        }
        string res="";
        vector<string>arr;
        splitStr(path,'/',arr);
        stack<string>st;
        for(int i=0;i<arr.size();i++){
            if(arr[i]==""){
                continue;
            }else if(arr[i]==".."){
                if(st.empty()==false){
                    st.pop();
                }
            }else if(arr[i]=="."){
                continue;
            }else{
                st.push(arr[i]);
            }
        }
        if(st.empty()){
            return "/";
        }
        vector<string>arrNew(st.size());
        int i=st.size()-1;
        while(st.empty()==false){
            auto cur=st.top();
            st.pop();
            arrNew[i]=cur;
            i--;
        }
        auto jstr=joinStr(arrNew,'/');
        
        return "/"+jstr;
    }
};
