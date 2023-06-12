class Solution {
    /*
    给定两个字符串 s 和 t ，判断它们是否是同构的。
如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
    */
public:
    bool isIsomorphic(string s, string t) {
        // 需2个方向都判断，如果只判断s->t
        // 则出现s=ab,t=cc情况此时明显不匹配
        return isIsomorphicCore(s,t)&&isIsomorphicCore(t,s);
    }
private:
    bool isIsomorphicCore(string s,string t){
        if(s.size()<1||t.size()<1||s.size()!=t.size()){
            return false;
        }
        unordered_map<char,char>table;
        for(int i=0;i<s.size();i++){
            char c1=s[i];
            char c2=t[i];
            if(table.count(c1)==0){
                // no key
                table[c1]=c2;
            }else{
                if(table[c1]!=c2){
                    return false;
                }
            }
        }
        return true;
    }
};
