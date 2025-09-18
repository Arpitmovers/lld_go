// https://akhileshmj.medium.com/lld-7-api-rate-limiter-2a41c327ac66



Fixed size window
 we will create fixed-size time frame windows, where each window has a counter responsible for counting the number of hits get in a particular time frame.
If the number of hits exceeds the threshold, all the requests in that window will be discarded.



Sliding window counter


