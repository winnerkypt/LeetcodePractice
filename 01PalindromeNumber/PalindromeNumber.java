
/*
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

For example, 121 is a palindrome while 123 is not.
 

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/
package PalindromeNumber;
public class PalindromeNumber {
    public static boolean isPalindrome(int x) {
        String num = String.valueOf(x);
        StringBuilder input1 = new StringBuilder();
        input1.append(num);
        input1.reverse();
        System.out.println(input1.toString().equals(num));
        return input1.toString().equals(num);
            
    }

    public static void main(String[] args) {
        isPalindrome(121);
        isPalindrome(-121);
        isPalindrome(100021);
    }
}

