class Solution {
    /*
    字符串模拟：
    1. 如果当前行只有一个单词，特殊处理为左对齐
    2. 如果当前行为最后一行，特殊处理为左对齐
    3. 其余为一般情况，分别计算：
        当前行单词总长度
        当前行空格总长度
        往下取整后的单位空格长度
       然后依次进行拼接。当空格无法均分时，每次靠左的间隙多加一个空格，
       直到剩余的空格能够被后面的间隙所均分.
    */
public:
    vector<string> fullJustify(vector<string>& words, int maxWidth) {
        vector<string> res;
        int n=words.size();
        vector<string>curRow;// 装载当前行的所有 word
        int i=0;
        while(i<n){
            curRow.clear();
            curRow.push_back(words[i]);
            int curSize=words[i].size();
            i++;
            while(i<n&&curSize+1+words[i].size()<=maxWidth){
                // 尝试是否能够放下下一个单词，单词之间空格隔开，故还需要+1
                curSize+=1+words[i].size();
                curRow.push_back(words[i]);
                i++;
            }
            int cnt=curRow.size();//注意curSize包括间隙的长度，curRow只有单词
            // 当前行为最后一行，特殊处理为左对齐
            if(i==n){
                // 首先将curRows每个单词间添加间隙
                string str=curRow[0];
                for(int k=1;k<curRow.size();k++){
                    str+=" ";
                    str+=curRow[k];
                }
                while(str.size()!=maxWidth){
                    str+=" ";
                }
                res.push_back(str);
                break;
            }
             // 如果当前行只有一个 word，特殊处理为左对齐
            if(cnt==1){
                string str=curRow[0];
                while(str.size()!=maxWidth){
                    str+=" ";
                }
                res.push_back(str);
                continue;
            }
            /**
            * 其余为一般情况
            * wordWidth : 当前行单词总长度;
            * spaceWidth : 当前行空格总长度;
            * spaceItemWidth : 往下取整后的单位空格长度
            */
            // cnt是当前行单词个数，cnt-1是当前行间隙个数
            // curSize是当前行包括间隙的长度
            int wordWidth=curSize-(cnt-1);//当前行单词总长度
            int spaceWidth=maxWidth-wordWidth;//当前行需要的总的空格长度
            int spaceItemWidth=spaceWidth/(cnt-1);// 每个间隙平均下来需要的空格长度
            string spaceItemStr="";
            for(int k=0;k<spaceItemWidth;k++){
                spaceItemStr+=" ";
            }
            string str="";
            int sum=0;// 当前空格长度
            for(int k=0;k<cnt;k++){
                string item=curRow[k];
                str+=item;
                if(k==cnt-1){
                    // 最后一个单词，后面无需间隙
                    break;
                }
                str+=spaceItemStr;
                sum+=spaceItemWidth;
                // 剩余的间隙数量（可填入空格的次数）
                int remain=cnt-1-k-1;
                // 剩余间隙数量 * 最小单位空格长度 + 当前空格长度 < 所需空格总长度，则在当前间隙多补充一个空格
                // 如果空格分配间隙不均匀，则从左侧开始多分配
                if(remain*spaceItemWidth+sum<spaceWidth){
                    str+=" ";
                    sum++;
                }
            }
            res.push_back(str);
        }
        return res;
    }
};
