import math
import random


class TheoryOfNumbers:
    # task 1
    def __init__(self, m):
        self.m = m

    @classmethod
    def __euler(cls, n):
        f = n
        if n % 2 == 0:
            while n % 2 == 0:
                n = n // 2
            f = f // 2
        i = 3
        while i * i <= n:
            if n % i == 0:
                while n % i == 0:
                    n = n // i
                f = f // i
                f = f * (i - 1)
            i = i + 2
        if n > 1:
            f = f // n
            f = f * (n - 1)
        return f

    def __gcd(self, a):
        if a < self.m:
            a, self.m = self.m, a

        while self.m != 0:
            a, self.m = self.m, a % self.m

        return a

    def task_2(self, a):
        return a % self.m

    def task_3(self, a, b):
        return pow(a, b) % self.m

    def task_4(self, a, b, m):
        x = (b * pow(a, self.__euler(m) - self.__gcd(a))) % m
        return x

    def task_5(self, a, b):
        res = []

        for i in range(a, b + 1):

            for j in range(2, i):
                if i % j == 0:
                    break

            else:
                res.append(i)
        return random.choice(res)


t = TheoryOfNumbers(53)
print(t.task_4(13, 2, 53))
print(t.task_5(24, 33))
