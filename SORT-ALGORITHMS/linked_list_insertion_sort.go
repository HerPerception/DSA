/*=================================================
Leetcode task for insertion sort for linked list.
===================================================*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    current := head
    dummy := &ListNode{}
    for current != nil {
        nextNode := current.Next
        search := dummy
        for search.Next != nil && search.Next.Val < current.Val {
            search = search.Next
        }
        current.Next = search.Next
        search.Next = current
        current = nextNode
    }
    return dummy.Next
}
