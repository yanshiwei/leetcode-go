class Solution {
    /*
    你一个字符串 s，「k 倍重复项删除操作」将会从 s 中选择 k 个相邻且相等的字母，并删除它们，使被删去的字符串的左侧和右侧连在一起。你需要对 s 重复进行无限次这样的删除操作，直到无法继续为止。在执行完所有删除操作后，返回最终得到的字符串。
    */
public:
    string removeDuplicates(string s, int k) {
        /*
        暴力求解
        int length=-1;// 第i次删除前字符串长度
        while(length!=s.size()){
            length=s.size();
            int cnt=1;
            for(int i=0;i<s.size();i++){
                if(i==0||s[i]!=s[i-1]){
                    cnt=1;
                }else{
                    cnt++;
                    if(cnt==k){
                        s.erase(i-k+1,k);
                        break; // 删除一个连续k次的字母后重新开始
                    }
                }
            }
        }
        return s;
        */
        /*
        stack规范法
        1.遍历字符串，对于每个字符与顶元素
        1.1 若空栈或者不相等，则直接入栈，次数为1
        1.2 若相等，次数+1
        1.2.1 次数等于k，满足删除条件，退栈
        1.2.2 次数小于k,需继续
        2.遍历栈，根据次数恢复字符串
        */
        stack<pair<char,int>>st;// 存储字符及对应出现的次数
        for(int i=0;i<s.size();i++){
            if(st.empty()==false&&st.top().first==s[i]){
                // 若栈非空且当前元素等于栈顶元素
                ++st.top().second;
                if(st.top().second==k){
                    // 若栈顶元素数目=k
                    st.pop();
                }
            }else{
                // 若栈空或者不等 直接入栈 次数设置1
                st.push(make_pair(s[i],1));
            }
        }
        string res="";
        while(st.empty()==false){
            int cnt=st.top().second;
            while(cnt>0){
                res+=st.top().first;
                cnt--;
            }
            st.pop();
        }
        reverse(res.begin(),res.end());
        return res;
    }
};
