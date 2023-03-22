class Solution
{
public:
    bool isValid(string s)
    {
        stack<char> check;
        for (int i = 0; i < s.length(); i++)
        {
            if (s[i] == '(' || s[i] == '[' || s[i] == '{')
            {
                check.push(s[i]);
            }
            else if (s[i] == ')')
            {
                if (i == 0 || check.empty() || check.top() != '(')
                {
                    return false;
                }
                check.pop();
            }
            else if (s[i] == ']')
            {
                if (i == 0 || check.empty() || check.top() != '[')
                {
                    return false;
                }
                check.pop();
            }
            else if (s[i] == '}')
            {
                if (i == 0 || check.empty() || check.top() != '{')
                {
                    return false;
                }
                check.pop();
            }
        }
        return check.empty();
    }
};