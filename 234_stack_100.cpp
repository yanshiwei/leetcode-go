/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    bool isPalindrome(ListNode* head) {
        stack<int>st;
        ListNode*cur=head;
        while(cur){
            st.push(cur->val);
            cur=cur->next;
        }
        cur=head;
        while(cur){
            if(cur->val!=st.top()){
                return false;
            }
            st.pop();
            cur=cur->next;
        }
        return true;
    }
};
