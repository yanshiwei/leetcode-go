/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
    //剑指 Offer II 022. 链表中环的入口节点
public:
    ListNode *detectCycle(ListNode *head) {
        if(head==nullptr||head->next==nullptr){
            return nullptr;
        }
        ListNode*slow=head;
        ListNode*fast=head;
        while(1){
            if(fast==nullptr||fast->next==nullptr){
                return nullptr;// no loop
            }
            slow=slow->next;
            fast=fast->next->next;
            if (fast == slow) break;
        }
        //相遇点即是入口位置，f=2s=s+nb,相遇点时从头节点出发的快指针走了a+nb步，
        //而慢指针走了nb步则慢指针只需再走a步即可到达头节点。
        //反过来，从头节点一步步走当等于入口节点时候，正好是a
        fast=head;
        while(slow!=fast){
            slow=slow->next;
            fast=fast->next;
        }
        return fast;
    }
};
