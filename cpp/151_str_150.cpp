class Solution {
    /*
    题目：
    给你一个字符串 s ，请你反转字符串中 单词 的顺序。
单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
    注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。也就是去掉前面/后面空格，以及单词之间的多个空格只保留一个空格。
    */
public:
    string reverseWords(string s) {
        if(s.size()<1){
            return s;
        }
        // 1, 移除多余空格，注意除了s开始的第一个单词外，其他单词开头都需要一个空格
        int fast=0;
        int slow=0;
        while(fast<s.size()){
            // 找到第一个不是空格的char
            while(fast<s.size()&&s[fast]==' '){
                fast++;
            }
            // 因为s至少含有一个字母，所以fast不会到末尾,fast指向第一个非空字符
            // 碰到下一个空格前都需要保留
            while(fast<s.size()&&s[fast]!=' '){
                s[slow++]=s[fast++];
            }
            // fast指向空格
            //除了s开始的第一个单词外，其他单词开头都需要一个空格
            if(fast<s.size()&&slow!=0){
                // slow!=0说明不是第一个元素
                s[slow++]=' ';
            }
        }
        // 移除后面空格
        if(s[slow-1]==' '){
            // 类似case "  hello world  "也就是最后单词后面还有空格
            s.resize(slow-1);//去掉空格
        }else {
            // 类似case "a good   example"也就是单词结尾没有空格
            s.resize(slow);
        }
        // 2 整个字符串反转
        reverse(s.begin(),s.end());
        // 3 逐个单词反转
        slow=0;
        fast=0;
        while(fast<s.size()){
            // 碰到下一个空格前都需要保留
            while(fast<s.size()&&s[fast]!=' '){
                fast++;
            }
            reverse(s.begin()+slow,s.begin()+fast);
            slow=fast+1;
            fast=slow;
        }
        return s;
    }
};
