class StackError(Exception):
    pass


class StackEmptyError(StackError):
    pass


class StackFullError(StackError):
    pass


class Stack(object):
    """ Stack class"""

    def __init__(self, size):
        """size: the maximum capacity of the stack
           ds: underlying data structure
        """
        self.size = size
        self.ds = []

    @property
    def length(self):
        """ the length of the stack"""
        return len(self.ds)

    def pop(self):
        """ return and remove the top element of the stack"""
        if not self.is_empty():
            return self.ds.pop()
        else:
            raise StackEmptyError("stack is empty!")

    def push(self, value):
        """ push an element to the top of the stack"""
        if self.length < self.size:
            self.ds.append(value)
        else:
            raise StackFullError("stack is full!")

    def push_many(self, *args):
        """ push many values all at once in sequential order"""
        for i in args:
            self.push(i)

    def is_empty(self):
        """ returns True if stack is empty"""
        return self.length == 0

    def peek(self):
        """ returns the top element of the stack"""
        if self.is_empty():
            raise StackEmptyError("stack is empty!")
        else:
            return self.ds[-1]

    def max(self):
        """ returns the element with the maximum value"""
        if not self.is_empty():
            cpy = self.ds.copy()
            cpy.sort()
            return cpy[-1]
        else:
            raise StackEmptyError("stack is empty")

    def sort(self, *args, **kwargs):
        """ sort the stack"""
        self.ds.sort(*args, **kwargs)   # pass whatever arguments passed in to list.sort

    def reverse(self):
        """ reverse the stack"""
        self.ds.reverse()

    def __str__(self):
        """ str representation of the stack"""
        fmt_strs = [("|    %+2s    |\n|__________|\n"
                     % element) for element in self.ds[::-1]]  # print the last element first
        return ''.join(fmt_strs)


def stack_test():
    my_stack = Stack(10)
    my_stack.push(2)
    my_stack.push(9)
    my_stack.push_many(4, 1, 51, 8, 10, -4)
    print("length of the stack:", my_stack.length)
    print("maximum in the stack: ", my_stack.max())
    print("top element: ", my_stack.peek())
    print(my_stack)
    my_stack.reverse()
    print(my_stack)
    my_stack.sort()
    print("sorted stack:")
    print(my_stack)
    try:
        for i in range(100):
            my_stack.pop()
    except StackEmptyError:
        print("too many elements to pop, stack is empty!")


if __name__ == "__main__":
    stack_test()

# |    51    |
# |__________|
# |    10    |
# |__________|
# |     9    |
# |__________|
# |     8    |
# |__________|
# |     4    |
# |__________|
# |     2    |
# |__________|
# |     1    |
# |__________|
# |    -4    |
# |__________|





