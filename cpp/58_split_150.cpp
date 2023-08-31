class Solution {
    // 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
    // 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
public:
    vector<string>split(string&str,string det){
        if(str.size()<1){
            return {};
        }
        string s=str;
        size_t pos=s.find(det);
        vector<string>res;
        while(pos!=string::npos){
            string tmp=s.substr(0,pos);
            if(tmp.size()>0){
                res.push_back(tmp);
            }
            s.erase(0,pos+det.size());
            pos=s.find(det);
        }
        if(s.size()>0){
            res.push_back(s);
        }
        return res;
    }
    int lengthOfLastWord(string s) {
        if(s.size()<1){
            return 0;
        }
        vector<string>res=split(s," ");
        return res[res.size()-1].size();
    }
};
