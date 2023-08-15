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
    //翻转一个子链表，并且返回新的头与尾
    vector<ListNode*>reverse(ListNode*head,ListNode*tail){
        ListNode*pre=nullptr;
        ListNode*cur=head;
        while(pre!=tail){
            ListNode*next=cur->next;
            cur->next=pre;
            pre=cur;
            cur=next;
        }
        vector<ListNode*>res{tail,head};
        return res;
    }
    ListNode* reverseKGroup(ListNode* head, int k) {
        ListNode*dummy=new ListNode();
        dummy->next=head;
        ListNode*pre=dummy;
        ListNode*cur=head;
        while(cur){
            //查看剩余部分长度是否大于等于 k
            ListNode*tail=pre;
            for(int i=0;i<k;i++){
                tail=tail->next;
                if(tail==nullptr){
                    return dummy->next;
                }
            }
            // k长度的子链表，链表真实头部是cur,尾部是tail
            ListNode*next=tail->next;
            // 反转k长度的子链表
            auto res = reverse(cur,tail);
            //把子链表重新接回原链表
            pre->next = res[0];
            res[1]->next=next;
            //从下一个k段开始
            pre=res[1];
            cur=res[1]->next;
        }
        return dummy->next;
    }
};
