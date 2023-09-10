class Solution {
    /*
    s =
"badc"
    t =
"baba"
    */
public:
    bool isIsomorphicCore(string s, string t){
        int lens=s.size();
        int lent=t.size();
        unordered_map<char,char>mmap;
        for(int i=0;i<lens;i++){
            char c=s[i];
            if(mmap.count(c)==0){
                // first
                mmap[c]=t[i];
            }else{
                // maped before
                if(mmap[c]!=t[i]){
                    return false;
                }
            }
        }
        return true;
    }
    bool isIsomorphic(string s, string t) {
        int lens=s.size();
        int lent=t.size();
        if(lens!=lent){
            return false;
        }
        return isIsomorphicCore(s,t)&&isIsomorphicCore(t,s);
    }
};
