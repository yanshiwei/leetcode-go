class Solution {
    /*
    先按照奇数位插入，保证最大程度字符分开，因为奇数出现个数>=偶数（当数组长是奇数时候，奇数出现个数>偶数；当数组长是偶数时候，奇数出现个数==偶数）
    */
    // sort 比较方法1, 仿函数类
    struct greater{
        bool operator()(pair<int,int>m,pair<int,int>n){
            return m.second>n.second;
        }
    };
    // sort 比较方法2, 函数，注意要静态
    bool static cmp(const pair<int, int>& a, const pair<int, int>& b) {
        return a.second > b.second; // 按照频率从大到小排序
    }
public:
    string reorganizeString(string s) {
        unordered_map<char,int>fre;
        int maxFre=0;
        // 1 统计频率，并求出最大次数
        for(char c: s){
            fre[c]++;
            maxFre=max(maxFre,fre[c]);
        }
        // 2 若单个字母出现数量大于一半则肯定会有相邻
        if(maxFre>(s.size()+1)/2){
            return "";
        }
        // 3 转为vector（即数组）按频率从大到小排序
        vector<pair<int,int>>arr(fre.begin(),fre.end());
        // sort(arr.begin(),arr.end(),greater()); // 比较方法1
        // sort(arr.begin(),arr.end(),cmp);// 比较方法2
        sort(arr.begin(),arr.end(),[](const pair<int, int>& a, const pair<int, int>& b){
            return a.second>b.second;
        }); // 比较方法3
        // 4 按奇数位顺序插入，插满之后按偶数位顺序插入
        int idx=0;// 先按奇数位散开
        string res(s);
        for(int i=0;i<arr.size();i++){
            while(arr[i].second){
                res[idx]=arr[i].first;
                idx+=2;
                if(idx>=s.size()){
                    idx=1;// 奇数位插满了插偶数位
                }
                arr[i].second--;
            }
        }
        return res;
    }
};
