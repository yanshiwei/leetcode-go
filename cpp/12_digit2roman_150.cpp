class Solution {
public:
    struct Info {
        int val;
        string str;
    };

    string intToRoman(int num) {
        string res="";
        vector<Info>dict;
        dict.push_back(Info{1000,"M"});
        dict.push_back(Info{900,"CM"});
        dict.push_back(Info{500,"D"});
        dict.push_back(Info{400,"CD"});
        dict.push_back(Info{100,"C"});
        dict.push_back(Info{90,"XC"});
        dict.push_back(Info{50,"L"});
        dict.push_back(Info{40,"XL"});
        dict.push_back(Info{10,"X"});
        dict.push_back(Info{9,"IX"});
        dict.push_back(Info{5,"V"});
        dict.push_back(Info{4,"IV"});
        dict.push_back(Info{1,"I"});
        for(auto &info:dict){
            int v=info.val;
            string s=info.str;
            while(num>=v){
                num-=v;
                res+=s;
            }
            if(num<=0){
                break;
            }
        }
        return res;
    }
};
