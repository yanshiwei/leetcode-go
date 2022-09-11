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
    /*
    给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
    */
public:
    ListNode* oddEvenList(ListNode* head) {
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        ListNode*p=head;
        ListNode*dummyEvenHead=new ListNode();
        ListNode*curEven=nullptr;
        ListNode*dummyOddHead=new ListNode();
        ListNode*curOdd=nullptr;
        int idx=0;
        while(p){
            idx++;
            ListNode *next=p->next;
            if((idx&1)==1){
                // 奇数
                if(curOdd==nullptr){
                    // first
                    dummyOddHead->next=p;
                    curOdd=p;
                }else{
                    // hase node
                    curOdd->next=p;
                    curOdd=p;
                }
            }else{
                // 偶数
                if(curEven==nullptr){
                    // first
                    dummyEvenHead->next=p;
                    curEven=p;
                }else{
                    // hase node
                    curEven->next=p;
                    curEven=p;
                }  
            }
            p->next=nullptr;
            p=next;
        }
        // merge tow lists
        ListNode* firstEvenNode=dummyEvenHead->next;
        curOdd->next=firstEvenNode;
        return dummyOddHead->next;
    }
};
