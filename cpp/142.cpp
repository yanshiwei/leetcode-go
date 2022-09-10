/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
/*
  设链表总长度L，头节点到入口长a，环r，入口到相遇点x
   slow走了s步，fast走了2s，且fast在相遇时候比slow多走nr圈故有
   s=nr,而s=a+x
   so a+x=nr=(n-1)r+r=(n-1)r+L-a
   so a=(n-1)r+L-a-x
   从头结点到入口的距离，等于转了(n-1)圈以后，相遇点到入口的距离
*/
    ListNode *detectCycle(ListNode *head) {
        if(head==NULL||head->next==NULL){
            return NULL;
        }
        // 先找到相遇点
        ListNode*slow=head;
        ListNode*fast=head;
        while(1){
            if(fast==NULL||fast->next==NULL){
                // no cycle
                return NULL;
            }
            slow=slow->next;
            fast=fast->next->next;
            if(slow==fast){
                break;
            }
        }
        //从头结点到入口的距离，等于转了(n-1)圈以后，相遇点到入口的距离
        ListNode*p=head;
        while(slow!=p){
            slow=slow->next;// 从头节点开始
            p=p->next;// 从相遇点开始
        }
        return p;
    }
};
