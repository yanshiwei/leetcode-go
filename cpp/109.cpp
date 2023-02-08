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
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    // 给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。
    // 类似108题，只不过此处是链表，故关键点怎么高效找链表的中间点
    // 快慢指针法，当快指针到末尾时，慢指针就是中间节点
    TreeNode*sortedListToBSTHelper(ListNode*head,ListNode*tail){
        if(head==tail){
            return nullptr;
        }
        // find mid
        ListNode* fast = head;
        ListNode* slow = head;
        while (fast != tail && fast->next != tail) {
            fast = fast->next;
            fast = fast->next;
            slow = slow->next;
        }
        TreeNode* root = new TreeNode(slow->val);
        root->left = sortedListToBSTHelper(head, slow);
        root->right = sortedListToBSTHelper(slow->next, tail);
        return root;
    }
    TreeNode* sortedListToBST(ListNode* head) {
        if(head==nullptr){
            return nullptr;
        }
        return sortedListToBSTHelper(head,nullptr);
    }
};
