class Solution {
    /*
    你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：
写出一个秘密数字，并请朋友猜这个数字是多少。朋友每猜测一次，你就会给他一个包含下述信息的提示：
猜测数字中有多少位属于数字和确切位置都猜对了（称为 "Bulls"，公牛），
有多少位属于数字猜对了但是位置不对（称为 "Cows"，奶牛）。也就是说，这次猜测中有多少位非公牛数字可以通过重新排列转换成公牛数字。
给你一个秘密数字 secret 和朋友猜测的数字 guess ，请你返回对朋友这次猜测的提示。
提示的格式为 "xAyB" ，x 是公牛个数， y 是奶牛个数，A 表示公牛，B 表示奶牛。
输入：secret = "1123", guess = "0111"
输出："1A1B"
    */
public:
    string getHint(string secret, string guess) {
        if(secret.size()<1||guess.size()<1||secret.size()!=guess.size()){
            return "";
        }
        int bullCnt=0;
        // secret 和 guess 仅由数字0-9组成
        vector<int> secrectCnt(10);
        vector<int> guessCnt(10);
        for(int i=0;i<secret.size();i++){
            if(secret[i]==guess[i]){
                bullCnt++;
            }else{
                secrectCnt[secret[i]-'0']++;
                guessCnt[guess[i]-'0']++;
            }
        }
        // 基于两个Cnt数组统计secret[i]！=guess[i]时各自数字出现次数，由于多余的数字无法匹配，对于0-9的每位数字，应取其在 secret 和guess 中的出现次数的最小值。将每位数字出现次数的最小值累加，即为奶牛的个数
        int cowCnt=0;
        for(int i=0;i<10;i++){
            cowCnt+=min(secrectCnt[i],guessCnt[i]);
        }
        return to_string(bullCnt)+"A"+to_string(cowCnt)+"B";
    }
};
