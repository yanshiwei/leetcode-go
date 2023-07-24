class Solution {
    /*
    给定一个包含大写字母和小写字母的字符串 s ，返回 通过这些字母构造成的 最长的回文串 
    */
public:
    int longestPalindrome(string s) {
        // ，回文字符串只有中间的那个字母可以是奇数次，其余的必须是偶数次。所以多的要删掉一个，使之成为偶数个。 这样讲可能有点含糊，举个例子： 如"abbccdddeee": 每个字母出现次数为：{ a: 1, b: 2, c: 2, d: 3, e: 3 } 出现奇数次的为：a,d,e，共3种字母。选一个在中间，另外两种字母都要各舍弃一个，成为偶数个。所以需要删掉的个数为：3-1=2。
        unordered_map<char,int>hash;
        for(auto &c:s){
            hash[c]++;
        }
        int numOfOdd=0;// 出现奇数次的字母有多少个
        for(auto it=hash.begin();it!=hash.end();it++){
            if(it->second%2!=0){
                numOfOdd++;
            }
        }
        // 如果出现奇数次的字母大于1，只有其中一个字母可以安排一次在中间，其余的都只能出现偶数次，所以删掉多余的。
        if(numOfOdd>1){
            numOfOdd--;
        }else{
            numOfOdd=0;
        }
        return s.size()-numOfOdd;
    }
};
