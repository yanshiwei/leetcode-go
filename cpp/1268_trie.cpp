class Solution {
    /*
    给你一个产品数组 products 和一个字符串 searchWord ，products  数组中每个产品都是一个字符串。
请你设计一个推荐系统，在依次输入单词 searchWord 的每一个字母后，推荐 products 数组中前缀与 searchWord 相同的最多三个产品。如果前缀相同的可推荐产品超过三个，请按字典序返回最小的三个。
请你以二维列表的形式，返回在输入 searchWord 每个字母后相应的推荐产品的列表。
    */
public:
    vector<vector<string>> suggestedProducts(vector<string>& products, string searchWord) {
        sort(products.begin(),products.end(),[](string&a,string&b){
            return a<b;
        });// 按照字典顺序排序
        Tire tt;
        for(const auto&p : products){
            tt.insert(p);
        }
        return tt.startWithList(searchWord);
    }
private:
    class Tire{
        public:
        Tire(){
            isEnd=false;
            memset(children,0,sizeof(children));
        }
        void insert(string word){
            Tire*node=this;
            for(char c : word){
                if(node->children[c-'a']==nullptr){
                    node->children[c-'a']=new Tire();
                }
                node=node->children[c-'a'];
                if(node->preWords.size()<3){
                    // 最多只需要3个
                    node->preWords.push_back(word);
                }
            }
            node->isEnd=true;
        }
        bool search(string word){
            Tire*node=this;
            for(char c: word){
                if(node->children[c-'a']==nullptr){
                    return false;
                }
                node=node->children[c-'a'];
            }
            return node->isEnd;
        }
        bool startWith(string prefix){
            Tire*node=this;
            for(char c: prefix){
                if(node->children[c-'a']==nullptr){
                    return false;
                }
                node=node->children[c-'a'];
            }
            return true;
        }
        // 新增 方便查询命中的结果
        vector<vector<string>> startWithList(string prefix){
            Tire*node=this;
            vector<vector<string>>res(prefix.size());
            for(int i=0;i<prefix.size();i++){
                char c=prefix[i];
                if(node->children[c-'a']==nullptr){
                    break;
                }
                node=node->children[c-'a'];
                for(const auto&word:node->preWords){
                    res[i].push_back(word);
                }
            }
            return res;
        }
        private:
            bool isEnd;
            Tire*children[26];
            vector<string>preWords;//前缀包含此节点的单词
    };
};
