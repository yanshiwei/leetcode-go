/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* next;
    Node* random;
    
    Node(int _val) {
        val = _val;
        next = NULL;
        random = NULL;
    }
};
*/

class Solution {
    /*
    给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。
    */
public:
    Node*splitLinkedList(Node*head){
        if(head==NULL){
            return NULL;
        }
        Node*cur=head;
        Node*cloneHead=head->next;
        Node*cloneCur=head->next;
        // handle first
        cur->next=cloneCur->next;
        cur=cur->next;
        while(cur){
            cloneCur->next=cur->next;
            cloneCur=cloneCur->next;

            cur->next=cloneCur->next;
            cur=cur->next;
        }
        return cloneHead;
    }
    Node*cloneLinkedList(Node*head){
        if(head==NULL){
            return NULL;
        }
        Node*cur=head;
        while(cur){
            Node*tmp=new Node(cur->val);
            tmp->next=cur->next;
            tmp->random=NULL;

            cur->next=tmp;
            cur=tmp->next;
        }
        cur=head;
        while(cur){
            if(cur->random){
                cur->next->random=cur->random->next;
            }
            cur=cur->next->next;
        }
        return head;
    }
    Node* copyRandomList(Node* head) {
        Node*tmpHead=cloneLinkedList(head);
        return splitLinkedList(tmpHead);
    }
};
