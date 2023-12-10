class WordDictionary {
    /*
    请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
实现词典类 WordDictionary ：
WordDictionary() 初始化词典对象
void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。
    相比较208 构造Trie树，仅查询部分复杂了一些，加上了'.'，可以表示任何一个字母。
    查询部分可以用递归回溯解决，其余部分和实现前缀树完全一样。
    */
public:
    WordDictionary() {
        isEnd=false;
        memset(next,0,sizeof(next));
    }
    
    void addWord(string word) {
        WordDictionary*node=this;
        for(char c: word){
            if(node->next[c-'a']==nullptr){
                node->next[c-'a']=new WordDictionary();
            }
            node=node->next[c-'a'];
        }
        node->isEnd=true;
    }
    
    bool search(string word) {
        return searchWord(this,word);
    }
private:
    bool searchWord(WordDictionary*node, string word){
        if(word.size()==1){
            // only size=1
            if(isalpha(word[0])){
                if(node->next[word[0]-'a']&&node->next[word[0]-'a']->isEnd){
                    return true;
                }
                return false;
            }else{
                // . 故只需要下一位是结束位就行
                for(int i=0;i<26;i++){
                    if(node->next[i]&&node->next[i]->isEnd){
                        return true;
                    }
                }
                return false;
            }
        }else{
            // muitiple char
            // handle first
            if(isalpha(word[0])){
                if(node->next[word[0]-'a']){
                    // 下一位存在
                    return searchWord(node->next[word[0]-'a'],word.substr(1,word.size()-1));
                }else{
                    // 下一位不存在
                    return false;
                }
            }else{
                // . 故只需要下一位存在则尝试判断能否一路到底
                bool b=false;
                for(int i=0;i<26;i++){
                    if(node->next[i]){
                        b=searchWord(node->next[i],word.substr(1,word.size()-1));
                        if(b){
                            return true;
                        }
                    }
                }
                return false;

            }
        }
    }
    bool isEnd;
    WordDictionary*next[26];
};

/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary* obj = new WordDictionary();
 * obj->addWord(word);
 * bool param_2 = obj->search(word);
 */
