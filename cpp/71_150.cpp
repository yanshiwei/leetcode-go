class Solution {
public:
    vector<string>split(string&str,const string&det){
        vector<string>res;
        if(str.size()<1){
            return res;
        }
        string strs=str+det;//方便统一处理
        size_t pos=strs.find(det);
        while(pos!=string::npos){
            string tmp=strs.substr(0,pos);
            res.push_back(tmp);
            strs=strs.substr(pos+1);
            pos=strs.find(det);
        }
        return res;
    }
    string join(vector<string>&strs,char op){
        if(strs.size()<1){
            return "";
        }
        string res="";
        int i=0;
        while(i<strs.size()-1){
            res+=strs[i];
            res+=op;
            i++;
        }
        res+=strs[i];//最后一位后面不加op
        return res;
    }
    string simplifyPath(string path) {
        if(path.size()<1){
            return "";
        }
        string res="";
        vector<string>arr=split(path,"/");
        stack<string>st;
        for(int i=0;i<arr.size();i++){
            if(arr[i]==""){
                continue;
            }
            if(arr[i]==".."){
                if(st.empty()==false){
                    st.pop();
                }
            }else if(arr[i]=="."){
                continue;
            }else{
                // 普通字符
                st.push(arr[i]);
            }
        }
        if(st.empty()){
            return "/";
        }
        vector<string>arrNew(st.size());
        int i=st.size()-1;
        while(!st.empty()){
            string cur=st.top();
            st.pop();
            arrNew[i]=cur;
            i--;
        }
        string strNew=join(arrNew,'/');
        return "/"+strNew;
    }
};
