class Solution {
    /*
    给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。
    */
public:
    vector<int> findSubstring(string s, vector<string>& words) {
        if(s.size()<1||words.size()<1){
            return {};
        }
        unordered_map<string,int>hash;
        for(int i=0;i<words.size();i++){
            hash[words[i]]++;
        }
        vector<int>res;
        // windows size
        int word_len=words[0].size();
        int len=word_len*words.size();
        // start match,因为我们已经是从一个窗口开始匹配，那么他循环的次数等于当前
        //主串的总长-窗口长+1次
        for(int i=0;i<s.size()-len+1;i++){
            // substr of windows
            string s1=s.substr(i,len);
            //新建一个map用来统计当前s1中串中每个单词大小的串的出现次数，如果出现的次数等于前面统计的
            //wordMap中统计的次数，那门他们绝逼相等
            //或者反过来想，就是我们把前面拿到的hashmap放入到这个临时的map中，此时tmpmap中记录了
            //words中每个单词出现的次数，我们只需要遍历当前串，将遍历到的单词在临时map中将其次数减少即可，最后，
            //我们判断临时map是否为空，如果为空，则表示完全匹配
            unordered_map<string,int> tmp(hash);
            //遍历这个取出的子串,这里我们规定遍历每次走一个单词长度
            for(int j=0;j<s1.size();j+=word_len){
                string word = s1.substr(j,word_len);
                if(tmp.count(word)>0){
                    // 存在则次数减1 当我们map减少至0时，这里要进行移除，
                    tmp[word]--;
                    if(tmp[word]==0){
                        tmp.erase(word);
                    }
                }else{
                    // 不存在 肯定不匹配，直接跳出循环即可
                    break;
                }
            }
            // 该窗口遍历结束
            if(tmp.empty()){
                res.push_back(i);
            }
        }
        return res;
    }
};
