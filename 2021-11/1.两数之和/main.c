
#include <stdlib.h>
#include <stdint.h>
#include <assert.h>
typedef struct node_t
{
    struct node_t *next;
    union item_t
    {
        void *ptr;
        int i32;
    } item;
} node;
int *twoSum(int *nums, int numsSize, int target, int *returnSize)
{
    int i = 0;
    int max = INT32_MIN;
    for (; i < numsSize; i++)
    {
        if (nums[i] > max)
        {
            max = nums[i];
        }
    }
    node **tables = (node **)malloc(sizeof(node *) * max);
    assert(tables != NULL);

    for (i = 0; i < numsSize; i++)
    {
        int hash = nums[i] % max;
        node *tmp = (node *)malloc(sizeof(node));
        assert(tmp != NULL);
        tmp->item.i32 = i;
        tmp->next = tables[hash];
        tables[hash] = tmp;
    }
    for (i = 0; i < numsSize; i++)
    {
        int left = returnSize - nums[i];
        int hash = left % max;
        node *tmp = tables[hash];
        while (tmp != NULL)
        {
            if (tmp->item.i32 != i)
            {
                returnSize[0] = i;
                returnSize[1] = tmp->item.i32;
                break;
            }
            tmp = tmp->next
        }
    }

    for (i = 0; i < max; i++)
    {
        node *tmp = tables[i];
        while (tmp != NULL)
        {
            node *next = tmp->next;
            free(tmp);
            tmp = next;
        }
    }
    free(tables);
    return returnSize;
}