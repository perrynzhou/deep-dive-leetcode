
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
int *twoSum(int *nums, int numsSize, int target, int *returnSize)
{
	int max = -1;
	for (int i = 0; i < numsSize; i++)
	{
		if (max == -1)
		{
			max = nums[i];
		}
		else
		{
			max = (nums[i] > max) ? nums[i] : max;
		}
	}
	size_t alloc_sz = max / sizeof(uint8_t);
	if (alloc_sz % sizeof(uint8_t) != 0)
	{
		++alloc_sz;
	}
	int *result = calloc(*returnSize,sizeof(int));
	uint8_t *bits = calloc(alloc_sz, sizeof(uint8_t));
	for (int i = 0; i < numsSize; i++)
	{
		int index = nums[i] >> 3;
		bits[index] |= (1 << (nums[i] & 7));
	}
	bool flag = false;
	int left = -1;
	for (int i = 0; i < numsSize; i++)
	{
		if (!flag)
		{
			left = target - nums[i];
			int index = left >> 3;
			int ret = bits[index] & (1 << (left & 7));
			if (ret == 1)
			{
				result[0]=i;
				flag = true;
			}
		}else {
			if(nums[i]==left) {
				result[1]=i;
				break;
			}
		}
	}
	if (bits != NULL)
	{
		free(bits);
	}
	return result;
}