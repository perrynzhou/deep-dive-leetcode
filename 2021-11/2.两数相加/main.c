
#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
struct ListNode
{
    int val;
    struct ListNode *next;
}__attribute__(8);

struct ListNode *init_list(int n,int *v)
{
    int i = 0;
    struct ListNode *head = NULL;
    struct ListNode *tmp = NULL;
    while (i < n)
    {
        struct ListNode *cur = malloc(sizeof(struct ListNode));
        assert(cur != NULL);
        int tmp_val = rand() % 10;
        if(v !=NULL) {
            tmp_val =  *v;
        }
        cur->val = tmp_val;
        if (head == NULL)
        {
            head = cur;
            tmp = cur;
        }
        else
        {
            tmp->next = cur;
            tmp = cur;
        }
        i++;
    }
    return head;
}
void print_list(struct ListNode *tmp, int n)
{
    int i = 0;
    struct ListNode *tail = NULL;
    while (i < n)
    {
        struct ListNode *next = tmp->next;
        fprintf(stdout, "%d ", tmp->val);
        tmp = next;
        tail = tmp;
        i++;
    }
    if(tail !=NULL) {
        fprintf(stdout, "%d ", tail->val);

    }
    fprintf(stdout, "\n");
}
struct ListNode *addTwoNumbers(struct ListNode *l1, struct ListNode *l2)
{
    if (l1 == NULL && l2 == NULL)
    {
        return NULL;
    }
    struct ListNode *head = NULL;
    struct ListNode *tmp = NULL;
    int step = 0;
    int sum = 0;
    while (l1 != NULL || l2 != NULL)
    {
        if (l1 != NULL)
        {
            sum += l1->val;
            l1 = l1->next;
        }
        if (l2 != NULL)
        {
            sum += l2->val;
            l2 = l2->next;
        }
        struct ListNode *cur = malloc(sizeof(struct ListNode));
        assert(cur!=NULL);
        // 9,9
        // 9,9,9,9
        int val = sum % 10 + step;
        if(sum>9) {
            step = 1;
        }
        cur->val = val%10;
        if(val>9) {
            step = 1;
        }
        fprintf(stdout,"val=%d,step=%d\n",cur->val,step);

        if (head == NULL)
        {
            head = cur;
            tmp = cur;
        }
        else
        {
            tmp->next = cur;
            tmp = cur;
        }
        sum = 0;
    }
    if(step ==1) {
         struct ListNode *cur = malloc(sizeof(struct ListNode));
         assert(cur!=NULL);
         cur->val = step;
         tmp->next = cur;
         tmp = cur;
    }
    return head;
}
int main(int argc, char *argv[])
{
    int m = atoi(argv[1]);
    int n = atoi(argv[2]);

    int tmp_val = 0;
    int *tmp_val_ptr =NULL;
    if(argv[3]!=NULL) {
        tmp_val = atoi(argv[3]);
        tmp_val_ptr  = &tmp_val;
    }
    int k = (m > n) ? m : n;
    struct ListNode *l1 = init_list(m,tmp_val_ptr);
    struct ListNode *l2 = init_list(n,tmp_val_ptr);
    print_list(l1, m);
    print_list(l2, n);
    struct ListNode *l3 = addTwoNumbers(l1, l2);
    print_list(l3,k);
}