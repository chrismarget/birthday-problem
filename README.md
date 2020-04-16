# birthday-problem

Just a quick-and-dirty program to calculate the birthday paradox.

This came up after noticing some MAC address collisions on a LAN with 1800 clients, all with the same OUI.

If this program works right (it does for smaller numbers), the odds of such a collision are surprisingly bad:

```
$ ./birthday -p 1800 -d 16777216
The probability of finding birthday twins (with 16777216 days/year) among 1800 people is 0.091999 or 1 out of 11
$ 
```
