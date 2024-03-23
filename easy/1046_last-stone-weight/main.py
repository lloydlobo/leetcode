class Solution(object):
    def lastStoneWeight(self, stones):
        """
        :type stones: List[int]
        :rtype: int
        """

        while len(stones)>1:
            stones.sort(reverse=True)

            if stones[0] == stones[1]:
                del stones[1]
                del stones[0]
            elif stones[0] != stones[1]:
                stones[0]-=stones[1]
                del stones[1]

        return stones[0] if stones else 0
        